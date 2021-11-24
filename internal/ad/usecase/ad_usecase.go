package usecase

import (
	"github.com/TechnoHandOver/backend/internal/ad"
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"
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

func (adUsecase *AdUsecase) Update(ad_ *models.Ad) *response.Response {
	existingAd, err := adUsecase.adRepository.Select(ad_.Id)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	if ad_.UserAuthorVkId != existingAd.UserAuthorVkId {
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

func (adUsecase *AdUsecase) Delete(userVkId uint32, id uint32) *response.Response {
	existingAd, err := adUsecase.adRepository.Select(id)
	if err != nil {
		if err == consts.RepErrNotFound {
			return response.NewEmptyResponse(consts.NotFound)
		}

		return response.NewErrorResponse(consts.InternalError, err)
	}

	if userVkId != existingAd.UserAuthorVkId {
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
	ads, err := adUsecase.adRepository.SelectArray(adsSearch)
	if err != nil {
		return response.NewErrorResponse(consts.InternalError, err)
	}

	return response.NewResponse(consts.OK, ads)
}
