package delivery

import (
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/middlewares"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/session"
	"github.com/TechnoHandOver/backend/internal/tools/parser"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/TechnoHandOver/backend/internal/tools/responser"
	"github.com/TechnoHandOver/backend/internal/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SessionDelivery struct {
	sessionUsecase session.Usecase
	userUsecase    user.Usecase
}

func NewSessionDelivery(sessionUsecase session.Usecase, userUsecase user.Usecase) *SessionDelivery {
	return &SessionDelivery{
		sessionUsecase: sessionUsecase,
		userUsecase:    userUsecase,
	}
}

func (sessionDelivery *SessionDelivery) Configure(echo_ *echo.Echo, _ *middlewares.Manager) {
	echo_.POST("/api/sessions", sessionDelivery.HandlerLogin())
}

func (sessionDelivery *SessionDelivery) HandlerLogin() echo.HandlerFunc {
	type LoginRequest struct {
		VkId   uint32 `json:"vkId" validate:"required"`
		Name   string `json:"name" validate:"required,gte=2,lte=100"`
		Avatar string `json:"avatar" validate:"required,url,lte=2000"`
	}

	return func(context echo.Context) error {
		userRequest := new(LoginRequest)
		if err := parser.ParseRequest(context, userRequest); err != nil {
			return responser.Respond(context, response.NewErrorResponse(consts.BadRequest, err))
		}

		user_ := &models.User{
			VkId:   userRequest.VkId,
			Name:   userRequest.Name,
			Avatar: userRequest.Avatar,
		}

		userResponse := sessionDelivery.userUsecase.Login(user_)
		if userResponse.Error != nil {
			return responser.Respond(context, userResponse)
		}

		sessionResponse := sessionDelivery.sessionUsecase.Create(user_.VkId)
		if sessionResponse.Code != consts.OK {
			return responser.Respond(context, sessionResponse)
		}

		session_ := sessionResponse.Data.(*models.Session)
		context.SetCookie(&http.Cookie{
			Name:  consts.EchoCookieAuthName,
			Value: session_.Id,
		})

		return responser.Respond(context, userResponse)
	}
}
