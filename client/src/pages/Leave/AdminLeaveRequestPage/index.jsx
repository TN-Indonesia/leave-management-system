import React, { Component } from "react";
import { bindActionCreators } from "redux";
import { connect } from "react-redux";
import update from "react-addons-update";
import moment from "moment-business-days";
import {
  formOnChange,
  AdminSumbitLeave
} from "../../../store/Actions/adminLeaveRequestAction";
import { typeLeaveFetchData } from "../../../store/Actions/typeLeaveAction";
import { userLoginFetchData } from "../../../store/Actions/userLoginAction";
import { publicHolidayFetchData } from "../../../store/Actions/publicHolidayAction";
import HeaderAdmin from "../../../pages/menu/HeaderAdmin";
import Footer from "../../../components/Footer";
import "./style.css";
import {
  Layout,
  Form,
  Input,
  Select,
  Button,
  Checkbox,
  DatePicker
} from "antd";
const { Content } = Layout;
const FormItem = Form.Item;
const { TextArea } = Input;
const Option = Select.Option;

class LeaveRequestPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      from: null,
      to: null,
      start: null,
      end: null,
      endOpen: false,
      contactID: "+62",
      halfDate: [],
      publicHolidayDates: [],
    };

    this.handleOnChange = this.handleOnChange.bind(this);
    this.handleChangeTypeOfLeave = this.handleChangeTypeOfLeave.bind(this);
    this.disabledDateBack = this.disabledDateBack.bind(this);
    this.handleOnChangeNumber = this.handleOnChangeNumber.bind(this);
    this.handleOnChangeID = this.handleOnChangeID.bind(this);
    this.onChangeIsHalfDay = this.onChangeIsHalfDay.bind(this);
    this.onChangeAddHalfDay = this.onChangeAddHalfDay.bind(this);
    this.handleOnChangeEmployeeNumber = this.handleOnChangeEmployeeNumber.bind(this);    
  }

  componentWillMount() {
    console.log(" ----------------- Form-Leave-Requessssssssssst ----------------- ");
    if (!localStorage.getItem("token")) {
      this.props.history.push("/");
    } else if (
      localStorage.getItem("role") !== "admin"
    ) {
      this.props.history.push("/");
    }

    this.props.typeLeaveFetchData();
    this.props.userLoginFetchData();
    this.props.publicHolidayFetchData();
  }

  componentDidMount(){
  }

  onChange = (field, value) => {
    this.setState({
      [field]: value
    });
  };

  handleOnChangeNumber = (value, field) => {
    this.onChange(field, Number(value));
    console.log("input=======>", value);
  };

  handleSubmitSupervisor = e => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        console.log("Received values of form: ", values);
      }
    });
    this.props.AdminSumbitLeave(this.props.leaveForm, url => {
      this.props.history.push(url);
    });
  };

  handleOnChange = e => {
    let newLeave = {
      ...this.props.leaveForm,
      [e.target.name]: e.target.value
    };
    this.props.formOnChange(newLeave);

    this.props.form.setFieldsValue({
      [e.target.name]: e.target.value
    });
  };

  handleChangeTypeOfLeave(value) {
    console.log("val", value);
    if (value === "11" || value === "44" || value === "55") {
      let typeLeave = {
        ...this.props.leaveForm,
        type_leave_id: Number(value),
        reason: ""
      };
      this.props.formOnChange(typeLeave);
    } else {
      let typeLeave = {
        ...this.props.leaveForm,
        type_leave_id: Number(value)
      };
      this.props.formOnChange(typeLeave);
    }
  }

  handleChangeSelect(value) {
    console.log("selected=======>", value);
  }

  disabledStartDate = startValue => {
    const endValue = this.state.to;
    if (!startValue || !endValue) {
      return false;
    }
    return startValue.valueOf() > endValue.valueOf();
  };

  disabledEndDate = endValue => {
    const startValue = this.state.from;
    if (!endValue || !startValue) {
      return false;
    }
    return endValue.valueOf() <= startValue.valueOf();
  };

  onStartChange = value => {
    if (value !== null) {
      const date = new Date(value._d),
        mnth = ("0" + (date.getMonth() + 1)).slice(-2),
        day = ("0" + date.getDate()).slice(-2);
      let newDate = [day, mnth, date.getFullYear()].join("-");
      let newStart = [mnth, day, date.getFullYear()].join("-");

      let dateFrom = {
        ...this.props.leaveForm,
        date_from: newDate
      };

      this.props.formOnChange(dateFrom);
      this.onChange("start", newStart);
    }
    this.onChange("from", value);
  };

  onEndChange = value => {
    if (value !== null) {
      const date = new Date(value._d),
        mnth = ("0" + (date.getMonth() + 1)).slice(-2),
        day = ("0" + date.getDate()).slice(-2);
      let newDate = [day, mnth, date.getFullYear()].join("-");
      let newEnd = [mnth, day, date.getFullYear()].join("-");
      let dateTo = {
        ...this.props.leaveForm,
        date_to: newDate
      };
      this.props.formOnChange(dateTo);
      this.onChange("end", newEnd);
    }
    this.onChange("to", value);
  };

  disabledDate(current) {
    //gimana ambil state ini disini: this.state.publicHolidayDates
    var publicHolidayDates = this.state.publicHolidayDates;
    // var publicHolidayDates = ['2018-12-29', '2018-12-30', '2018-12-31'];
    return current < moment().startOf("day") || publicHolidayDates.find( d => moment(d).format("DDMMYYYY") === moment(current._d).format("DDMMYYYY"))
  }

  disabledDateSick(current) {
    return (
      current &&
      current <
      moment()
        .subtract(7, "days")
        .startOf("day")
    );
  }

  disabledDateBack(current) {
    return this.state.to && this.state.to > current;
  }

  handleStartOpenChange = open => {
    if (!open) {
      this.setState({ endOpen: true });
    }
  };

  handleEndOpenChange = open => {
    this.setState({ endOpen: open });
  };

  getDates(start, end) {
    let startDate = new Date(start);
    let endDate = new Date(end);
    let dates = [];
    while (startDate <= endDate) {
      let weekDay = startDate.getDay();
      if (weekDay < 6 && weekDay > 0) {
        let month = startDate.getMonth() + 1;
        if (month <= 9) {
          month = "0" + month;
        }
        let day = startDate.getDate();
        if (day <= 9) {
          day = "0" + day;
        }
        dates.push(day + "-" + month + "-" + startDate.getFullYear());
      }
      startDate.setDate(startDate.getDate() + 1);
    }
    return dates;
  }

  onChangeAddHalfDay(e) {
    let hiddenDiv = document.getElementById("halfDay");
    if (e.target.checked === true) {
      hiddenDiv.style.display = "block";
    } else {
      hiddenDiv.style.display = "none";
    }
    console.log(`checked add hald day = ${e.target.checked}`);
  }

  onChangeIsHalfDay(e, value) {
    console.log(`${e.target.value} checked is ${e.target.checked}`);

    if (e.target.checked) {
      this.setState(prevState => ({
        halfDate: update(prevState.halfDate, { $push: [e.target.value] })
      }));
    } else {
      let array = this.state.halfDate;
      let index = array.indexOf(e.target.value);
      this.setState(prevState => ({
        halfDate: update(prevState.halfDate, { $splice: [[index, 1]] })
      }));
    }
  }

  onBackOn = value => {
    if (value !== null) {
      const date = new Date(value._d),
        mnth = ("0" + (date.getMonth() + 1)).slice(-2),
        day = ("0" + date.getDate()).slice(-2);
      let newDate = [day, mnth, date.getFullYear()].join("-");
      let backOn = {
        ...this.props.leaveForm,
        back_on: newDate,
        half_dates: this.state.halfDate
      };
      this.props.formOnChange(backOn);
    }
  };

  handleOnChangeID = value => {
    this.onChange("contactID", value);
  };

  handleOnChangeNumber = e => {
    let newLeave = {
      ...this.props.leaveForm,
      contact_number: `${this.state.contactID}${e.target.value}`
    };
    this.props.formOnChange(newLeave);
  };

  handleOnChangeEmployeeNumber = e => {
    let employee_num = {
      ...this.props.leaveForm,
      employee_number: Number(e.target.value)
    };
    this.props.formOnChange(employee_num);
  };

  getWorkingDate(startWorkingDate) {
    let today = new Date();
    let dd = today.getDate();
    let mm = today.getMonth() + 1;
    let yyyy = today.getFullYear();

    if (dd < 10) {
      dd = '0' + dd
    }
    if (mm < 10) {
      mm = '0' + mm
    }

    let dateNow = `${dd}-${mm}-${yyyy}`
    let start = moment(`${startWorkingDate}`, "DD-MM-YYYY");
    let end = moment(`${dateNow}`, "DD-MM-YYYY");
    let diffrent = end.diff(start, 'days')

    return diffrent
  }

  handleBlur() {
    console.log("blur");
  }

  handleFocus() {
    console.log("focus");
  }

  render() {
    const { from, to, start, end, endOpen } = this.state;
    const { getFieldDecorator } = this.props.form;
    const dates = this.getDates(start, end);
    const elements = [];
    const dateFormat = "DD-MM-YYYY";    

    const formItemLayout = {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 8 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 16 }
      },
      style: {}
    };

    let result = this.getWorkingDate("02-05-2018")

    const formStyle = {
      width: "100%"
    };

    const prefixSelector = getFieldDecorator("prefix", {
      initialValue: "+62"
    })(
      <Select onChange={this.handleOnChangeID} style={{ width: 70 }}>
        <Option value="+62">+62</Option>
        <Option value="+66">+66</Option>
      </Select>
    );

    for (let i = 0; i < dates.length; i++) {
      elements.push(
        <Checkbox
          key={i}
          id="is_half_day"
          name="is_half_day"
          onChange={e => this.onChangeIsHalfDay(e, dates[i])}
          value={dates[i]}
        >
          {dates[i]}
        </Checkbox>,
        <br />
      );
    }

    console.log("========>", this.state)
    
    console.log("this.props.publicHoliday")
    console.log(this.props.publicHoliday)

    function pad(num, size) {
      var s = num+"";
      while (s.length < size) s = "0" + s;
      return s;
    }

    for (let i = 0; i < this.props.publicHoliday.length; i++) {
      var holidayDateToAppend = []
      var dateStart = this.props.publicHoliday[i].date_start.split("-").reverse().join("-")
      var dateEnd = this.props.publicHoliday[i].date_end.split("-").reverse().join("-")
      holidayDateToAppend.push(dateStart)
      if (dateStart!=dateEnd) { //if date_start and date_end is different
        var dateStartInt = parseInt(dateStart.substring(dateStart.length-2, dateStart.length))
        var dateEndInt = parseInt(dateEnd.substring(dateEnd.length-2, dateEnd.length))
        if (dateStartInt>dateEndInt) { //if date_start is higher than dateEnd, holiday dates is within 2 month
          if (dateStartInt+1==dateEndInt) {
            holidayDateToAppend.push(dateEnd)
          }else{
            var dateStartMoment = moment(dateStart);
            var dateEndMoment = moment(dateEnd);
            var diffDays = dateEndMoment.diff(dateStartMoment, 'days');
            
            for (let j = 1; j <= diffDays; j++) {
              holidayDateToAppend.push(moment().add(j, 'd').format('YYYY-MM-DD'))
            }
          }
        }else{ //both date in the same month
          var suffixMonthYear = dateStart.substring(0, dateStart.length-2)
          for (let j = dateStartInt+1; j <= dateEndInt; j++) {
            holidayDateToAppend.push(suffixMonthYear+pad(j,2))
          }
        }
      }
      for (let k = 0; k < holidayDateToAppend.length; k++) {
        this.state.publicHolidayDates.push(holidayDateToAppend[k])
      }
      console.log(this.state.publicHolidayDates)
    }

    return (
      <Layout>
        <HeaderAdmin />
        <Content
          className="container"
          style={{
            display: "flex",
            margin: "20px 16px 0",
            justifyContent: "center",
            paddingBottom: "146px"
          }}
        >
          <div
            style={{
              padding: 100,
              paddingBottom: 50,
              paddingTop: 50,
              background: "#fff",
              minHeight: 360
            }}
          >
            <h1> Form Leave Request </h1>

            <Form onSubmit={this.handleSubmit} className="login-form">
              <FormItem {...formItemLayout} label="Employee ID">
                {getFieldDecorator("Employee ID", {
                  rules: [
                    {
                      required: true
                    }
                  ]
                })(
                  <Input
                    type="number"
                    id="employee_number"
                    name="employee_number"
                    placeholder="Employee ID"
                    onChange={this.handleOnChangeEmployeeNumber}
                  />
                )}
              </FormItem>
              <FormItem {...formItemLayout} label="Type Of Leave">
                {getFieldDecorator("type_leave_id", {
                  rules: [
                    {
                      required: true
                    }
                  ]
                })(
                  <Select
                    id="type_leave_id"
                    name="type_leave_id"
                    placeholder="Select type of leave"
                    optionFilterProp="children"
                    onChange={this.handleChangeTypeOfLeave}
                    onSelect={(value, event) =>
                      this.handleChangeSelect(value, event)
                    }
                    showSearch
                    filterOption={(input, option) =>
                      option.props.children
                        .toLowerCase()
                        .indexOf(input.toLowerCase()) >= 0
                    }
                    onFocus={this.handleFocus}
                    onBlur={this.handleBlur}
                    style={formStyle}
                  >
                    {result < 365 ?
                      this.props.typeLeave.map(d => (
                        <Option key={d.id} value={d.id}>{d.type_name}</Option>
                      ))
                      :
                      this.props.typeLeave.map(d => (
                        <Option key={d.id} value={d.id}>{d.type_name}</Option>
                      ))
                    }

                  </Select>
                )}
              </FormItem>

              {this.props.leaveForm.type_leave_id === 22 ||
                this.props.leaveForm.type_leave_id === 33 ||
                this.props.leaveForm.type_leave_id === 66 ? (
                  <FormItem {...formItemLayout} label="Reason">
                    <Input
                      type="text"
                      id="reason"
                      name="reason"
                      placeholder="reason"
                      onChange={this.handleOnChange}
                      style={formStyle}
                    />
                  </FormItem>
                ) : (
                  ""
                )}

              {this.props.leaveForm.type_leave_id === 22 ||
                this.props.leaveForm.type_leave_id === 33 ? (
                  <FormItem {...formItemLayout} label="From">
                    {getFieldDecorator("start date", {
                      rules: [
                        {
                          required: true
                        }
                      ]
                    })(
                      <DatePicker
                        id="date_from"
                        name="date_from"
                        disabledDate={this.disabledDateSick}
                        format={dateFormat}
                        value={from}
                        placeholder="Start"
                        onChange={this.onStartChange}
                        onOpenChange={this.handleStartOpenChange}
                        style={formStyle}
                      />
                    )}
                  </FormItem>
                ) : (
                  <FormItem {...formItemLayout} label="From">
                    {getFieldDecorator("start date", {
                      rules: [
                        {
                          required: true
                        }
                      ]
                    })(
                      <DatePicker
                        id="date_from"
                        name="date_from"
                        disabledDate={this.disabledDate}
                        format={dateFormat}
                        value={from}
                        placeholder="Start"
                        onChange={this.onStartChange}
                        onOpenChange={this.handleStartOpenChange}
                        style={formStyle}
                      />
                    )}
                  </FormItem>
                )}

              <FormItem {...formItemLayout} label="To">
                {getFieldDecorator("end date", {
                  rules: [
                    {
                      required: true
                    }
                  ]
                })(
                  <DatePicker
                    id="date_to"
                    name="date_to"
                    disabledDate={this.disabledEndDate}
                    format={dateFormat}
                    value={to}
                    placeholder="End"
                    onChange={this.onEndChange}
                    open={endOpen}
                    onOpenChange={this.handleEndOpenChange}
                    style={formStyle}
                  />
                )}
              </FormItem>
              <FormItem>
                <Checkbox
                  id="add_half_day"
                  name="add_half_day"
                  onChange={this.onChangeAddHalfDay}
                  style={formStyle}
                >
                  Add Half Day
                </Checkbox>
              </FormItem>

              <div id="halfDay">
                <FormItem {...formItemLayout} label="Half Day">
                  {elements}
                </FormItem>
              </div>

              <FormItem {...formItemLayout} label="Back to work on">
                {getFieldDecorator("back to work", {
                  rules: [
                    {
                      required: true
                    }
                  ]
                })(
                  <DatePicker
                    id="back_on"
                    name="back_on"
                    disabledDate={this.disabledDateBack}
                    onChange={this.onBackOn}
                    format={dateFormat}
                    placeholder="Back to work"
                    style={formStyle}
                  />
                )}
              </FormItem>

              <FormItem {...formItemLayout} label="Contact Address">
                {getFieldDecorator("contact address", {
                  rules: [
                    {
                      required: true
                    }
                  ]
                })(
                  <TextArea
                    type="text"
                    id="contact_address"
                    name="contact_address"
                    placeholder="address, email, etc"
                    onChange={this.handleOnChange}
                    autosize={{ minRows: 2, maxRows: 8 }}
                    style={formStyle}
                  />
                )}
              </FormItem>

              <FormItem {...formItemLayout} label="Contact Number">
                {getFieldDecorator("contact number", {
                  rules: [
                    {
                      required: true
                    }
                  ]
                })(
                  <Input
                    type="text"
                    id="contact_number"
                    name="contact_number"
                    placeholder="Phone number"
                    addonBefore={prefixSelector}
                    onChange={this.handleOnChangeNumber}
                    style={formStyle}
                  />
                )}
              </FormItem>

              <FormItem {...formItemLayout} label="Notes">
                {getFieldDecorator("notes", {
                  rules: [
                    {
                      required: true
                    }
                  ]
                })(
                  <TextArea
                    type="text"
                    id="notes"
                    name="notes"
                    placeholder="notes"
                    onChange={this.handleOnChange}
                    autosize={{ minRows: 2, maxRows: 8 }}
                    style={formStyle}
                  />
                )}
              </FormItem>


              <FormItem>
                <Button
                  onClick={this.handleSubmitSupervisor}
                  htmlType="submit"
                  type="primary"
                  style={{
                    width: "35%"
                  }}
                >
                  CREATE
                  </Button>
              </FormItem>
            </Form>
          </div>
        </Content>

        <Footer />
      </Layout>
    );
  }
}

const mapStateToProps = state => ({
  leaveForm: state.adminLeaveRequestReducer,
  typeLeave: state.fetchTypeLeaveReducer.typeLeave,
  user: state.fetchUserLoginReducer.user,
  publicHoliday: state.fetchPublicHolidayReducer.publicHoliday,
});

const WrappedLeaveForm = Form.create()(LeaveRequestPage);

const mapDispatchToProps = dispatch =>
  bindActionCreators(
    {
      formOnChange,      
      AdminSumbitLeave,
      typeLeaveFetchData,
      userLoginFetchData,
      publicHolidayFetchData
    },
    dispatch
  );

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(WrappedLeaveForm);
