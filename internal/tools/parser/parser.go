package parser

import "github.com/labstack/echo/v4"

func ParseRequest(context echo.Context, object interface{}) error {
	if err := context.Bind(object); err != nil {
		return err
	}

	return context.Validate(object)
}
