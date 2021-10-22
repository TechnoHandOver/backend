package usecase

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/TechnoHandOver/backend/internal/user"
	"log"
	"net/http"
)

type UserUsecase struct {
	userRepository user.UserRepository
}

func NewUserUsecase(userRepository user.UserRepository) *UserUsecase {
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
				log.Println(err)
				return response.NewErrorResponse(http.StatusInternalServerError, nil)
			}
		} else {
			log.Println(err)
			return response.NewErrorResponse(http.StatusInternalServerError, nil)
		}
	} else if user2.Name != user_.Name || user2.Avatar != user_.Avatar {
		userUpdate := new(models.UserUpdate)

		if user2.Name != user_.Name {
			userUpdate.Name = user_.Name
		}

		if user2.Avatar != user_.Avatar {
			userUpdate.Avatar = user_.Avatar
		}

		user2, err = userUsecase.userRepository.Update(user2.Id, userUpdate)
		if err != nil {
			log.Println(err)
			return response.NewErrorResponse(http.StatusInternalServerError, nil)
		}
	}

	return response.NewResponse(http.StatusOK, user2)
}
