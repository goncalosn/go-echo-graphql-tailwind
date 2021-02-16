package main

import (
	"log"

	. "github.com/goncalosn/go-echo-graphql-tailwind/pkg/handlers"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := mgm.SetDefaultConfig(nil, "menu", options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	e := echo.New()

	GetHandlers(e)

	e.Logger.Fatal(e.Start(":1323"))

}
