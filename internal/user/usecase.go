package user

import (
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"
)

type Usecase interface {
	Login(user *models.User) *response.Response
	Get(id uint32) *response.Response
	CreateRouteTmp(routeTmp *models.RouteTmp) *response.Response
	GetRouteTmp(routeTmpId uint32) *response.Response
	UpdateRouteTmp(routeTmp *models.RouteTmp) *response.Response
	DeleteRouteTmp(userId uint32, routeTmpId uint32) *response.Response
	ListRouteTmp(userId uint32) *response.Response
	CreateRoutePerm(routePerm *models.RoutePerm) *response.Response
	GetRoutePerm(routePermId uint32) *response.Response
	UpdateRoutePerm(routePerm *models.RoutePerm) *response.Response
	DeleteRoutePerm(userId uint32, routePermId uint32) *response.Response
	ListRoutePerm(userId uint32) *response.Response
}
