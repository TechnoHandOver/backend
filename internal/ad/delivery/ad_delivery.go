package delivery

import (
	"github.com/TechnoHandOver/backend/internal/ad"
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/middlewares"
	"github.com/TechnoHandOver/backend/internal/models"
	. "github.com/TechnoHandOver/backend/internal/models/timestamps"
	"github.com/TechnoHandOver/backend/internal/tools/parser"
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
	echo_.POST("/api/ads", adDelivery.HandlerAdCreate(), middlewaresManager.AuthMiddleware.CheckAuth())
	echo_.GET("/api/ads/:id", adDelivery.HandlerAdGet())
	echo_.PUT("/api/ads/:id", adDelivery.HandlerAdUpdate(), middlewaresManager.AuthMiddleware.CheckAuth())
	echo_.GET("/api/ads/search", adDelivery.HandlerAdsSearch())
}

func (adDelivery *AdDelivery) HandlerAdCreate() echo.HandlerFunc {
	type AdCreateRequest struct {
		LocDep      *string   `json:"locDep" validate:"required,gte=2,lte=100"`
		LocArr      *string   `json:"locArr" validate:"required,gte=2,lte=100"`
		DateTimeArr *DateTime `json:"dateTimeArr" validate:"required"`
		Item        *string   `json:"item" validate:"required,gte=3,lte=50"`
		MinPrice    *uint32   `json:"minPrice" validate:"required"`
		Comment     *string   `json:"comment" validate:"required,lte=100"`
	}

	return func(context echo.Context) error {
		adCreateRequest := new(AdCreateRequest)
		if err := parser.ParseRequest(context, adCreateRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		ad_ := &models.Ad{
			UserAuthorVkId: context.Get(consts.EchoContextKeyUserVkId).(uint32),
			LocDep:         *adCreateRequest.LocDep,
			LocArr:         *adCreateRequest.LocArr,
			DateTimeArr:    *adCreateRequest.DateTimeArr,
			Item:           *adCreateRequest.Item,
			MinPrice:       *adCreateRequest.MinPrice,
			Comment:        *adCreateRequest.Comment,
		}

		return responser.Respond(context, adDelivery.adUsecase.Create(ad_))
	}
}

func (adDelivery *AdDelivery) HandlerAdGet() echo.HandlerFunc {
	type AdGetRequest struct {
		Id *uint32 `param:"id" validate:"required"`
	}

	return func(context echo.Context) error {
		adGetRequest := new(AdGetRequest)
		if err := parser.ParseRequest(context, adGetRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		id := *adGetRequest.Id

		return responser.Respond(context, adDelivery.adUsecase.Get(id))
	}
}

func (adDelivery *AdDelivery) HandlerAdUpdate() echo.HandlerFunc {
	type AdUpdateRequest struct {
		Id          *uint32   `param:"id" validate:"required"`
		LocDep      *string   `json:"locDep" validate:"required,gte=2,lte=100"`
		LocArr      *string   `json:"locArr" validate:"required,gte=2,lte=100"`
		DateTimeArr *DateTime `json:"dateTimeArr" validate:"required"`
		Item        *string   `json:"item" validate:"required,gte=3,lte=50"`
		MinPrice    *uint32   `json:"minPrice" validate:"required"`
		Comment     *string   `json:"comment" validate:"required,lte=100"`
	}

	return func(context echo.Context) error {
		adUpdateRequest := new(AdUpdateRequest)
		if err := parser.ParseRequest(context, adUpdateRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		ad_ := &models.Ad{
			Id:          *adUpdateRequest.Id,
			LocDep:      *adUpdateRequest.LocDep,
			LocArr:      *adUpdateRequest.LocArr,
			DateTimeArr: *adUpdateRequest.DateTimeArr,
			Item:        *adUpdateRequest.Item,
			MinPrice:    *adUpdateRequest.MinPrice,
			Comment:     *adUpdateRequest.Comment,
		}

		return responser.Respond(context, adDelivery.adUsecase.Update(ad_))
	}
}

func (adDelivery *AdDelivery) HandlerAdsSearch() echo.HandlerFunc {
	type AdsSearchRequest struct {
		LocDep      *string   `query:"loc_dep" validate:"omitempty,gte=2,lte=100"`
		LocArr      *string   `query:"loc_arr" validate:"omitempty,gte=2,lte=100"`
		DateTimeArr *DateTime `query:"date_time_arr" validate:"omitempty"`
		MaxPrice    *uint32   `query:"max_price" validate:"omitempty"`
	}

	return func(context echo.Context) error {
		adsSearchRequest := new(AdsSearchRequest)
		if err := parser.ParseRequest(context, adsSearchRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		adsSearch := &models.AdsSearch{
			LocDep:      parser.GetOrDefault(adsSearchRequest.LocDep, "").(string),
			LocArr:      parser.GetOrDefault(adsSearchRequest.LocArr, "").(string),
			DateTimeArr: parser.GetOrDefault(adsSearchRequest.DateTimeArr, DateTime{}).(DateTime),
			MaxPrice:    parser.GetOrDefault(adsSearchRequest.MaxPrice, 0).(uint32),
		}

		return responser.Respond(context, adDelivery.adUsecase.Search(adsSearch))
	}
}
