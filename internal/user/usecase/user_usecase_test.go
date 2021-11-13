package usecase_test

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/models/timestamps"
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

func TestUserUsecase_Get(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	expectedUser := &models.User{
		Id:     1,
		VkId:   2,
		Name:   "Василий Петров",
		Avatar: "https://mail.ru/vasiliy_petrov_avatar.jpg",
	}

	mockUserRepository.
		EXPECT().
		SelectByVkId(gomock.Eq(expectedUser.VkId)).
		Return(expectedUser, nil)

	response_ := userUsecase.Get(expectedUser.VkId)
	assert.Equal(t, response.NewResponse(consts.OK, expectedUser), response_)
}

func TestUserUsecase_Get_notFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	const vkId uint32 = 2

	mockUserRepository.
		EXPECT().
		SelectByVkId(gomock.Eq(vkId)).
		Return(nil, sql.ErrNoRows)

	response_ := userUsecase.Get(vkId)
	assert.Equal(t, response.NewEmptyResponse(consts.NotFound), response_)
}

func TestUserUsecase_CreateRouteTmp(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	dateTimeDep, err := timestamps.NewDateTime("10.11.2021 18:10")
	assert.Nil(t, err)
	dateTimeArr, err := timestamps.NewDateTime("10.11.2021 18:15")
	assert.Nil(t, err)
	routeTmp := &models.RouteTmp{
		UserAuthorVkId: 2,
		LocDep:         "Корпус Энерго",
		LocArr:         "Корпус УЛК",
		MinPrice:       500,
		DateTimeDep:    *dateTimeDep,
		DateTimeArr:    *dateTimeArr,
	}
	expectedRouteTmp := &models.RouteTmp{
		Id:             1,
		UserAuthorVkId: routeTmp.UserAuthorVkId,
		LocDep:         routeTmp.LocDep,
		LocArr:         routeTmp.LocArr,
		MinPrice:       routeTmp.MinPrice,
		DateTimeDep:    routeTmp.DateTimeDep,
		DateTimeArr:    routeTmp.DateTimeArr,
	}

	mockUserRepository.
		EXPECT().
		InsertRouteTmp(gomock.Eq(routeTmp)).
		DoAndReturn(func(routeTmp *models.RouteTmp) (*models.RouteTmp, error) {
			routeTmp.Id = expectedRouteTmp.Id
			return routeTmp, nil
		})

	response_ := userUsecase.CreateRouteTmp(routeTmp)
	assert.Equal(t, response.NewResponse(consts.Created, expectedRouteTmp), response_)
}

func TestUserUsecase_GetRouteTmp(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	dateTimeDep, err := timestamps.NewDateTime("13.11.2021 11:45")
	assert.Nil(t, err)
	dateTimeArr, err := timestamps.NewDateTime("13.11.2021 11:50")
	assert.Nil(t, err)
	expectedRouteTmp := &models.RouteTmp{
		Id:             1,
		UserAuthorVkId: 2,
		LocDep:         "Корпус Энерго",
		LocArr:         "Корпус УЛК",
		MinPrice:       500,
		DateTimeDep:    *dateTimeDep,
		DateTimeArr:    *dateTimeArr,
	}

	mockUserRepository.
		EXPECT().
		SelectRouteTmp(gomock.Eq(expectedRouteTmp.Id)).
		Return(expectedRouteTmp, nil)

	response_ := userUsecase.GetRouteTmp(expectedRouteTmp.Id)
	assert.Equal(t, response.NewResponse(consts.OK, expectedRouteTmp), response_)
}

func TestUserUsecase_GetRouteTmp_notFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	const routeTmpId uint32 = 1

	mockUserRepository.
		EXPECT().
		SelectRouteTmp(gomock.Eq(routeTmpId)).
		Return(nil, sql.ErrNoRows)

	response_ := userUsecase.GetRouteTmp(routeTmpId)
	assert.Equal(t, response.NewEmptyResponse(consts.NotFound), response_)
}
