import React, { Component } from "react";
import { bindActionCreators } from "redux";
import { connect } from "react-redux";
import update from "react-addons-update";
import moment from "moment-business-days";
import {
  formOnChange,
  SumbitLeave,
  SumbitLeaveSupervisor
} from "../../../store/Actions/leaveRequestAction";
import { typeLeaveFetchData } from "../../../store/Actions/typeLeaveAction";
import { userLoginFetchData } from "../../../store/Actions/userLoginAction";
import { publicHolidayFetchData } from "../../../store/Actions/publicHolidayAction";
import HeaderNav from "../../../pages/menu/HeaderNav";
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
let totalDays;

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
      publicHolidayDates: null,
      totalDays: 0,
      backOnDate : false,
      dateBackOn : "",
    };

    this.handleOnChange = this.handleOnChange.bind(this);
    this.handleOnChangeNumber = this.handleOnChangeNumber.bind(this);
    this.handleOnChangeEmployeeNumber = this.handleOnChangeEmployeeNumber.bind(this);
    this.handleChangeTypeOfLeave = this.handleChangeTypeOfLeave.bind(this);
    this.handleOnChangeID = this.handleOnChangeID.bind(this);
    this.onChangeIsHalfDay = this.onChangeIsHalfDay.bind(this);
    this.onChangeAddHalfDay = this.onChangeAddHalfDay.bind(this);
    this.disabledDate = this.disabledDate.bind(this);
    this.disabledDateSick = this.disabledDateSick.bind(this);
    this.disabledDateBack = this.disabledDateBack.bind(this);
  }

  componentWillMount() {
    console.log(" ----------------- Form-Leave-Request ----------------- ");
    if (!localStorage.getItem("token")) {
      this.props.history.push("/");
    } else if (
      localStorage.getItem("role") !== "employee" &&
      localStorage.getItem("role") !== "supervisor"
    ) {
      this.props.history.push("/");
    }
  }

  componentDidMount() {
    this.props.typeLeaveFetchData();
    this.props.userLoginFetchData();
    this.props.publicHolidayFetchData();
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.publicHoliday !== this.props.publicHoliday) {
      this.setState({ publicHolidayDates: nextProps.publicHoliday });
    }
  }

  componentDidUpdate(prevProps, prevState) {
    if (totalDays) {
      if (prevState.totalDays !== totalDays) {
        this.setState({ totalDays: totalDays });
      }
    }
  };

  onChange = (field, value) => {
    this.setState({
      [field]: value
    });
  };

  handleOnChangeNumber = (value, field) => {
    this.onChange(field, Number(value));
    console.log("input=======>", value);
  };

  handleOnChange = e => {
    let newLeave = {
      ...this.props.leaveForm,
      [e.target.name]: e.target.value
    };
    this.props.formOnChange(newLeave);
  };

  handleSubmit = e => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        console.log("Received values of form: ", values);
      }
    });
    this.props.SumbitLeave(this.props.leaveForm, url => {
      this.props.history.push(url);
    });
  };

  handleSubmitSupervisor = e => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        console.log("Received values of form: ", values);
      }
    });
    this.props.SumbitLeaveSupervisor(this.props.leaveForm, url => {
      this.props.history.push(url);
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

  handleStartOpenChange = open => {
    if (!open) {
      this.setState({ endOpen: true });
    }
  };

  handleEndOpenChange = open => {
    this.setState({ endOpen: open });
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

  onChangeAddHalfDay(e) {
    let hiddenDiv = document.getElementById("halfDay");
    if (e.target.checked === true) {
      hiddenDiv.style.display = "block";
    } else {
      hiddenDiv.style.display = "none";
    }
    console.log(`checked add half day = ${e.target.checked}`);
  }

  onChangeBackoptionMorning(e, max) {
    if (e.target.checked === true) {
        let backOn = {
          ...this.props.leaveForm,
          back_on: max,
          half_dates: this.state.halfDate,
          total: Number(this.state.totalDays)
        };
        this.props.formOnChange(backOn);
        this.setState( {dateBackOn: max})
        const myDate = moment(max, 'YYYY-MM-DD').toDate();
      console.log("cekilis",myDate);
      // this.onBackOn;
    } else {
      console.log("uncekilis");
      this.setState({dateBackOn:""})
    }
  }
  //TODO Make weekend and holliday
  onChangeBackoptionNoon(e, max) {
    if (e.target.checked === true) {
        
        var myDate = moment(max, 'YYYY-MM-DD');
        console.log("masuk", myDate)
        myDate.add(1,"days")
        console.log("masuk 1", myDate)

        let time = moment(myDate).format('DD-MM-YYYY');
        console.log("masuk2", time)

      let backOn = {
          ...this.props.leaveForm,
          back_on: time,
          half_dates: this.state.halfDate,
          total: Number(this.state.totalDays)
        };
        this.props.formOnChange(backOn);
        this.setState( {dateBackOn: time})
        console.log("cekilis asoy",this.state.dateBackOn);

    } else {
      console.log("uncekilis");
      this.setState({dateBackOn:""})
    }
  }

  onChangeIsHalfDay(e, value, last) {
    console.log(`${e.target.value} checked is ${e.target.checked}`);
    console.log("udah masuk", last)
    let parentDiv = document.getElementById("add_half_day");
    if (e.target.checked) {
      parentDiv.disabled = true
      if(e.target.value === last ) {
        console.log("masuk pak eko")
        this.setState({ backOnDate: true });
      }
      this.setState(prevState => ({
        halfDate: update(prevState.halfDate, { $push: [e.target.value] })
      }));
    } else {
      parentDiv.disabled = false
      let array = this.state.halfDate;
      let index = array.indexOf(e.target.value);
      if(e.target.value === last ) {
        this.setState({ backOnDate: false });
      }
      this.setState(prevState => ({
        halfDate: update(prevState.halfDate, { $splice: [[index, 1]] })
      }));
    }
  }

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
        date_to: newDate,

      };
      this.props.formOnChange(dateTo);
      this.onChange("end", newEnd);
    }
    this.onChange("to", value);
  };

  onBackOn = value => {
    if (value !== null) {
      const date = new Date(value._d),
        mnth = ("0" + (date.getMonth() + 1)).slice(-2),
        day = ("0" + date.getDate()).slice(-2);
      let newDate = [day, mnth, date.getFullYear()].join("-");
      let backOn = {
        ...this.props.leaveForm,
        back_on: newDate,
        half_dates: this.state.halfDate,
        total: Number(this.state.totalDays)
      };
      this.props.formOnChange(backOn);
    }
  };

  disabledStartDate = startValue => {
    const endValue = this.state.to;
    if (!startValue || !endValue) {
      return false;
    }
    return startValue.valueOf() > endValue.valueOf();
  };

  disabledEndDate = endValue => {
    const publicHolidayDates = this.state.publicHolidayDates;
    const startValue = this.state.from;
    if (!endValue || !startValue) {
      return false;
    }

    return endValue.valueOf() <= startValue.valueOf()
      || publicHolidayDates.find(d => moment(d).format("DDMMYYYY") === moment(endValue).format("DDMMYYYY"))
      || moment(endValue).format("dddd") === "Saturday"
      || moment(endValue).format("dddd") === "Sunday";
  };

  disabledDate(current) {
    const publicHolidayDates = this.state.publicHolidayDates;
    return current < moment().startOf("day")
      || publicHolidayDates.find(d => moment(d).format("DDMMYYYY") === moment(current._d).format("DDMMYYYY"))
      || moment(current._d).format("dddd") === "Saturday"
      || moment(current._d).format("dddd") === "Sunday";
  }

  disabledDateSick(current) {
    const publicHolidayDates = this.state.publicHolidayDates;
    return (
      current &&
      current <
      moment()
        .subtract(7, "days")
        .startOf("day")
    )
      || publicHolidayDates.find(d => moment(d).format("DDMMYYYY") === moment(current._d).format("DDMMYYYY"))
      || moment(current._d).format("dddd") === "Saturday"
      || moment(current._d).format("dddd") === "Sunday";
  }

  disabledDateBack(current) {
    const publicHolidayDates = this.state.publicHolidayDates;
    return this.state.to > current
      || publicHolidayDates.find(d => moment(d).format("DDMMYYYY") === moment(current._d).format("DDMMYYYY"))
      || moment(current._d).format("dddd") === "Saturday"
      || moment(current._d).format("dddd") === "Sunday";
  }

  getDates(start, end) {
    let publicHolidayDates = this.state.publicHolidayDates;
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

    if (publicHolidayDates) {
      let newDate = []
      for (let i = 0; i < publicHolidayDates.length; i++) {
        let date = publicHolidayDates[i].split("-").reverse().join("-")
        newDate.push(date)
      }

      for (let i = 0; i < dates.length; i++) {
        for (let j = 0; j < newDate.length; j++) {
          if (dates[i] === newDate[j]) {
            dates.splice(i, 1);
          }
        }
      }
    }

    return dates;
  }

  countTotalDay(startDate, endDate) {
    let disabledDays = this.state.publicHolidayDates;
    let start = new Date(startDate);
    let end = new Date(endDate);
    let weekend_count = 0;
    for (let i = start.valueOf(); i <= end.valueOf(); i += 86400000) {
      let temp = new Date(i);
      let holiday;
      for (let j = 0; j < disabledDays.length; j++) {
        holiday = disabledDays[j];
        if (!(temp < new Date(holiday)) && !(temp > new Date(holiday))) {
          weekend_count++
        }
      }

      if (temp.getDay() === 0 || temp.getDay() === 6) {
        weekend_count++;
      }
    }

    let result = (((end - start) / 86400000) - weekend_count) + 1;
    return result
  }

  handleBlur() {
    console.log("blur");
  }

  handleFocus() {
    console.log("focus");
  }


  render() {
    const { from, to, start, end, endOpen ,backOnDate} = this.state;
    const { getFieldDecorator } = this.props.form;
    const dates = this.getDates(start, end);
    const elements = [];
    const dateFormat = "DD-MM-YYYY";
    const role = localStorage.getItem("role");


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


    //TODO make backon for half leave
    for (let i = 0; i < dates.length; i++) {
      elements.push(
        <Checkbox
          key={i}
          id="is_half_day"
          name="is_half_day"
          onChange={e => this.onChangeIsHalfDay(e, dates[i], dates[dates.length-1])}
          value={dates[i]}
        >
          {dates[i]}
        </Checkbox>,
        <br />
      );
    }

    if (this.state.start !== null && this.state.end !== null) {
      let dateStart = new Date(this.state.start);
      let dateEnd = new Date(this.state.end);
      totalDays = this.countTotalDay(dateStart, dateEnd)
    }

    console.log("========>", this.state)

    return (
      <Layout>
        <HeaderNav />
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
                    {this.props.typeLeave.map(d => (
                      <Option key={d.id} value={d.id}>{d.type_name}</Option>
                    ))}

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
                { backOnDate ? (
                  <div>
                  <FormItem >
                  <Checkbox
                    id="on_back_date"
                    name="on_back_date"
                    style={formStyle}
                    onChange= {e => this.onChangeBackoptionMorning(e, dates[dates.length-1])}
                   
                  >
                  Morning
                  </Checkbox>
                </FormItem>
                <FormItem>
                  <Checkbox
                    id="on_back_date"
                    name="on_back_date"
                    style={formStyle}
                    onChange= {e => this.onChangeBackoptionNoon(e, dates[dates.length-1])}
                  >
                  Afternoon
                  </Checkbox>
                </FormItem>
                </div>
              ) : (
                  <div></div>
                )}
             
              </div>
              <FormItem {...formItemLayout} label="Back to work on">
                {/* {getFieldDecorator("back to work", {
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
                )} */}
              <Input
                    type="text"
                    id="back_on"
                    name="back_on"
                    disabled
                    value={this.state.dateBackOn}
                    style={formStyle}
                  />
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

              {role === "employee" ? (
                <FormItem>
                  <Button
                    onClick={this.handleSubmit}
                    htmlType="submit"
                    type="primary"
                    style={{
                      width: "35%"
                    }}
                  >
                    Create
                  </Button>
                </FormItem>
              ) : (
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
                )}
            </Form>
          </div>
        </Content>

        <Footer />
      </Layout>
    );
  }
}

const mapStateToProps = state => ({
  leaveForm: state.leaveRequestReducer,
  typeLeave: state.fetchTypeLeaveReducer.typeLeave,
  user: state.fetchUserLoginReducer.user,
  publicHoliday: state.fetchPublicHolidayReducer.publicHoliday
});

const WrappedLeaveForm = Form.create()(LeaveRequestPage);

const mapDispatchToProps = dispatch =>
  bindActionCreators(
    {
      formOnChange,
      SumbitLeave,
      SumbitLeaveSupervisor,
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
