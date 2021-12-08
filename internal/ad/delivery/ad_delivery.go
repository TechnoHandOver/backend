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
	echo_.DELETE("/api/ads/:id", adDelivery.HandlerAdDelete(), middlewaresManager.AuthMiddleware.CheckAuth())
	echo_.GET("/api/ads/list", adDelivery.HandlerAdsList(), middlewaresManager.AuthMiddleware.CheckAuth())
	echo_.GET("/api/ads/search", adDelivery.HandlerAdsSearch(), middlewaresManager.AuthMiddleware.CheckAuth())
	echo_.POST("/api/ads/:id/execution", adDelivery.HandlerAdExecutionCreate(), middlewaresManager.AuthMiddleware.CheckAuth())
	echo_.DELETE("/api/ads/:id/execution", adDelivery.HandlerAdExecutionDelete(), middlewaresManager.AuthMiddleware.CheckAuth())
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
			UserAuthorId: context.Get(consts.EchoContextKeyUserId).(uint32),
			LocDep:       *adCreateRequest.LocDep,
			LocArr:       *adCreateRequest.LocArr,
			DateTimeArr:  *adCreateRequest.DateTimeArr,
			Item:         *adCreateRequest.Item,
			MinPrice:     *adCreateRequest.MinPrice,
			Comment:      *adCreateRequest.Comment,
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
			Id:           *adUpdateRequest.Id,
			UserAuthorId: context.Get(consts.EchoContextKeyUserId).(uint32),
			LocDep:       *adUpdateRequest.LocDep,
			LocArr:       *adUpdateRequest.LocArr,
			DateTimeArr:  *adUpdateRequest.DateTimeArr,
			Item:         *adUpdateRequest.Item,
			MinPrice:     *adUpdateRequest.MinPrice,
			Comment:      *adUpdateRequest.Comment,
		}

		return responser.Respond(context, adDelivery.adUsecase.Update(ad_))
	}
}

func (adDelivery *AdDelivery) HandlerAdDelete() echo.HandlerFunc {
	type AdDeleteRequest struct {
		Id *uint32 `param:"id" validate:"required"`
	}

	return func(context echo.Context) error {
		adDeleteRequest := new(AdDeleteRequest)
		if err := parser.ParseRequest(context, adDeleteRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		id := *adDeleteRequest.Id
		userId := context.Get(consts.EchoContextKeyUserId).(uint32)

		return responser.Respond(context, adDelivery.adUsecase.Delete(userId, id))
	}
}

func (adDelivery *AdDelivery) HandlerAdsList() echo.HandlerFunc {
	return func(context echo.Context) error {
		userId := context.Get(consts.EchoContextKeyUserId).(uint32)
		adsSearch := &models.AdsSearch{
			UserAuthorId: &userId,
		}

		return responser.Respond(context, adDelivery.adUsecase.Search(adsSearch))
	}
}

func (adDelivery *AdDelivery) HandlerAdsSearch() echo.HandlerFunc {
	type AdsSearchRequest struct {
		LocDep      *string                `query:"loc_dep" validate:"omitempty,lte=100"`
		LocArr      *string                `query:"loc_arr" validate:"omitempty,lte=100"`
		DateTimeArr *DateTime              `query:"date_time_arr" validate:"omitempty"` //TODO: а точно нужен поиск по дате? как он будет работать?
		MaxPrice    *uint32                `query:"max_price" validate:"omitempty"`
		Order       *models.AdsSearchOrder `query:"order" validate:"omitempty"`
	}

	return func(context echo.Context) error {
		adsSearchRequest := new(AdsSearchRequest)
		if err := parser.ParseRequest(context, adsSearchRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		userId := context.Get(consts.EchoContextKeyUserId).(uint32)
		adsSearch := &models.AdsSearch{
			NotUserAuthorId: &userId,
			LocDep:          adsSearchRequest.LocDep,
			LocArr:          adsSearchRequest.LocArr,
			DateTimeArr:     adsSearchRequest.DateTimeArr,
			MaxPrice:        adsSearchRequest.MaxPrice,
			Order:           adsSearchRequest.Order,
		}

		return responser.Respond(context, adDelivery.adUsecase.Search(adsSearch))
	}
}

func (adDelivery *AdDelivery) HandlerAdExecutionCreate() echo.HandlerFunc {
	type AdExecutionRequest struct {
		Id *uint32 `param:"id" validate:"required"`
	}

	return func(context echo.Context) error {
		adExecutionRequest := new(AdExecutionRequest)
		if err := parser.ParseRequest(context, adExecutionRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		adId := *adExecutionRequest.Id
		userId := context.Get(consts.EchoContextKeyUserId).(uint32)

		return responser.Respond(context, adDelivery.adUsecase.SetAdUserExecutor(userId, adId))
	}
}

func (adDelivery *AdDelivery) HandlerAdExecutionDelete() echo.HandlerFunc {
	type AdExecutionRequest struct {
		Id *uint32 `param:"id" validate:"required"`
	}

	return func(context echo.Context) error {
		adExecutionRequest := new(AdExecutionRequest)
		if err := parser.ParseRequest(context, adExecutionRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		adId := *adExecutionRequest.Id
		userId := context.Get(consts.EchoContextKeyUserId).(uint32)

		return responser.Respond(context, adDelivery.adUsecase.UnsetAdUserExecutor(userId, adId))
	}
}
