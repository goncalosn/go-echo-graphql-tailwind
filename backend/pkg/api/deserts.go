package api

import (
	"log"
	"net/http"

	. "github.com/goncalosn/go-echo-graphql-tailwind/internal/models"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

//GetDeserts returns every desert
func GetDeserts(c echo.Context) error {
	results := []Desert{}

	err := mgm.Coll(&Desert{}).SimpleFind(&results, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, results)
}

//GetDesertByID returns the desert with the id from the parameter
func GetDesertByID(c echo.Context) error {
	//Get document's collection
	desert := &Desert{}
	coll := mgm.Coll(desert)

	// Find and decode doc to the desert model.
	err := coll.FindByID(c.Param("id"), desert)

	if err != nil {
		if desert == nil {
			return c.JSON(http.StatusNotFound, nil)
		}

		log.Println(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, desert)

}
