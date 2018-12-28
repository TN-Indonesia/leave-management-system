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
				let publicHoliday = seperateDate(body)
				let payload = {
					publicHoliday
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

function pad(num, size) {
	let s = num + "";
	while (s.length < size) s = "0" + s;
	return s;
}

function seperateDate(arrPublicHoliday) {
	let publicHolidayArr = []
	arrPublicHoliday && arrPublicHoliday.map((val, idx) => {
		let dateStart = val.date_start.split("-").reverse().join("-")
		let dateEnd = val.date_end.split("-").reverse().join("-")
		publicHolidayArr.push(dateStart)

		if (dateStart !== dateEnd) { //if date_start and date_end is different
			let dateStartInt = parseInt(dateStart.substring(dateStart.length - 2, dateStart.length))
			let dateEndInt = parseInt(dateEnd.substring(dateEnd.length - 2, dateEnd.length))
			if (dateStartInt > dateEndInt) { //if date_start is higher than dateEnd, holiday dates is within 2 month
				if (dateStartInt + 1 === dateEndInt) {
					publicHolidayArr.push(dateEnd)
				}
			} else { //both date in the same month
				let suffixMonthYear = dateStart.substring(0, dateStart.length - 2)
				for (let j = dateStartInt + 1; j <= dateEndInt; j++) {
					publicHolidayArr.push(suffixMonthYear + pad(j, 2))
				}
			}
		}
		return val
	})

	return publicHolidayArr
}