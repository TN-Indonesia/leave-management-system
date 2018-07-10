export function handleEdit(payload) {
	return (dispatch) => {
		dispatch({
			type: 'UPDATE_NEW_PASSWORD',
			payload: payload
		})
	}
}

export function updateNewPassword(savePassword,pusher ) {
	const employeeNumber = localStorage.getItem('id')
	console.log("saved", savePassword)
	return (dispatch) => {
		fetch(`http://localhost:8080/api/user/update/${employeeNumber}`, {
				method: 'PUT',
				body: JSON.stringify(savePassword)
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {				
				if (body === "update password success") {
					alert("update password success")
					pusher('/profile')
				} else if (error === "wrong old password") {
					alert("wrong old password")
				} else if (error === "wrong confirm password") {
					alert("wrong confirm password")
				}

			}).catch(err => {
				console.log(err)
			})
	}
}