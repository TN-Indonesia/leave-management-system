package api

type (
	// ReqLeave ...
	ReqLeave struct {
		TypeLeaveID    int64    `json:"type_leave_id"`
		Reason         string   `json:"reason"`
		DateFrom       string   `json:"date_from"`
		DateTo         string   `json:"date_to"`
		HalfDates      []string `json:"half_dates"`
		Total          float64  `json:"total"`
		BackOn         string   `json:"back_on"`
		ContactAddress string   `json:"contact_address"`
		ContactNumber  string   `json:"contact_number"`
		Notes          string   `json:"notes"`
	}

	// ReqLeaveAdmin ...
	ReqLeaveAdmin struct {
		EmployeeNumber int64    `json:"employee_number"`
		TypeLeaveID    int64    `json:"type_leave_id"`
		Reason         string   `json:"reason"`
		DateFrom       string   `json:"date_from"`
		DateTo         string   `json:"date_to"`
		HalfDates      []string `json:"half_dates"`
		Total          float64  `json:"total"`
		BackOn         string   `json:"back_on"`
		ContactAddress string   `json:"contact_address"`
		ContactNumber  string   `json:"contact_number"`
		Notes          string   `json:"notes"`
	}

	// UpdateLeaveRequest ...
	UpdateLeaveRequest struct {
		TypeLeaveID    int64    `json:"type_leave_id"`
		Reason         string   `json:"reason"`
		DateFrom       string   `json:"date_from"`
		DateTo         string   `json:"date_to"`
		HalfDates      []string `json:"half_dates"`
		Total          float64  `json:"total" orm:"column(total)"`
		BackOn         string   `json:"back_on"`
		ContactAddress string   `json:"contact_address"`
		ContactNumber  string   `json:"contact_number"`
		Notes          string   `json:"notes"`
	}

	// UpdateLeaveBalance ...
	UpdateLeaveBalance struct {
		TypeID         string  `json:"type_id"`
		TypeName       string  `json:"type_name"`
		LeaveRemaining float64 `json:"leave_remaining"`
	}

	// BackToWork ...
	BackToWork struct {
		ToDate    string
		HalfDates string
	}

	// BackToWorkDate ...
	BackToWorkDate struct {
		// Status       string `json:"status" orm:"column(status)"`
		BackToWorkDate string `json:"back_to_work_date" orm:"column(reject_reason)"`
	}
)
