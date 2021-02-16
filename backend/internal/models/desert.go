package models

import (
	"github.com/kamva/mgm/v3"
)

type Desert struct {
	mgm.DefaultModel `bson:",inline"`

	Name  string `bson:"name"`
	Price string `bson:"price"`
	Image string `bson:"image"`
}
