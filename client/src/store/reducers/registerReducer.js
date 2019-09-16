import {
	REGISTER_USER,
	CLEAR_FIELD
} from "../Actions/types"

const registerState = {
	employee_number: '',
	name: '',
	gender: '',
	position: '',
	start_working_date: '',
	mobile_phone: '',
	email: '',
	password: '',
	role: '',
	employee_number: null
}

export default function registerReducer(state = registerState, action) {
	switch (action.type) {
		case REGISTER_USER:
			return {
				...action.payload
			}
		case CLEAR_FIELD:
			return {
				...registerState
			}
		default:
			return state
	}
}