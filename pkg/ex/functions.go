package ex

import (
	"errors"
	"fmt"
	"github.com/byte-engine/bason/pkg/stacks"
)

type Option func(businessError *BusinessError)

func NewFrom(err error, options ...Option) error {
	switch e := err.(type) {
	case TemplateError:
		return NewBy(err, e.Template(), options...)
	default:
		return NewBy(err, "", options...)
	}
}

func NewBy(err error, message string, options ...Option) error {
	return errors.Join(err, New(message, options...))
}

func New(message string, options ...Option) *BusinessError {
	err := &BusinessError{
		Message: message,
	}
	for _, option := range options {
		if option != nil {
			option(err)
		}
	}
	return err
}

func By(businessError *BusinessError, options ...Option) *BusinessError {
	return businessError.new(options...)
}

func Code(code int) Option {
	return func(businessError *BusinessError) {
		businessError.Code = code
	}
}

func M(args ...any) Option {
	return func(businessError *BusinessError) {
		businessError.Message = fmt.Sprintf(businessError.Message, args...)
	}
}

func P(payload any) Option {
	return func(businessError *BusinessError) {
		businessError.Payload = payload
	}
}

func L(level int) Option {
	return func(businessError *BusinessError) {
		businessError.Level = level
	}
}

func T(errorType string) Option {
	return func(businessError *BusinessError) {
		businessError.ErrorType = errorType
	}
}

func C(cause error) Option {
	return func(businessError *BusinessError) {
		businessError.Cause = cause
	}
}

func S(skip int) Option {
	return func(businessError *BusinessError) {
		businessError.stackInfo = stacks.Frames(skip + 1)
	}
}
