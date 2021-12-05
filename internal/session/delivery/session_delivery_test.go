package delivery_test

import (
	"encoding/json"
	"fmt"
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/middlewares"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/session/delivery"
	"github.com/TechnoHandOver/backend/internal/session/mock_session"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/TechnoHandOver/backend/internal/tools/responser"
	HandoverValidator "github.com/TechnoHandOver/backend/internal/tools/validator"
	"github.com/TechnoHandOver/backend/internal/user/mock_user"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSessionDelivery_HandlerLogin(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockSessionUsecase := mock_session.NewMockUsecase(controller)
	mockUserUsecase := mock_user.NewMockUsecase(controller)
	sessionDelivery := delivery.NewSessionDelivery(mockSessionUsecase, mockUserUsecase)
	echo_ := echo.New()
	echo_.Validator = HandoverValidator.NewRequestValidator()
	sessionDelivery.Configure(echo_, &middlewares.Manager{})

	user := &models.User{
		VkId:   201,
		Name:   "Vasiliy Pupkin",
		Avatar: "https://yandex.ru/logo.png",
	}
	expectedUser := &models.User{
		Id:     1,
		VkId:   user.VkId,
		Name:   user.Name,
		Avatar: user.Avatar,
	}
	session := &models.Session{
		Id:     uuid.NewString(),
		UserId: expectedUser.Id,
	}

	loginCall := mockUserUsecase.
		EXPECT().
		Login(gomock.Eq(user)).
		DoAndReturn(func(user *models.User) *response.Response {
			user.Id = expectedUser.Id
			return response.NewResponse(consts.OK, user)
		})
	mockSessionUsecase.
		EXPECT().
		Create(gomock.Eq(expectedUser.Id)).
		Return(response.NewResponse(consts.OK, session)).
		After(loginCall)

	jsonRequest, err := json.Marshal(user)
	assert.Nil(t, err)

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedUser,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodPost, "/api/sessions", strings.NewReader(string(jsonRequest)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)

	handler := sessionDelivery.HandlerLogin()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)

	assert.Equal(t, recorder.Result().Cookies(), []*http.Cookie{
		{
			Name:     consts.EchoCookieAuthName,
			Raw:      fmt.Sprintf("%s=%s; Secure; SameSite=None", consts.EchoCookieAuthName, session.Id),
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
			Value:    session.Id,
		},
	})
}
