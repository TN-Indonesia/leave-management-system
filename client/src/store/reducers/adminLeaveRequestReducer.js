import {
	CREATE_LEAVE_REQUEST,
	CLEAR_FIELD
} from "../Actions/types"

const leaveState = {
	employee_number: null,
	type_leave_id: null,
	reason: '',
	date_from: '',
	date_to: '',
	half_dates: [],
	total: 0,
	back_on: null,
	contact_address: '',
	contact_number: '',
	notes: '',
}

export default function adminLeaveRequestReducer(state = leaveState, action) {
	switch (action.type) {
		case CREATE_LEAVE_REQUEST:
			return {
				...action.payload,
			}
		case CLEAR_FIELD:
			return {
				...leaveState
			}
		default:
			return state
	}
}