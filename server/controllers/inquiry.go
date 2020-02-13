package controllers

import (
	"errors"
	"server/helpers"
	logicMeal "server/models/logic/meal"
	structAPI "server/structs/api"
	"strconv"
	"strings"
)

// GetMealRequestForEmployeeInquiry ...
func (c *MealController) GetMealRequestForEmployeeInquiry() {
	var resp structAPI.RespData

	idStr := c.Ctx.Input.Param(":id")
	ID, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @GetMealRequestForEmployeeInquiry - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	status := strings.ToUpper(c.Ctx.Input.Param(":status"))

	resGet, errGetReject := logicMeal.GetMealRequestForEmployeeInquiry(ID, status)
	if errGetReject != nil {
		resp.Error = errGetReject.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @GetMealRequestForEmployeeInquiry - controller", err)
	}
}

// GetMealRequestForSupervisorInquiry ...
func (c *MealController) GetMealRequestForSupervisorInquiry() {
	var resp structAPI.RespData

	idStr := c.Ctx.Input.Param(":id")
	ID, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @GetMealRequestForSupervisorInquiry - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	status := strings.ToUpper(c.Ctx.Input.Param(":status"))

	resGet, errGetReject := logicMeal.GetMealRequestForSupervisorInquiry(ID, status)
	if errGetReject != nil {
		resp.Error = errGetReject.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @GetMealRequestForSupervisorInquiry - controller", err)
	}
}

// GetMealRequestForAdminInquiry ...
func (c *MealController) GetMealRequestForAdminInquiry() {
	var resp structAPI.RespData

	resGet, errGetReject := logicMeal.GetMealRequestForAdminInquiry()
	if errGetReject != nil {
		resp.Error = errGetReject.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @GetMealRequestForAdminInquiry - controller", err)
	}
}
