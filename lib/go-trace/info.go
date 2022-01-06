package trace

import (
	"context"
)

type infoOfContext struct{}

var (
	keyForInfoOfContext = infoOfContext{}
)

// Info 存储了 trace 的必要信息。
//
// 注意：这里并不包含任何跟超时相关的信息，
// 如果要设置超时，应该使用 `context.WithDeadline` 或 `context.WithTimeout`。
type Info struct {
	Traceid Traceid
	Spanid  Spanid
	Logid   Logid

	SrcMethod string
	Caller    string
	Callee    string
}

// Merge 合并 target 到 info，只有 target 里面有值的字段才会被复制过来。
func (info *Info) Merge(target *Info) {
	if target == nil {
		return
	}

	if target.Traceid != "" {
		info.Traceid = target.Traceid
	}

	if target.Spanid != "" {
		info.Spanid = target.Spanid
	}

	if target.Logid != "" {
		info.Logid = target.Logid
	}

	if target.SrcMethod != "" {
		info.SrcMethod = target.SrcMethod
	}

	if target.Caller != "" {
		info.Caller = target.Caller
	}

	if target.Callee != "" {
		info.Callee = target.Callee
	}
}

// WithInfo 将自定义的 trace 信息加入到 ctx 里面。
func WithInfo(ctx context.Context, info *Info) context.Context {
	return context.WithValue(ctx, keyForInfoOfContext, info)
}

func parseInfo(ctx context.Context) *Info {
	info, ok := ctx.Value(keyForInfoOfContext).(*Info)

	if !ok {
		return nil
	}

	return info
}
