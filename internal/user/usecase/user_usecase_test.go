package usecase_test

import (
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
		Return(nil, consts.RepErrNotFound)

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
		Return(nil, consts.RepErrNotFound)

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
		Return(nil, consts.RepErrNotFound)

	response_ := userUsecase.GetRouteTmp(routeTmpId)
	assert.Equal(t, response.NewEmptyResponse(consts.NotFound), response_)
}

func TestUserUsecase_UpdateRouteTmp(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	dateTimeDep1, err := timestamps.NewDateTime("13.11.2021 13:50")
	assert.Nil(t, err)
	dateTimeArr1, err := timestamps.NewDateTime("13.11.2021 13:55")
	assert.Nil(t, err)
	routeTmp := &models.RouteTmp{
		Id:             1,
		UserAuthorVkId: 2,
		LocDep:         "Общежитие №10",
		LocArr:         "УЛК",
		MinPrice:       500,
		DateTimeDep:    *dateTimeDep1,
		DateTimeArr:    *dateTimeArr1,
	}
	dateTimeDep2, err := timestamps.NewDateTime("13.11.2021 14:00")
	assert.Nil(t, err)
	dateTimeArr2, err := timestamps.NewDateTime("13.11.2021 14:05")
	assert.Nil(t, err)
	expectedRouteTmp := &models.RouteTmp{
		Id:             routeTmp.Id,
		UserAuthorVkId: 2,
		LocDep:         "Общежитие №9",
		LocArr:         "СК",
		MinPrice:       600,
		DateTimeDep:    *dateTimeDep2,
		DateTimeArr:    *dateTimeArr2,
	}

	call := mockUserRepository.
		EXPECT().
		SelectRouteTmp(gomock.Eq(routeTmp.Id)).
		Return(routeTmp, nil)

	mockUserRepository.
		EXPECT().
		UpdateRouteTmp(gomock.Eq(expectedRouteTmp)).
		Return(expectedRouteTmp, nil).
		After(call)

	response_ := userUsecase.UpdateRouteTmp(expectedRouteTmp)
	assert.Equal(t, response.NewResponse(consts.OK, expectedRouteTmp), response_)
}

func TestUserUsecase_UpdateRouteTmp_forbidden(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	dateTimeDep1, err := timestamps.NewDateTime("13.11.2021 13:50")
	assert.Nil(t, err)
	dateTimeArr1, err := timestamps.NewDateTime("13.11.2021 13:55")
	assert.Nil(t, err)
	routeTmp := &models.RouteTmp{
		Id:             1,
		UserAuthorVkId: 2,
		LocDep:         "Общежитие №10",
		LocArr:         "УЛК",
		MinPrice:       500,
		DateTimeDep:    *dateTimeDep1,
		DateTimeArr:    *dateTimeArr1,
	}
	dateTimeDep2, err := timestamps.NewDateTime("13.11.2021 14:00")
	assert.Nil(t, err)
	dateTimeArr2, err := timestamps.NewDateTime("13.11.2021 14:05")
	assert.Nil(t, err)
	expectedRouteTmp := &models.RouteTmp{
		Id:             routeTmp.Id,
		UserAuthorVkId: 3,
		LocDep:         "Общежитие №9",
		LocArr:         "СК",
		MinPrice:       600,
		DateTimeDep:    *dateTimeDep2,
		DateTimeArr:    *dateTimeArr2,
	}

	mockUserRepository.
		EXPECT().
		SelectRouteTmp(gomock.Eq(routeTmp.Id)).
		Return(routeTmp, nil)

	response_ := userUsecase.UpdateRouteTmp(expectedRouteTmp)
	assert.Equal(t, response.NewEmptyResponse(consts.Forbidden), response_)
}

func TestUserUsecase_UpdateRouteTmp_notFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	dateTimeDep, err := timestamps.NewDateTime("13.11.2021 13:55")
	assert.Nil(t, err)
	dateTimeArr, err := timestamps.NewDateTime("13.11.2021 14:00")
	assert.Nil(t, err)
	routeTmp := &models.RouteTmp{
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
		SelectRouteTmp(gomock.Eq(routeTmp.Id)).
		Return(nil, consts.RepErrNotFound)

	response_ := userUsecase.UpdateRouteTmp(routeTmp)
	assert.Equal(t, response.NewEmptyResponse(consts.NotFound), response_)
}

func TestUserUsecase_DeleteRouteTmp(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	dateTimeDep, err := timestamps.NewDateTime("13.11.2021 17:20")
	assert.Nil(t, err)
	dateTimeArr, err := timestamps.NewDateTime("13.11.2021 17:25")
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

	call := mockUserRepository.
		EXPECT().
		SelectRouteTmp(gomock.Eq(expectedRouteTmp.Id)).
		Return(expectedRouteTmp, nil)

	mockUserRepository.
		EXPECT().
		DeleteRouteTmp(gomock.Eq(expectedRouteTmp.Id)).
		Return(expectedRouteTmp, nil).
		After(call)

	response_ := userUsecase.DeleteRouteTmp(expectedRouteTmp.UserAuthorVkId, expectedRouteTmp.Id)
	assert.Equal(t, response.NewResponse(consts.OK, expectedRouteTmp), response_)
}

func TestUserUsecase_DeleteRouteTmp_forbidden(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	dateTimeDep, err := timestamps.NewDateTime("13.11.2021 17:20")
	assert.Nil(t, err)
	dateTimeArr, err := timestamps.NewDateTime("13.11.2021 17:25")
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

	response_ := userUsecase.DeleteRouteTmp(expectedRouteTmp.UserAuthorVkId+1, expectedRouteTmp.Id)
	assert.Equal(t, response.NewEmptyResponse(consts.Forbidden), response_)
}

