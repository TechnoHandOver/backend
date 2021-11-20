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
	echo_.GET("/api/users/routes-tmp/list", userDelivery.HandlerRouteTmpList())
	echo_.POST("/api/users/routes-perm", userDelivery.HandlerRoutePermCreate(),
		middlewaresManager.AuthMiddleware.CheckAuth())
	echo_.GET("/api/users/routes-perm/:id", userDelivery.HandlerRoutePermGet())
	echo_.PUT("/api/users/routes-perm/:id", userDelivery.HandlerRoutePermUpdate(),
		middlewaresManager.AuthMiddleware.CheckAuth())
	echo_.DELETE("/api/users/routes-perm/:id", userDelivery.HandlerRoutePermDelete(),
		middlewaresManager.AuthMiddleware.CheckAuth())
	echo_.GET("/api/users/routes-perm/list", userDelivery.HandlerRoutePermList())
}

func (userDelivery *UserDelivery) HandlerRouteTmpCreate() echo.HandlerFunc {
	type RouteTmpCreateRequest struct {
		LocDep      *string   `json:"locDep" validate:"required,gte=2,lte=100"`
		LocArr      *string   `json:"locArr" validate:"required,gte=2,lte=100"`
		MinPrice    *uint32   `json:"minPrice" validate:"required"`
		DateTimeDep *DateTime `json:"dateTimeDep" validate:"required"`
		DateTimeArr *DateTime `json:"dateTimeArr" validate:"required"`
	}

	return func(context echo.Context) error {
		routeTmpCreateRequest := new(RouteTmpCreateRequest)
		if err := parser.ParseRequest(context, routeTmpCreateRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		routeTmp := &models.RouteTmp{
			UserAuthorVkId: context.Get(consts.EchoContextKeyUserVkId).(uint32),
			LocDep:         *routeTmpCreateRequest.LocDep,
			LocArr:         *routeTmpCreateRequest.LocArr,
			MinPrice:       *routeTmpCreateRequest.MinPrice,
			DateTimeDep:    *routeTmpCreateRequest.DateTimeDep,
			DateTimeArr:    *routeTmpCreateRequest.DateTimeArr,
		}

		return responser.Respond(context, userDelivery.userUsecase.CreateRouteTmp(routeTmp))
	}
}

func (userDelivery *UserDelivery) HandlerRouteTmpGet() echo.HandlerFunc {
	type RouteTmpGetRequest struct {
		Id *uint32 `param:"id" validate:"required"`
	}

	return func(context echo.Context) error {
		routeTmpGetRequest := new(RouteTmpGetRequest)
		if err := parser.ParseRequest(context, routeTmpGetRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		id := *routeTmpGetRequest.Id

		return responser.Respond(context, userDelivery.userUsecase.GetRouteTmp(id))
	}
}

func (userDelivery *UserDelivery) HandlerRouteTmpUpdate() echo.HandlerFunc {
	type RouteTmpUpdateRequest struct {
		Id          *uint32   `param:"id" validate:"required"`
		LocDep      *string   `json:"locDep" validate:"required,gte=2,lte=100"`
		LocArr      *string   `json:"locArr" validate:"required,gte=2,lte=100"`
		MinPrice    *uint32   `json:"minPrice" validate:"required"`
		DateTimeDep *DateTime `json:"dateTimeDep" validate:"required"`
		DateTimeArr *DateTime `json:"dateTimeArr" validate:"required"`
	}

	return func(context echo.Context) error {
		routeTmpUpdateRequest := new(RouteTmpUpdateRequest)
		if err := parser.ParseRequest(context, routeTmpUpdateRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		routeTmp := &models.RouteTmp{
			Id:             *routeTmpUpdateRequest.Id,
			UserAuthorVkId: context.Get(consts.EchoContextKeyUserVkId).(uint32),
			LocDep:         *routeTmpUpdateRequest.LocDep,
			LocArr:         *routeTmpUpdateRequest.LocArr,
			MinPrice:       *routeTmpUpdateRequest.MinPrice,
			DateTimeDep:    *routeTmpUpdateRequest.DateTimeDep,
			DateTimeArr:    *routeTmpUpdateRequest.DateTimeArr,
		}

		return responser.Respond(context, userDelivery.userUsecase.UpdateRouteTmp(routeTmp))
	}
}

func (userDelivery *UserDelivery) HandlerRouteTmpDelete() echo.HandlerFunc {
	type RouteTmpDeleteRequest struct {
		Id *uint32 `param:"id" validate:"required"`
	}

	return func(context echo.Context) error {
		routeTmpDeleteRequest := new(RouteTmpDeleteRequest)
		if err := parser.ParseRequest(context, routeTmpDeleteRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		id := *routeTmpDeleteRequest.Id
		userVkId := context.Get(consts.EchoContextKeyUserVkId).(uint32)

		return responser.Respond(context, userDelivery.userUsecase.DeleteRouteTmp(userVkId, id))
	}
}

func (userDelivery *UserDelivery) HandlerRouteTmpList() echo.HandlerFunc {
	return func(context echo.Context) error {
		return responser.Respond(context, userDelivery.userUsecase.ListRouteTmp())
	}
}

func (userDelivery *UserDelivery) HandlerRoutePermCreate() echo.HandlerFunc {
	type RoutePermCreateRequest struct {
		LocDep    *string    `json:"locDep" validate:"required,gte=2,lte=100"`
		LocArr    *string    `json:"locArr" validate:"required,gte=2,lte=100"`
		MinPrice  *uint32    `json:"minPrice" validate:"required"`
		EvenWeek  *bool      `json:"evenWeek" validate:"required"`
		OddWeek   *bool      `json:"oddWeek" validate:"required"`
		DayOfWeek *DayOfWeek `json:"dayOfWeek" validate:"required,eq=Mon|eq=Tue|eq=Wed|eq=Thu|eq=Fri|eq=Sat|eq=Sun"`
		TimeDep   *Time      `json:"timeDep" validate:"required"`
		TimeArr   *Time      `json:"timeArr" validate:"required"`
	}

	return func(context echo.Context) error {
		routePermCreateRequest := new(RoutePermCreateRequest)
		if err := parser.ParseRequest(context, routePermCreateRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		routePerm := &models.RoutePerm{
			UserAuthorVkId: context.Get(consts.EchoContextKeyUserVkId).(uint32),
			LocDep:         *routePermCreateRequest.LocDep,
			LocArr:         *routePermCreateRequest.LocArr,
			MinPrice:       parser.GetOrDefault(routePermCreateRequest.MinPrice, 0).(uint32),
			EvenWeek:       parser.GetOrDefault(routePermCreateRequest.EvenWeek, true).(bool),
			OddWeek:        parser.GetOrDefault(routePermCreateRequest.OddWeek, true).(bool),
			DayOfWeek:      *routePermCreateRequest.DayOfWeek,
			TimeDep:        *routePermCreateRequest.TimeDep,
			TimeArr:        *routePermCreateRequest.TimeArr,
		}

		return responser.Respond(context, userDelivery.userUsecase.CreateRoutePerm(routePerm))
	}
}

func (userDelivery *UserDelivery) HandlerRoutePermGet() echo.HandlerFunc {
	type RoutePermGetRequest struct {
		Id *uint32 `param:"id" validate:"required"`
	}

	return func(context echo.Context) error {
		routePermGetRequest := new(RoutePermGetRequest)
		if err := parser.ParseRequest(context, routePermGetRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		id := *routePermGetRequest.Id

		return responser.Respond(context, userDelivery.userUsecase.GetRoutePerm(id))
	}
}

func (userDelivery *UserDelivery) HandlerRoutePermUpdate() echo.HandlerFunc {
	type RoutePermUpdateRequest struct {
		Id        *uint32    `param:"id" validate:"required"`
		LocDep    *string    `json:"locDep" validate:"required,gte=2,lte=100"`
		LocArr    *string    `json:"locArr" validate:"required,gte=2,lte=100"`
		MinPrice  *uint32    `json:"minPrice" validate:"required"`
		EvenWeek  *bool      `json:"evenWeek" validate:"required"`
		OddWeek   *bool      `json:"oddWeek" validate:"required"`
		DayOfWeek *DayOfWeek `json:"dayOfWeek" validate:"required,eq=Mon|eq=Tue|eq=Wed|eq=Thu|eq=Fri|eq=Sat|eq=Sun"`
		TimeDep   *Time      `json:"timeDep" validate:"required"`
		TimeArr   *Time      `json:"timeArr" validate:"required"`
	}

	return func(context echo.Context) error {
		routePermUpdateRequest := new(RoutePermUpdateRequest)
		if err := parser.ParseRequest(context, routePermUpdateRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		routePerm := &models.RoutePerm{
			Id:             *routePermUpdateRequest.Id,
			UserAuthorVkId: context.Get(consts.EchoContextKeyUserVkId).(uint32),
			LocDep:         *routePermUpdateRequest.LocDep,
			LocArr:         *routePermUpdateRequest.LocArr,
			MinPrice:       *routePermUpdateRequest.MinPrice,
			EvenWeek:       *routePermUpdateRequest.EvenWeek,
			OddWeek:        *routePermUpdateRequest.OddWeek,
			DayOfWeek:      *routePermUpdateRequest.DayOfWeek,
			TimeDep:        *routePermUpdateRequest.TimeDep,
			TimeArr:        *routePermUpdateRequest.TimeArr,
		}

		return responser.Respond(context, userDelivery.userUsecase.UpdateRoutePerm(routePerm))
	}
}

func (userDelivery *UserDelivery) HandlerRoutePermDelete() echo.HandlerFunc {
	type RoutePermDeleteRequest struct {
		Id *uint32 `param:"id" validate:"required"`
	}

	return func(context echo.Context) error {
		routePermDeleteRequest := new(RoutePermDeleteRequest)
		if err := parser.ParseRequest(context, routePermDeleteRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		id := *routePermDeleteRequest.Id
		userVkId := context.Get(consts.EchoContextKeyUserVkId).(uint32)

		return responser.Respond(context, userDelivery.userUsecase.DeleteRoutePerm(userVkId, id))
	}
}

func (userDelivery *UserDelivery) HandlerRoutePermList() echo.HandlerFunc {
	return func(context echo.Context) error {
		return responser.Respond(context, userDelivery.userUsecase.ListRoutePerm())
	}
}
