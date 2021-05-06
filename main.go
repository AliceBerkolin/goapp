package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	serverPort = 8080
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	server := echo.New()
	serverAddress := fmt.Sprintf(":%d", serverPort)
	server.POST("/user", AddUser)
	err := server.Start(serverAddress)
	if err != nil {
		panic(err)
	}
}

func AddUser(ctx echo.Context) error {
	reqBody := ctx.Request().Body
	if reqBody == nil {
		return ctx.String(http.StatusNotAcceptable, "Body is required")
	}
	user := &User{}
	err := ctx.Bind(user)
	if err != nil {
		msg := fmt.Sprintf("Failed to parse json. Error: %s", err)
		return ctx.String(http.StatusInternalServerError, msg)
	}
	fmt.Printf("Got request to add user:\n%+v\n", user)
	return ctx.String(http.StatusOK, "User was successfully added")
}
