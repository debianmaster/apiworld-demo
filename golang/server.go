package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Props struct {
	Name     string `json:"name" xml:"name" form:"name" query:"name"`
	FavColor string `json:"favColor" xml:"favColor" form:"favColor" query:"favColor"`
	Hostname string `json:"hostname" xml:"hostname" form:"hostname" query:"hostname"`
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	props := new(Props)
	props.Name = "golang"
	props.FavColor = "yellow"
	hostname, err := os.Hostname()

	if err != nil {
		panic(err)
	}
	props.Hostname = hostname
	e.GET("/golang/props", func(c echo.Context) error {
		return c.JSON(http.StatusOK, props)
	})
	e.Logger.Fatal(e.Start(":3000"))
}
