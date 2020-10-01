import {
	FETCH_USER,
	DELETE_USER,
	FETCH_LEAVE_PENDING,
	FETCH_LEAVE_APPROVE,
	FETCH_LEAVE_REJECT,
	CANCEL_LEAVE_REQUEST,
	FETCH_LEAVE_BALANCES,
	EDIT_BALANCES
} from "../Actions/types"

let adminState = {
	loading: true,
	users: [],
	leaves: [],
	balances: [
		{ type_name: "Annual Leave", leave_remaining: 12 },
		{ type_name: "Sick Leave", leave_remaining: 30 },
		{ type_name: "Marriage Leave", leave_remaining: 2 },
		{ type_name: "Other Leave", leave_remaining: 99 }
	],
}

export default function adminReducer(state = adminState, action) {
	switch (action.type) {
		case FETCH_USER:
			return {
				...action.payload
			}
		case DELETE_USER:
			return {
				...action.payload
			}
		case FETCH_LEAVE_PENDING:
			return {
				...action.payload
			}
		case FETCH_LEAVE_APPROVE:
			return {
				...action.payload
			}
		case FETCH_LEAVE_REJECT:
			return {
				...action.payload
			}
		case CANCEL_LEAVE_REQUEST:
			return {
				...action.payload
			}
		case FETCH_LEAVE_BALANCES:
			return {
				...action.payload
			}
		case EDIT_BALANCES:
			return {
				...action.payload
			}
		default:
			return state
	}
}