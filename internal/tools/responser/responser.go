package responser

import (
	"github.com/TechnoHandOver/backend/internal/consts"
	"github.com/TechnoHandOver/backend/internal/tools/response"
	"github.com/labstack/echo/v4"
	"log"
)

type DataResponse struct {
	Data interface{} `json:"data"`
}

func Respond(context echo.Context, response_ *response.Response) error {
	if response_.Error != nil {
		log.Println(response_.Error)
	}

	if response_.Data == nil {
		return context.NoContent(consts.StatusCodes[response_.Code])
	}
	return context.JSON(consts.StatusCodes[response_.Code], DataResponse{
		Data: response_.Data,
	}) //TODO: возможно, нужно кастомно отдавать ввиду UTF-8...
}
