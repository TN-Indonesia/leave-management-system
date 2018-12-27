import {
	ROOT_API,
	FETCH_PUBLIC_HOLIDAY
} from "./types"

function publicHolidayFetch(payload) {
	return {
		type: FETCH_PUBLIC_HOLIDAY,
		payload: payload
	}
}

export function publicHolidayFetchData() {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/holidays/public`, {
				method: 'GET',
			})
			.then((resp) => resp.json())  
			.then(({
				body,
				error
			}) => {
				let payload = {
					publicHoliday: body
				}
				dispatch(publicHolidayFetch(payload))

				if (error !== null) {
					console.error("error not null @publicHolidayFetchData: ", error)
				}  
			})
			.catch(error => {
				console.error("error @publicHolidayFetchData: ", error)
			})
	}
}