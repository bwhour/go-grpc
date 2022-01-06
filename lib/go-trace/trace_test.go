package trace

import (
	"context"
	"fmt"
	"testing"
	"time"

	log "gitlab.mydadao.com/brain-infrastructure/go-log"
)

func TestTrace(t *testing.T) {
	traceid := "abcdef"
	spanid := "123456"
	now := time.Now()
	timeout := 10 * time.Millisecond
	time.Sleep(10 * time.Millisecond)

	ctx := NewContext(context.Background(), Trace{
		keyTraceid: traceid,
		keySpanid:  spanid,
		keyTimeout: fmt.Sprint(int64(timeout)),
		keyNow:     fmt.Sprint(now.UnixNano()),
	})
	tr := FromContext(ctx)

	if tr.Traceid().String() != traceid {
		t.Fatalf("trace id should be derived. [expected:%v] [actual:%v]", traceid, tr.Traceid())
	}

	if tr.Spanid().String() == spanid || !tr.Spanid().IsValid() {
		t.Fatalf("span id should be not derived. [spanid:%v]", spanid)
	}

	if actual := tr.Timeout(); actual > timeout {
		t.Fatalf("timeout should be derived. [expected:%v] [actual:%v]", timeout, actual)
	}

	if n := tr.Now(); n.Equal(now) || n.Before(now) {
		t.Fatalf("now should not be derived. [prev:%v] [cur:%v]", now, n)
	} else if deadline, ok := tr.Deadline(); !ok {
		t.Fatalf("deadline is not available.")
	} else if expected := n.Add(tr.timeout() - tr.ElapsedTime()); deadline != expected {
		t.Fatalf("deadline is invalid. [expected:%v] [actual:%v]", expected, deadline)
	}

	if deadline, ok := ctx.Deadline(); !ok || deadline.Add(-timeout).After(time.Now()) {
		t.Fatalf("deadline is too large or invalid. [deadline:%v]", deadline)
	}

	select {
	case <-ctx.Done():
	case <-time.After(timeout * 2):
		t.Fatalf("fail to wait for ctx timeout.")
	}

	if err := ctx.Err(); err != context.DeadlineExceeded {
		t.Fatalf("ctx timeout err is invalid. [expected:%v] [actual:%v]", context.DeadlineExceeded, err)
	}
}

func TestEmptyTrace(t *testing.T) {
	tr := Trace{}

	if delta := tr.Now().Sub(time.Now()); delta > 0 || delta < -10*time.Microsecond {
		t.Fatalf("trace now should be the same as time.Now(). [delta:%v]", delta)
	}

	if timeout := tr.Timeout(); timeout != maxTimeout {
		t.Fatalf("trace timeout should be an invalid value. [timeout:%v]", timeout)
	}

	if elapsed := tr.ElapsedTime(); elapsed != 0 {
		t.Fatalf("trace elapsed should be 0. [elapsed:%v]", elapsed)
	}
}

