package usecase_test

import (
	"github.com/TechnoHandOver/backend/internal/ad/mock_ad"
	"github.com/TechnoHandOver/backend/internal/ad/usecase"
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/models/timestamps"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	HandoverTesting "github.com/TechnoHandOver/backend/internal/tools/testing"
	"github.com/golang/mock/gomock"
	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdUsecase_Create(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdRepository := mock_ad.NewMockRepository(controller)
	adUsecase := usecase.NewAdUsecaseImpl(mockAdRepository)

	dateTimeArr, err := timestamps.NewDateTime("04.11.2021 19:20")
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

	mockAdRepository.
		EXPECT().
		Insert(gomock.Eq(ad)).
		DoAndReturn(func(ad *models.Ad) (*models.Ad, error) {
			ad.Id = expectedAd.Id
			ad.UserAuthorVkId = expectedAd.UserAuthorVkId
			return ad, nil
		})

	response_ := adUsecase.Create(ad)
	assert.Equal(t, response.NewResponse(consts.Created, expectedAd), response_)
}

func TestAdUsecase_Get(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdRepository := mock_ad.NewMockRepository(controller)
	adUsecase := usecase.NewAdUsecaseImpl(mockAdRepository)

	dateTimeArr, err := timestamps.NewDateTime("04.11.2021 19:20")
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

	mockAdRepository.
		EXPECT().
		Select(gomock.Eq(expectedAd.Id)).
		Return(expectedAd, nil)

	response_ := adUsecase.Get(expectedAd.Id)
	assert.Equal(t, response.NewResponse(consts.OK, expectedAd), response_)
}

func TestAdUsecase_Get_notFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdRepository := mock_ad.NewMockRepository(controller)
	adUsecase := usecase.NewAdUsecaseImpl(mockAdRepository)

	const id uint32 = 1

	mockAdRepository.
		EXPECT().
		Select(gomock.Eq(id)).
		Return(nil, consts.RepErrNotFound)

	response_ := adUsecase.Get(id)
	assert.Equal(t, response.NewEmptyResponse(consts.NotFound), response_)
}

func TestAdUsecase_Update(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdRepository := mock_ad.NewMockRepository(controller)
	adUsecase := usecase.NewAdUsecaseImpl(mockAdRepository)

	dateTimeArr1, err := timestamps.NewDateTime("04.11.2021 19:40")
	assert.Nil(t, err)
	ad := &models.Ad{
		Id:           1,
		UserAuthorId: 101,
		LocDep:       "Общежитие №10",
		LocArr:       "УЛК",
		DateTimeArr:  *dateTimeArr1,
		Item:         "Зачётная книжка",
		MinPrice:     500,
		Comment:      "Поеду на велосипеде",
	}
	dateTimeArr2, err := timestamps.NewDateTime("04.11.2021 19:45")
	assert.Nil(t, err)
	expectedAd := &models.Ad{
		Id:             ad.Id,
		UserAuthorId:   ad.UserAuthorId,
		UserAuthorVkId: 201,
		LocDep:         "Общежитие №9",
		LocArr:         "СК",
		DateTimeArr:    *dateTimeArr2,
		Item:           "Спортивная форма",
		MinPrice:       600,
		Comment:        "Поеду на роликах :)",
	}

	call := mockAdRepository.
		EXPECT().
		Select(gomock.Eq(ad.Id)).
		Return(expectedAd, nil)

	mockAdRepository.
		EXPECT().
		Update(gomock.Eq(ad)).
		Return(expectedAd, nil).
		After(call)

	response_ := adUsecase.Update(ad)
	assert.Equal(t, response.NewResponse(consts.OK, expectedAd), response_)
}

