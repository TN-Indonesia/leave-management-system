package meal

import (
	"errors"
	"server/helpers"
	"server/helpers/constant"
	"strconv"
	"time"

	structAPI "server/structs/api"
	structDB "server/structs/db"
	structLogic "server/structs/logic"

	"github.com/astaxie/beego/orm"
)

//Meal ...
type Meal struct{}

//GetMealRequestByID ...
func (m *Meal) GetMealRequestByID(id int64) (
	result structLogic.GetMealRequest,
	err error,
) {
	var dbMeal structDB.MealRequest

	o := orm.NewOrm()
	o.Using(dbMeal.OrmDBName())
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @GetMealRequestByID - modelsPGMeal", errQB)
		return result, errQB
	}

	qb.Select(dbMeal.TableName() + ".*").
		From(dbMeal.TableName()).
		Where(dbMeal.TableName() + `.id = ? `)
	qb.Limit(1)
	sql := qb.String()
	// log.Println(sql)

	errRaw := o.Raw(sql, id).QueryRow(&result)
	if errRaw != nil {
		helpers.CheckErr("Failed query select @GetMealRequestByID - modelsPGMeal", errRaw)
		return result, errors.New("ID not exist")
	}

	return result, err
}

// PostMealRequest ...
func (m *Meal) PostMealRequest(
	RequestorID int64,
	OtherRequestorID string,
	SupervisorID int64,
	Amount float64,
	BriefDescription string,
	Notes string,
	ReceiptUploadPath string,
	RequestDate time.Time,
	Status string,
	CreatedAt time.Time,
	UpdatedAt time.Time,
) (ID int64, err error) {
	var dbMeal structDB.MealRequest

	o := orm.NewOrm()
	o.Using(dbMeal.OrmDBName())
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @PostMealRequest - modelsPGMeal", errQB)
		err = errQB
		return
	}

	qb.InsertInto(
		dbMeal.TableName(),
		"requestor_id",
		"other_requestor_id",
		"supervisor_id",
		"amount",
		"brief_description",
		"notes",
		"receipt_upload_path",
		"request_date",
		"status",
		"created_at",
		"updated_at",
	).Values(`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?`)
	sql := qb.String() + ` returning id`

	values := []interface{}{
		RequestorID,
		OtherRequestorID,
		SupervisorID,
		Amount,
		BriefDescription,
		Notes,
		ReceiptUploadPath,
		RequestDate.Format(constant.FormatDateTime),
		Status,
		CreatedAt.Format(constant.FormatDateTime),
		UpdatedAt.Format(constant.FormatDateTime),
	}

	err = o.Raw(sql, values).QueryRow(&ID)
	if err != nil {
		helpers.CheckErr("Error insert meal request @PostMealRequest - modelsPGMeal", err)
		err = errors.New("Insert meal request failed")
		return
	}

	return
}

// ApprovalMealRequest ...
func (m *Meal) ApprovalMealRequest(req structAPI.PostApprovalMealRequest) error {
	var dbMeal structDB.MealRequest

	o := orm.NewOrm()
	o.Using(dbMeal.OrmDBName())
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		helpers.CheckErr("Query builder failed @ApprovalMealRequest - modelsPGMeal", errQB)
		return errQB
	}

	qb.Update(dbMeal.TableName()).Set(
		"status = ?",
		"reject_reason = ?",
		"action_date = ?",
		"updated_at = ?",
	).Where("id = ?")
	sql := qb.String()

	var status string
	if req.Result {
		status = constant.StatusApproved
	} else {
		status = constant.StatusRejected
	}

	values := []interface{}{
		status,
		req.RejectReason,
		time.Now().Format(constant.FormatDateTime),
		time.Now().Format(constant.FormatDateTime),
		req.ID,
	}

	_, errUpdate := o.Raw(sql, values).Exec()
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

// GetMealRequestForEmployeeInquiry ...
func (m *Meal) GetMealRequestForEmployeeInquiry(ID int64, status string) (resDB []structDB.MealRequest, err error) {
	var dbMeal structDB.MealRequest

	o := orm.NewOrm()
	o.Using(dbMeal.OrmDBName())
	qb, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		helpers.CheckErr("Query builder failed @GetMealRequestForEmployeeInquiry - modelsPGMeal", err)
		return
	}

	qb.Select("*").From(dbMeal.TableName()).Where("requestor_id = ?").And("status = ?").OrderBy("created_at desc")
	sql := qb.String()

	_, err = o.Raw(sql, ID, status).QueryRows(&resDB)
	if err != nil {
		helpers.CheckErr("Failed to get meal request for requestor_id = "+strconv.FormatInt(ID, 10)+" and status = "+status, err)
		return
	}

	return
}

// GetMealRequestForSupervisorInquiry ...
func (m *Meal) GetMealRequestForSupervisorInquiry(ID int64, status string) (resDB []structDB.MealRequest, err error) {
	var dbMeal structDB.MealRequest

	o := orm.NewOrm()
	o.Using(dbMeal.OrmDBName())
	qb, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		helpers.CheckErr("Query builder failed @GetMealRequestForSupervisorInquiry - modelsPGMeal", err)
		return
	}

	qb.Select("*").From(dbMeal.TableName()).Where("supervisor_id = ?").And("status = ?").OrderBy("created_at desc")
	sql := qb.String()

	_, err = o.Raw(sql, ID, status).QueryRows(&resDB)
	if err != nil {
		helpers.CheckErr("Failed to get meal request for supervisor_id = "+strconv.FormatInt(ID, 10)+" and status = "+status, err)
		return
	}

	return
}

// GetMealRequestForAdminInquiry ...
func (m *Meal) GetMealRequestForAdminInquiry() (resDB []structDB.MealRequest, err error) {
	var dbMeal structDB.MealRequest

	o := orm.NewOrm()
	o.Using(dbMeal.OrmDBName())
	qb, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		helpers.CheckErr("Query builder failed @GetMealRequestForAdminInquiry - modelsPGMeal", err)
		return
	}

	qb.Select("*").From(dbMeal.TableName()).Where("status = ?").OrderBy("action_date desc")
	sql := qb.String()

	_, err = o.Raw(sql, constant.StatusApproved).QueryRows(&resDB)
	if err != nil {
		helpers.CheckErr("Failed to get meal request for status = "+constant.StatusApproved, err)
		return
	}

	return
}
