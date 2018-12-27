package holiday

import (
	structLogic "server/structs/logic"
)

// IBasePublicHoliday

type IBasePublicHoliday interface {
	//Get All Public Holiday
	GetAllPublicHoliday() (
		holidays []structLogic.GetAllPublicHoliday,
		err error,
	)
}
