import {
	ROOT_API,
	GET_SUPERVISORS
} from "./types"

function fetchSupervisors(payload) {
	return {
		type: GET_SUPERVISORS,
		payload: payload
	}
}

export function getSupervisors() {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/user/supervisor`, {
				method: 'GET',
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				const data = body.map(({ employee_number, name }) => ({ key:employee_number, id: employee_number, name: name }));
				let payload = {
					supervisor: data
				}
				dispatch(fetchSupervisors(payload))
				if (error !== null) {
					console.error("error not null @getSupervisors: ", error)
				}
			})
			.catch(error => {
				console.error("error @getSupervisors: ", error)
			})
	}
}