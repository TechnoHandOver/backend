package consts

import "net/http"

type Code int

const (
	OK Code = iota
	Created
	Unauthorized
	NotFound
	InternalError
)

var StatusCodes = map[Code]int{
	OK:            http.StatusOK,
	Created:       http.StatusCreated,
	Unauthorized:  http.StatusUnauthorized,
	NotFound:      http.StatusNotFound,
	InternalError: http.StatusInternalServerError,
}
