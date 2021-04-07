package rerrors

import (
	"fmt"
	"log"
)

const (
	InvalidRequestErr = 101
	InternalErr       = 102
)

var (
	internalErrMsg = "failed to process the request at this time, please try again later."

	errTypes = map[int]string{
		InvalidRequestErr: "InvalidRequestErr",
		InternalErr:       "InternalErr",
	}

	errMessages = map[int]string{
		InvalidRequestErr: "invalid request error",
		InternalErr:       internalErrMsg,
	}

	errDetails = map[int]string{
		InvalidRequestErr: "invalid request parameters",
		InternalErr:       internalErrMsg,
	}
)

// Type returns the mapped error type for the given error code
func errorType(code int) string {
	if value, ok := errTypes[code]; ok {
		return value
	}
	return "UnKnownError"
}

// message returns the mapped message for the given error code
func message(code int) string {
	if value, ok := errMessages[code]; ok {
		return value
	}
	return internalErrMsg
}

// detail returns a formatted string containing the string from the concrete error type
func detail(code int, err error) string {
	if value, ok := errDetails[code]; ok {
		return fmt.Sprintf("%s: %v", value, err)
	}
	return "unknown"
}

// Format Returns a formatted error type
func Format(code int, err error) error {
	return NewError(code, errorType(code), message(code), detail(code, err))
}

// LogFormat Returns a formatted error type and logs it on the standard output
func LogFormat(code int, err error) error {
	e := Format(code, err)
	log.Println(e)
	return e
}
