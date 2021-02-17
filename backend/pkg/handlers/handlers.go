package handlers

import (
	"net/http"

	. "github.com/goncalosn/go-echo-graphql-tailwind/pkg/api"
	"github.com/labstack/echo/v4"
)

func GetHandlers(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//foods
	e.GET("/foods", GetFoods)
	e.GET("/food/:id", GetFoodByID)
	//drinks
	e.GET("/drinks", GetDrinks)
	e.GET("/drinks/:id", GetDrinkByID)
	//deserts
	e.GET("/deserts", GetDeserts)
	e.GET("/desert/:id", GetDesertByID)
	//orders
	e.GET("/orders", GetOrders)
	e.GET("/order/:id", GetOrderByID)
	e.POST("/order/new", PostOrder)
	e.PUT("/order/:id", UpdateOrder)
	e.DELETE("/order/:id", DeleteOrder)
}
