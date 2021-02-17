package models

import (
	"github.com/kamva/mgm/v3"
)

type Order struct {
	mgm.IDField `bson:",inline"`

	Foods   []map[string]interface{} `json:foods bson:"foods"`
	Drinks  []map[string]interface{} `json:drinks bson:"drinks"`
	Deserts []map[string]interface{} `json:desert bson:"deserts"`
	Price   float64                  `json:price bson:"price"`
}

func NewOrder(foods []map[string]interface{}, drinks []map[string]interface{}, deserts []map[string]interface{}, price float64) *Order {
	return &Order{
		Foods:   foods,
		Drinks:  drinks,
		Deserts: deserts,
		Price:   price,
	}
}
