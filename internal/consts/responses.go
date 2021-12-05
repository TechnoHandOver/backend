package consts

import "net/http"

type Code int

const (
	OK Code = iota
	Created
	BadRequest
	Unauthorized
	Forbidden
	NotFound
	Conflict
	InternalError
)

var StatusCodes = map[Code]int{
	OK:            http.StatusOK,
	Created:       http.StatusCreated,
	BadRequest:    http.StatusBadRequest,
	Unauthorized:  http.StatusUnauthorized,
	Forbidden:     http.StatusForbidden,
	NotFound:      http.StatusNotFound,
	Conflict:      http.StatusConflict,
	InternalError: http.StatusInternalServerError,
}
