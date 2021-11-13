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
