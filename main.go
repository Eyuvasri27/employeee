package main

import (
	"goecho/common"
	"goecho/routers"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	common.Database()
	routers.Setemproute(e)
	e.Logger.Fatal(e.Start(":1323"))

}