func TestAdUsecase_Update_forbidden(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdRepository := mock_ad.NewMockRepository(controller)
	adUsecase := usecase.NewAdUsecaseImpl(mockAdRepository)

	dateTimeArr1, err := timestamps.NewDateTime("24.11.2021 13:50")
	assert.Nil(t, err)
	ad := &models.Ad{
		Id:           1,
		UserAuthorId: 101,
		LocDep:       "Общежитие №10",
		LocArr:       "УЛК",
		DateTimeArr:  *dateTimeArr1,
		Item:         "Зачётная книжка",
		MinPrice:     500,
		Comment:      "Поеду на велосипеде",
	}
	dateTimeArr2, err := timestamps.NewDateTime("24.11.2021 13:55")
	assert.Nil(t, err)
	expectedAd := &models.Ad{
		Id:             ad.Id,
		UserAuthorId:   102,
		UserAuthorVkId: 202,
		LocDep:         "Общежитие №9",
		LocArr:         "СК",
		DateTimeArr:    *dateTimeArr2,
		Item:           "Спортивная форма",
		MinPrice:       600,
		Comment:        "Поеду на роликах :)",
	}

	mockAdRepository.
		EXPECT().
		Select(gomock.Eq(ad.Id)).
		Return(expectedAd, nil)

	response_ := adUsecase.Update(ad)
	assert.Equal(t, response.NewEmptyResponse(consts.Forbidden), response_)
}

func TestAdUsecase_Update_notFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdRepository := mock_ad.NewMockRepository(controller)
	adUsecase := usecase.NewAdUsecaseImpl(mockAdRepository)

	dateTimeArr, err := timestamps.NewDateTime("04.11.2021 19:35")
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

	mockAdRepository.
		EXPECT().
		Select(gomock.Eq(ad.Id)).
		Return(nil, consts.RepErrNotFound)

	response_ := adUsecase.Update(ad)
	assert.Equal(t, response.NewEmptyResponse(consts.NotFound), response_)
}

func TestAdUsecase_Delete(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdRepository := mock_ad.NewMockRepository(controller)
	adUsecase := usecase.NewAdUsecaseImpl(mockAdRepository)

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

	call := mockAdRepository.
		EXPECT().
		Select(gomock.Eq(expectedAd.Id)).
		Return(expectedAd, nil)

	mockAdRepository.
		EXPECT().
		Delete(gomock.Eq(expectedAd.Id)).
		Return(expectedAd, nil).
		After(call)

	response_ := adUsecase.Delete(expectedAd.UserAuthorId, expectedAd.Id)
	assert.Equal(t, response.NewResponse(consts.OK, expectedAd), response_)
}

func TestAdUsecase_Delete_forbidden(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdRepository := mock_ad.NewMockRepository(controller)
	adUsecase := usecase.NewAdUsecaseImpl(mockAdRepository)

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

	mockAdRepository.
		EXPECT().
		Select(gomock.Eq(expectedAd.Id)).
		Return(expectedAd, nil)

	response_ := adUsecase.Delete(expectedAd.UserAuthorId+1, expectedAd.Id)
	assert.Equal(t, response.NewEmptyResponse(consts.Forbidden), response_)
}

func TestAdUsecase_Delete_notFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdRepository := mock_ad.NewMockRepository(controller)
	adUsecase := usecase.NewAdUsecaseImpl(mockAdRepository)

	const id uint32 = 1
	const userAuthorId uint32 = 101

	mockAdRepository.
		EXPECT().
		Select(gomock.Eq(id)).
		Return(nil, consts.RepErrNotFound)

	response_ := adUsecase.Delete(userAuthorId, id)
	assert.Equal(t, response.NewEmptyResponse(consts.NotFound), response_)
}

func TestAdUsecase_Search(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdRepository := mock_ad.NewMockRepository(controller)
	adUsecase := usecase.NewAdUsecaseImpl(mockAdRepository)

	dateTimeArr1, err := timestamps.NewDateTime("04.11.2021 19:40")
	assert.Nil(t, err)
	dateTimeArr2, err := timestamps.NewDateTime("04.11.2021 19:45")
	assert.Nil(t, err)
	adsSearch := HandoverTesting.NewAdsSearch(101, "Общежитие", "СК", *dateTimeArr1, 1000)
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

	mockAdRepository.
		EXPECT().
		SelectArray(gomock.Eq(adsSearch)).
		Return(expectedAds, nil)

	response_ := adUsecase.Search(adsSearch)
	assert.Equal(t, response.NewResponse(consts.OK, expectedAds), response_)
}

