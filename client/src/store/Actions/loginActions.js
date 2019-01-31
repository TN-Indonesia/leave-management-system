// import jwtDecode from 'jwt-decode';
import {
	ROOT_API,
	LOGIN_FORM_ONCHANGE,
	CLEAR_FIELD
} from "./types"

import {
	message
} from "antd";
import swal from 'sweetalert';

export function handleFormInput(payload) {
	return {
		type: LOGIN_FORM_ONCHANGE,
		payload: payload
	}
}

function clearField() {
	return {
		type: CLEAR_FIELD
	}
}

export function submitLogin(payload, pusher) {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/login`, {
				method: 'POST',
				body: JSON.stringify(payload)
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				if (error !== null) {
					swal({
						title: "Sorry",
						text: `${error}`,
						icon: "error",
						button: false,
					});
				} else {
					const token = body['Token']
					const id = body['ID']
					const role = body['Role']

					if (role === 'admin') {
						localStorage.setItem('token', token)
						localStorage.setItem('role', role)
						localStorage.setItem('id', id)
						pusher('/admin')
						dispatch(clearField())
						swal({
							title: "Welcome",
							text: "Login success!",
							icon: "success",
							button: false,
						});
					} else if (role === 'director') {
						localStorage.setItem('token', token)
						localStorage.setItem('role', role)
						localStorage.setItem('id', id)
						pusher('/director')
						dispatch(clearField())
						swal({
							title: "Welcome",
							text: "Login success!",
							icon: "success",
							button: false,
						});
					} else if (role === 'supervisor') {
						localStorage.setItem('token', token)
						localStorage.setItem('role', role)
						localStorage.setItem('id', id)
						pusher('/supervisor')
						dispatch(clearField())
						swal({
							title: "Welcome",
							text: "Login success!",
							icon: "success",
							button: false,
						});
					} else if (role === 'employee') {
						localStorage.setItem('token', token)
						localStorage.setItem('role', role)
						localStorage.setItem('id', id)
						pusher('/employee')
						dispatch(clearField())
						swal({
							title: "Welcome",
							text: "Login success!",
							icon: "success",
							button: false,
						});
					} else if (role !== 'admin' || role !== 'director' || role !== 'supervisor' || role !== 'employee') {
						pusher('/')
						dispatch(clearField())
						swal({
							title: "Sorry",
							text: `Login failed!`,
							icon: "error",
							button: false,
						});
					}
				}
			})
			.catch(error => {
				console.error("error @submitLogin: ", error)
			})
	}
}