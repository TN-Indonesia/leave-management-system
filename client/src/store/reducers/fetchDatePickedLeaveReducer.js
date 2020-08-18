import {
    FETCH_PICKED_DATE_LEAVE
} from "../Actions/types"

let pickedLeaveState = {
    pickedLeave: [],
}


export default function fetchDatePickedLeaveReducer(state = pickedLeaveState, action) {
    switch (action.type) {
        case FETCH_PICKED_DATE_LEAVE:
            return {
                ...action.payload
            }
        default:
            return state
    }
}
