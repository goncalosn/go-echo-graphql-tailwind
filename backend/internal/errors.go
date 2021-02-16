package internal

type M map[string]interface{}

// Define our errors:
var InternalError = M{"message": "internal error"}
var FoodNotFound = M{"message": "food not found"}
var DrinkNotFound = M{"message": "drink not found"}
var DesertNotFound = M{"message": "desert not found"}
