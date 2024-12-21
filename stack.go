package kleverr

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

type StackError struct {
	cause  error
	frames []StackFrame
}

type StackFrame struct {
	Func string
	File string
	Line int
}

func newStackError(err error, diff int) error {
	if err == nil {
		return nil
	}

	serr := &StackError{cause: err}

	var pcs = make([]uintptr, 32)
	if n := runtime.Callers(3+diff, pcs); n > 0 {
		frames := runtime.CallersFrames(pcs[:n])

		for {
			frame, more := frames.Next()

			serr.frames = append(serr.frames, StackFrame{
				Func: frame.Function,
				File: frame.File,
				Line: frame.Line,
			})

			if !more {
				break
			}
		}
	}

	return serr
}

func (e *StackError) Error() string {
	return e.cause.Error()
}

func (e *StackError) Unwrap() error {
	return e.cause
}

func (e *StackError) Print() string {
	var b = new(strings.Builder)
	fmt.Fprintf(b, "%s", e.cause.Error())
	for _, frame := range e.frames {
		fmt.Fprintf(b, "\n%s\n  %s:%d", frame.Func, frame.File, frame.Line)
	}

	if serr := Get(e.cause); serr != nil {
		fmt.Fprintln(b)
		fmt.Fprintln(b, serr.Print())
	}

	return b.String()
}

func Get(err error) *StackError {
	var e *StackError
	if errors.As(err, &e) {
		return e
	}
	return nil
}

func Ret(err error) error {
	return newStackError(err, 0)
}

func New(m string) error {
	return newStackError(errors.New(m), 0)
}

func Newf(m string, args ...any) error {
	return newStackError(fmt.Errorf(m, args...), 0)
}

func Ret1[X any](err error) (x X, serr error) {
	return x, newStackError(err, 0)
}

func New1[X any](m string) (x X, serr error) {
	return x, newStackError(errors.New(m), 0)
}

func New1f[X any](m string, args ...any) (x X, serr error) {
	return x, newStackError(fmt.Errorf(m, args...), 0)
}

func Ret2[X any, Y any](err error) (x X, y Y, serr error) {
	return x, y, newStackError(err, 0)
}

func New2[X any, Y any](m string) (x X, y Y, serr error) {
	return x, y, newStackError(errors.New(m), 0)
}

func New2f[X any, Y any](m string, args ...any) (x X, y Y, serr error) {
	return x, y, newStackError(fmt.Errorf(m, args...), 0)
}

func Ret3[X any, Y any, Z any](err error) (x X, y Y, z Z, serr error) {
	return x, y, z, newStackError(err, 0)
}

func New3[X any, Y any, Z any](m string) (x X, y Y, z Z, serr error) {
	return x, y, z, newStackError(errors.New(m), 0)
}

func New3f[X any, Y any, Z any](m string, args ...any) (x X, y Y, z Z, serr error) {
	return x, y, z, newStackError(fmt.Errorf(m, args...), 0)
}
