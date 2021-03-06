package delivery_test

import (
	"encoding/json"
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/middlewares"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/models/timestamps"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/TechnoHandOver/backend/internal/tools/responser"
	HandoverValidator "github.com/TechnoHandOver/backend/internal/tools/validator"
	"github.com/TechnoHandOver/backend/internal/user/delivery"
	"github.com/TechnoHandOver/backend/internal/user/mock_user"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestUserDelivery_HandlerRouteTmpCreate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserUsecase := mock_user.NewMockUsecase(controller)
	userDelivery := delivery.NewUserDelivery(mockUserUsecase)
	echo_ := echo.New()
	echo_.Validator = HandoverValidator.NewRequestValidator()
	userDelivery.Configure(echo_, &middlewares.Manager{})

	dateTimeDep, err := timestamps.NewDateTime("10.11.2021 18:10")
	assert.Nil(t, err)
	dateTimeArr, err := timestamps.NewDateTime("10.11.2021 18:15")
	assert.Nil(t, err)
	routeTmp := &models.RouteTmp{
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		DateTimeDep:  *dateTimeDep,
		DateTimeArr:  *dateTimeArr,
	}
	expectedRouteTmp := &models.RouteTmp{
		Id:           1,
		UserAuthorId: routeTmp.UserAuthorId,
		LocDep:       routeTmp.LocDep,
		LocArr:       routeTmp.LocArr,
		MinPrice:     routeTmp.MinPrice,
		DateTimeDep:  routeTmp.DateTimeDep,
		DateTimeArr:  routeTmp.DateTimeArr,
	}

	mockUserUsecase.
		EXPECT().
		CreateRouteTmp(gomock.Eq(routeTmp)).
		DoAndReturn(func(routeTmp *models.RouteTmp) *response.Response {
			routeTmp.Id = expectedRouteTmp.Id
			return response.NewResponse(consts.Created, routeTmp)
		})

	jsonRequest, err := json.Marshal(routeTmp)
	assert.Nil(t, err)

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedRouteTmp,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodPost, "/api/users/routes-tmp", strings.NewReader(string(jsonRequest)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.Set(consts.EchoContextKeyUserId, routeTmp.UserAuthorId)

	handler := userDelivery.HandlerRouteTmpCreate()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestUserDelivery_HandlerRouteTmpGet(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserUsecase := mock_user.NewMockUsecase(controller)
	userDelivery := delivery.NewUserDelivery(mockUserUsecase)
	echo_ := echo.New()
	echo_.Validator = HandoverValidator.NewRequestValidator()
	userDelivery.Configure(echo_, &middlewares.Manager{})

	dateTimeDep, err := timestamps.NewDateTime("13.11.2021 11:30")
	assert.Nil(t, err)
	dateTimeArr, err := timestamps.NewDateTime("13.11.2021 11:35")
	assert.Nil(t, err)
	expectedRouteTmp := &models.RouteTmp{
		Id:           1,
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		DateTimeDep:  *dateTimeDep,
		DateTimeArr:  *dateTimeArr,
	}

	mockUserUsecase.
		EXPECT().
		GetRouteTmp(gomock.Eq(expectedRouteTmp.UserAuthorId), gomock.Eq(expectedRouteTmp.Id)).
		Return(response.NewResponse(consts.OK, expectedRouteTmp))

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedRouteTmp,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/users/routes-tmp/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(expectedRouteTmp.Id), 10))
	context.Set(consts.EchoContextKeyUserId, expectedRouteTmp.UserAuthorId)

	handler := userDelivery.HandlerRouteTmpGet()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestUserDelivery_HandlerRouteTmpUpdate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserUsecase := mock_user.NewMockUsecase(controller)
	userDelivery := delivery.NewUserDelivery(mockUserUsecase)
	echo_ := echo.New()
	echo_.Validator = HandoverValidator.NewRequestValidator()
	userDelivery.Configure(echo_, &middlewares.Manager{})

	dateTimeDep, err := timestamps.NewDateTime("13.11.2021 13:30")
	assert.Nil(t, err)
	dateTimeArr, err := timestamps.NewDateTime("13.11.2021 13:35")
	assert.Nil(t, err)
	expectedRouteTmp := &models.RouteTmp{
		Id:           1,
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		DateTimeDep:  *dateTimeDep,
		DateTimeArr:  *dateTimeArr,
	}

	mockUserUsecase.
		EXPECT().
		UpdateRouteTmp(gomock.Eq(expectedRouteTmp)).
		Return(response.NewResponse(consts.OK, expectedRouteTmp))

	jsonRequest, err := json.Marshal(expectedRouteTmp)
	assert.Nil(t, err)

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedRouteTmp,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(jsonRequest)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/users/routes-tmp/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(expectedRouteTmp.Id), 10))
	context.Set(consts.EchoContextKeyUserId, expectedRouteTmp.UserAuthorId)

	handler := userDelivery.HandlerRouteTmpUpdate()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestUserDelivery_HandlerRouteTmpDelete(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserUsecase := mock_user.NewMockUsecase(controller)
	userDelivery := delivery.NewUserDelivery(mockUserUsecase)
	echo_ := echo.New()
	echo_.Validator = HandoverValidator.NewRequestValidator()
	userDelivery.Configure(echo_, &middlewares.Manager{})

	dateTimeDep, err := timestamps.NewDateTime("13.11.2021 13:30")
	assert.Nil(t, err)
	dateTimeArr, err := timestamps.NewDateTime("13.11.2021 13:35")
	assert.Nil(t, err)
	expectedRouteTmp := &models.RouteTmp{
		Id:           1,
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		DateTimeDep:  *dateTimeDep,
		DateTimeArr:  *dateTimeArr,
	}

	mockUserUsecase.
		EXPECT().
		DeleteRouteTmp(gomock.Eq(expectedRouteTmp.UserAuthorId), gomock.Eq(expectedRouteTmp.Id)).
		Return(response.NewResponse(consts.OK, expectedRouteTmp))

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedRouteTmp,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodDelete, "/", nil)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/users/routes-tmp/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(expectedRouteTmp.Id), 10))
	context.Set(consts.EchoContextKeyUserId, expectedRouteTmp.UserAuthorId)

	handler := userDelivery.HandlerRouteTmpDelete()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestUserDelivery_HandlerRouteTmpList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserUsecase := mock_user.NewMockUsecase(controller)
	userDelivery := delivery.NewUserDelivery(mockUserUsecase)
	echo_ := echo.New()
	echo_.Validator = HandoverValidator.NewRequestValidator()
	userDelivery.Configure(echo_, &middlewares.Manager{})

	const userId uint32 = 101
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
			Id:           1,
			UserAuthorId: userId,
			LocDep:       "Общежитие №10",
			LocArr:       "УЛК",
			MinPrice:     500,
			DateTimeDep:  *dateTimeDep1,
			DateTimeArr:  *dateTimeArr1,
		},
		&models.RouteTmp{
			Id:           2,
			UserAuthorId: userId,
			LocDep:       "Общежитие №9",
			LocArr:       "СК",
			MinPrice:     600,
			DateTimeDep:  *dateTimeDep2,
			DateTimeArr:  *dateTimeArr2,
		},
	}

	mockUserUsecase.
		EXPECT().
		ListRouteTmp(gomock.Eq(userId)).
		Return(response.NewResponse(consts.OK, expectedRoutesTmp))

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedRoutesTmp,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodGet, "/api/users/routes-tmp/list", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.Set(consts.EchoContextKeyUserId, userId)

	handler := userDelivery.HandlerRouteTmpList()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestUserDelivery_HandlerRoutePermCreate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserUsecase := mock_user.NewMockUsecase(controller)
	userDelivery := delivery.NewUserDelivery(mockUserUsecase)
	echo_ := echo.New()
	echo_.Validator = HandoverValidator.NewRequestValidator()
	userDelivery.Configure(echo_, &middlewares.Manager{})

	timeDep, err := timestamps.NewTime("12:30")
	assert.Nil(t, err)
	timeArr, err := timestamps.NewTime("12:35")
	assert.Nil(t, err)
	routePerm := &models.RoutePerm{
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		EvenWeek:     true,
		OddWeek:      false,
		DayOfWeek:    timestamps.DayOfWeekWednesday,
		TimeDep:      *timeDep,
		TimeArr:      *timeArr,
	}
	expectedRoutePerm := &models.RoutePerm{
		Id:           1,
		UserAuthorId: routePerm.UserAuthorId,
		LocDep:       routePerm.LocDep,
		LocArr:       routePerm.LocArr,
		MinPrice:     routePerm.MinPrice,
		EvenWeek:     routePerm.EvenWeek,
		OddWeek:      routePerm.OddWeek,
		DayOfWeek:    routePerm.DayOfWeek,
		TimeDep:      routePerm.TimeDep,
		TimeArr:      routePerm.TimeArr,
	}

	mockUserUsecase.
		EXPECT().
		CreateRoutePerm(gomock.Eq(routePerm)).
		DoAndReturn(func(routePerm *models.RoutePerm) *response.Response {
			routePerm.Id = expectedRoutePerm.Id
			return response.NewResponse(consts.Created, routePerm)
		})

	jsonRequest, err := json.Marshal(routePerm)
	assert.Nil(t, err)

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedRoutePerm,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodPost, "/api/users/routes-perm", strings.NewReader(string(jsonRequest)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.Set(consts.EchoContextKeyUserId, routePerm.UserAuthorId)

	handler := userDelivery.HandlerRoutePermCreate()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestUserDelivery_HandlerRoutePermGet(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserUsecase := mock_user.NewMockUsecase(controller)
	userDelivery := delivery.NewUserDelivery(mockUserUsecase)
	echo_ := echo.New()
	echo_.Validator = HandoverValidator.NewRequestValidator()
	userDelivery.Configure(echo_, &middlewares.Manager{})

	timeDep, err := timestamps.NewTime("15:00")
	assert.Nil(t, err)
	timeArr, err := timestamps.NewTime("15:05")
	assert.Nil(t, err)
	expectedRoutePerm := &models.RoutePerm{
		Id:           1,
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		EvenWeek:     true,
		OddWeek:      false,
		DayOfWeek:    timestamps.DayOfWeekWednesday,
		TimeDep:      *timeDep,
		TimeArr:      *timeArr,
	}

	mockUserUsecase.
		EXPECT().
		GetRoutePerm(gomock.Eq(expectedRoutePerm.UserAuthorId), gomock.Eq(expectedRoutePerm.Id)).
		Return(response.NewResponse(consts.OK, expectedRoutePerm))

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedRoutePerm,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/users/routes-perm/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(expectedRoutePerm.Id), 10))
	context.Set(consts.EchoContextKeyUserId, expectedRoutePerm.UserAuthorId)

	handler := userDelivery.HandlerRoutePermGet()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestUserDelivery_HandlerRoutePermUpdate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserUsecase := mock_user.NewMockUsecase(controller)
	userDelivery := delivery.NewUserDelivery(mockUserUsecase)
	echo_ := echo.New()
	echo_.Validator = HandoverValidator.NewRequestValidator()
	userDelivery.Configure(echo_, &middlewares.Manager{})

	timeDep, err := timestamps.NewTime("16:15")
	assert.Nil(t, err)
	timeArr, err := timestamps.NewTime("16:20")
	assert.Nil(t, err)
	expectedRoutePerm := &models.RoutePerm{
		Id:           1,
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		EvenWeek:     true,
		OddWeek:      false,
		DayOfWeek:    timestamps.DayOfWeekWednesday,
		TimeDep:      *timeDep,
		TimeArr:      *timeArr,
	}

	mockUserUsecase.
		EXPECT().
		UpdateRoutePerm(gomock.Eq(expectedRoutePerm)).
		Return(response.NewResponse(consts.OK, expectedRoutePerm))

	jsonRequest, err := json.Marshal(expectedRoutePerm)
	assert.Nil(t, err)

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedRoutePerm,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(jsonRequest)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/users/routes-perm/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(expectedRoutePerm.Id), 10))
	context.Set(consts.EchoContextKeyUserId, expectedRoutePerm.UserAuthorId)

	handler := userDelivery.HandlerRoutePermUpdate()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestUserDelivery_HandlerRoutePermDelete(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserUsecase := mock_user.NewMockUsecase(controller)
	userDelivery := delivery.NewUserDelivery(mockUserUsecase)
	echo_ := echo.New()
	echo_.Validator = HandoverValidator.NewRequestValidator()
	userDelivery.Configure(echo_, &middlewares.Manager{})

	timeDep, err := timestamps.NewTime("16:15")
	assert.Nil(t, err)
	timeArr, err := timestamps.NewTime("16:20")
	assert.Nil(t, err)
	expectedRoutePerm := &models.RoutePerm{
		Id:           1,
		UserAuthorId: 101,
		LocDep:       "Корпус Энерго",
		LocArr:       "Корпус УЛК",
		MinPrice:     500,
		EvenWeek:     true,
		OddWeek:      false,
		DayOfWeek:    timestamps.DayOfWeekWednesday,
		TimeDep:      *timeDep,
		TimeArr:      *timeArr,
	}

	mockUserUsecase.
		EXPECT().
		DeleteRoutePerm(gomock.Eq(expectedRoutePerm.UserAuthorId), gomock.Eq(expectedRoutePerm.Id)).
		Return(response.NewResponse(consts.OK, expectedRoutePerm))

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedRoutePerm,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodDelete, "/", nil)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/users/routes-perm/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(expectedRoutePerm.Id), 10))
	context.Set(consts.EchoContextKeyUserId, expectedRoutePerm.UserAuthorId)

	handler := userDelivery.HandlerRoutePermDelete()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestUserDelivery_HandlerRoutePermList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserUsecase := mock_user.NewMockUsecase(controller)
	userDelivery := delivery.NewUserDelivery(mockUserUsecase)
	echo_ := echo.New()
	echo_.Validator = HandoverValidator.NewRequestValidator()
	userDelivery.Configure(echo_, &middlewares.Manager{})

	const userId uint32 = 101
	timeDep1, err := timestamps.NewTime("17:30")
	assert.Nil(t, err)
	timeArr1, err := timestamps.NewTime("17:35")
	assert.Nil(t, err)
	timeDep2, err := timestamps.NewTime("17:40")
	assert.Nil(t, err)
	timeArr2, err := timestamps.NewTime("17:45")
	assert.Nil(t, err)
	expectedRoutesPerm := &models.RoutesPerm{
		&models.RoutePerm{
			Id:           1,
			UserAuthorId: userId,
			LocDep:       "Общежитие №10",
			LocArr:       "УЛК",
			MinPrice:     500,
			EvenWeek:     true,
			OddWeek:      false,
			DayOfWeek:    timestamps.DayOfWeekWednesday,
			TimeDep:      *timeDep1,
			TimeArr:      *timeArr1,
		},
		&models.RoutePerm{
			Id:           2,
			UserAuthorId: userId,
			LocDep:       "Общежитие №9",
			LocArr:       "СК",
			MinPrice:     600,
			EvenWeek:     false,
			OddWeek:      true,
			DayOfWeek:    timestamps.DayOfWeekSaturday,
			TimeDep:      *timeDep2,
			TimeArr:      *timeArr2,
		},
	}

	mockUserUsecase.
		EXPECT().
		ListRoutePerm(gomock.Eq(userId)).
		Return(response.NewResponse(consts.OK, expectedRoutesPerm))

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedRoutesPerm,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodGet, "/api/users/routes-perm/list", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.Set(consts.EchoContextKeyUserId, userId)

	handler := userDelivery.HandlerRoutePermList()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}
