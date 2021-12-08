package usecase

import (
	"fmt"
	"github.com/TechnoHandOver/backend/internal/ad"
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/models/timestamps"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/openlyinc/pointy"
	"time"
)

type AdUsecase struct {
	adRepository ad.Repository
}

func NewAdUsecaseImpl(repository ad.Repository) ad.Usecase {
	return &AdUsecase{
		adRepository: repository,
	}
}

func (adUsecase *AdUsecase) Create(ad_ *models.Ad) *response.Response {
	//TODO: assert.IsZero(ad_.Id) ?
	ad_, err := adUsecase.adRepository.Insert(ad_)
	if err != nil {
		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.Created, ad_)
}

func (adUsecase *AdUsecase) Get(id uint32) *response.Response {
	ad_, err := adUsecase.adRepository.Select(id)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.OK, ad_)
}

func (adUsecase *AdUsecase) Update(ad_ *models.Ad) *response.Response { //TODO: bug, only author can update it!
	existingAd, err := adUsecase.adRepository.Select(ad_.Id)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	if ad_.UserAuthorId != existingAd.UserAuthorId {
		return response.NewEmptyResponse(consts.Forbidden)
	}

	ad_, err = adUsecase.adRepository.Update(ad_)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.OK, ad_)
}

func (adUsecase *AdUsecase) Delete(userId uint32, id uint32) *response.Response {
	existingAd, err := adUsecase.adRepository.Select(id)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	if userId != existingAd.UserAuthorId {
		return response.NewEmptyResponse(consts.Forbidden)
	}

	ad_, err := adUsecase.adRepository.Delete(id)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.OK, ad_)
}

func (adUsecase *AdUsecase) Search(adsSearch *models.AdsSearch) *response.Response {
	adsSearch.NullUserExecutorVkId = pointy.Bool(true)
	if adsSearch.MinDateTimeArr == nil {
		minDateTimeArr := timestamps.DateTime(time.Now().Add(time.Duration(-12) * time.Hour))
		fmt.Println(minDateTimeArr.String())
		adsSearch.MinDateTimeArr = &minDateTimeArr
	}

	ads, err := adUsecase.adRepository.SelectArray(adsSearch)
	if err != nil {
		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.OK, ads)
}

func (adUsecase *AdUsecase) SetAdUserExecutor(userId uint32, adId uint32) *response.Response {
	ad_, err := adUsecase.adRepository.Select(adId)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	if ad_.UserAuthorId == userId {
		return response.NewEmptyResponse(consts.Forbidden)
	}
	if ad_.UserExecutorVkId != nil {
		return response.NewEmptyResponse(consts.Conflict)
	}

	adUserExecution := &models.AdUserExecution{
		AdId:           adId,
		UserExecutorId: userId,
	}
	adUserExecution, err = adUsecase.adRepository.InsertAdUserExecution(adUserExecution)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	updatedAd, err := adUsecase.adRepository.Select(adUserExecution.AdId)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.OK, updatedAd)
}

func (adUsecase *AdUsecase) UnsetAdUserExecutor(userId uint32, adId uint32) *response.Response {
	adUserExecution, err := adUsecase.adRepository.SelectAdUserExecution(adId)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	if adUserExecution.UserExecutorId != userId {
		return response.NewEmptyResponse(consts.Forbidden)
	}

	adUserExecution, err = adUsecase.adRepository.DeleteAdUserExecution(adId)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	updatedAd, err := adUsecase.adRepository.Select(adUserExecution.AdId)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.OK, updatedAd)
}
