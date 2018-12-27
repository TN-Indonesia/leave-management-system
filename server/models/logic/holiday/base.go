package holiday

import (
	dbInterfaceHoliday "server/models/db/interfaces/holiday"
	dbLayerHoliday "server/models/db/pgsql/holiday"
)

// constant var
var (
	DBHoliday dbInterfaceHoliday.IBasePublicHoliday
)

func init() {
	DBHoliday = new(dbLayerHoliday.PublicHoliday)
}
