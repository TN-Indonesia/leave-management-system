import {	
	combineReducers,		
} from 'redux';

import loginReducer from './loginReducer';
import resetPasswordReducer from './resetPasswordReducer';

import adminReducer from './adminReducer';
import registerReducer from './registerReducer';
import editUserReducer from './editUserReducer';

import leaveRequestReducer from './leaveRequestReducer';
import editRequestReducer from './editRequestReducer';

import fetchDirectorReducer from './fetchDirectorReducer';
import fetchSupervisorReducer from './fetchSupervisorReducer';
import fetchEmployeeReducer from './fetchEmployeeReducer';

import fetchUserSummaryReducer from './fetchUserSummaryReducer';
import profileReducer from './profileReducer';
import passwordReducer from './passwordReducer';

import fetchTypeLeaveReducer from './fetchTypeLeaveReducer';
import AddSupervisorReducer from './AddSupervisorReducer';


const appStore = combineReducers({
	loginReducer,
	resetPasswordReducer,

	adminReducer,
	registerReducer,
	editUserReducer,

	leaveRequestReducer,
	editRequestReducer,

	fetchDirectorReducer,
	fetchSupervisorReducer,
	fetchEmployeeReducer,

	fetchUserSummaryReducer,
	profileReducer,
	passwordReducer,

	fetchTypeLeaveReducer,
	AddSupervisorReducer,
});

export default appStore