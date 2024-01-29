package cerror

type Option func(*Error)

func WithReasonCode(rc ReasonCode) Option {
	return func(e *Error) {
		e.reasonCodes = append(e.reasonCodes, rc)
	}
}
