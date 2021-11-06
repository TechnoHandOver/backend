package delivery

import (
	"github.com/TechnoHandOver/backend/internal/ad"
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/middlewares"
	"github.com/TechnoHandOver/backend/internal/models"
	. "github.com/TechnoHandOver/backend/internal/models/timestamps"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/TechnoHandOver/backend/internal/tools/responser"
	"github.com/labstack/echo/v4"
)

type AdDelivery struct {
	adUsecase ad.Usecase
}

func NewAdDelivery(usecase ad.Usecase) *AdDelivery {
	return &AdDelivery{
		adUsecase: usecase,
	}
}

func (adDelivery *AdDelivery) Configure(echo_ *echo.Echo, middlewaresManager *middlewares.Manager) {
	echo_.POST("/api/ad", adDelivery.HandlerAdCreate(), middlewaresManager.AuthMiddleware.CheckAuth())
	echo_.GET("/api/ad/:id", adDelivery.HandlerAdGet())
	echo_.PUT("/api/ad/:id", adDelivery.HandlerAdUpdate(), middlewaresManager.AuthMiddleware.CheckAuth())
	echo_.GET("/api/ad/search", adDelivery.HandlerAdSearch())
}

func (adDelivery *AdDelivery) HandlerAdCreate() echo.HandlerFunc {
	return func(context echo.Context) error {
		ad_ := new(models.Ad)
		if err := context.Bind(ad_); err != nil {
			return responser.Respond(context, response.NewErrorResponse(err))
		}

		ad_.UserAuthorVkId = context.Get(consts.EchoContextKeyUserVkId).(uint32)

		return responser.Respond(context, adDelivery.adUsecase.Create(ad_))
	}
}

func (adDelivery *AdDelivery) HandlerAdGet() echo.HandlerFunc {
	type AdGetRequest struct {
		Id uint32 `param:"id"`
	}

	return func(context echo.Context) error {
		adGetRequest := new(AdGetRequest)
		if err := context.Bind(adGetRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(err))
		}

		return responser.Respond(context, adDelivery.adUsecase.Get(adGetRequest.Id))
	}
}

func (adDelivery *AdDelivery) HandlerAdUpdate() echo.HandlerFunc {
	type AdUpdateRequest struct {
		Id          uint32   `param:"id"`
		LocDep      string   `json:"locDep,omitempty"`
		LocArr      string   `json:"locArr,omitempty"`
		DateTimeArr DateTime `json:"dateTimeArr,omitempty"`
		Item        string   `json:"item"`
		MinPrice    uint32   `json:"minPrice,omitempty"`
		Comment     string   `json:"comment,omitempty"` //TODO: нужны валидаторы моделей (length(str) <= 100 и подобное...)
	}

	return func(context echo.Context) error {
		adUpdateRequest := new(AdUpdateRequest)
		if err := context.Bind(adUpdateRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(err))
		}

		ad_ := &models.Ad{
			Id:          adUpdateRequest.Id,
			LocDep:      adUpdateRequest.LocDep,
			LocArr:      adUpdateRequest.LocArr,
			DateTimeArr: adUpdateRequest.DateTimeArr,
			Item:        adUpdateRequest.Item,
			MinPrice:    adUpdateRequest.MinPrice,
			Comment:     adUpdateRequest.Comment,
		}

		return responser.Respond(context, adDelivery.adUsecase.Update(ad_))
	}
}

func (adDelivery *AdDelivery) HandlerAdSearch() echo.HandlerFunc {
	return func(context echo.Context) error {
		adsSearch := new(models.AdsSearch)
		if err := context.Bind(adsSearch); err != nil {
			return responser.Respond(context, response.NewErrorResponse(err))
		}

		return responser.Respond(context, adDelivery.adUsecase.Search(adsSearch))
	}
}
