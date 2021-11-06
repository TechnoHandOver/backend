package user

import (
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"
)

type Usecase interface {
	Login(user *models.User) *response.Response
	Get(vkId uint32) *response.Response
}
