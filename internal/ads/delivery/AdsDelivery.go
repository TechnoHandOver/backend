package delivery

import (
	"github.com/TechnoHandOver/backend/internal/ads/usecase"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/responser"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AdsDelivery struct {
	adsUsecase *usecase.AdsUsecase
}

func NewAdsDelivery(adsUsecase *usecase.AdsUsecase) *AdsDelivery {
	return &AdsDelivery{
		adsUsecase: adsUsecase,
	}
}

func (adsDelivery *AdsDelivery) Configure(echo_ *echo.Echo) {
	echo_.POST("/ads/create", adsDelivery.HandlerAdsCreate())
	echo_.GET("/ads/list", adsDelivery.HandlerAdsList())
}

func (adsDelivery *AdsDelivery) HandlerAdsCreate() echo.HandlerFunc {
	return func(context echo.Context) error {
		ads := new(models.Ads)
		if err := context.Bind(ads); err != nil {
			return context.NoContent(http.StatusInternalServerError)
		}

		//ads.authorUserId = ...

		return responser.Respond(context, adsDelivery.adsUsecase.Create(ads))
	}
}

func (adsDelivery *AdsDelivery) HandlerAdsList() echo.HandlerFunc {
	return func(context echo.Context) error {
		return responser.Respond(context, adsDelivery.adsUsecase.List())
	}
}
