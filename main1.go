package main

import (
"github.com/labstack/echo"
"net/http"
"fmt"
"log"
"encoding/json"
"io/ioutil"
"github.com/go-playground/validator/v10"
)



type (
	User struct {
		ID    int    `json:"id"`
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
		Age   int    `json:"age" validate:"required"`
	}

	CustomValidator struct {
		validator *validator.Validate
	
	}
)

func (cv *CustomValidator) validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
	
	
func Welcome(c echo.Context) error{
	return c.String(http.StatusOK, "welcome to my website")
}


func AddUser(c echo.Context) error {
	user := User{}

	defer c.Request().Body.Close()

	b,err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
	   log.Printf("Failed reading the request body %s", err)
	   return c.String(http.StatusInternalServerError, "")
	}

	err=json.Unmarshal(b,&user)
	if err!=nil{
		log.Printf("Faildunmarshal in addCats: %s\n, err")
		return c.String (http.StatusInternalServerError, "")
    }

	log.Printf("Got request to add user:\n%+v\n", user)
	return c.String(http.StatusOK, "User was successfully added")
}


func main() {
	fmt.Println("welcome to the server")

	e:= echo.New()
	e.validator = &CustomValidator{validator:validator.New()}

	e.GET("/", Welcome)

	e.POST("/users", AddUser(c echo.Context)(err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, u)
	})
	e.Start(":8000")

}