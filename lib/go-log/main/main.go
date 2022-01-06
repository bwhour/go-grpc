package main

import (
	"bytes"
	"context"
	"strings"

	log "github.com/bwhour/go-grpc/lib/go-log"
	trace "github.com/bwhour/go-grpc/lib/go-trace"
)

const (
	defaultLoggerFormatArgsLogSize = 1 << 10
)

func main() {
	log.Init(&log.Config{
		Output:       "std",
		Formatter:    "json",
		Level:        "debug",
		ShowFileLine: true,
	})
	defer log.Close()

	ctx := trace.NewContext(context.Background(), trace.Trace{})

	log.Infof("infof=%s", "this is a log info message")

	log.Debugf("infof=%s", "this is a log debug message")

	log.Infof("main_message||%v||user_id=%d||errno=%d||err_msg=%s", trace.FromContext(ctx), 123, 0, "ok")

	log.Infof("test")

	log.Infof("test||test=1")

	log.Infof("test||saa")

	log.Infof("test||sdfsdf||")
	// testTrace(ctx, 123)
	// msg := "dltag_xxxx=sdsf||key1=value1||s||key2=value2||key3=value3||test=s"
	// msg := "key=||"

	// str := msgToJsonstring(msg)
	// fmt.Println(str)
}

func testTrace(ctx context.Context, userID int64) {
	log.Infof("main_message||%v||user_id=%d||errno=%d||err_msg=%s", trace.FromContext(ctx), userID, 0, "ok")
}

func msgToJsonstring(str string) string {
	var buf = &bytes.Buffer{}
	buf.Grow(defaultLoggerFormatArgsLogSize)

	buf.WriteString(`{`)

	var isDltag = true
	for {
		if str == "" {
			break
		}
		n := strings.Index(str, "||")
		if n == -1 {
			e := strings.Index(str, "=")
			buf.WriteString(`"`)
			if e == -1 {
				if isDltag {
					buf.WriteString(`dltag":"`)
					buf.WriteString(str)
					buf.WriteString(`"`)
				} else {
					buf.WriteString(str)
					buf.WriteString(`":""`)
				}
			} else {
				buf.WriteString(str[0:e])
				buf.WriteString(`":"`)
				buf.WriteString(str[e+1:])
				buf.WriteString(`"`)
			}

			isDltag = false
			str = str[n+2:]
			break
		}

		sp := str[0:n]
		if sp == "" {
			str = str[n+2:]
			isDltag = false
			continue
		}

		e := strings.Index(sp, "=")
		buf.WriteString(`"`)
		if e == -1 {
			if isDltag {
				buf.WriteString(`dltag":"`)
				buf.WriteString(sp)
				buf.WriteString(`"`)
			} else {
				buf.WriteString(sp)
				buf.WriteString(`":""`)
			}
		} else {
			buf.WriteString(sp[0:e])
			buf.WriteString(`":"`)
			buf.WriteString(sp[e+1:])
			buf.WriteString(`"`)
		}

		str = str[n+2:]
		if str != "" {
			buf.WriteString(`,`)
		}

		isDltag = false
	}

	buf.WriteString("}")

	return buf.String()
}
