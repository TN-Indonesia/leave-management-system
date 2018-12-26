package holiday

import (
	"server/helpers"

	structLogic "server/structs/logic"
)

// GetPendingRequest ...
func GetAllPublicHoliday() ([]structLogic.GetAllPublicHoliday, error) {
	respGet, errGet := DBHoliday.GetAllPublicHoliday()
	if errGet != nil {
		helpers.CheckErr("Error get pending request @Get All Public Holiday - logicHoliday", errGet)
	}

	return respGet, errGet
}
