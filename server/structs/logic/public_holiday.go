package logic

// GetAllPublicHoliday
type GetAllPublicHoliday struct {
	ID          int64  `json:"id" orm:"column(id);pk"`
	DateStart   string `json:"date_start" orm:"column(date_start)"`
	DateEnd     string `json:"date_end" orm:"column(date_end)"`
	Description string `json:"description" orm:"column(description)"`
}
