package user

import (
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"
)

type Usecase interface {
	Login(user *models.User) *response.Response
	Get(vkId uint32) *response.Response
	CreateRouteTmp(routeTmp *models.RouteTmp) *response.Response
	GetRouteTmp(routeTmpId uint32) *response.Response
	UpdateRouteTmp(routeTmp *models.RouteTmp) *response.Response
	DeleteRouteTmp(userVkId uint32, routeTmpId uint32) *response.Response
	ListRouteTmp() *response.Response
}
