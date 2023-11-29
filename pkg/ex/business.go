package ex

import (
	"github.com/byte-engine/bason/pkg/buffer"
	"github.com/byte-engine/bason/pkg/stacks"
	"runtime"
)

var (
	_pool = buffer.NewPool()
	// Get retrieves a buffer from the pool, creating one if necessary.
	get = _pool.Get
)

type BusinessError struct {
	Code      int
	Message   string
	Payload   interface{}
	Level     int
	ErrorType string
	Cause     error
	stackInfo *runtime.Frames
}

func (err *BusinessError) Error() string {
	return err.String()
}

func (err *BusinessError) Is(target error) bool {
	if x, ok := target.(*BusinessError); ok {
		return x.Code == err.Code
	}
	return false
}

func (err *BusinessError) StackTrack() *runtime.Frames {
	return err.stackInfo
}

func (err *BusinessError) Unwrap() error {
	return err.Cause
}

func (err *BusinessError) clone() *BusinessError {
	return &BusinessError{
		Code:      err.Code,
		Message:   err.Message,
		Level:     err.Level,
		ErrorType: err.ErrorType,
	}
}

func (err *BusinessError) new(options ...Option) *BusinessError {
	other := err.clone()

	for _, option := range options {
		if option != nil {
			option(other)
		}
	}

	return other
}

func (err *BusinessError) String() string {
	var (
		message       = err.Message
		buf           = get()
		ex      error = err
	)

	defer buf.Free()
	for ex != nil {
		buf.AppendString(message)
		if x, ok := ex.(interface{ StackTrack() *runtime.Frames }); ok {
			buf.AppendString("\n")
			buf.AppendString(stacks.StringFrames(x.StackTrack()))
			buf.AppendString("\n")
		}
		if x, ok := ex.(interface{ Unwarp() error }); ok {
			ex = x.Unwarp()
			buf.AppendString("C by: ")
		} else {
			ex = nil
		}
	}

	return buf.String()
}
