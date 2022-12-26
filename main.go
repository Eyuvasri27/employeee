package main

import (
	"fmt"
	"goecho/common"
	"goecho/routers"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	common.Database()
	routers.Setemproute(e)
	fmt.Println("employee details")
	e.Logger.Fatal(e.Start(":1323"))

}
