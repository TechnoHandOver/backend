package consts

import "errors"

const (
	EchoCookieAuthName     = "handover_auth_session_id"
	EchoContextKeyUserVkId = "userVkId"
)

type RepositoryError error

var (
	RepErrNotFound RepositoryError = errors.New("Not found\n")
)
