package logic

// GetAllPublicHoliday
type GetAllPublicHoliday struct {
	ID          int64  `json:"id" orm:"column(id);pk"`
	DateStart   string `json:"date_start" orm:"column(date_start)"`
	DateEnd     string `json:"date_end" orm:"column(date_end)"`
	Description string `json:"description" orm:"column(description)"`
}

// GetPickedDateLeave ...
type GetPickedDateLeave struct {
	ID       int64  `json:"id" orm:"column(id);pk"`
	DateFrom string `json:"date_from" orm:"column(date_from)"`
	DateTo   string `json:"date_to" orm:"column(date_to)"`
}

type GetPickedDateLeaveConverted struct {
	ID        int64  `json:"id" orm:"column(id);pk"`
	DateStart string `json:"date_start" orm:"column(date_start)"`
	DateEnd   string `json:"date_end" orm:"column(date_end)"`
}
