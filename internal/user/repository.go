package user

import "github.com/TechnoHandOver/backend/internal/models"

type Repository interface {
	Insert(user *models.User) (*models.User, error)
	Select(id uint32) (*models.User, error)
	SelectByVkId(vkId uint32) (*models.User, error)
	Update(user *models.User) (*models.User, error)
}