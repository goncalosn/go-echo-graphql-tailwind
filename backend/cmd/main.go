package main

import (
	"github.com/goncalosn/go-echo-graphql-tailwind/internal/db"
	"github.com/goncalosn/go-echo-graphql-tailwind/pkg/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	handlers.GetHandlers(e)

	db.Open()

	e.Logger.Fatal(e.Start(":1323"))

}
