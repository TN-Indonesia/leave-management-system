import {
    FETCH_USER_LOGIN
} from "../Actions/types"

let userLoginState = {
    user: {},
    loading: true,
}

export default function fetchUserLoginReducer(state = userLoginState, action) {
    switch (action.type) {
        case FETCH_USER_LOGIN:
            return {
                ...action.payload
            }
        default:
            return state
    }
}