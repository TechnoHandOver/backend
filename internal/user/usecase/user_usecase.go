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

func (userUsecase *UserUsecase) Get(vkId uint32) *response.Response {
	user_, err := userUsecase.userRepository.SelectByVkId(vkId)
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

func (userUsecase *UserUsecase) GetRouteTmp(routeTmpId uint32) *response.Response {
	routeTmp, err := userUsecase.userRepository.SelectRouteTmp(routeTmpId)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
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

	if routeTmp.UserAuthorVkId != existingRouteTmp.UserAuthorVkId {
		return response.NewErrorResponse(consts.Forbidden, err)
	}

	routeTmp, err = userUsecase.userRepository.UpdateRouteTmp(routeTmp)
	if err != nil {
		switch err {
		case consts.RepErrNotFound:
			return response.NewEmptyResponse(consts.NotFound)
		case consts.RepErrNothingToUpdate:
			return response.NewEmptyResponse(consts.BadRequest)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.OK, routeTmp)
}
