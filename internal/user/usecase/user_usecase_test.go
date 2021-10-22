package usecase_test

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/TechnoHandOver/backend/internal/user/mock_user"
	"github.com/TechnoHandOver/backend/internal/user/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUserUsecase_Login(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockUserRepository(controller)
	mockUserUsecase := usecase.NewUserUsecase(mockUserRepository)

	user := &models.User{
		VkId: 2,
		Name: "Василий Петров",
		Avatar: "https://mail.ru/vasiliy_petrov_avatar.jpg",
	}
	expectedUser := &models.User {
		Id: 1,
		VkId: user.VkId,
		Name: user.Name,
		Avatar: user.Avatar,
	}

	selectByIdCall := mockUserRepository.
		EXPECT().
		SelectByVkId(user.VkId).
		Return(nil, sql.ErrNoRows)

	mockUserRepository.
		EXPECT().
		Insert(user).
		Return(expectedUser, nil).
		After(selectByIdCall)

	response_ := mockUserUsecase.Login(user)
	assert.Equal(t, response.NewResponse(http.StatusOK, expectedUser), response_)
}

/*func TestUserUsecase_Login_Update(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockUserRepository(controller)
	mockUserUsecase := usecase.NewUserUsecase(mockUserRepository)

	oldUser := &models.User{
		Id: 1,
		VkId: 2,
		Name: "Василий Петров",
		Avatar: "https://mail.ru/vasiliy_petrov_avatar.jpg",
	}
	newUser := &models.User {
		VkId: oldUser.VkId,
		Name: "Василий Данилов",
		Avatar: "https://mail.ru/vasiliy_danilov_avatar.jpg",
	}
	newUser2 := &models.User {
		Id: oldUser.Id,
		VkId: oldUser.VkId,
		Name: "Василий Данилов",
		Avatar: "https://mail.ru/vasiliy_danilov_avatar.jpg",
	}
	newUserUpdate := &models.UserUpdate{
		Name: newUser.Name,
		Avatar: newUser.Avatar,
	}

	selectByIdCall := mockUserRepository.
		EXPECT().
		SelectByVkId(vkId). //TODO: gomock.equals?
		Return(oldUser)

	mockUserRepository.
		EXPECT().
		Update().
		Return(expectedUser, nil).
		After(selectByIdCall)

	response_ := mockUserUsecase.Login(newUser)
	assert.Equal(t, response.NewResponse(http.StatusOK, expectedUser), response_)
}*/
