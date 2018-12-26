import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import { Provider } from 'react-redux';
import store from './store';
import './App.css';

import Notfound from './components/NotFound';

import LoginPage from './pages/LoginPage';
import ProfilePage from './pages/ProfilePage';
import ResetPasswordPage from './pages/password/ResetPasswordPage';
import PasswordPage from './pages/password/PasswordPage';

import LeaveRequestPage from './pages/Leave/LeaveRequestPage';
import AdminLeaveRequestPage from './pages/Leave/AdminLeaveRequestPage';
import LeaveEditPage from './pages/Leave/LeaveEditPage'; //todo:not finish

import AdminLandingPage from './pages/admin/AdminLandingPage';
import AdminRegisterPage from './pages/admin/AdminRegisterPage';
import AdminEditPage from './pages/admin/AdminEditPage';
import AdminEditBalancePage from './pages/admin/AdminEditBalancePage';
import AdminPendingPage from './pages/admin/AdminPendingPage';
import AdminApprovePage from './pages/admin/AdminApprovePage';
import AdminRejectPage from './pages/admin/AdminRejectPage';
import AdminPasswordPage from './pages/admin/AdminPasswordPage';

import DirectorLandingPage from './pages/director/DirectorLandingPage';
import DirectorPendingPage from './pages/director/DirectorPendingPage';
import DirectorApprovePage from './pages/director/DirectorApprovePage';
import DirectorRejectPage from './pages/director/DirectorRejectPage';

import SupervisorLandingPage from './pages/supervisor/SupervisorLandingPage';
import SupervisorPendingPage from './pages/supervisor/SupervisorPendingPage';
import SupervisorApprovePage from './pages/supervisor/SupervisorApprovePage';
import SupervisorRejectPage from './pages/supervisor/SupervisorRejectPage';

import EmployeeLandingPage from './pages/employee/EmployeeLandingPage';
import EmployeePendingPage from './pages/employee/EmployeePendingPage';
import EmployeeApprovePage from './pages/employee/EmployeeApprovePage';
import EmployeeRejectPage from './pages/employee/EmployeeRejectPage';

class App extends Component {
  render() {
    return (
      <Provider store={store}>
        <div className="App">
          <BrowserRouter>
            <Switch>
              <Route exact path="/" component={LoginPage} />
              <Route exact path="/login" component={LoginPage} />
              <Route exact path="/reset-password" component={ResetPasswordPage} />

              {/* admin */}
              <Route exact path="/admin" component={AdminLandingPage} />
              <Route exact path="/admin/register-user" component={AdminRegisterPage} />            
              <Route exact path={`/admin/edit-user/:id`} component={AdminEditPage}/>
              <Route exact path={`/admin/edit-balance/:id`} component={AdminEditBalancePage}/>
              
              <Route exact path="/admin/list-pending-request" component={AdminPendingPage} />
              <Route exact path="/admin/list-approve-request" component={AdminApprovePage} />
              <Route exact path="/admin/list-reject-request" component={AdminRejectPage} />
              <Route path={`/admin/edit-password`} component={AdminPasswordPage}/>
                           
              <Route exact path="/employee" component={EmployeeLandingPage} />

              <Route exact path="/profile" component={ProfilePage} />
              <Route path={`/profile/:id`} component={PasswordPage}/>

              <Route exact path="/employee/leave-request" component={LeaveRequestPage} />
              <Route exact path="/admin/leave-request" component={AdminLeaveRequestPage} />
              <Route exact path={`/editrequest/:id`} component={LeaveEditPage}/>
              <Route exact path="/employee/list-pending-request" component={EmployeePendingPage} />
              <Route exact path="/employee/list-approve-request" component={EmployeeApprovePage} />
              <Route exact path="/employee/list-reject-request" component={EmployeeRejectPage} />

              <Route exact path="/supervisor" component={SupervisorLandingPage} />
              <Route exact path="/supervisor/list-pending-request" component={SupervisorPendingPage} />
              <Route exact path="/supervisor/list-approve-request" component={SupervisorApprovePage} />
              <Route exact path="/supervisor/list-reject-request" component={SupervisorRejectPage} />

              <Route exact path="/director" component={DirectorLandingPage} />
              <Route exact path="/director/list-pending-request" component={DirectorPendingPage} />
              <Route exact path="/director/list-approve-request" component={DirectorApprovePage} />
              <Route exact path="/director/list-reject-request" component={DirectorRejectPage} />
              
              <Route component={Notfound} />
            </Switch>
          </BrowserRouter>
        </div>
      </Provider>
    );
  }
}
export default App;
