package logic

type (
	// InfoMailSupervisorApproval ...
	InfoMailSupervisorApproval struct {
		SupervisorName string
		RequestorName  string

		RequestDate        string
		RequestTime        string
		OtherRequestorName []GetOtherRequestorNameByEmployeeNumbers
		Amount             string
		BriefDescription   string
		Notes              string

		ApproveLink string
		RejectLink  string
	}

	// InfoMailApprovedRequest ...
	InfoMailApprovedRequest struct {
		HRName         string
		RequestorName  string
		SupervisorName string

		RequestDate        string
		RequestTime        string
		OtherRequestorName []GetOtherRequestorNameByEmployeeNumbers
		Amount             string
		BriefDescription   string
		Notes              string
	}

	// InfoMailSubmittedRequest ...
	InfoMailSubmittedRequest struct {
		RequestorName  string
		SupervisorName string

		RequestDate        string
		RequestTime        string
		OtherRequestorName []GetOtherRequestorNameByEmployeeNumbers
		Amount             string
		BriefDescription   string
		Notes              string
		Reason             string
		ActionDate         string
		Status             string
	}
)
