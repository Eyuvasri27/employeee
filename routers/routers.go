package routers

import (
	"goecho/controlers"

	"github.com/labstack/echo"
)

func Setemproute(e *echo.Echo) {
	e.POST("/employee", controlers.Postemployee)
	e.GET("/employee", controlers.Listemployee)
	e.GET("/employee/:id", controlers.Getemployeebyid)
	e.PUT("/employee/:id", controlers.Putemployee)
	e.DELETE("/employee/:id", controlers.Deleteemp)
}
