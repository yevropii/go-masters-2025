package errs

import (
	"fmt"
	"net/http"
)

type HTTPError interface {
	error
	StatusCode() int
}

type BadRequest struct {
	msg string
}

func (e BadRequest) Error() string {
	return e.msg
}

func (e BadRequest) StatusCode() int {
	return http.StatusBadRequest
}

func NewBadRequest(format string, a ...any) error {
	return BadRequest{msg: fmt.Sprintf(format, a...)}
}

type NotFound struct {
	msg string
}

func (e NotFound) Error() string {
	return e.msg
}

func (e NotFound) StatusCode() int {
	return http.StatusNotFound
}

func NewNotFound(format string, a ...any) error {
	return NotFound{msg: fmt.Sprintf(format, a...)}
}
