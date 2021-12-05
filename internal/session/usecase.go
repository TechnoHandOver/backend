package session

import "github.com/TechnoHandOver/backend/internal/tools/response"

type Usecase interface {
	Create(userId uint32) *response.Response
	Get(id string) *response.Response
}
