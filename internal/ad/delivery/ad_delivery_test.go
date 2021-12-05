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
	HandoverTesting "github.com/TechnoHandOver/backend/internal/tools/testing"
	HandoverValidator "github.com/TechnoHandOver/backend/internal/tools/validator"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/openlyinc/pointy"
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
	echo_.Validator = HandoverValidator.NewRequestValidator()
	adDelivery.Configure(echo_, &middlewares.Manager{})

	dateTimeArr, err := timestamps.NewDateTime("27.10.2021 19:31")
	assert.Nil(t, err)
	ad := &models.Ad{
		UserAuthorId: 101,
		LocDep:       "Общежитие №10",
		LocArr:       "УЛК",
		DateTimeArr:  *dateTimeArr,
		Item:         "Зачётная книжка",
		MinPrice:     500,
		Comment:      "Поеду на велосипеде",
	}
	expectedAd := &models.Ad{
		Id:             1,
		UserAuthorId:   ad.UserAuthorId,
		UserAuthorVkId: 201,
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
			ad.UserAuthorVkId = expectedAd.UserAuthorVkId
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
	context.Set(consts.EchoContextKeyUserId, ad.UserAuthorId)

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
	echo_.Validator = HandoverValidator.NewRequestValidator()
	adDelivery.Configure(echo_, &middlewares.Manager{})

	dateTimeArr, err := timestamps.NewDateTime("27.10.2021 19:44")
	assert.Nil(t, err)
	expectedAd := &models.Ad{
		Id:             1,
		UserAuthorId:   101,
		UserAuthorVkId: 201,
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
	echo_.Validator = HandoverValidator.NewRequestValidator()
	adDelivery.Configure(echo_, &middlewares.Manager{})

	dateTimeArr, err := timestamps.NewDateTime("27.10.2021 19:50")
	assert.Nil(t, err)
	ad := &models.Ad{
		Id:           1,
		UserAuthorId: 101,
		LocDep:       "Общежитие №10",
		LocArr:       "УЛК",
		DateTimeArr:  *dateTimeArr,
		Item:         "Зачётная книжка",
		MinPrice:     500,
		Comment:      "Поеду на велосипеде",
	}
	expectedAd := &models.Ad{
		Id:             ad.Id,
		UserAuthorId:   ad.UserAuthorId,
		UserAuthorVkId: 201,
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
		DoAndReturn(func(ad *models.Ad) *response.Response {
			ad.UserAuthorVkId = expectedAd.UserAuthorVkId
			return response.NewResponse(consts.OK, ad)
		})

	jsonRequest, err := json.Marshal(expectedAd)
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
	context.SetParamValues(strconv.FormatUint(uint64(expectedAd.Id), 10))
	context.Set(consts.EchoContextKeyUserId, expectedAd.UserAuthorId)

	handler := adDelivery.HandlerAdUpdate()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestAdDelivery_HandlerAdDelete(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdUsecase := mock_ad.NewMockUsecase(controller)
	adDelivery := delivery.NewAdDelivery(mockAdUsecase)
	echo_ := echo.New()
	echo_.Validator = HandoverValidator.NewRequestValidator()
	adDelivery.Configure(echo_, &middlewares.Manager{})

	dateTimeArr, err := timestamps.NewDateTime("22.11.2021 16:55")
	assert.Nil(t, err)
	expectedAd := &models.Ad{
		Id:             1,
		UserAuthorId:   101,
		UserAuthorVkId: 201,
		LocDep:         "Общежитие №10",
		LocArr:         "УЛК",
		DateTimeArr:    *dateTimeArr,
		Item:           "Зачётная книжка",
		MinPrice:       500,
		Comment:        "Поеду на велосипеде",
	}

	mockAdUsecase.
		EXPECT().
		Delete(gomock.Eq(expectedAd.UserAuthorId), gomock.Eq(expectedAd.Id)).
		Return(response.NewResponse(consts.OK, expectedAd))

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedAd,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodDelete, "/", nil)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/ads/:id")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(expectedAd.Id), 10))
	context.Set(consts.EchoContextKeyUserId, expectedAd.UserAuthorId)

	handler := adDelivery.HandlerAdDelete()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestAdDelivery_HandlerAdsList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdUsecase := mock_ad.NewMockUsecase(controller)
	adDelivery := delivery.NewAdDelivery(mockAdUsecase)
	echo_ := echo.New()
	echo_.Validator = HandoverValidator.NewRequestValidator()
	adDelivery.Configure(echo_, &middlewares.Manager{})

	dateTimeArr1, err := timestamps.NewDateTime("04.11.2021 19:40")
	assert.Nil(t, err)
	dateTimeArr2, err := timestamps.NewDateTime("04.11.2021 19:45")
	assert.Nil(t, err)
	const userAuthorId uint32 = 101
	const userAuthorVkId uint32 = 201
	expectedAds := &models.Ads{
		&models.Ad{
			Id:             1,
			UserAuthorId:   userAuthorId,
			UserAuthorVkId: userAuthorVkId,
			LocDep:         "Общежитие №10",
			LocArr:         "УЛК",
			DateTimeArr:    *dateTimeArr1,
			Item:           "Тубус",
			MinPrice:       500,
			Comment:        "Поеду на коньках",
		},
		&models.Ad{
			Id:             2,
			UserAuthorId:   userAuthorId,
			UserAuthorVkId: userAuthorVkId,
			LocDep:         "Общежитие №9",
			LocArr:         "СК",
			DateTimeArr:    *dateTimeArr2,
			Item:           "Спортивная форма",
			MinPrice:       600,
			Comment:        "Поеду на роликах :)",
		},
	}
	adsSearch := HandoverTesting.NewAdsSearchByUserAuthorId(userAuthorId)

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

	request := httptest.NewRequest(http.MethodGet, "/api/ads/list", strings.NewReader(string(jsonRequest)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.Set(consts.EchoContextKeyUserId, userAuthorId)

	handler := adDelivery.HandlerAdsList()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestAdDelivery_HandlerAdsSearch(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdUsecase := mock_ad.NewMockUsecase(controller)
	adDelivery := delivery.NewAdDelivery(mockAdUsecase)
	echo_ := echo.New()
	echo_.Validator = HandoverValidator.NewRequestValidator()
	adDelivery.Configure(echo_, &middlewares.Manager{})

	dateTimeArr1, err := timestamps.NewDateTime("04.11.2021 19:40")
	assert.Nil(t, err)
	dateTimeArr2, err := timestamps.NewDateTime("04.11.2021 19:45")
	assert.Nil(t, err)
	adsSearch := HandoverTesting.NewAdsSearchBySecondaryFields("Общежитие", "СК", *dateTimeArr1, 1000)
	expectedAds := &models.Ads{
		&models.Ad{
			Id:             1,
			UserAuthorId:   101,
			UserAuthorVkId: 201,
			LocDep:         "Общежитие №10",
			LocArr:         "УЛК",
			DateTimeArr:    *dateTimeArr1,
			Item:           "Тубус",
			MinPrice:       500,
			Comment:        "Поеду на коньках",
		},
		&models.Ad{
			Id:             2,
			UserAuthorId:   102,
			UserAuthorVkId: 202,
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

	request := httptest.NewRequest(http.MethodGet, "/api/ads/search", strings.NewReader(string(jsonRequest)))
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

func TestAdDelivery_HandlerAdExecutionCreate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdUsecase := mock_ad.NewMockUsecase(controller)
	adDelivery := delivery.NewAdDelivery(mockAdUsecase)
	echo_ := echo.New()
	echo_.Validator = HandoverValidator.NewRequestValidator()
	adDelivery.Configure(echo_, &middlewares.Manager{})

	const userId uint32 = 102
	dateTimeArr, err := timestamps.NewDateTime("05.12.2021 19:50")
	assert.Nil(t, err)
	expectedAd := &models.Ad{
		Id:             1,
		UserAuthorId:   101,
		UserAuthorVkId: 201,
		LocDep:         "Общежитие №10",
		LocArr:         "УЛК",
		DateTimeArr:    *dateTimeArr,
		Item:           "Зачётная книжка",
		MinPrice:       500,
		Comment:        "Поеду на велосипеде",
	}

	mockAdUsecase.
		EXPECT().
		SetAdUserExecutor(gomock.Eq(userId), gomock.Eq(expectedAd.Id)).
		Return(response.NewResponse(consts.Created, expectedAd))

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedAd,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodPost, "/", nil)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/ads/:id/execution")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(expectedAd.Id), 10))
	context.Set(consts.EchoContextKeyUserId, userId)

	handler := adDelivery.HandlerAdExecutionCreate()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}

func TestAdDelivery_HandlerAdExecutionDelete(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdUsecase := mock_ad.NewMockUsecase(controller)
	adDelivery := delivery.NewAdDelivery(mockAdUsecase)
	echo_ := echo.New()
	echo_.Validator = HandoverValidator.NewRequestValidator()
	adDelivery.Configure(echo_, &middlewares.Manager{})

	const userId uint32 = 102
	dateTimeArr, err := timestamps.NewDateTime("05.12.2021 20:35")
	assert.Nil(t, err)
	expectedAd := &models.Ad{
		Id:               1,
		UserAuthorId:     101,
		UserAuthorVkId:   201,
		UserExecutorVkId: pointy.Uint32(userId),
		LocDep:           "Общежитие №10",
		LocArr:           "УЛК",
		DateTimeArr:      *dateTimeArr,
		Item:             "Зачётная книжка",
		MinPrice:         500,
		Comment:          "Поеду на велосипеде",
	}

	mockAdUsecase.
		EXPECT().
		UnsetAdUserExecutor(gomock.Eq(userId), gomock.Eq(expectedAd.Id)).
		Return(response.NewResponse(consts.OK, expectedAd))

	jsonExpectedResponse, err := json.Marshal(responser.DataResponse{
		Data: expectedAd,
	})
	assert.Nil(t, err)
	jsonExpectedResponse = append(jsonExpectedResponse, '\n')

	request := httptest.NewRequest(http.MethodDelete, "/", nil)

	recorder := httptest.NewRecorder()
	context := echo_.NewContext(request, recorder)
	context.SetPath("/api/ads/:id/execution")
	context.SetParamNames("id")
	context.SetParamValues(strconv.FormatUint(uint64(expectedAd.Id), 10))
	context.Set(consts.EchoContextKeyUserId, userId)

	handler := adDelivery.HandlerAdExecutionDelete()

	err = handler(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	responseBody, err := ioutil.ReadAll(recorder.Body)
	assert.Nil(t, err)
	assert.Equal(t, jsonExpectedResponse, responseBody)
}
