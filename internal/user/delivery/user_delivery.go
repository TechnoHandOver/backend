package delivery

import (
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/middlewares"
	"github.com/TechnoHandOver/backend/internal/models"
	. "github.com/TechnoHandOver/backend/internal/models/timestamps"
	"github.com/TechnoHandOver/backend/internal/tools/parser"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/TechnoHandOver/backend/internal/tools/responser"
	"github.com/TechnoHandOver/backend/internal/user"
	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUsecase user.Usecase
}

func NewUserDelivery(userUsecase user.Usecase) *UserDelivery {
	return &UserDelivery{
		userUsecase: userUsecase,
	}
}

func (userDelivery *UserDelivery) Configure(echo_ *echo.Echo, middlewaresManager *middlewares.Manager) {
	echo_.POST("/api/users/routes-tmp", userDelivery.HandlerRouteTmpCreate(),
		middlewaresManager.AuthMiddleware.CheckAuth())
	echo_.GET("/api/users/routes-tmp/:id", userDelivery.HandlerRouteTmpGet())
	echo_.PUT("/api/users/routes-tmp/:id", userDelivery.HandlerRouteTmpUpdate(),
		middlewaresManager.AuthMiddleware.CheckAuth())
	echo_.DELETE("/api/users/routes-tmp/:id", userDelivery.HandlerRouteTmpDelete(),
		middlewaresManager.AuthMiddleware.CheckAuth())
}

func (userDelivery *UserDelivery) HandlerRouteTmpCreate() echo.HandlerFunc {
	type RouteTmpCreateRequest struct {
		LocDep      string   `json:"locDep" validate:"required,gte=2,lte=100"`
		LocArr      string   `json:"locArr" validate:"required,gte=2,lte=100"`
		MinPrice    uint32   `json:"minPrice,omitempty" validate:"omitempty"`
		DateTimeDep DateTime `json:"dateTimeDep" validate:"required"`
		DateTimeArr DateTime `json:"dateTimeArr" validate:"required"`
	}

	return func(context echo.Context) error {
		routeTmpCreateRequest := new(RouteTmpCreateRequest)
		if err := parser.ParseRequest(context, routeTmpCreateRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		routeTmp := &models.RouteTmp{
			UserAuthorVkId: context.Get(consts.EchoContextKeyUserVkId).(uint32),
			LocDep:         routeTmpCreateRequest.LocDep,
			LocArr:         routeTmpCreateRequest.LocArr,
			MinPrice:       routeTmpCreateRequest.MinPrice,
			DateTimeDep:    routeTmpCreateRequest.DateTimeDep,
			DateTimeArr:    routeTmpCreateRequest.DateTimeArr,
		}

		return responser.Respond(context, userDelivery.userUsecase.CreateRouteTmp(routeTmp))
	}
}

func (userDelivery *UserDelivery) HandlerRouteTmpGet() echo.HandlerFunc {
	type RouteTmpGetRequest struct {
		Id uint32 `param:"id" validate:"required"`
	}

	return func(context echo.Context) error {
		routeTmpGetRequest := new(RouteTmpGetRequest)
		if err := parser.ParseRequest(context, routeTmpGetRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		return responser.Respond(context, userDelivery.userUsecase.GetRouteTmp(routeTmpGetRequest.Id))
	}
}

func (userDelivery *UserDelivery) HandlerRouteTmpUpdate() echo.HandlerFunc {
	type RouteTmpUpdateRequest struct {
		Id          uint32   `param:"id" validate:"required"`
		LocDep      string   `json:"locDep,omitempty" validate:"omitempty,gte=2,lte=100"`
		LocArr      string   `json:"locArr,omitempty" validate:"omitempty,gte=2,lte=100"`
		MinPrice    uint32   `json:"minPrice,omitempty" validate:"omitempty"`
		DateTimeDep DateTime `json:"dateTimeDep,omitempty" validate:"omitempty"`
		DateTimeArr DateTime `json:"dateTimeArr,omitempty" validate:"omitempty"`
	}

	return func(context echo.Context) error {
		routeTmpUpdateRequest := new(RouteTmpUpdateRequest)
		if err := parser.ParseRequest(context, routeTmpUpdateRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		routeTmp := &models.RouteTmp{
			Id:             routeTmpUpdateRequest.Id,
			UserAuthorVkId: context.Get(consts.EchoContextKeyUserVkId).(uint32),
			LocDep:         routeTmpUpdateRequest.LocDep,
			LocArr:         routeTmpUpdateRequest.LocArr,
			MinPrice:       routeTmpUpdateRequest.MinPrice,
			DateTimeDep:    routeTmpUpdateRequest.DateTimeDep,
			DateTimeArr:    routeTmpUpdateRequest.DateTimeArr,
		}

		return responser.Respond(context, userDelivery.userUsecase.UpdateRouteTmp(routeTmp))
	}
}

func (userDelivery *UserDelivery) HandlerRouteTmpDelete() echo.HandlerFunc {
	type RouteTmpDeleteRequest struct {
		Id uint32 `param:"id" validate:"required"`
	}

	return func(context echo.Context) error {
		routeTmpDeleteRequest := new(RouteTmpDeleteRequest)
		if err := parser.ParseRequest(context, routeTmpDeleteRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		userVkId := context.Get(consts.EchoContextKeyUserVkId).(uint32)

		return responser.Respond(context, userDelivery.userUsecase.DeleteRouteTmp(userVkId, routeTmpDeleteRequest.Id))
	}
}
