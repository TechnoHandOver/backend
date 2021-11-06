package delivery

import (
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

func (userDelivery *UserDelivery) Configure(echo_ *echo.Echo, _ *middlewares.Manager) {
	echo_.POST("/api/user", userDelivery.HandlerUserLogin())
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