func TestUserUsecase_DeleteRouteTmp_notFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	dateTimeDep, err := timestamps.NewDateTime("13.11.2021 17:20")
	assert.Nil(t, err)
	dateTimeArr, err := timestamps.NewDateTime("13.11.2021 17:25")
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
		Return(nil, consts.RepErrNotFound)

	response_ := userUsecase.DeleteRouteTmp(expectedRouteTmp.UserAuthorVkId, expectedRouteTmp.Id)
	assert.Equal(t, response.NewEmptyResponse(consts.NotFound), response_)
}

func TestUserUsecase_ListRouteTmp(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	dateTimeDep1, err := timestamps.NewDateTime("17.11.2021 10:25")
	assert.Nil(t, err)
	dateTimeArr1, err := timestamps.NewDateTime("17.11.2021 10:30")
	assert.Nil(t, err)
	dateTimeDep2, err := timestamps.NewDateTime("17.11.2021 10:35")
	assert.Nil(t, err)
	dateTimeArr2, err := timestamps.NewDateTime("17.11.2021 10:40")
	assert.Nil(t, err)
	expectedRoutesTmp := &models.RoutesTmp{
		&models.RouteTmp{
			Id:             1,
			UserAuthorVkId: 3,
			LocDep:         "Общежитие №10",
			LocArr:         "УЛК",
			MinPrice:       500,
			DateTimeDep:    *dateTimeDep1,
			DateTimeArr:    *dateTimeArr1,
		},
		&models.RouteTmp{
			Id:             2,
			UserAuthorVkId: 4,
			LocDep:         "Общежитие №9",
			LocArr:         "СК",
			MinPrice:       600,
			DateTimeDep:    *dateTimeDep2,
			DateTimeArr:    *dateTimeArr2,
		},
	}

	mockUserRepository.
		EXPECT().
		SelectRouteTmpArray().
		Return(expectedRoutesTmp, nil)

	response_ := userUsecase.ListRouteTmp()
	assert.Equal(t, response.NewResponse(consts.OK, expectedRoutesTmp), response_)
}

func TestUserUsecase_CreateRoutePerm(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	timeDep, err := timestamps.NewTime("12:30")
	assert.Nil(t, err)
	timeArr, err := timestamps.NewTime("12:35")
	assert.Nil(t, err)
	routePerm := &models.RoutePerm{
		UserAuthorVkId: 2,
		LocDep:         "Корпус Энерго",
		LocArr:         "Корпус УЛК",
		MinPrice:       500,
		EvenWeek:       true,
		OddWeek:        false,
		DayOfWeek:      timestamps.DayOfWeekWednesday,
		TimeDep:        *timeDep,
		TimeArr:        *timeArr,
	}
	expectedRoutePerm := &models.RoutePerm{
		Id:             1,
		UserAuthorVkId: routePerm.UserAuthorVkId,
		LocDep:         routePerm.LocDep,
		LocArr:         routePerm.LocArr,
		MinPrice:       routePerm.MinPrice,
		EvenWeek:       routePerm.EvenWeek,
		OddWeek:        routePerm.OddWeek,
		DayOfWeek:      routePerm.DayOfWeek,
		TimeDep:        routePerm.TimeDep,
		TimeArr:        routePerm.TimeArr,
	}

	mockUserRepository.
		EXPECT().
		InsertRoutePerm(gomock.Eq(routePerm)).
		DoAndReturn(func(routePerm *models.RoutePerm) (*models.RoutePerm, error) {
			routePerm.Id = expectedRoutePerm.Id
			return routePerm, nil
		})

	response_ := userUsecase.CreateRoutePerm(routePerm)
	assert.Equal(t, response.NewResponse(consts.Created, expectedRoutePerm), response_)
}

func TestUserUsecase_GetRoutePerm(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	timeDep, err := timestamps.NewTime("15:00")
	assert.Nil(t, err)
	timeArr, err := timestamps.NewTime("15:05")
	assert.Nil(t, err)
	expectedRoutePerm := &models.RoutePerm{
		Id: 1,
		UserAuthorVkId: 2,
		LocDep:         "Корпус Энерго",
		LocArr:         "Корпус УЛК",
		MinPrice:       500,
		EvenWeek:       true,
		OddWeek:        false,
		DayOfWeek:      timestamps.DayOfWeekWednesday,
		TimeDep:        *timeDep,
		TimeArr:        *timeArr,
	}

	mockUserRepository.
		EXPECT().
		SelectRoutePerm(gomock.Eq(expectedRoutePerm.Id)).
		Return(expectedRoutePerm, nil)

	response_ := userUsecase.GetRoutePerm(expectedRoutePerm.Id)
	assert.Equal(t, response.NewResponse(consts.OK, expectedRoutePerm), response_)
}

func TestUserUsecase_GetRoutePerm_notFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserRepository := mock_user.NewMockRepository(controller)
	userUsecase := usecase.NewUserUsecaseImpl(mockUserRepository)

	const routePermId uint32 = 1

	mockUserRepository.
		EXPECT().
		SelectRoutePerm(gomock.Eq(routePermId)).
		Return(nil, consts.RepErrNotFound)

	response_ := userUsecase.GetRoutePerm(routePermId)
	assert.Equal(t, response.NewEmptyResponse(consts.NotFound), response_)
}
