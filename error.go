package cevir

import (
	"errors"
	"fmt"
)

const EFILEEXTNOTMATCH = "file_ext_does_not_match"

type Error struct {
	Code    string
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("cevir error: code=%s message=%s", e.Code, e.Message)
}

func ErrorCode(err error) string {
	var e *Error
	if err == nil {
		return ""
	} else if errors.As(err, &e) {
		return e.Message
	}
	return "Invalid error"
}
