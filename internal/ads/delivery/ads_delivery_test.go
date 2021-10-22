package delivery_test

import (
	"encoding/json"
	"github.com/TechnoHandOver/backend/internal/ads/delivery"
	"github.com/TechnoHandOver/backend/internal/ads/mock_ads"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/models/timestamps"
	"github.com/TechnoHandOver/backend/internal/tools/response"
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

func TestAdsDelivery_HandlerAdsCreate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdsUsecase := mock_ads.NewMockAdsUsecase(controller)

	const id uint32 = 1
	const userAuthorVkId uint32 = 2
	dateTimeArr, err := timestamps.NewDateTime("27.10.2021 19:31")
	assert.Nil(t, err)

	ads := &models.Ads{
		UserAuthorVkId: userAuthorVkId,
		LocDep:         "Общежитие №10",
		LocArr:         "УЛК",
		DateTimeArr:    *dateTimeArr,
		MinPrice:       500,
		Comment:        "Поеду на велосипеде",
	}
	expectedAds := &models.Ads{
		Id:             id,
		UserAuthorVkId: userAuthorVkId,
		LocDep:         ads.LocDep,
		LocArr:         ads.LocArr,
		DateTimeArr:    ads.DateTimeArr,
		MinPrice:       ads.MinPrice,
		Comment:        ads.Comment,
	}

	mockAdsUsecase.
		EXPECT().
		Create(ads).
		DoAndReturn(func(ads *models.Ads) *response.Response {
			ads.Id = id
			ads.UserAuthorVkId = userAuthorVkId
			return response.NewResponse(http.StatusCreated, ads)
		})

	adsDelivery := delivery.NewAdsDelivery(mockAdsUsecase)
	echo_ := echo.New()
	adsDelivery.Configure(echo_)

	jsonRequest, err := json.Marshal(ads)
	assert.Nil(t, err)

	jsonExpectedResponse, err := json.Marshal(response.DataResponse{
		Data: expectedAds,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodPost, "/api/ads", strings.NewReader(string(jsonRequest)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)

	handler := adsDelivery.HandlerAdsCreate()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestAdsDelivery_HandlerAdsGet(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdsUsecase := mock_ads.NewMockAdsUsecase(controller)

	const id uint32 = 1
	const userAuthorVkId uint32 = 2
	dateTimeArr, err := timestamps.NewDateTime("27.10.2021 19:44")
	assert.Nil(t, err)

	expectedAds := &models.Ads{
		Id:             id,
		UserAuthorVkId: userAuthorVkId,
		LocDep:         "Общежитие №10",
		LocArr:         "УЛК",
		DateTimeArr:    *dateTimeArr,
		MinPrice:       500,
		Comment:        "Поеду на велосипеде",
	}

	mockAdsUsecase.
		EXPECT().
		Get(id).
		Return(response.NewResponse(http.StatusOK, expectedAds))

	adsDelivery := delivery.NewAdsDelivery(mockAdsUsecase)
	echo_ := echo.New()
	adsDelivery.Configure(echo_)

	jsonExpectedResponse, err := json.Marshal(response.DataResponse{
		Data: expectedAds,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/ads/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(id), 10))

	handler := adsDelivery.HandlerAdsGet()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestAdsDelivery_HandlerAdsGet_Failed(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdsUsecase := mock_ads.NewMockAdsUsecase(controller)

	const id uint32 = 1

	mockAdsUsecase.
		EXPECT().
		Get(id).
		Return(response.NewErrorResponse(http.StatusNotFound, nil))

	adsDelivery := delivery.NewAdsDelivery(mockAdsUsecase)
	echo_ := echo.New()
	adsDelivery.Configure(echo_)

	jsonExpectedResponse, err := json.Marshal(response.ErrorResponse{})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/ads/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(id), 10))

	handler := adsDelivery.HandlerAdsGet()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestAdsDelivery_HandlerAdsUpdate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdsUsecase := mock_ads.NewMockAdsUsecase(controller)

	const id uint32 = 1
	const userAuthorVkId uint32 = 2
	dateTimeArr, err := timestamps.NewDateTime("27.10.2021 19:50")
	assert.Nil(t, err)

	adsUpdate := &models.AdsUpdate{
		LocDep:         "Общежитие №10",
		LocArr:         "УЛК",
		DateTimeArr:    *dateTimeArr,
		MinPrice:       500,
		Comment:        "Поеду на велосипеде",
	}
	expectedAds := &models.Ads{
		Id:             id,
		UserAuthorVkId: userAuthorVkId,
		LocDep:         adsUpdate.LocDep,
		LocArr:         adsUpdate.LocArr,
		DateTimeArr:    adsUpdate.DateTimeArr,
		MinPrice:       adsUpdate.MinPrice,
		Comment:        adsUpdate.Comment,
	}

	mockAdsUsecase.
		EXPECT().
		Update(id, adsUpdate).
		Return(response.NewResponse(http.StatusOK, expectedAds))

	adsDelivery := delivery.NewAdsDelivery(mockAdsUsecase)
	echo_ := echo.New()
	adsDelivery.Configure(echo_)

	jsonRequest, err := json.Marshal(adsUpdate)
	assert.Nil(t, err)

	jsonExpectedResponse, err := json.Marshal(response.DataResponse{
		Data: expectedAds,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(jsonRequest)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/ads/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(id), 10))

	handler := adsDelivery.HandlerAdsUpdate()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestAdsDelivery_HandlerAdsUpdate_Failed(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdsUsecase := mock_ads.NewMockAdsUsecase(controller)

	const id uint32 = 1
	dateTimeArr, err := timestamps.NewDateTime("27.10.2021 19:50")
	assert.Nil(t, err)

	adsUpdate := &models.AdsUpdate{
		LocDep:         "Общежитие №10",
		LocArr:         "УЛК",
		DateTimeArr:    *dateTimeArr,
		MinPrice:       500,
		Comment:        "Поеду на велосипеде",
	}

	mockAdsUsecase.
		EXPECT().
		Update(id, adsUpdate).
		Return(response.NewErrorResponse(http.StatusNotFound, nil))

	adsDelivery := delivery.NewAdsDelivery(mockAdsUsecase)
	echo_ := echo.New()
	adsDelivery.Configure(echo_)

	jsonRequest, err := json.Marshal(adsUpdate)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(jsonRequest)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/ads/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(id), 10))

	handler := adsDelivery.HandlerAdsUpdate()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, recorder.Code)
}
