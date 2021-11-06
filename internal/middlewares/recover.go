package middlewares

import (
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/TechnoHandOver/backend/internal/tools/responser"
	"github.com/labstack/echo/v4"
	"log"
)

type RecoverMiddleware struct{}

func NewRecoverMiddleware() *RecoverMiddleware {
	return &RecoverMiddleware{}
}

func (recoverMiddleware *RecoverMiddleware) Recover() echo.MiddlewareFunc {
	return recoverMiddleware.recover
	//return middleware.Recover() //TODO: сделать кастомный recover, который будет сохранять в логи всё как нужно и отправлять клиенту нужные ответы без доп. полей...
}

func (recoverMiddleware *RecoverMiddleware) recover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		defer func() error {
			if err := recover(); err != nil {
				log.Println(err)
				return responser.Respond(context, response.NewEmptyResponse(consts.InternalError))
			}
			return nil
		}()

		return next(context)
	}
}
