package api

import (
	"log"
	"net/http"

	. "github.com/goncalosn/go-echo-graphql-tailwind/internal/models"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

//GetFoods returns every food
func GetFoods(c echo.Context) error {
	results := []Food{}

	err := mgm.Coll(&Food{}).SimpleFind(&results, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, results)
}

//GetFoodByID returns the food with the id from the parameter
func GetFoodByID(c echo.Context) error {
	//Get document's collection
	food := &Food{}
	coll := mgm.Coll(food)

	// Find and decode doc to the food model.
	err := coll.FindByID(c.Param("id"), food)

	if err != nil {
		if food == nil {
			return c.JSON(http.StatusNotFound, nil)
		}

		log.Println(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, food)

}
