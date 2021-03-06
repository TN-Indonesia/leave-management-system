package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	structLogic "server/structs/logic"

	"github.com/astaxie/beego"
)

// Date ...
type Date struct {
	day   int
	month time.Month
	year  int
}

// CheckErr ...
func CheckErr(errMsg string, err error) {
	if err != nil {
		beego.Warning(errMsg, err)
	}
}

// BytesToString ...
func BytesToString(data []byte) string {
	return string(data[:])
}

// ArrayToString ...
func ArrayToString(arr []string, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(arr), " ", delim, -1), "[]")
}

// Multiply ...
func Multiply(x, y, z float64) float64 {
	return (x - (y * z))
}

func stringToInt(value string) int {
	result, err := strconv.Atoi(value)
	CheckErr("err", err)
	return result
}

func stringToMonth(value string) time.Month {
	tm, err := time.Parse("01", value)
	CheckErr("err", err)
	return tm.Month()
}

func date2unix(d Date, loc *time.Location) int64 {
	return time.Date(d.year, d.month, d.day, 0, 0, 0, 0, loc).Unix()
}

func primitive(d Date, loc *time.Location) int64 {
	base := Date{3, time.January, 2000}
	seconds := date2unix(d, loc) - date2unix(base, loc)
	weeks := seconds / (7 * 24 * 60 * 60)
	secondIntoWeek := seconds % (7 * 24 * 60 * 60)
	workdays := secondIntoWeek / (24 * 60 * 60)
	if workdays > 5 {
		workdays = 5
	}
	return 5*weeks + workdays
}

func dayCountExcludingWeekends(from, to Date, loc *time.Location) int {
	return int(primitive(to, loc) - primitive(from, loc))
}

// GetTotalDay ...
func GetTotalDay(from string, to string) int {
	loc, err := time.LoadLocation("Asia/Jakarta")
	CheckErr("err", err)

	f := strings.Split(from, "-")
	t := strings.Split(to, "-")

	dateFrom := Date{stringToInt(f[0]), stringToMonth(f[1]), stringToInt(f[2])}
	dateTo := Date{stringToInt(t[0]), stringToMonth(t[1]), stringToInt(t[2])}
	result := dayCountExcludingWeekends(dateFrom, dateTo, loc)
	fmt.Println(from)
	fmt.Println("arr", f[0])
	return result + 1
}

// GetDay ...
func GetDay(date string) int {
	dateDay, _ := time.Parse("2006-01-02", date)
	m := dateDay.Month()
	var i = int(m)
	return i
}

var addedTimeNow = 0

// Now replace helper.Now()
func Now() time.Time {
	return time.Now().AddDate(0, 0, addedTimeNow)
}

// NowLoc ...
func NowLoc(timeLoc string) (
	dateTime time.Time,
	err error,
) {
	// "Asia/Bangkok"
	loc, err := time.LoadLocation(timeLoc)
	if err != nil {
		beego.Warning("failed load location", err)
	}
	dateTime = Now().In(loc)
	return
}

//PredictBackOn ...
func PredictBackOn(date string, allPublicHoliday []structLogic.GetAllPublicHoliday) string {
	result := ""
	t, _ := time.Parse("02-01-2006", date)
	t = t.AddDate(0, 0, 1)

	split := t.String()[0:10]
	result = split[8:10] + "-" + split[5:7] + "-" + split[0:4]

	if t.Weekday() == 6 || t.Weekday() == 0 {
		beego.Warning(time.Weekday(6))
		return PredictBackOn(result, allPublicHoliday)
	}

	//Check if back to work date is within public holiday
	for _, holiday := range allPublicHoliday {
		allDayWithinRange := GetAllDateWithinRange(holiday.DateStart, holiday.DateEnd)

		for _, d := range allDayWithinRange { // iterate every date inside range of public holiday
			if result == d {
				return PredictBackOn(result, allPublicHoliday)
			}
		}
	}

	return result

}

// InTimeSpan ...
func InTimeSpan(start, end, check string) bool {
	tanggaldari, _ := time.Parse("02-01-2006", start)
	tanggalsampai, _ := time.Parse("02-01-2006", end)
	tanggalcheck, _ := time.Parse("02-01-2006", check)
	return tanggalcheck.After(tanggaldari) && tanggalcheck.Before(tanggalsampai)
}

//TODO: ADD PUBLIC HOLIDAY AND MAKE IT RECURSIVE

//return slice of string contain all dates of given range

// GetAllDateWithinRange ...
func GetAllDateWithinRange(start, end string) []string {
	var allDayWithinRange []string

	startDate, _ := time.Parse("02-01-2006", start)
	endDate, _ := time.Parse("02-01-2006", end)

	days := endDate.Sub(startDate).Hours() / 24 // get number of days between 2 dates
	day := int(days)

	dayAfter := startDate

	fmt.Println(days)

	allDayWithinRange = append(allDayWithinRange, start)
	for i := 0; i < day; i++ {
		dayAfter = dayAfter.AddDate(0, 0, 1)
		dayAfterStr := dayAfter.Format("02-01-2006")
		allDayWithinRange = append(allDayWithinRange, dayAfterStr)
	}

	return allDayWithinRange
}
