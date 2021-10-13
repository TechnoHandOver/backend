package delivery

import (
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/tools/responser"
	"github.com/TechnoHandOver/backend/internal/user/usecase"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type UserDelivery struct {
	userUsecase *usecase.UserUsecase
}

func NewUserDelivery(userUsecase *usecase.UserUsecase) *UserDelivery {
	return &UserDelivery{
		userUsecase: userUsecase,
	}
}

func (userDelivery *UserDelivery) Configure(echo_ *echo.Echo) {
	echo_.POST("/api/user", userDelivery.HandlerUserLogin())
}

func (userDelivery *UserDelivery) HandlerUserLogin() echo.HandlerFunc {
	return func(context echo.Context) error {
		user := new(models.User)
		if err := context.Bind(user); err != nil {
			log.Println(err)
			return context.NoContent(http.StatusInternalServerError)
		}

		return responser.Respond(context, userDelivery.userUsecase.Login(user))
	}
}
