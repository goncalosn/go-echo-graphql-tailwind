package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//GetFoods returns every food
func GetFoods(c echo.Context) error {
	// Food ID from path `foods`
	comidas := [2]string{"Batatas fritas", "hamburger"}

	return c.JSON(http.StatusOK, &comidas)
}
