// Package trace 实现了符合公司规范的 trace/span/log id 的生成，并存储在一个标准的 context 里面透明传输。
package trace

import (
	"bytes"
	"context"
	"math"
	"strconv"
	"time"
)

const (
	keyTraceid     = "traceid"
	keySpanid      = "spanid"
	keyLogid       = "logid"
	keySrcMethod   = "src_method"
	keyCaller      = "caller"
	keyCallee      = "callee"
	keyNow         = "now"
	keyTimeout     = "timeout"
	keyElapsedTime = "elapsed_time"
)

const maxTimeout time.Duration = math.MaxInt64

// NewContext 创建新的 Context，并继承上游的 trace 信息。
func NewContext(ctx context.Context, tr Trace) context.Context {
	v := newValue(ctx, tr)
	return newTimeoutContext(ctx, v)
}

// FromContext 从 ctx 中获得 Trace 信息，如果 trace 信息不存在则会创建一个全新的 trace 信息。
func FromContext(ctx context.Context) Trace {
	v := parseValue(ctx)

	if v == nil {
		v = newValue(ctx, nil)
	}

	return v.Trace(ctx)
}

// FromContextMayEmpty 从 ctx 中获得 Trace 信息，如果 ctx 里面没有存储 trace 信息会返回 nil。
func FromContextMayEmpty(ctx context.Context) Trace {
	v := parseValue(ctx)

	if v == nil {
		return nil
	}

	return v.Trace(ctx)
}

var (
	emptyValue value

	emptyContextString = makeLogString("", "", "")
)

// ContextString 返回 ctx 的 traceid 等信息，用于输出日志。
// 基本等价于 trace.FromContext(ctx).String()，但是性能高很多。
func ContextString(ctx context.Context) string {
	v, okValue := ctx.Value(keyForValueOfContext).(*value)
	info, okInfo := ctx.Value(keyForInfoOfContext).(*Info)

	if !okValue && !okInfo {
		return emptyContextString
	}

	if !okValue {
		v = &emptyValue
	}

	vInfo := v.Info
	vInfo.Merge(info)
	return makeLogString(vInfo.Traceid, vInfo.Spanid, vInfo.Logid)
}

// ContextCaller 返回 ctx 的 caller 信息，用于高效的。
// 基本等价于 trace.FromContext(ctx).Caller()，但是性能高很多。
func ContextCaller(ctx context.Context) string {
	info, okInfo := ctx.Value(keyForInfoOfContext).(*Info)

	if okInfo && info.Caller != "" {
		return info.Caller
	}

	v, okValue := ctx.Value(keyForValueOfContext).(*value)

	if okValue {
		return v.Caller
	}

	return ""
}

// ContextCallee 返回 ctx 的 callee 信息，用于高效的。
// 基本等价于 trace.FromContext(ctx).Callee()，但是性能高很多。
func ContextCallee(ctx context.Context) string {
	info, okInfo := ctx.Value(keyForInfoOfContext).(*Info)

	if okInfo && info.Callee != "" {
		return info.Callee
	}

	v, okValue := ctx.Value(keyForValueOfContext).(*value)

	if okValue {
		return v.Callee
	}

	return ""
}

func makeLogString(traceid Traceid, spanid Spanid, logid Logid) string {
	buf := &bytes.Buffer{}
	buf.Grow(len(traceid) + len(spanid) + len(logid) + len(traceidLogKeyBytes) + len(spanidLogKeyBytes) + len(logidLogKeyBytes) + 2*len(logidLogKeyBytes))
	buf.Write(traceidLogKeyBytes)
	buf.WriteString(string(traceid))
	buf.Write(logKeySeparator)
	buf.Write(spanidLogKeyBytes)
	buf.WriteString(string(spanid))
	buf.Write(logKeySeparator)
	buf.Write(logidLogKeyBytes)
	buf.WriteString(string(logid))

	return makeString(buf.Bytes())
}

// Trace 用来记录各种 trace 信息。
type Trace map[string]string

var (
	traceidLogKeyBytes = []byte("traceid=")
	spanidLogKeyBytes  = []byte("spanid=")
	logidLogKeyBytes   = []byte("logid=")
	logKeySeparator    = []byte("||")
)

