package usecase_test

import (
	"github.com/TechnoHandOver/backend/internal/ad/mock_ad"
	"github.com/TechnoHandOver/backend/internal/ad/usecase"
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/models/timestamps"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/golang/mock/gomock"
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

	mockAdRepository.
		EXPECT().
		Insert(gomock.Eq(ad)).
		DoAndReturn(func(ad *models.Ad) (*models.Ad, error) {
			ad.Id = expectedAd.Id
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
		UserAuthorVkId: 2,
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
	dateTimeArr2, err := timestamps.NewDateTime("04.11.2021 19:45")
	assert.Nil(t, err)
	ad := &models.Ad{
		Id:             1,
		UserAuthorVkId: 2,
		LocDep:         "Общежитие №10",
		LocArr:         "УЛК",
		DateTimeArr:    *dateTimeArr1,
		Item:           "Зачётная книжка",
		MinPrice:       500,
		Comment:        "Поеду на велосипеде",
	}
	expectedAd := &models.Ad{
		Id:             ad.Id,
		UserAuthorVkId: ad.UserAuthorVkId,
		LocDep:         "Общежитие №9",
		LocArr:         "СК",
		DateTimeArr:    *dateTimeArr2,
		Item:           "Спортивная форма",
		MinPrice:       600,
		Comment:        "Поеду на роликах :)",
	}

	mockAdRepository.
		EXPECT().
		Update(gomock.Eq(ad)).
		Return(expectedAd, nil)

	response_ := adUsecase.Update(ad)
	assert.Equal(t, response.NewResponse(consts.OK, expectedAd), response_)
}

func TestAdUsecase_Update_notFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdRepository := mock_ad.NewMockRepository(controller)
	adUsecase := usecase.NewAdUsecaseImpl(mockAdRepository)

	dateTimeArr, err := timestamps.NewDateTime("04.11.2021 19:35")
	assert.Nil(t, err)
	ad := &models.Ad{
		Id:             1,
		UserAuthorVkId: 2,
		LocDep:         "Общежитие №10",
		LocArr:         "УЛК",
		DateTimeArr:    *dateTimeArr,
		Item:           "Зачётная книжка",
		MinPrice:       500,
		Comment:        "Поеду на велосипеде",
	}

	mockAdRepository.
		EXPECT().
		Update(gomock.Eq(ad)).
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
		UserAuthorVkId: 2,
		LocDep:         "Общежитие №10",
		LocArr:         "УЛК",
		DateTimeArr:    *dateTimeArr,
		Item:           "Зачётная книжка",
		MinPrice:       500,
		Comment:        "Поеду на велосипеде",
	}

	mockAdRepository.
		EXPECT().
		Delete(gomock.Eq(expectedAd.Id)).
		Return(expectedAd, nil)

	response_ := adUsecase.Delete(expectedAd.Id)
	assert.Equal(t, response.NewResponse(consts.OK, expectedAd), response_)
}

func TestAdUsecase_Delete_notFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAdRepository := mock_ad.NewMockRepository(controller)
	adUsecase := usecase.NewAdUsecaseImpl(mockAdRepository)

	const id uint32 = 1

	mockAdRepository.
		EXPECT().
		Delete(gomock.Eq(id)).
		Return(nil, consts.RepErrNotFound)

	response_ := adUsecase.Delete(id)
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
	adsSearch := &models.AdsSearch{
		LocDep:      "Общежитие",
		LocArr:      "СК",
		DateTimeArr: *dateTimeArr1,
		MaxPrice:    1000,
	}
	expectedAds := &models.Ads{
		&models.Ad{
			Id:             1,
			UserAuthorVkId: 2,
			LocDep:         "Общежитие №10",
			LocArr:         "УЛК",
			DateTimeArr:    *dateTimeArr1,
			Item:           "Тубус",
			MinPrice:       500,
			Comment:        "Поеду на коньках",
		},
		&models.Ad{
			Id:             1,
			UserAuthorVkId: 3,
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
