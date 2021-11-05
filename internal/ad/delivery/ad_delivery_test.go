package delivery_test

import (
	"encoding/json"
	"github.com/TechnoHandOver/backend/internal/ad/delivery"
	"github.com/TechnoHandOver/backend/internal/ad/mock_ad"
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

func TestAdDelivery_HandlerAdCreate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdUsecase := mock_ad.NewMockUsecase(controller)
	adDelivery := delivery.NewAdDelivery(mockAdUsecase)
	echo_ := echo.New()
	adDelivery.Configure(echo_)

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
			return response.NewResponse(http.StatusCreated, ad)
		})

	jsonRequest, err := json.Marshal(ad)
	assert.Nil(t, err)

	jsonExpectedResponse, err := json.Marshal(response.DataResponse{
		Data: expectedAd,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodPost, "/api/ad", strings.NewReader(string(jsonRequest)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)

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
	adDelivery.Configure(echo_)

	type AdGetRequest struct {
		Id uint32 `param:"id"`
	}

	adGetRequest := &AdGetRequest{
		Id: 1,
	}
	dateTimeArr, err := timestamps.NewDateTime("27.10.2021 19:44")
	assert.Nil(t, err)
	expectedAd := &models.Ad{
		Id:             adGetRequest.Id,
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
		Get(gomock.Eq(adGetRequest.Id)).
		Return(response.NewResponse(http.StatusOK, expectedAd))

	jsonExpectedResponse, err := json.Marshal(response.DataResponse{
		Data: expectedAd,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/ad/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(adGetRequest.Id), 10))

	handler := adDelivery.HandlerAdGet()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestAdDelivery_HandlerAdGet_notFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdUsecase := mock_ad.NewMockUsecase(controller)
	adDelivery := delivery.NewAdDelivery(mockAdUsecase)
	echo_ := echo.New()
	adDelivery.Configure(echo_)

	type AdGetRequest struct {
		Id uint32
	}

	adGetRequest := &AdGetRequest{
		Id: 1,
	}

	mockAdUsecase.
		EXPECT().
		Get(gomock.Eq(adGetRequest.Id)).
		Return(response.NewEmptyResponse(http.StatusNotFound))

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/ad/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(adGetRequest.Id), 10))

	handler := adDelivery.HandlerAdGet()

	err := handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestAdDelivery_HandlerAdUpdate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdUsecase := mock_ad.NewMockUsecase(controller)
	adDelivery := delivery.NewAdDelivery(mockAdUsecase)
	echo_ := echo.New()
	adDelivery.Configure(echo_)

	type AdUpdateRequest struct {
		Id          uint32
		LocDep      string              `json:"locDep,omitempty"`
		LocArr      string              `json:"locArr,omitempty"`
		DateTimeArr timestamps.DateTime `json:"dateTimeArr,omitempty"`
		Item        string              `json:"item,omitempty"`
		MinPrice    uint32              `json:"minPrice,omitempty"`
		Comment     string              `json:"comment,omitempty"`
	}

	dateTimeArr, err := timestamps.NewDateTime("27.10.2021 19:50")
	assert.Nil(t, err)
	adUpdateRequest := &AdUpdateRequest{
		Id: 1,
		LocDep:         "Общежитие №10",
		LocArr:         "УЛК",
		DateTimeArr:    *dateTimeArr,
		Item:           "Зачётная книжка",
		MinPrice:       500,
		Comment:        "Поеду на велосипеде",
	}
	ad := &models.Ad{
		Id: adUpdateRequest.Id,
		LocDep: adUpdateRequest.LocDep,
		LocArr: adUpdateRequest.LocArr,
		DateTimeArr: adUpdateRequest.DateTimeArr,
		Item: adUpdateRequest.Item,
		MinPrice: adUpdateRequest.MinPrice,
		Comment: adUpdateRequest.Comment,
	}
	expectedAd := &models.Ad{
		Id:             ad.Id,
		UserAuthorVkId: 1,
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
		Return(response.NewResponse(http.StatusOK, expectedAd))

	jsonRequest, err := json.Marshal(ad)
	assert.Nil(t, err)

	jsonExpectedResponse, err := json.Marshal(response.DataResponse{
		Data: expectedAd,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(jsonRequest)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/ad/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(adUpdateRequest.Id), 10))

	handler := adDelivery.HandlerAdUpdate()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestAdDelivery_HandlerAdUpdate_notFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdUsecase := mock_ad.NewMockUsecase(controller)
	adDelivery := delivery.NewAdDelivery(mockAdUsecase)
	echo_ := echo.New()
	adDelivery.Configure(echo_)

	type AdUpdateRequest struct {
		Id          uint32
		LocDep      string              `json:"locDep,omitempty"`
		LocArr      string              `json:"locArr,omitempty"`
		DateTimeArr timestamps.DateTime `json:"dateTimeArr,omitempty"`
		Item        string              `json:"item,omitempty"`
		MinPrice    uint32              `json:"minPrice,omitempty"`
		Comment     string              `json:"comment,omitempty"`
	}

	dateTimeArr, err := timestamps.NewDateTime("27.10.2021 19:50")
	assert.Nil(t, err)
	adUpdateRequest := &AdUpdateRequest{
		Id: 1,
		LocDep:         "Общежитие №10",
		LocArr:         "УЛК",
		DateTimeArr:    *dateTimeArr,
		Item:           "Зачётная книжка",
		MinPrice:       500,
		Comment:        "Поеду на велосипеде",
	}
	ad := &models.Ad{
		Id: adUpdateRequest.Id,
		LocDep: adUpdateRequest.LocDep,
		LocArr: adUpdateRequest.LocArr,
		DateTimeArr: adUpdateRequest.DateTimeArr,
		Item: adUpdateRequest.Item,
		MinPrice: adUpdateRequest.MinPrice,
		Comment: adUpdateRequest.Comment,
	}

	mockAdUsecase.
		EXPECT().
		Update(gomock.Eq(ad)).
		Return(response.NewEmptyResponse(http.StatusNotFound))

	jsonRequest, err := json.Marshal(adUpdateRequest)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(jsonRequest)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/ad/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(adUpdateRequest.Id), 10))

	handler := adDelivery.HandlerAdUpdate()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestAdDelivery_HandlerAdSearch(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdUsecase := mock_ad.NewMockUsecase(controller)
	adDelivery := delivery.NewAdDelivery(mockAdUsecase)
	echo_ := echo.New()
	adDelivery.Configure(echo_)

	dateTimeArr1, err := timestamps.NewDateTime("04.11.2021 19:40")
	assert.Nil(t, err)
	dateTimeArr2, err := timestamps.NewDateTime("04.11.2021 19:45")
	assert.Nil(t, err)
	adsSearch := &models.AdsSearch{
		LocDep: "Общежитие",
		LocArr: "СК",
		DateTimeArr: *dateTimeArr1,
		MaxPrice: 1000,
	}
	expectedAds := &models.Ads{
		&models.Ad{
			Id: 1,
			LocDep: "Общежитие №10",
			LocArr: "УЛК",
			DateTimeArr: *dateTimeArr1,
			Item: "Тубус",
			MinPrice: 500,
			Comment: "Поеду на коньках",
		},
		&models.Ad{
			Id: 1,
			LocDep: "Общежитие №9",
			LocArr: "СК",
			DateTimeArr: *dateTimeArr2,
			Item: "Спортивная форма",
			MinPrice: 600,
			Comment: "Поеду на роликах :)",
		},
	}

	mockAdUsecase.
		EXPECT().
		Search(gomock.Eq(adsSearch)).
		Return(response.NewResponse(http.StatusOK, expectedAds))

	jsonRequest, err := json.Marshal(adsSearch)
	assert.Nil(t, err)

	jsonExpectedResponse, err := json.Marshal(response.DataResponse{
		Data: expectedAds,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodPost, "/api/ad/search", strings.NewReader(string(jsonRequest)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)

	handler := adDelivery.HandlerAdSearch()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}