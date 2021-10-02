package usecase

import (
	"database/sql"
	"github.com/TechnoHandOver/backend/internal/ads/repository"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/labstack/gommon/log"
	"net/http"
)

type AdsUsecase struct {
	adsRepository *repository.AdsRepository //TODO: возможно, здесь не указатель, а копия
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
			return response.NewResponse(http.StatusUnauthorized, models.Error{
				Message: "Not authorized",
			})
		}
		log.Error(err)
		return response.NewResponse(http.StatusInternalServerError, nil)
	}
	return response.NewResponse(http.StatusCreated, createdAds)
}

func (adsUsecase *AdsUsecase) List() *response.Response {
	adses, err := adsUsecase.adsRepository.Select()
	if err != nil {
		log.Error(err)
		return response.NewResponse(http.StatusInternalServerError, nil)
	}
	return response.NewResponse(http.StatusOK, adses)
}