func TestAdUsecase_SetAdUserExecutor(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdRepository := mock_ad.NewMockRepository(controller)
	adUsecase := usecase.NewAdUsecaseImpl(mockAdRepository)

	dateTimeArr, err := timestamps.NewDateTime("05.12.2021 20:00")
	assert.Nil(t, err)
	ad := &models.Ad{
		Id:             1,
		UserAuthorId: 101,
		UserAuthorVkId: 201,
		LocDep:       "Общежитие №10",
		LocArr:       "УЛК",
		DateTimeArr:  *dateTimeArr,
		Item:         "Зачётная книжка",
		MinPrice:     500,
		Comment:      "Поеду на велосипеде",
	}
	adUserExecution := &models.AdUserExecution{
		AdId: ad.Id,
		UserExecutorId: ad.UserAuthorId + 1,
	}
	expectedAd := &models.Ad{
		Id:             ad.Id,
		UserAuthorId:   ad.UserAuthorId,
		UserAuthorVkId: ad.UserAuthorVkId,
		UserExecutorVkId: &adUserExecution.UserExecutorId,
		LocDep:         ad.LocDep,
		LocArr:         ad.LocArr,
		DateTimeArr:    ad.DateTimeArr,
		Item:           ad.Item,
		MinPrice:       ad.MinPrice,
		Comment:        ad.Comment,
	}

	callSelect1 := mockAdRepository.
		EXPECT().
		Select(gomock.Eq(ad.Id)).
		Return(ad, nil)
	callInsertAdUserExecution := mockAdRepository.
		EXPECT().
		InsertAdUserExecution(gomock.Eq(adUserExecution)).
		Return(adUserExecution, nil).
		After(callSelect1)
	mockAdRepository.
		EXPECT().
		Select(gomock.Eq(ad.Id)).
		Return(expectedAd, nil).
		After(callInsertAdUserExecution)

	response_ := adUsecase.SetAdUserExecutor(adUserExecution.UserExecutorId, adUserExecution.AdId)
	assert.Equal(t, response.NewResponse(consts.OK, expectedAd), response_)
}

func TestAdUsecase_SetAdUserExecutor_self(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdRepository := mock_ad.NewMockRepository(controller)
	adUsecase := usecase.NewAdUsecaseImpl(mockAdRepository)

	dateTimeArr, err := timestamps.NewDateTime("05.12.2021 20:05")
	assert.Nil(t, err)
	ad := &models.Ad{
		Id:             1,
		UserAuthorId: 101,
		UserAuthorVkId: 201,
		LocDep:       "Общежитие №10",
		LocArr:       "УЛК",
		DateTimeArr:  *dateTimeArr,
		Item:         "Зачётная книжка",
		MinPrice:     500,
		Comment:      "Поеду на велосипеде",
	}
	adUserExecution := &models.AdUserExecution{
		AdId: ad.Id,
		UserExecutorId: ad.UserAuthorId,
	}

	mockAdRepository.
		EXPECT().
		Select(gomock.Eq(ad.Id)).
		Return(ad, nil)

	response_ := adUsecase.SetAdUserExecutor(adUserExecution.UserExecutorId, adUserExecution.AdId)
	assert.Equal(t, response.NewEmptyResponse(consts.Forbidden), response_)
}

func TestAdUsecase_SetAdUserExecutor_conflict(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdRepository := mock_ad.NewMockRepository(controller)
	adUsecase := usecase.NewAdUsecaseImpl(mockAdRepository)

	dateTimeArr, err := timestamps.NewDateTime("05.12.2021 20:10")
	assert.Nil(t, err)
	ad := &models.Ad{
		Id:             1,
		UserAuthorId: 101,
		UserAuthorVkId: 201,
		UserExecutorVkId: pointy.Uint32(202),
		LocDep:       "Общежитие №10",
		LocArr:       "УЛК",
		DateTimeArr:  *dateTimeArr,
		Item:         "Зачётная книжка",
		MinPrice:     500,
		Comment:      "Поеду на велосипеде",
	}
	adUserExecution := &models.AdUserExecution{
		AdId: ad.Id,
		UserExecutorId: ad.UserAuthorId + 1,
	}

	mockAdRepository.
		EXPECT().
		Select(gomock.Eq(ad.Id)).
		Return(ad, nil)

	response_ := adUsecase.SetAdUserExecutor(adUserExecution.UserExecutorId, adUserExecution.AdId)
	assert.Equal(t, response.NewEmptyResponse(consts.Conflict), response_)
}
