package delivery_test

import (
	"encoding/json"
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/middlewares"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/models/timestamps"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/TechnoHandOver/backend/internal/tools/responser"
	RequestValidator "github.com/TechnoHandOver/backend/internal/tools/validator"
	"github.com/TechnoHandOver/backend/internal/user/delivery"
	"github.com/TechnoHandOver/backend/internal/user/mock_user"
	Validator "github.com/go-playground/validator/v10"
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
	echo_.Validator = RequestValidator.NewRequestValidator(Validator.New())
	userDelivery.Configure(echo_, &middlewares.Manager{})

	const vkId uint32 = 2
	dateTimeDep, err := timestamps.NewDateTime("10.11.2021 18:10")
	assert.Nil(t, err)
	dateTimeArr, err := timestamps.NewDateTime("10.11.2021 18:15")
	assert.Nil(t, err)
	routeTmp := &models.RouteTmp{
		UserAuthorVkId: vkId,
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

	mockUserUsecase.
		EXPECT().
		CreateRouteTmp(gomock.Eq(routeTmp)).
		DoAndReturn(func(routeTmp *models.RouteTmp) *response.Response {
			routeTmp.Id = expectedRouteTmp.Id
			return response.NewResponse(http.StatusOK, routeTmp)
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
	context.Set(consts.EchoContextKeyUserVkId, vkId)

	handler := userDelivery.HandlerRouteTmpCreate()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

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
	echo_.Validator = RequestValidator.NewRequestValidator(Validator.New())
	userDelivery.Configure(echo_, &middlewares.Manager{})

	type RouteTmpGetRequest struct {
		Id uint32 `param:"id"`
	}

	routeTmpGetRequest := &RouteTmpGetRequest{
		Id: 1,
	}
	dateTimeDep, err := timestamps.NewDateTime("13.11.2021 11:30")
	assert.Nil(t, err)
	dateTimeArr, err := timestamps.NewDateTime("13.11.2021 11:35")
	assert.Nil(t, err)
	expectedRouteTmp := &models.RouteTmp{
		Id:             routeTmpGetRequest.Id,
		UserAuthorVkId: 2,
		LocDep:         "Корпус Энерго",
		LocArr:         "Корпус УЛК",
		MinPrice:       500,
		DateTimeDep:    *dateTimeDep,
		DateTimeArr:    *dateTimeArr,
	}

	mockUserUsecase.
		EXPECT().
		GetRouteTmp(gomock.Eq(routeTmpGetRequest.Id)).
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
	context.SetParamValues(strconv.FormatUint(uint64(routeTmpGetRequest.Id), 10))

	handler := userDelivery.HandlerRouteTmpGet()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestUserDelivery_HandlerRouteTmpGet_notFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserUsecase := mock_user.NewMockUsecase(controller)
	userDelivery := delivery.NewUserDelivery(mockUserUsecase)
	echo_ := echo.New()
	echo_.Validator = RequestValidator.NewRequestValidator(Validator.New())
	userDelivery.Configure(echo_, &middlewares.Manager{})

	type RouteTmpGetRequest struct {
		Id uint32 `param:"id"`
	}

	routeTmpGetRequest := &RouteTmpGetRequest{
		Id: 1,
	}

	mockUserUsecase.
		EXPECT().
		GetRouteTmp(gomock.Eq(routeTmpGetRequest.Id)).
		Return(response.NewEmptyResponse(consts.NotFound))

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/users/routes-tmp/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(routeTmpGetRequest.Id), 10))

	handler := userDelivery.HandlerRouteTmpGet()

	err := handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, recorder.Code)
}
