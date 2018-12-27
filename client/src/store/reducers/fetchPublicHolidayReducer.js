import {
    FETCH_PUBLIC_HOLIDAY
} from "../Actions/types"

let publicHolidayState = {
    publicHoliday: [],
}

export default function fetchPublicHolidayReducer(state = publicHolidayState, action) {
    switch (action.type) {
        case FETCH_PUBLIC_HOLIDAY:
            return {
                ...action.payload
            }
        default:
            return state
    }
}