package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/labstack/echo/v4/middleware"
    "github.com/proapi/go-embed-spa/frontend"
)

func main() {
	app := echo.New()
	app.GET("/hello.json", handleHello)
	app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: frontend.BuildHTTPFS(),
		HTML5:      true,
	}))
	log.Fatal(app.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}

func handleHello(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "hello from the echo server",
	})
}
