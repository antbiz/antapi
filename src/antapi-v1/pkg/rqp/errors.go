package rqp

import "errors"

var (
	ErrRequired           = errors.New("required")
	ErrBadFormat          = errors.New("bad format")
	ErrInvalidValue       = errors.New("invalid value")
	ErrEmptyValue         = errors.New("empty value")
	ErrUnknownMethod      = errors.New("unknown method")
	ErrNotInScope         = errors.New("not in scope")
	ErrSimilarNames       = errors.New("similar names of keys are not allowed")
	ErrMethodNotAllowed   = errors.New("method are not allowed")
	ErrFilterNotAllowed   = errors.New("filter are not allowed")
	ErrFilterNotFound     = errors.New("filter not found")
	ErrValidationNotFound = errors.New("validation not found")
)
