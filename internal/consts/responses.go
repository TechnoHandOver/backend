package consts

import "net/http"

type Code int

const (
	OK Code = iota
	Created
	BadRequest
	Unauthorized
	NotFound
	InternalError
)

var StatusCodes = map[Code]int{
	OK:            http.StatusOK,
	Created:       http.StatusCreated,
	BadRequest:    http.StatusBadRequest,
	Unauthorized:  http.StatusUnauthorized,
	NotFound:      http.StatusNotFound,
	InternalError: http.StatusInternalServerError,
}
