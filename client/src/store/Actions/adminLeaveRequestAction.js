import {
	ROOT_API,
	CREATE_LEAVE_REQUEST,
	CLEAR_FIELD
} from "./types";

import {
	message
} from "antd";

export function formOnChange(payload) {
	return (dispach) => {
		dispach({
			type: CREATE_LEAVE_REQUEST,
			payload: payload
		})
	}
}

function clearField() {
	return {
		type: CLEAR_FIELD
	}
}

export function AdminSumbitLeave(payload, pusher) {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/admin/leave`, {
				method: 'POST',
				body: JSON.stringify(payload)
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				if (body !== null) {
					dispatch(clearField())
					pusher('/admin/list-pending-request')
					message.success('Create leave request success')
				} else if (error === "Type request malform") {
					message.error('Create failed, please check all field!')
				} else {
					message.error(error)
				}
			}).catch(error => {
				console.error("error @AdminSumbitLeave: ", error)
			})
	}
}