package usecase_test

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/TechnoHandOver/backend/internal/user/mock_user"
	"github.com/TechnoHandOver/backend/internal/user/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserUsecase_Login(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	user := &models.User{
		VkId:   2,
		Name:   "Василий Петров",
		Avatar: "https://mail.ru/vasiliy_petrov_avatar.jpg",
	}
	expectedUser := &models.User{
		Id:     1,
		VkId:   user.VkId,
		Name:   user.Name,
		Avatar: user.Avatar,
	}

	mockUserRepository.
		EXPECT().
		SelectByVkId(gomock.Eq(user.VkId)).
		Return(expectedUser, nil)

	response_ := userUsecase.Login(user)
	assert.Equal(t, response.NewResponse(consts.OK, expectedUser), response_)
}

func TestUserUsecase_Login_create(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	user := &models.User{
		VkId:   2,
		Name:   "Василий Петров",
		Avatar: "https://mail.ru/vasiliy_petrov_avatar.jpg",
	}
	expectedUser := &models.User{
		Id:     1,
		VkId:   user.VkId,
		Name:   user.Name,
		Avatar: user.Avatar,
	}

	selectByIdCall := mockUserRepository.
		EXPECT().
		SelectByVkId(gomock.Eq(user.VkId)).
		Return(nil, sql.ErrNoRows)

	mockUserRepository.
		EXPECT().
		Insert(gomock.Eq(user)).
		DoAndReturn(func(user *models.User) (*models.User, error) {
			user.Id = expectedUser.Id
			return user, nil
		}).
		After(selectByIdCall)

	response_ := userUsecase.Login(user)
	assert.Equal(t, response.NewResponse(consts.OK, expectedUser), response_)
}

func TestUserUsecase_Login_update(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	existingUser := &models.User{
		Id:     1,
		VkId:   2,
		Name:   "Василий Петров",
		Avatar: "https://mail.ru/vasiliy_petrov_avatar.jpg",
	}
	updatedUser := &models.User{
		VkId:   existingUser.VkId,
		Name:   "Пётр Васильев",
		Avatar: "https://mail.ru/petr_vasiliev_avatar.jpg",
	}
	expectedUser := &models.User{
		Id:     existingUser.Id,
		VkId:   updatedUser.VkId,
		Name:   updatedUser.Name,
		Avatar: updatedUser.Avatar,
	}

	selectByIdCall := mockUserRepository.
		EXPECT().
		SelectByVkId(gomock.Eq(updatedUser.VkId)).
		Return(existingUser, nil)

	mockUserRepository.
		EXPECT().
		Update(gomock.Eq(expectedUser)).
		Return(expectedUser, nil).
		After(selectByIdCall)

	response_ := userUsecase.Login(updatedUser)
	assert.Equal(t, response.NewResponse(consts.OK, expectedUser), response_)
}
