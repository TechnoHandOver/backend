package usecase

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/ad"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"

	"net/http"
)

type AdUsecase struct {
	adRepository ad.Repository
}

func NewAdUsecaseImpl(adRepository ad.Repository) ad.Usecase {
	return &AdUsecase{
		adRepository: adRepository,
	}
}

func (adUsecase *AdUsecase) Create(ad_ *models.Ad) *response.Response {
	//TODO: assert.IsZero(ad_.Id) ?
	ad_, err := adUsecase.adRepository.Insert(ad_)
	if err != nil {
		return response.NewErrorResponse(err)
	}

	return response.NewResponse(http.StatusCreated, ad_)
}

func (adUsecase *AdUsecase) Get(id uint32) *response.Response {
	ad_, err := adUsecase.adRepository.Select(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return response.NewEmptyResponse(http.StatusNotFound) //TODO: вообще, тут отдавать не http.StatusNotFound, а SomeMapClass.NOT_FOUND_ERROR...; везде так
		}

		return response.NewErrorResponse(err)
	}

	return response.NewResponse(http.StatusOK, ad_)
}

func (adUsecase *AdUsecase) Update(ad_ *models.Ad) *response.Response {
	ad_, err := adUsecase.adRepository.Update(ad_)
	if err != nil {
		if err == sql.ErrNoRows {
			return response.NewEmptyResponse(http.StatusNotFound)
		}

		return response.NewErrorResponse(err)
	}

	return response.NewResponse(http.StatusOK, ad_)
}

func (adUsecase *AdUsecase) Search(adsSearch *models.AdsSearch) *response.Response {
	ads, err := adUsecase.adRepository.SelectArray(adsSearch)
	if err != nil {
		return response.NewErrorResponse(err)
	}

	return response.NewResponse(http.StatusOK, ads)
}
