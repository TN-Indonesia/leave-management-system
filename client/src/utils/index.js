export function MakeID() {
    let x = 10;
    let text = "";
    let possible =
        "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*?";

    for (let i = 0; i < 7; i++)
        text += possible.charAt(Math.floor(Math.random() * possible.length));

    return text;
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
    let result = ((end - start) / 86400000 - weekend_count);
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