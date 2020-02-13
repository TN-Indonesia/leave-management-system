package logic

// GetOtherRequestorNameByEmployeeNumbers ...
type GetOtherRequestorNameByEmployeeNumbers struct {
	Name string `json:"name" orm:"column(name)"`
}
