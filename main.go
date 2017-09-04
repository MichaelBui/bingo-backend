package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/michaelbui/bingo-backend/controllers"
	"io/ioutil"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	defineRoutes(e)

	logRoutesToFiles(e)

	e.Logger.Fatal(e.Start(":1323"))
}

func defineRoutes(e *echo.Echo) {

	e.File("/routes", "routes.json")

	// Normal routes
	e.GET("/", controllers.Default().Index)
	e.GET("/numbers", controllers.Number().List)
	e.POST("/users", controllers.User().Add)
	e.GET("/users/:email", controllers.User().Get)
	e.PATCH("/users/:email/numbers", controllers.User().UpdateNumbers)

	// Secured routes
	secured := e.Group("/secured")
	secured.Use(middleware.KeyAuth(func(key string, e echo.Context) (bool, error) {
		return key == "080917", nil
	}))
	secured.POST("/numbers", controllers.Number().Next)
	secured.POST("/reset", controllers.Admin().Reset)
	secured.POST("/activate", controllers.Admin().Activate)
}

func logRoutesToFiles(e *echo.Echo) error {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		return err
	}
	ioutil.WriteFile("routes.json", data, 0644)
	return nil
}
