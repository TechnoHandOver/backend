package middlewares

import (
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/models"
	"github.com/TechnoHandOver/backend/internal/session"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/TechnoHandOver/backend/internal/tools/responser"
	"github.com/TechnoHandOver/backend/internal/user"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	sessionUsecase session.Usecase
	userUsecase    user.Usecase
}

func NewAuthMiddleware(sessionUsecase session.Usecase, userUsecase user.Usecase) *AuthMiddleware {
	return &AuthMiddleware{
		sessionUsecase: sessionUsecase,
		userUsecase:    userUsecase,
	}
}

func (authMiddleware *AuthMiddleware) CheckAuth() echo.MiddlewareFunc {
	return authMiddleware.checkAuth
}

func (authMiddleware *AuthMiddleware) checkAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		cookie, err := context.Cookie(consts.EchoCookieAuthName)
		if err != nil {
			return responser.Respond(context, response.NewEmptyResponse(consts.Unauthorized))
		}

		response_ := authMiddleware.sessionUsecase.Get(cookie.Value)
		if response_.Code != consts.OK {
			if response_.Code == consts.NotFound {
				return responser.Respond(context, response.NewEmptyResponse(consts.Unauthorized))
			}

			return responser.Respond(context, response_)
		}

		session_ := response_.Data.(*models.Session)
		if response_ := authMiddleware.userUsecase.Get(session_.UserId); response_.Code != consts.OK {
			return responser.Respond(context, response_)
		}

		context.Set(consts.EchoContextKeyUserId, session_.UserId)

		return next(context)
	}
}
