package controllers

import (
	"server/helpers"
	logicHoliday "server/models/logic/holiday"
	structAPI "server/structs/api"

	"github.com/astaxie/beego"
)

// Public Holiday Controller
type PublicHolidayController struct {
	beego.Controller
}

// GetAllPublicHoliday
func (c *PublicHolidayController) GetAllPublicHoliday() {
	var resp structAPI.RespData

	resGet, errGetHol := logicHoliday.GetAllPublicHoliday()
	if errGetHol != nil {
		resp.Error = errGetHol.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @GetAllPublicHoliday - controller", err)
	}
}
