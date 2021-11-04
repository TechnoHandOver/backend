package usecase

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/TechnoHandOver/backend/internal/user"
	"net/http"
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
		if err == sql.ErrNoRows {
			user2, err = userUsecase.userRepository.Insert(user_)
			if err != nil {
				return response.NewErrorResponse(err)
			}
		} else {
			return response.NewErrorResponse(err)
		}
	} else if user2.Name != user_.Name || user2.Avatar != user_.Avatar {
		user2.Name = user_.Name
		user2.Avatar = user_.Avatar

		user2, err = userUsecase.userRepository.Update(user2)
		if err != nil {
			return response.NewErrorResponse(err)
		}
	}

	return response.NewResponse(http.StatusOK, user2)
}
