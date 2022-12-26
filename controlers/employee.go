package controlers

import (
	"fmt"
	"goecho/models"
	"goecho/storages"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func Postemployee(c echo.Context) error {
	u := new(models.Employee) //structure
	if err := c.Bind(u); err != nil {
		return err
	}
	cur := storages.Getcursor()
	cur.Createemp(u)
	return c.JSON(http.StatusCreated, u)
	//return c.String(http.StatusOK, "ok")
}

// list
func Listemployee(c echo.Context) error {

	cur := storages.Getcursor()
	result, _ := cur.Listemployee()

	return c.JSON(http.StatusCreated, result)

}

func Getemployeebyid(c echo.Context) error {

	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	cur := storages.Getcursor()
	result, _ := cur.Geteempbyid(Id)

	return c.JSON(http.StatusOK, result)

}

func Putemployee(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	cur := storages.Getcursor()
	fmt.Println("id", id)
	u := new(models.Employee)
	if err := c.Bind(u); err != nil {
		return err
	}
	fmt.Println("u", u)
	cur.Updateeempbyid(Id, u)
	return c.JSON(http.StatusCreated, u)
}

func Deleteemp(c echo.Context) error {
	id := c.Param("id")
	Id, _ := strconv.Atoi(id)
	cur := storages.Getcursor()
	cur.Deleteemp(Id)

	return c.JSON(http.StatusOK, "Deleted")

}
