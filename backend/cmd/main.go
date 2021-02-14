package main

import (
	"github.com/goncalosn/go-echo-graphql-tailwind/pkg/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	handlers.GetHandlers(e)

	e.Logger.Fatal(e.Start(":1323"))
}
