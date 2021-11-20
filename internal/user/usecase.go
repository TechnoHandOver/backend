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
	CreateRoutePerm(routePerm *models.RoutePerm) *response.Response
	GetRoutePerm(routePermId uint32) *response.Response
	UpdateRoutePerm(routePerm *models.RoutePerm) *response.Response
	DeleteRoutePerm(userVkId uint32, routePermId uint32) *response.Response
}
