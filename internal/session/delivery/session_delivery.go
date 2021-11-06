package delivery

import (
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/middlewares"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/session"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/TechnoHandOver/backend/internal/tools/responser"
	"github.com/TechnoHandOver/backend/internal/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SessionDelivery struct {
	sessionUsecase session.Usecase
	userUsecase user.Usecase
}

func NewSessionDelivery(sessionUsecase session.Usecase, userUsecase user.Usecase) *SessionDelivery {
	return &SessionDelivery{
		sessionUsecase: sessionUsecase,
		userUsecase: userUsecase,
	}
}

func (sessionDelivery *SessionDelivery) Configure(echo_ *echo.Echo, _ *middlewares.Manager) {
	echo_.POST("/api/session", sessionDelivery.HandlerLogin())
}

func (sessionDelivery *SessionDelivery) HandlerLogin() echo.HandlerFunc {
	return func(context echo.Context) error {
		user_ := new(models.User)
		if err := context.Bind(user_); err != nil {
			return responser.Respond(context, response.NewErrorResponse(err))
		}

		if err := responser.Respond(context, sessionDelivery.userUsecase.Login(user_)); err != nil {
			return err
		}

		response_ := sessionDelivery.sessionUsecase.Create(user_.VkId)
		if response_.Code != consts.OK {
			return responser.Respond(context, response_)
		}

		session_ := response_.Data.(*models.Session)
		cookie := &http.Cookie{
			Name: consts.EchoCookieAuthName,
			Value: session_.Id,
			Secure: true,
			SameSite: http.SameSiteNoneMode,
		}
		context.SetCookie(cookie)

		return nil
	}
}
