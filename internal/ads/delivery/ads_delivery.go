package delivery

import (
	"github.com/TechnoHandOver/backend/internal/ads/usecase"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/responser"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
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
	echo_.POST("/api/ads/create", adsDelivery.HandlerAdsCreate())
	echo_.GET("/api/ads/:id", adsDelivery.HandlerAdsGet())
	echo_.PUT("/api/ads/:id", adsDelivery.HandlerAdsUpdate())
	echo_.GET("/api/ads/list", adsDelivery.HandlerAdsList())
}

func (adsDelivery *AdsDelivery) HandlerAdsCreate() echo.HandlerFunc {
	return func(context echo.Context) error {
		ads := new(models.Ads)
		if err := context.Bind(ads); err != nil {
			return context.NoContent(http.StatusInternalServerError)
		}

		ads.UserAuthorId = 1 //TODO: убрать, когда будет реализована авторизация

		return responser.Respond(context, adsDelivery.adsUsecase.Create(ads))
	}
}

func (adsDelivery *AdsDelivery) HandlerAdsGet() echo.HandlerFunc {
	return func(context echo.Context) error {
		var id uint32
		if idUint64, err := strconv.ParseUint(context.Param("id"), 10, 32); err == nil {
			id = uint32(idUint64)
		} else {
			log.Error(err)
			return context.NoContent(http.StatusInternalServerError)
		}

		return responser.Respond(context, adsDelivery.adsUsecase.Get(id))
	}
}

func (adsDelivery *AdsDelivery) HandlerAdsUpdate() echo.HandlerFunc {
	return func(context echo.Context) error {
		var id uint32
		if idUint64, err := strconv.ParseUint(context.Param("id"), 10, 32); err == nil {
			id = uint32(idUint64)
		} else {
			log.Error(err)
			return context.NoContent(http.StatusInternalServerError)
		}

		updatedAds := new(models.AdsUpdate)
		if err := context.Bind(updatedAds); err != nil {
			return context.NoContent(http.StatusInternalServerError)
		}

		return responser.Respond(context, adsDelivery.adsUsecase.Update(id, updatedAds))
	}
}

func (adsDelivery *AdsDelivery) HandlerAdsList() echo.HandlerFunc {
	return func(context echo.Context) error {
		return responser.Respond(context, adsDelivery.adsUsecase.List())
	}
}