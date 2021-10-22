package user

import (
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"
)

type UserUsecase interface {
	Login(user *models.User) *response.Response
}

type UserRepository interface {
	Insert(user *models.User) (*models.User, error)
	Select(id uint32) (*models.User, error)
	SelectByVkId(vkId uint32) (*models.User, error)
	Update(id uint32, userUpdate *models.UserUpdate) (*models.User, error)
}
