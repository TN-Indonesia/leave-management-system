package meal

import (
	dbInterfaceMeal "server/models/db/interfaces/meal_request"
	dbInterfaceUser "server/models/db/interfaces/user"

	dbLayerMeal "server/models/db/pgsql/meal_request"
	dbLayerUser "server/models/db/pgsql/user"
)

// constant var
var (
	DBMeal dbInterfaceMeal.IBaseMeal
	DBUser dbInterfaceUser.IBaseUser
)

func init() {
	DBMeal = new(dbLayerMeal.Meal)
	DBUser = new(dbLayerUser.User)
}
