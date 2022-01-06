package trace

import (
	"context"
	"testing"
	"time"
)

func TestZeroValue(t *testing.T) {
	v := &value{}

	v.AdjustTime()

	if !v.Now.IsZero() {
		t.Fatalf("now should be zero. [v:%v]", v)
	}

	if _, ok := v.Deadline(); ok {
		t.Fatalf("deadline should be invalid.")
	}
}

func TestValue(t *testing.T) {
	now := time.Now()
	timeout := 25 * time.Millisecond
	expected := &value{
		Info: Info{
			Traceid:   MakeTraceid(0),
			Spanid:    MakeSpanid(0, 0),
			Logid:     MakeLogid(0),
			SrcMethod: "src/method",
			Caller:    "caller",
			Callee:    "callee",
		},
		Now:         now,
		Timeout:     timeout * 4,
		ElapsedTime: timeout,
	}

	if deadline, ok := expected.Deadline(); !ok || !deadline.Equal(now.Add(expected.Timeout-expected.ElapsedTime)) {
		t.Fatalf("invalid deadline. [expected:%v] [actual:%v]", now.Add(expected.Timeout-expected.ElapsedTime), deadline)
	}

	ctx, cancel := context.WithDeadline(context.Background(), now.Add(timeout))
	defer cancel()
	tr := expected.Trace(ctx)

	if tr.Traceid() != expected.Traceid {
		t.Fatalf("traceid is invalid. [expected:%v] [actual:%v]", expected.Traceid, tr.Traceid())
	}

	if tr.Spanid() != expected.Spanid {
		t.Fatalf("spanid is invalid. [expected:%v] [actual:%v]", expected.Spanid, tr.Spanid())
	}

	if tr.Logid() != expected.Logid {
		t.Fatalf("logid is invalid. [expected:%v] [actual:%v]", expected.Logid, tr.Logid())
	}

	if tr.SrcMethod() != expected.SrcMethod {
		t.Fatalf("src method is invalid. [expected:%v] [actual:%v]", expected.SrcMethod, tr.SrcMethod())
	}

	if tr.Caller() != expected.Caller {
		t.Fatalf("caller is invalid. [expected:%v] [actual:%v]", expected.Caller, tr.Caller())
	}

	if tr.Callee() != expected.Callee {
		t.Fatalf("callee is invalid. [expected:%v] [actual:%v]", expected.Callee, tr.Callee())
	}

	if now := tr.Now(); !now.Equal(expected.Now) {
		t.Fatalf("now is invalid. [expected:%v] [actual:%v]", expected.Now, now)
	}

	if tr.timeout() != timeout {
		t.Fatalf("timeout is invalid. [expected:%v] [actual:%v]", timeout, tr.timeout())
	}

	if tr.ElapsedTime() != 0 {
		t.Fatalf("elapsed time is invalid. [expected:%v] [actual:%v]", 0, tr.ElapsedTime())
	}
}
