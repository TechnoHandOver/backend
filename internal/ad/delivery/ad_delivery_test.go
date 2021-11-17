package delivery_test

import (
	"encoding/json"
	"github.com/TechnoHandOver/backend/internal/ad/delivery"
	"github.com/TechnoHandOver/backend/internal/ad/mock_ad"
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/middlewares"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/models/timestamps"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/TechnoHandOver/backend/internal/tools/responser"
	RequestValidator "github.com/TechnoHandOver/backend/internal/tools/validator"
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

func TestAdDelivery_HandlerAdCreate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdUsecase := mock_ad.NewMockUsecase(controller)
	adDelivery := delivery.NewAdDelivery(mockAdUsecase)
	echo_ := echo.New()
	echo_.Validator = RequestValidator.NewRequestValidator()
	adDelivery.Configure(echo_, &middlewares.Manager{})

	dateTimeArr, err := timestamps.NewDateTime("27.10.2021 19:31")
	assert.Nil(t, err)

	ad := &models.Ad{
		UserAuthorVkId: 2,
		LocDep:         "Общежитие №10",
		LocArr:         "УЛК",
		DateTimeArr:    *dateTimeArr,
		Item:           "Зачётная книжка",
		MinPrice:       500,
		Comment:        "Поеду на велосипеде",
	}
	expectedAd := &models.Ad{
		Id:             1,
		UserAuthorVkId: ad.UserAuthorVkId,
		LocDep:         ad.LocDep,
		LocArr:         ad.LocArr,
		DateTimeArr:    ad.DateTimeArr,
		Item:           ad.Item,
		MinPrice:       ad.MinPrice,
		Comment:        ad.Comment,
	}

	mockAdUsecase.
		EXPECT().
		Create(gomock.Eq(ad)).
		DoAndReturn(func(ad *models.Ad) *response.Response {
			ad.Id = expectedAd.Id
			return response.NewResponse(consts.Created, ad)
		})

	jsonRequest, err := json.Marshal(ad)
	assert.Nil(t, err)

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedAd,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodPost, "/api/ads", strings.NewReader(string(jsonRequest)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.Set(consts.EchoContextKeyUserVkId, ad.UserAuthorVkId)

	handler := adDelivery.HandlerAdCreate()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestAdDelivery_HandlerAdGet(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdUsecase := mock_ad.NewMockUsecase(controller)
	adDelivery := delivery.NewAdDelivery(mockAdUsecase)
	echo_ := echo.New()
	echo_.Validator = RequestValidator.NewRequestValidator()
	adDelivery.Configure(echo_, &middlewares.Manager{})

	dateTimeArr, err := timestamps.NewDateTime("27.10.2021 19:44")
	assert.Nil(t, err)
	expectedAd := &models.Ad{
		Id:             1,
		UserAuthorVkId: 2,
		LocDep:         "Общежитие №10",
		LocArr:         "УЛК",
		DateTimeArr:    *dateTimeArr,
		Item:           "Зачётная книжка",
		MinPrice:       500,
		Comment:        "Поеду на велосипеде",
	}

	mockAdUsecase.
		EXPECT().
		Get(gomock.Eq(expectedAd.Id)).
		Return(response.NewResponse(consts.OK, expectedAd))

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedAd,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/ads/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(expectedAd.Id), 10))

	handler := adDelivery.HandlerAdGet()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestAdDelivery_HandlerAdUpdate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdUsecase := mock_ad.NewMockUsecase(controller)
	adDelivery := delivery.NewAdDelivery(mockAdUsecase)
	echo_ := echo.New()
	echo_.Validator = RequestValidator.NewRequestValidator()
	adDelivery.Configure(echo_, &middlewares.Manager{})

	dateTimeArr, err := timestamps.NewDateTime("27.10.2021 19:50")
	assert.Nil(t, err)
	ad := &models.Ad{
		Id:          1,
		LocDep:      "Общежитие №10",
		LocArr:      "УЛК",
		DateTimeArr: *dateTimeArr,
		Item:        "Зачётная книжка",
		MinPrice:    500,
		Comment:     "Поеду на велосипеде",
	}
	expectedAd := &models.Ad{
		Id:             ad.Id,
		UserAuthorVkId: 2,
		LocDep:         ad.LocDep,
		LocArr:         ad.LocArr,
		DateTimeArr:    ad.DateTimeArr,
		Item:           ad.Item,
		MinPrice:       ad.MinPrice,
		Comment:        ad.Comment,
	}

	mockAdUsecase.
		EXPECT().
		Update(gomock.Eq(ad)).
		Return(response.NewResponse(consts.OK, expectedAd))

	jsonRequest, err := json.Marshal(ad)
	assert.Nil(t, err)

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedAd,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(jsonRequest)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/ads/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(ad.Id), 10))
	context.Set(consts.EchoContextKeyUserVkId, expectedAd.UserAuthorVkId)

	handler := adDelivery.HandlerAdUpdate()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestAdDelivery_HandlerAdSearch(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdUsecase := mock_ad.NewMockUsecase(controller)
	adDelivery := delivery.NewAdDelivery(mockAdUsecase)
	echo_ := echo.New()
	echo_.Validator = RequestValidator.NewRequestValidator()
	adDelivery.Configure(echo_, &middlewares.Manager{})

	dateTimeArr1, err := timestamps.NewDateTime("04.11.2021 19:40")
	assert.Nil(t, err)
	dateTimeArr2, err := timestamps.NewDateTime("04.11.2021 19:45")
	assert.Nil(t, err)
	adsSearch := &models.AdsSearch{
		LocDep:      "Общежитие",
		LocArr:      "СК",
		DateTimeArr: *dateTimeArr1,
		MaxPrice:    1000,
	}
	expectedAds := &models.Ads{
		&models.Ad{
			Id:             1,
			UserAuthorVkId: 10,
			LocDep:         "Общежитие №10",
			LocArr:         "УЛК",
			DateTimeArr:    *dateTimeArr1,
			Item:           "Тубус",
			MinPrice:       500,
			Comment:        "Поеду на коньках",
		},
		&models.Ad{
			Id:             2,
			UserAuthorVkId: 20,
			LocDep:         "Общежитие №9",
			LocArr:         "СК",
			DateTimeArr:    *dateTimeArr2,
			Item:           "Спортивная форма",
			MinPrice:       600,
			Comment:        "Поеду на роликах :)",
		},
	}

	mockAdUsecase.
		EXPECT().
		Search(gomock.Eq(adsSearch)).
		Return(response.NewResponse(consts.OK, expectedAds))

	jsonRequest, err := json.Marshal(adsSearch)
	assert.Nil(t, err)

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedAds,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodPost, "/api/ads/search", strings.NewReader(string(jsonRequest)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)

	handler := adDelivery.HandlerAdsSearch()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}
