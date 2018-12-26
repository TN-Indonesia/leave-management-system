package db

// PublicHoliday ...
type PublicHoliday struct {
	ID          int64  `json:"id" orm:"column(id);pk"`
	DateStart   int64  `json:"date_start" orm:"column(date_start)"`
	DateEnd     int64  `json:"date_end" orm:"column(date_end)"`
	Description string `json:"description" orm:"column(description)"`
}

// TableName ...
func (u *PublicHoliday) TableName() string {
	return "public_holiday"
}
