package main

import (
	"github.com/labstack/echo"
	"github.com/michaelbui/bingo-backend/controllers"
	"encoding/json"
	"io/ioutil"
)

func main () {
	e := echo.New()

	defindRoutes(e)

	logRoutesToFiles(e)

	e.Logger.Fatal(e.Start(":1323"))
}

func defindRoutes(e *echo.Echo) {
	e.File("/", "routes.json")
	e.GET("/reset", controllers.Default().Reset)
}

func logRoutesToFiles(e *echo.Echo) error {
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		return err
	}
	ioutil.WriteFile("routes.json", data, 0644)
	return nil
}