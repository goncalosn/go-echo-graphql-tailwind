package api

import (
	"encoding/json"
	"log"
	"net/http"

	. "github.com/goncalosn/go-echo-graphql-tailwind/internal/models"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

//GetOrders returns every order
func GetOrders(c echo.Context) error {
	results := []Order{}

	err := mgm.Coll(&Order{}).SimpleFind(&results, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, results)
}

//GetOrderByID returns the order with the id from the parameter
func GetOrderByID(c echo.Context) error {
	//Get orders's collection
	order := &Order{}
	coll := mgm.Coll(order)

	// Find and decode doc to the order model.
	err := coll.FindByID(c.Param("id"), order)

	if err != nil {
		if order == nil {
			return c.JSON(http.StatusNotFound, nil)
		}

		log.Println(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, order)
}

//PostOrder creates an order for the data recieved
func PostOrder(c echo.Context) error {
	var price float64 = 0

	decodeJSON := json.NewDecoder(c.Request().Body)

	order := &Order{}
	err := decodeJSON.Decode(order)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	for i := 0; i < len(order.Foods); i++ {
		id := order.Foods[i]["id"]
		quantity := order.Foods[i]["quantity"]

		//Get foods collection
		food := &Food{}
		coll := mgm.Coll(food)

		// Find and decode doc to the food model.
		err = coll.FindByID(id, food)

		if err != nil {
			if food == nil {
				return c.JSON(http.StatusNotFound, nil)
			}

			log.Println(err)
			return c.JSON(http.StatusInternalServerError, nil)
		}

		price += food.Price * quantity.(float64)
	}

	for i := 0; i < len(order.Drinks); i++ {
		id := order.Drinks[i]["id"]
		quantity := order.Drinks[i]["quantity"]

		//Get drinks collection
		drink := &Drink{}
		coll := mgm.Coll(drink)

		// Find and decode doc to the drink model.
		err = coll.FindByID(id, drink)

		if err != nil {
			if drink == nil {
				return c.JSON(http.StatusNotFound, nil)
			}

			log.Println(err)
			return c.JSON(http.StatusInternalServerError, nil)
		}

		price += drink.Price * quantity.(float64)
	}

	for i := 0; i < len(order.Deserts); i++ {
		id := order.Deserts[i]["id"]
		quantity := order.Deserts[i]["quantity"]

		//Get deserts collection
		desert := &Desert{}
		coll := mgm.Coll(desert)

		// Find and decode doc to the desert model.
		err = coll.FindByID(id, desert)

		if err != nil {
			if desert == nil {
				return c.JSON(http.StatusNotFound, nil)
			}

			log.Println(err)
			return c.JSON(http.StatusInternalServerError, nil)
		}

		price += desert.Price * quantity.(float64)
	}

	//criar novo pedido de comida/bebida/sobremesa
	newOrder := NewOrder(order.Foods, order.Drinks, order.Deserts, price)

	// Make sure pass the model by reference.
	err = mgm.Coll(order).Create(newOrder)

	return c.JSON(http.StatusOK, nil)
}

//DeleteOrder returns 201 if order deleted else returns 500
func DeleteOrder(c echo.Context) error {
	//Get order's collection
	order := &Order{}
	coll := mgm.Coll(order)

	// Find and decode doc to the order model.
	err := coll.FindByID(c.Param("id"), order)

	if err != nil {
		if order == nil {
			return c.JSON(http.StatusNotFound, nil)
		}

		log.Println(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, order)
}
