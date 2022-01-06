package trace

import (
	"context"
	"math/rand"
	"sync"
	"time"
)

var (
	traceRand      = rand.New(rand.NewSource(time.Now().UnixNano()))
	traceRandMutex = sync.Mutex{}
)

// value 记录了 context 里面的 trace 信息。
//
// 这里为了能够实现跨越 rpc 边界传递 trace 信息，特别是时间相关信息，
// 专门设计了 Now、Timeout、ElapsedTime 这三个字段来持久化这些信息。
// 对于任何一个新创建的 trace context 来说：
//
//     * Now 记录的是创建 context的时间。
//     * Timeout 是最上游设定的超时时间，层层传递到最底层服务。
//     * ElapsedTime 是上游已经消耗掉的时间，这样就可以实现跨越服务边界
//       的超时控制，下游服务的 deadline 始终是 now + timeout - elapsed。
type value struct {
	mu sync.Mutex

	Info
	infoString string // 缓存了 Info 的 String 结果。

	Now         time.Time
	Timeout     time.Duration
	ElapsedTime time.Duration
}

func newValue(ctx context.Context, tr Trace) *value {
	v, ok := ctx.Value(keyForValueOfContext).(*value)

	if !ok || v == nil {
		v = &value{}
	}

	v.merge(tr)
	v.normalize()
	v.infoString = makeLogString(v.Traceid, v.Spanid, v.Logid)
	return v
}

func parseValue(ctx context.Context) *value {
	v, ok := ctx.Value(keyForValueOfContext).(*value)

	if !ok {
		return nil
	}

	v.AdjustTime()
	return v
}

// Merge 从 tr 里继承部分数据。
func (v *value) merge(tr Trace) {
	if tr == nil {
		return
	}

	if traceid := tr.Traceid(); traceid != "" {
		v.Traceid = traceid
	}

	if timeout := tr.timeout(); timeout > 0 {
		v.Timeout = timeout
	}

	if elapsed := tr.ElapsedTime(); elapsed > 0 {
		v.ElapsedTime = elapsed
	}

	// tr 里面的 callee 就是当前服务，所以设置为新 value 的 caller。
	if callee := tr.Callee(); callee != "" {
		tr.setCaller(callee)
	}
}

func (v *value) normalize() {
	var rnd int64

	if needRnd := v.Spanid == ""; needRnd {
		traceRandMutex.Lock()
		rnd = traceRand.Int63()
		traceRandMutex.Unlock()
	}

	now := time.Now()
	nowNano := now.UnixNano()

	if v.Traceid == "" {
		v.Traceid = MakeTraceid(nowNano)
	}

	if v.Logid == "" {
		v.Logid = MakeLogid(nowNano)
	}

	if v.Spanid == "" {
		v.Spanid = MakeSpanid(nowNano, rnd)
	}

	if v.Now.IsZero() {
		v.Now = now
	}
}

func (v *value) AdjustTime() {
	if v.Now.IsZero() {
		return
	}

	v.mu.Lock()
	defer v.mu.Unlock()
	now := time.Now()
	v.ElapsedTime += now.Sub(v.Now)
	v.Now = now
}

func (v *value) Deadline() (deadline time.Time, ok bool) {
	if v.Timeout <= 0 {
		return
	}

	deadline = v.Now.Add(v.Timeout - v.ElapsedTime)
	ok = true
	return
}

func (v *value) Trace(ctx context.Context) Trace {
	info := parseInfo(ctx)
	vInfo := v.Info
	vInfo.Merge(info)

	tr := Trace{}

	if vInfo.Traceid != "" {
		tr.setTraceid(vInfo.Traceid)
	}

	if vInfo.Spanid != "" {
		tr.setSpanid(vInfo.Spanid)
	}

	if vInfo.Logid != "" {
		tr.setLogid(vInfo.Logid)
	}

	if vInfo.SrcMethod != "" {
		tr.setSrcMethod(vInfo.SrcMethod)
	}

	if vInfo.Caller != "" {
		tr.setCaller(vInfo.Caller)
	}

	if vInfo.Callee != "" {
		tr.setCallee(vInfo.Callee)
	}

	deadline, ok := ctx.Deadline()
	vDeadline, vOK := v.Deadline()
	vNow := v.Now
	timeout := v.Timeout
	elapsed := v.ElapsedTime

	// ctx 的超时时间比 v 还短，那么使用 ctx 的时间。
	if ok && (!vOK || deadline.Before(vDeadline)) {
		if vNow.IsZero() {
			vNow = time.Now()
		}

		timeout = deadline.Sub(vNow)
		elapsed = 0
	}

	if timeout > 0 {
		tr.setTimeout(timeout)

		if !vNow.IsZero() {
			tr.setNow(vNow)
		}

		if elapsed > 0 {
			tr.setElapsedTime(elapsed)
		}
	}

	return tr
}

func (v *value) LogString() string {
	return v.infoString
}
