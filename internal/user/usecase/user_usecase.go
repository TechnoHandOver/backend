package usecase

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/TechnoHandOver/backend/internal/user/repository"
	"log"
	"net/http"
)

type UserUsecase struct {
	userRepository *repository.UserRepository
}

func NewUserUsecase(userRepository *repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (userUsecase *UserUsecase) Login(user *models.User) *response.Response {
	user2, err := userUsecase.userRepository.SelectByVkId(user.VkId)
	if err != nil {
		if err == sql.ErrNoRows {
			user2, err = userUsecase.userRepository.Insert(user)
			if err != nil {
				log.Println(err)
				return response.NewErrorResponse(http.StatusInternalServerError, nil)
			}
		} else {
			log.Println(err)
			return response.NewErrorResponse(http.StatusInternalServerError, nil)
		}
	} else if user2.Name != user.Name || user2.Avatar != user.Avatar {
		userUpdate := new(models.UserUpdate)

		if user2.Name != user.Name {
			userUpdate.Name = user.Name
		}

		if user2.Avatar != user.Avatar {
			userUpdate.Avatar = user.Avatar
		}

		user2, err = userUsecase.userRepository.Update(user2.Id, userUpdate)
		if err != nil {
			log.Println(err)
			return response.NewErrorResponse(http.StatusInternalServerError, nil)
		}
	}

	return response.NewResponse(http.StatusOK, user2)
}