// String 返回 t 的标准日志格式。
func (tr Trace) String() string {
	traceid, spanid, logid := tr.Traceid(), tr.Spanid(), tr.Logid()
	return makeLogString(traceid, spanid, logid)
}

// Traceid 返回 trace id 的值。
func (tr Trace) Traceid() Traceid {
	return Traceid(tr[keyTraceid])
}

// setTraceid 设置新的 trace id。
func (tr Trace) setTraceid(traceid Traceid) {
	tr[keyTraceid] = traceid.String()
}

// Spanid 返回 span id 的值。
func (tr Trace) Spanid() Spanid {
	return Spanid(tr[keySpanid])
}

// setSpanid 设置新的 span id。
func (tr Trace) setSpanid(spanid Spanid) {
	tr[keySpanid] = spanid.String()
}

// Logid 返回 log id 的值。
func (tr Trace) Logid() Logid {
	return Logid(tr[keyLogid])
}

// setLogid 设置新的 log id。
func (tr Trace) setLogid(logid Logid) {
	tr[keyLogid] = logid.String()
}

// SrcMethod 返回 src method 的值。
func (tr Trace) SrcMethod() string {
	return tr[keySrcMethod]
}

// setSrcMethod 设置新的 src method。
func (tr Trace) setSrcMethod(srcMethod string) {
	tr[keySrcMethod] = srcMethod
}

// Caller 返回 caller 的值。
func (tr Trace) Caller() string {
	return tr[keyCaller]
}

// setCaller 设置新的 caller。
func (tr Trace) setCaller(caller string) {
	tr[keyCaller] = caller
}

// Callee 返回 callee 的值。
func (tr Trace) Callee() string {
	return tr[keyCallee]
}

// setCallee 设置新的 callee。
func (tr Trace) setCallee(callee string) {
	tr[keyCallee] = callee
}

// Now 返回 trace 里记录的当前时间。
func (tr Trace) Now() time.Time {
	if now, ok := tr[keyNow]; ok && now != "" {
		nano := fastAtoi(now)
		return time.Unix(nano/int64(time.Second), nano%int64(time.Second))
	}

	return time.Now()
}

func (tr Trace) setNow(now time.Time) {
	tr[keyNow] = strconv.FormatInt(now.UnixNano(), 10)
}

// Timeout 返回 trace 里记录的超时时间，
// 这个值减去从 trace 创建开始到现在流逝的时间（ElapsedTime），如果当前已经超过超时时间，则会返回 0。
func (tr Trace) Timeout() time.Duration {
	if timeout, ok := tr[keyTimeout]; ok && timeout != "" {
		elapsed := tr.ElapsedTime()
		traceNow := tr.Now()
		now := time.Now()
		return time.Duration(fastAtoi(timeout)) - elapsed - now.Sub(traceNow)
	}

	return maxTimeout
}

func (tr Trace) timeout() time.Duration {
	if timeout, ok := tr[keyTimeout]; ok && timeout != "" {
		return time.Duration(fastAtoi(timeout))
	}

	return 0
}

func (tr Trace) setTimeout(timeout time.Duration) {
	tr[keyTimeout] = strconv.FormatInt(timeout.Nanoseconds(), 10)
}

// ElapsedTime 返回 trace 里记录的从最上游服务发起调用开始，已经经过的时间。
func (tr Trace) ElapsedTime() time.Duration {
	if elapsed, ok := tr[keyElapsedTime]; ok && elapsed != "" {
		return time.Duration(fastAtoi(elapsed))
	}

	return 0
}

func (tr Trace) setElapsedTime(elapsed time.Duration) {
	tr[keyElapsedTime] = strconv.FormatInt(elapsed.Nanoseconds(), 10)
}

// Deadline 返回当前 t 的超时时间，如果没有设置超时则 ok 为 false。
func (tr Trace) Deadline() (deadline time.Time, ok bool) {
	timeout := tr.timeout()

	if timeout != 0 {
		now := tr.Now()
		elapsed := tr.ElapsedTime()
		deadline = now.Add(timeout - elapsed)
		ok = true
		return
	}

	return
}

// fastAtoi 假定 s 一定是正确的数字字符串，因此可以更快的解析出 s 的值。
func fastAtoi(s string) (n int64) {
	for i := 0; i < len(s); i++ {
		n *= 10
		n += int64(s[i] - '0')
	}

	return
}
