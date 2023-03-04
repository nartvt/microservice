package uerror

import (
	"fmt"
	"net/http"
	"strings"
)

type StatusError struct {
	Err     error  `json:"err,omitempty"`
	Status  int    `json:"status,omitempty"`
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
	File    string `json:"file,omitempty"`
	Line    int    `json:"line,omitempty"`
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("Status %d at %s : %s", e.Status, e.FileLine(), e.Err)
}

func (e *StatusError) String() string {
	return fmt.Sprintf("Status %d at %s : %s %s %s", e.Status, e.FileLine(), e.Title, e.Message, e.Err)
}
func (e *StatusError) FileLine() string {
	parts := strings.Split(e.File, "/")
	f := strings.Join(parts[len(parts)-4:len(parts)], "/")
	return fmt.Sprintf("%s:%d", f, e.Line)
}

func (e *StatusError) setupFromArgs(args ...string) *StatusError {
	if e.Err == nil {
		e.Err = fmt.Errorf("Error:%d", e.Status)
	}
	if len(args) > 0 {
		e.Title = args[0]
	}
	if len(args) > 1 {
		e.Message = args[1]
	}
	return e
}

func NotFoundError(e error, args ...string) *StatusError {
	err := Error(e, http.StatusNotFound, "Not Found", "Sorry, the page you're looking for couldn't be found.")
	return err.setupFromArgs(args...)
}

func InternalError(e error, args ...string) *StatusError {
	err := Error(e, http.StatusInternalServerError, "Server Error", "Sorry, something went wrong, please let us know.")
	return err.setupFromArgs(args...)
}

func BadRequestError(e error, args ...string) *StatusError {
	err := Error(e, http.StatusBadRequest, "Bad Request", "Sorry, there was an error processing your request, please check your data.")
	return err.setupFromArgs(args...)
}

func ForbiddenError(e error, args ...string) *StatusError {
	err := Error(e, http.StatusForbidden, "Forbidden Request", "Sorry, you don't have permission to access this.")
	return err.setupFromArgs(args...)
}
func Error(e error, s int, t string, m string) *StatusError {
	// Get runtime info - use zero values if none available
	err := &StatusError{
		Status:  s,
		Title:   t,
		Message: m,
	}
	if e != nil {
		err.Err = e
	}
	return err
}
