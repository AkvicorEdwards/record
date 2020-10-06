package handler

import "time"

func GetBetweenDates(sdate, edate string) []string {
	var d []string
	if sdate == edate {
		return []string{sdate}
	}
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		return d
	}
	if date2.Before(date) {
		return d
	}
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}
