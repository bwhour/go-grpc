package trace

import (
	"context"
	"reflect"
	"testing"
)

func TestContextWithInfo(t *testing.T) {
	ctx := context.Background()
	expected := &Info{
		Traceid:   MakeTraceid(0),
		Spanid:    MakeSpanid(0, 0),
		Logid:     MakeLogid(0),
		SrcMethod: "test",
		Caller:    "caller",
		Callee:    "callee",
	}
	ctx = WithInfo(ctx, expected)
	actual := parseInfo(ctx)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("fail to set info into context. [expected:%v] [actual:%v]", expected, actual)
	}
}

func TestInfoMerge(t *testing.T) {
	expected := &Info{
		Traceid:   MakeTraceid(0),
		Spanid:    MakeSpanid(0, 0),
		Logid:     MakeLogid(0),
		SrcMethod: "test",
		Caller:    "caller",
		Callee:    "callee",
	}
	actual := &Info{}
	actual.Merge(expected)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("fail to merge info. [expected:%v] [actual:%v]", expected, actual)
	}
}
