package handler

import (
	"fmt"
	"net/http"
	"record/dam"
	"record/tpl"
	"strings"
)

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
