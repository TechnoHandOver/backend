package usecase

import (
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/TechnoHandOver/backend/internal/user"
)

type UserUsecase struct {
	userRepository user.Repository
}

func NewUserUsecaseImpl(userRepository user.Repository) user.Usecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (userUsecase *UserUsecase) Login(user_ *models.User) *response.Response {
	user2, err := userUsecase.userRepository.SelectByVkId(user_.VkId)
	if err != nil {
		if err == consts.RepErrNotFound {
			user2, err = userUsecase.userRepository.Insert(user_)
			if err != nil {
				return response.NewErrorResponse(consts.InternalError, err)
			}
		} else {
			return response.NewErrorResponse(consts.InternalError, err)
		}
	} else if user2.Name != user_.Name || user2.Avatar != user_.Avatar {
		user2.Name = user_.Name
		user2.Avatar = user_.Avatar

		user2, err = userUsecase.userRepository.Update(user2)
		if err != nil {
			return response.NewErrorResponse(consts.InternalError, err)
		}
	}

	return response.NewResponse(consts.OK, user2)
}

func (userUsecase *UserUsecase) Get(id uint32) *response.Response {
	user_, err := userUsecase.userRepository.Select(id)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.OK, user_)
}

func (userUsecase *UserUsecase) CreateRouteTmp(routeTmp *models.RouteTmp) *response.Response {
	routeTmp, err := userUsecase.userRepository.InsertRouteTmp(routeTmp)
	if err != nil {
		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.Created, routeTmp)
}

func (userUsecase *UserUsecase) GetRouteTmp(userId uint32, routeTmpId uint32) *response.Response {
	routeTmp, err := userUsecase.userRepository.SelectRouteTmp(routeTmpId)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	if userId != routeTmp.UserAuthorId {
		return response.NewEmptyResponse(consts.Forbidden)
	}

	return response.NewResponse(consts.OK, routeTmp)
}

func (userUsecase *UserUsecase) UpdateRouteTmp(routeTmp *models.RouteTmp) *response.Response {
	existingRouteTmp, err := userUsecase.userRepository.SelectRouteTmp(routeTmp.Id)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	if routeTmp.UserAuthorId != existingRouteTmp.UserAuthorId {
		return response.NewEmptyResponse(consts.Forbidden)
	}

	routeTmp, err = userUsecase.userRepository.UpdateRouteTmp(routeTmp)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.OK, routeTmp)
}

func (userUsecase *UserUsecase) DeleteRouteTmp(userId uint32, routeTmpId uint32) *response.Response {
	existingRouteTmp, err := userUsecase.userRepository.SelectRouteTmp(routeTmpId)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	if userId != existingRouteTmp.UserAuthorId {
		return response.NewEmptyResponse(consts.Forbidden)
	}

	routeTmp, err := userUsecase.userRepository.DeleteRouteTmp(routeTmpId)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.OK, routeTmp)
}

func (userUsecase *UserUsecase) ListRouteTmp(userId uint32) *response.Response {
	routesTmp, err := userUsecase.userRepository.SelectRouteTmpArrayByUserAuthorId(userId)
	if err != nil {
		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.OK, routesTmp)
}

func (userUsecase *UserUsecase) CreateRoutePerm(routePerm *models.RoutePerm) *response.Response {
	routePerm, err := userUsecase.userRepository.InsertRoutePerm(routePerm)
	if err != nil {
		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.Created, routePerm)
}

func (userUsecase *UserUsecase) GetRoutePerm(userId uint32, routePermId uint32) *response.Response {
	routePerm, err := userUsecase.userRepository.SelectRoutePerm(routePermId)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	if userId != routePerm.UserAuthorId {
		return response.NewEmptyResponse(consts.Forbidden)
	}

	return response.NewResponse(consts.OK, routePerm)
}

func (userUsecase *UserUsecase) UpdateRoutePerm(routePerm *models.RoutePerm) *response.Response {
	existingRoutePerm, err := userUsecase.userRepository.SelectRoutePerm(routePerm.Id)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	if routePerm.UserAuthorId != existingRoutePerm.UserAuthorId {
		return response.NewEmptyResponse(consts.Forbidden)
	}

	routePerm, err = userUsecase.userRepository.UpdateRoutePerm(routePerm)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.OK, routePerm)
}

func (userUsecase *UserUsecase) DeleteRoutePerm(userId uint32, routePermId uint32) *response.Response {
	existingRoutePerm, err := userUsecase.userRepository.SelectRoutePerm(routePermId)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	if userId != existingRoutePerm.UserAuthorId {
		return response.NewEmptyResponse(consts.Forbidden)
	}

	routePerm, err := userUsecase.userRepository.DeleteRoutePerm(routePermId)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.OK, routePerm)
}

func (userUsecase *UserUsecase) ListRoutePerm(userId uint32) *response.Response {
	routesPerm, err := userUsecase.userRepository.SelectRoutePermArrayByUserAuthorId(userId)
	if err != nil {
		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.OK, routesPerm)
}
