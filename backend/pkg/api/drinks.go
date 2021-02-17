package api

import (
	"log"
	"net/http"

	. "github.com/goncalosn/go-echo-graphql-tailwind/internal/models"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

//GetDrinks returns every drink
func GetDrinks(c echo.Context) error {
	results := []Drink{}

	err := mgm.Coll(&Drink{}).SimpleFind(&results, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, results)
}

//GetDrinkByID returns the drink with the id from the parameter
func GetDrinkByID(c echo.Context) error {
	//Get document's collection
	drink := &Drink{}
	coll := mgm.Coll(drink)

	// Find and decode doc to the drink model.
	err := coll.FindByID(c.Param("id"), drink)

	if err != nil {
		if drink == nil {
			return c.JSON(http.StatusNotFound, nil)
		}

		log.Println(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, drink)

}
