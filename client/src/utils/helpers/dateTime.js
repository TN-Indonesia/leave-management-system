import moment from "moment";

export function DisabledDate(current, publicHolidayDates) {
  // const publicHolidayDates = this.state.publicHolidayDates;
  return current < moment().startOf("day") ||
  publicHolidayDates.find(d => moment(d).format("DDMMYYYY") === moment(current._d).format("DDMMYYYY")) ||
    moment(current._d).format("dddd") === "Saturday" ||
    moment(current._d).format("dddd") === "Sunday";
}

export function DisabledDateSick(current, publicHolidayDates) {
  // const publicHolidayDates = this.state.publicHolidayDates;
  return (
      current &&
      current <
      moment()
      .subtract(7, "days")
      .startOf("day")
    ) ||
    publicHolidayDates.find(d => moment(d).format("DDMMYYYY") === moment(current._d).format("DDMMYYYY")) ||
    moment(current._d).format("dddd") === "Saturday" ||
    moment(current._d).format("dddd") === "Sunday";
}

export function DisabledDateBack(current, publicHolidayDates) {
  // const publicHolidayDates = this.state.publicHolidayDates;
  return this.state.to > current ||
    publicHolidayDates.find(d => moment(d).format("DDMMYYYY") === moment(current._d).format("DDMMYYYY")) ||
    moment(current._d).format("dddd") === "Saturday" ||
    moment(current._d).format("dddd") === "Sunday";
}

export function GetDates(start, end, publicHolidayDates) {
  // const publicHolidayDates = this.state.publicHolidayDates;
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

export function CountTotalDay(startDate, endDate, disabledDays) {
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

  let result = ((end - start) / 86400000 - weekend_count) + 1;
  return result
}

export function GetWorkingDate(startWorkingDate) {
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