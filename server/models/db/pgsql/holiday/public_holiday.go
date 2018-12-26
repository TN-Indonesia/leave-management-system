package holiday

import (
	"errors"
	"server/helpers"
	structDB "server/structs/db"
	structLogic "server/structs/logic"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PublicHoliday struct{}

// Get All Public Holiday
func (c *PublicHoliday) GetAllPublicHoliday() (holidays []structLogic.GetAllPublicHoliday, err error) {

	var publicHoliday structDB.PublicHoliday

	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetAllPublicHoliday", errQB)
		return holidays, errQB
	}

	qb.Select(
		publicHoliday.TableName()+".id",
		publicHoliday.TableName()+".date_start",
		publicHoliday.TableName()+".date_end",
		publicHoliday.TableName()+".description",
	).From(
		publicHoliday.TableName(),
	)

	sql := qb.String()

	count, errRaw := o.Raw(sql).QueryRows(&holidays)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetAllPublicHoliday", errRaw)
		return holidays, errors.New("Error get all public holiday")
	}
	beego.Debug("Total public holiday=", count)

	return holidays, errRaw
}
