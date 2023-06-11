package internal

import (
	"bytes"
	"net/http"
	"strconv"
)

type ErrorCode string

const (
	ErrorUnknown  ErrorCode = "ErrUnknown"
	DecodingError ErrorCode = "DecodingError"
)

type Error struct {
	Code       ErrorCode
	StatusCode int
	Message    string
}

func (e *Error) Error() string {
	buf := bytes.NewBuffer(make([]byte, 0, 192))

	buf.WriteString(strconv.Itoa(e.StatusCode) + " " + http.StatusText(e.StatusCode))

	if e.Message != "" {
		buf.WriteString(": " + e.Message)
	}

	if e.Code != "" {
		buf.WriteString(" (" + string(e.Code) + ")")
	}

	return buf.String()
}
