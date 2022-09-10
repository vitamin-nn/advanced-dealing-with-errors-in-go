package stacktrace

import (
	"fmt"
	"runtime"
	"strings"
)

const maxStacktraceDepth = 32

type Frame uintptr

func (f Frame) pc() uintptr {
	return uintptr(f) - 1
}

func (f Frame) String() string {
	fnc := runtime.FuncForPC(f.pc())
	if fnc == nil {
		return "unknown"
	}

	fncName := fnc.Name()
	if n := strings.LastIndex(fncName, "/"); n != -1 {
		fncName = fncName[n+1:]
	}

	fileStrFull, line := fnc.FileLine(f.pc())

	var fileStr string
	if n := strings.LastIndex(fileStrFull, "/"); n != -1 {
		if m := strings.LastIndex(fileStrFull[:n], "/"); m != -1 {
			fileStr = fileStrFull[m+1:]
		} else {
			fileStr = fileStrFull[n+1:]
		}
	} else {
		fileStr = fileStrFull
	}

	if fileStr == "" || fncName == "" {
		return "unknown"
	}

	return fmt.Sprintf("%s\n%s:%d", fncName, fileStr, line)
}

type StackTrace []Frame

func (s StackTrace) String() string {
	var b strings.Builder
	for _, f := range s {
		if _, err := fmt.Fprintf(&b, "%s\n", f); err != nil {
			panic("preparing stack trace")
		}
	}

	return b.String()
}

// Trace возвращает стектрейс глубиной не более maxStacktraceDepth.
// Возвращаемый стектрейс начинается с того места, где была вызвана Trace.
func Trace() StackTrace {
	pc := make([]uintptr, maxStacktraceDepth)
	n := runtime.Callers(2, pc)
	st := make(StackTrace, 0, n)

	for _, c := range pc[:n] {
		st = append(st, Frame(c))
	}

	return st
}
