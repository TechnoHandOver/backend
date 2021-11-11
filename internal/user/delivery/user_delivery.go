package delivery

import (
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/middlewares"
	"github.com/TechnoHandOver/backend/internal/models"
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
	echo_.POST("/api/users", userDelivery.HandlerUserLogin())
	echo_.POST("/api/users/routes-tmp", userDelivery.HandlerRouteTmpCreate(),
		middlewaresManager.AuthMiddleware.CheckAuth())
}

func (userDelivery *UserDelivery) HandlerUserLogin() echo.HandlerFunc {
	return func(context echo.Context) error {
		user_ := new(models.User)
		if err := context.Bind(user_); err != nil {
			return responser.Respond(context, response.NewErrorResponse(err))
		}

		return responser.Respond(context, userDelivery.userUsecase.Login(user_))
	}
}

func (userDelivery *UserDelivery) HandlerRouteTmpCreate() echo.HandlerFunc {
	return func(context echo.Context) error {
		routeTmp := new(models.RouteTmp)
		if err := context.Bind(routeTmp); err != nil {
			return responser.Respond(context, response.NewErrorResponse(err))
		}

		routeTmp.UserAuthorVkId = context.Get(consts.EchoContextKeyUserVkId).(uint32)

		return responser.Respond(context, userDelivery.userUsecase.CreateRouteTmp(routeTmp))
	}
}
