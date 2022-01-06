package log

import (
	"bytes"
	"strings"
	"testing"
)

func ExampleLogger() {
	// myLogger 可以是一个全局变量。
	myLogger := New("./log/mylogger.log")

	// 使用这个 logger。
	myLogger.Debug("k1=v1", "k2=v2")

	// 应该在 main 里面初始化 logger。
	Init(&Config{
		FilePath:     "./log/all.log",
		ShowFileLine: true,
		Debug:        true,
	})

	Debugf("default log in ./log/all.log")
	Errorf("error log in both ./log/all.log and ./log/error.log")
	Printf("print some%v", "thing")
	myLogger.Infof("my log||foo=%v", "bar")
	myLogger.Errorf("my another error log||foo=%v", "bar")
	myLogger.Print("some", "log")

	Close()

	// Output:
}

func TestLogNew(t *testing.T) {
	swapCreatedLoggers(func() {
		New("")
		New("./log/foo.log")
		New("./log/all.log")
		New("./log/all.log")
		New("./log/foo.log")

		Init(&Config{
			FilePath:     "./log/all.log",
			ShowFileLine: true,
			Debug:        true,
		})

		if expected := 3; len(createdLoggers) != expected {
			t.Fatalf("invalid count of created loggers. [expected:%v] [actual:%v]", expected, createdLoggers)
		}

		New("./log/foo.log")
		New("./log/bar.log")

		if expected := 4; len(createdLoggers) != expected {
			t.Fatalf("invalid count of created loggers. [expected:%v] [actual:%v]", expected, createdLoggers)
		}
	})
}

func TestStdLog(t *testing.T) {
	Init(&Config{
		Output:    "std",
		Formatter: "json",
		Level:     "debug",
	})
	// defer Close()

	Infof("infof=%s", "sdfsd")
}

func BenchmarkInfo(b *testing.B) {
	Init(&Config{
		Output:       "std",
		Formatter:    "json",
		Level:        "debug",
		ShowFileLine: true,
	})
	defer Close()
	for i := 0; i < b.N; i++ {
		Infof("this is a log info message")
	}
}

func TestLogString(t *testing.T) {
	Init(&Config{
		Output:       "std",
		Formatter:    "json",
		Level:        "debug",
		ShowFileLine: true,
	})
	defer Close()
}

func msgToJsonstring(str string) string {
	var buf *bytes.Buffer
	buf.Grow(defaultLoggerFormatArgsLogSize)

	buf.WriteString(`{`)

	for {
		if str == "" {
			break
		}
		n := strings.Index(str, "||")
		if n == -1 {
			e := strings.Index(str, "=")
			buf.WriteString(`"`)
			if e == -1 {
				buf.WriteString(str)
				buf.WriteString(`":""`)
			} else {
				buf.WriteString(str[0:e])
				buf.WriteString(`":"`)
				buf.WriteString(str[e+1:])
				buf.WriteString(`"`)
			}

			break
		}

		sp := str[0:n]
		if sp == "" {
			str = str[n+1:]
			continue
		}

		e := strings.Index(sp, "=")
		buf.WriteString(`"`)
		if e == -1 {
			buf.WriteString(sp)
			buf.WriteString(`":""`)
		} else {
			buf.WriteString(sp[0:e])
			buf.WriteString(`":"`)
			buf.WriteString(sp[e+1:])
			buf.WriteString(`"`)
		}

		str = str[n+1:]
		if str != "" {
			buf.WriteString(`,`)
		}
	}

	buf.WriteString("}")

	return buf.String()
}
