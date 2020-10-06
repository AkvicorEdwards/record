package handler

import (
	"fmt"
	"net/http"
	"record/def"
	"record/session"
	"record/tpl"
)

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

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := tpl.Index.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
		return
	}
}
