package handlers

import (
	"net/http"

	"github.com/goncalosn/go-echo-graphql-tailwind/pkg/api"
	"github.com/labstack/echo/v4"
)

func GetHandlers(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/foods", api.GetFoods)
	// 	e.GET("/food/:id", GetFoodById)
	// 	e.GET("/drinks", GetDrinks)
	// 	e.GET("/drinks/:id", GetDrinkById)
	// 	e.GET("/deserts", GetDeserts)
	// 	e.GET("/desert/:id", GetDesertById)
	// 	e.GET("/orders", GetOrders)
	// 	e.POST("/order/:food&:drink&:desert", GostOrder)
	// 	e.DELETE("/order/:id", GeleteOrder)
}
