package delivery_test

import (
	"encoding/json"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/TechnoHandOver/backend/internal/user/delivery"
	"github.com/TechnoHandOver/backend/internal/user/mock_user"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserDelivery_HandlerUserLogin(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUserUsecase := mock_user.NewMockUserUsecase(controller)

	const id uint32 = 1

	user := &models.User{
		VkId: 2,
		Name: "Василий Петров",
	}
	expectedUser := &models.User {
		Id: id,
		VkId: user.VkId,
		Name: user.Name,
	}

	mockUserUsecase.
		EXPECT().
		Login(user).
		DoAndReturn(func(user *models.User) *response.Response {
			user.Id = id
			return response.NewResponse(http.StatusOK, user)
		})

	userDelivery := delivery.NewUserDelivery(mockUserUsecase)
	echo_ := echo.New()
	userDelivery.Configure(echo_)

	jsonRequest, err := json.Marshal(user)
	assert.Nil(t, err)

	jsonExpectedResponse, err := json.Marshal(response.DataResponse{
		Data: expectedUser,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodPost, "/api/user", strings.NewReader(string(jsonRequest)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)

	handler := userDelivery.HandlerUserLogin()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}
