package responser

import (
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/labstack/echo/v4"
)

func Respond(context echo.Context, response *response.Response) error {
	if response.JSONObject.Data == nil {
		return context.NoContent(response.Code)
	}
	return context.JSON(response.Code, response.JSONObject)
}
