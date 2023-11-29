package ex

type TemplateError interface {
	Error() string
	Template() string
}

type templateError struct {
	message  string
	template string
}

func (t *templateError) Error() string {
	return t.message
}

func (t *templateError) Template() string {
	return t.template
}

func NewBase(message string, template string) TemplateError {
	return &templateError{
		message:  message,
		template: template,
	}
}
