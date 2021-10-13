package usecase

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/ads/repository"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"log"
	"net/http"
)

type AdsUsecase struct {
	adsRepository *repository.AdsRepository
}

func NewAdsUsecase(adsRepository *repository.AdsRepository) *AdsUsecase {
	return &AdsUsecase{
		adsRepository: adsRepository,
	}
}

func (adsUsecase *AdsUsecase) Create(ads *models.Ads) *response.Response {
	createdAds, err := adsUsecase.adsRepository.Insert(ads)
	if err != nil {
		if err == sql.ErrNoRows {
			return response.NewErrorResponse(http.StatusUnauthorized, &models.Error{
				Message: "Not authorized",
			})
		}

		log.Println(err)
		return response.NewErrorResponse(http.StatusInternalServerError, nil)
	}

	return response.NewResponse(http.StatusCreated, createdAds)
}

func (adsUsecase *AdsUsecase) Get(id uint32) *response.Response {
	ads, err := adsUsecase.adsRepository.Select(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return response.NewErrorResponse(http.StatusNotFound, &models.Error{
				Message: "Not found",
			})
		}

		log.Println(err)
		return response.NewErrorResponse(http.StatusInternalServerError, nil)
	}

	return response.NewResponse(http.StatusOK, ads)
}

func (adsUsecase *AdsUsecase) Update(id uint32, adsUpdate *models.AdsUpdate) *response.Response {
	updatedAds, err := adsUsecase.adsRepository.Update(id, adsUpdate)
	if err != nil {
		if err == sql.ErrNoRows {
			return response.NewErrorResponse(http.StatusNotFound, &models.Error{
				Message: "Not found",
			})
		}

		log.Println(err)
		return response.NewErrorResponse(http.StatusInternalServerError, nil)
	}

	return response.NewResponse(http.StatusOK, updatedAds)
}

func (adsUsecase *AdsUsecase) List() *response.Response {
	adses, err := adsUsecase.adsRepository.SelectArray()
	if err != nil {
		log.Println(err)
		return response.NewErrorResponse(http.StatusInternalServerError, nil)
	}

	return response.NewResponse(http.StatusOK, adses)
}
