package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"record/dam"
	"record/def"
	"record/session"
	"record/tpl"
	"regexp"
	"strings"
	"time"
)

type str2func map[string]func(http.ResponseWriter, *http.Request)

var public str2func

func ParsePrefix() {
	public = make(str2func)

	public["/login"] = login
	public["/heartbeat"] = heartbeat
}

type MyHandler struct{}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if session.GetPer(r) != 7 {
		public["/login"](w, r)
		return
	}

	if h, ok := public[r.URL.Path]; ok {
		h(w, r)
		return
	}

	match := func(pattern string) (matched bool) {
		matched, _ = regexp.MatchString(pattern, r.URL.String())
		return
	}

	if match("/favicon.ico") {
		download(w, "./record.ico")
	} else if match("/vue.js") {
		download(w, "./tpl/js/vue.js")
	} else if match("/echarts.min.js") {
		download(w, "./tpl/js/echarts.min.js")
	}

}

func download(w http.ResponseWriter, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		_, _ = fmt.Fprintln(w, "File Not Found")
		return
	}
	defer func() { _ = file.Close() }()
	data := make([]byte, 1024)
	for {
		n, err1 := file.Read(data)
		if err1 != nil && err1 != io.EOF {
			_, _ = fmt.Fprintln(w, "File Read Error")
			return
		}
		nn, err2 := w.Write(data[:n])
		if err2 != nil || nn != n {
			_, _ = fmt.Fprintln(w, "File Write Error")
			return
		}
		if err1 == io.EOF {
			return
		}
	}
}

func Fprint(w http.ResponseWriter, a ...interface{}) {
	_, _ = fmt.Fprint(w, a...)
}

func Fprintf(w http.ResponseWriter, format string, a ...interface{}) {
	_, _ = fmt.Fprintf(w, format, a...)
}

func Fprintln(w http.ResponseWriter, a ...interface{}) {
	_, _ = fmt.Fprintln(w, a...)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := tpl.Login.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
		return
	}

	password := r.FormValue("password")

	if def.Password != password {
		_, _ = fmt.Fprintln(w, "密码错误")
		return
	}

	session.SetPer(w, r, 7)

	http.Redirect(w, r, "/", 302)
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		date := r.FormValue("date")
		data := ""
		if date ==  "all" {
			d := dam.HeartbeatGetAll()
			for k, v := range d {
				if k != 0 {
					data += "\n"
				}
				data += v.Data
			}
		} else if len(strings.Split(date, ",")) == 2 {
			date := GetBetweenDates(strings.Split(date, ",")[0], strings.Split(date, ",")[1])
			for _, v := range date {
				d := dam.HeartbeatGetByDate(v).Data
				if len(d) <= 2 {
					continue
				}
				if len(data) != 0 {
					data += "\n"
				}
				data += d
			}
		} else {
			d := dam.HeartbeatGetByDate(date)
			data = d.Data
		}
		//fmt.Println(data)
		phrase := map[string]interface{}{
			"DATA": data,
		}
		if err := tpl.Heartbeat.Execute(w, phrase); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
		return
	} else if r.Method == "POST" {
		date := strings.TrimSpace(r.FormValue("date"))
		data := strings.TrimSpace(r.FormValue("data"))
		dataFirst := strings.Split(data, "\n")
		data = ""
		for i := len(dataFirst)-1; i >= 0; i-- {
			if dataFirst[i] == "" {
				continue
			}
			if len(data) != 0 {
				data += "\n"
			}
			data += dataFirst[i]
		}
		if _, ok := dam.HeartbeatAddRecord(date, data); !ok {
			Fprintf(w, "Error")
			return
		}
		http.Redirect(w, r, "/heartbeat", 302)
	}
}

func GetBetweenDates(sdate, edate string) []string {
	var d []string
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