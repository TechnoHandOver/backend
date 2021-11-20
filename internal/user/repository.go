package user

import "github.com/TechnoHandOver/backend/internal/models"

type Repository interface {
	Insert(user *models.User) (*models.User, error)
	Select(id uint32) (*models.User, error)
	SelectByVkId(vkId uint32) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	InsertRouteTmp(routeTmp *models.RouteTmp) (*models.RouteTmp, error)
	SelectRouteTmp(routeTmpId uint32) (*models.RouteTmp, error)
	SelectRouteTmpArray() (*models.RoutesTmp, error)
	UpdateRouteTmp(routeTmp *models.RouteTmp) (*models.RouteTmp, error)
	DeleteRouteTmp(routeTmpId uint32) (*models.RouteTmp, error)
	InsertRoutePerm(routePerm *models.RoutePerm) (*models.RoutePerm, error)
	SelectRoutePerm(routePermId uint32) (*models.RoutePerm, error)
	UpdateRoutePerm(routePerm *models.RoutePerm) (*models.RoutePerm, error)
}
