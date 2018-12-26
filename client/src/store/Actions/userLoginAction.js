import {
	ROOT_API,
	FETCH_USER_LOGIN
} from "./types"

function fetchUserLogin(payload) {
	return {
		type: FETCH_USER_LOGIN,
		payload: payload
	}
}


export function userLoginFetchData() {
	const id = localStorage.getItem('id')
	return (dispatch) => {
		fetch(`${ROOT_API}/api/admin/user/${id}`, {
				method: 'GET',
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				let payload = {
					loading: false,
					user: body,
				}				
				dispatch(fetchUserLogin(payload))
				if (error !== null) {
					console.error("error not null @userLoginFetchData: ", error)
				}
			})
			.catch(error => {
				console.error("error @userLoginFetchData: ", error)
			})
	}
}