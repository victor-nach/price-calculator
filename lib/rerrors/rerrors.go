package rerrors

import "encoding/json"

// Internal error type
type Err struct {
	Code      int
	ErrorType string
	Message   string
	Detail    string
}

// Ensure Customized error type implements error interface
var _ error = &Err{}

// NewError returns a new customized error type
func NewError(code int, errorType, message, detail string) *Err {
	return &Err{
		Code:      code,
		ErrorType: errorType,
		Message:   message,
		Detail:    detail,
	}
}

// Error returns a json string representation of our customized error type
// this is the only method required to implement the error interface
func (e *Err) Error() string {
	err := Err{
		Code:      e.Code,
		ErrorType: e.ErrorType,
		Message:   e.Message,
		Detail:    e.Detail,
	}
	b, _ := json.Marshal(err)

	return string(b)
}
