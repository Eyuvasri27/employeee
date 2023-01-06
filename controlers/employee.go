package controlers

import (
	"fmt"
	"goecho/models"
	"goecho/storages"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func CreateEmpUser(c echo.Context) error {

	el := new(models.EmpUserDetail)

	if err := c.Bind(el); err != nil {
		return err
	}
	cur := storages.Getcursor()
	cur.CreateEmpUser(el)
	return c.JSON(http.StatusCreated, el)
}

func EmpLogin(c echo.Context) error {
	l := new(models.Logincheck) //structure
	if err := c.Bind(l); err != nil {
		return err
	}
	cur := storages.Getcursor()

	result, err := cur.GetToken(l.Name, l.Password)
	fmt.Println("result:", result)
	fmt.Println("err:", err)
	if result != nil {

		expirationtime := time.Now().Add(time.Hour * 24)

		claims := &jwt.StandardClaims{
			Issuer:    "eyuvasri",
			Subject:   "dummy",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expirationtime.Unix(),
		}
		fmt.Println(claims)

		var sign = "random"

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		signed := []byte(sign)
		ss, err := token.SignedString(signed)
		if err != nil {
			fmt.Println(err)
			fmt.Println("tokenstring:  ", ss)

			return c.JSON(http.StatusOK, ss)
		}
		return c.JSON(http.StatusOK, ss)
	}
	return c.JSON(http.StatusBadRequest, "loginfailed")
}

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