func TestTraceGetterSetters(t *testing.T) {
	if tr := FromContextMayEmpty(context.Background()); tr != nil {
		t.Fatalf("tr should be nil in an empty context. [tr:%v]", tr)
	}

	ctx := NewContext(context.Background(), nil)

	if tr := FromContextMayEmpty(ctx); tr == nil {
		t.Fatalf("tr should not be nil in a ctx created by NewContext.")
	}

	tr := FromContext(context.Background())

	if tr == nil {
		t.Fatalf("tr should be not nil as normalized.")
	}

	if deadline, ok := tr.Deadline(); ok || !deadline.IsZero() {
		t.Fatalf("tr should not have deadline. [ok:%v] [deadline:%v]", ok, deadline)
	}

	traceid := MakeTraceid(0)
	tr.setTraceid(traceid)

	if actual := tr.Traceid(); actual != traceid {
		t.Fatalf("fail to set traceid. [expected:%v] [actual:%v]", traceid, actual)
	}

	spanid := MakeSpanid(0, 0)
	tr.setSpanid(spanid)

	if actual := tr.Spanid(); actual != spanid {
		t.Fatalf("fail to set spanid. [expected:%v] [actual:%v]", spanid, actual)
	}

	logid := MakeLogid(0)
	tr.setLogid(logid)

	if actual := tr.Logid(); actual != logid {
		t.Fatalf("fail to set logid. [expected:%v] [actual:%v]", logid, actual)
	}

	srcMethod := "foo"
	tr.setSrcMethod(srcMethod)

	if actual := tr.SrcMethod(); actual != srcMethod {
		t.Fatalf("fail to set srcMethod. [expected:%v] [actual:%v]", srcMethod, actual)
	}

	caller := "bar"
	tr.setCaller(caller)

	if actual := tr.Caller(); actual != caller {
		t.Fatalf("fail to set caller. [expected:%v] [actual:%v]", caller, actual)
	}

	callee := "player"
	tr.setCallee(callee)

	if actual := tr.Callee(); actual != callee {
		t.Fatalf("fail to set callee. [expected:%v] [actual:%v]", callee, actual)
	}

	if expected := fmt.Sprintf("traceid=%v||spanid=%v||logid=%v", traceid, spanid, logid); expected != tr.String() {
		t.Fatalf("invalid tr.String(). [expected:%v] [actual:%v]", expected, tr)
	}

	tr = FromContext(ctx)

	if expected, actual := fmt.Sprintf("traceid=%v||spanid=%v||logid=%v", tr.Traceid(), tr.Spanid(), tr.Logid()), ContextString(ctx); expected != actual {
		t.Fatalf("invalid trace.ContextString(ctx). [expected:%v] [actual:%v]", expected, actual)
	}

	if expected, actual := tr.Caller(), ContextCaller(ctx); expected != actual {
		t.Fatalf("invalid trace.ContextCaller(ctx). [expected:%v] [actual:%v]", expected, actual)
	}

	if expected, actual := "context-caller", ContextCaller(WithInfo(ctx, &Info{Caller: "context-caller"})); expected != actual {
		t.Fatalf("invalid trace.ContextCaller(ctx). [expected:%v] [actual:%v]", expected, actual)
	}

	if expected, actual := tr.Callee(), ContextCallee(ctx); expected != actual {
		t.Fatalf("invalid trace.ContextCallee(ctx). [expected:%v] [actual:%v]", expected, actual)
	}

	if expected, actual := "context-callee", ContextCallee(WithInfo(ctx, &Info{Callee: "context-callee"})); expected != actual {
		t.Fatalf("invalid trace.ContextCallee(ctx). [expected:%v] [actual:%v]", expected, actual)
	}

	ctx = context.Background()

	if expected, actual := emptyContextString, ContextString(ctx); expected != actual {
		t.Fatalf("invalid trace.ContextString(ctx). [expected:%v] [actual:%v]", expected, actual)
	}

	if expected, actual := "", ContextCaller(ctx); expected != actual {
		t.Fatalf("invalid trace.ContextCaller(ctx). [expected:%v] [actual:%v]", expected, actual)
	}

	if expected, actual := "", ContextCallee(ctx); expected != actual {
		t.Fatalf("invalid trace.ContextCallee(ctx). [expected:%v] [actual:%v]", expected, actual)
	}
}

func BenchmarkTraceString(b *testing.B) {
	ctx := NewContext(context.Background(), nil)

	for i := 0; i < b.N; i++ {
		tr := FromContext(ctx)
		_ = tr.String()
	}
}

func BenchmarkTraceContextString(b *testing.B) {
	ctx := NewContext(context.Background(), nil)

	for i := 0; i < b.N; i++ {
		ContextString(ctx)
	}
}

func TestTraceDetail(t *testing.T) {
	initLog()
	defer CloseLog()

	ctx := context.Background()

	tr := FromContext(ctx)

	log.Debugf("trace 1 ||%v", tr)
	log.Debugf("trace 2 ||%v", tr)
	log.Debugf("trace 3 ||%v", tr)

	parentCtx := NewContext(ctx, tr)

	log.Debugf("trace 4 ||%v", FromContext(parentCtx))

	chileCtx, cancel := context.WithTimeout(parentCtx, 100*time.Second)
	defer cancel()

	log.Debugf("trace 4 ||%v", FromContext(chileCtx))

	name(chileCtx)
}

func name(ctx context.Context) {
	log.Debugf("trace name||%v", FromContext(ctx))
	return
}
