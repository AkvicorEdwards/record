package handler

import (
	"fmt"
	"net/http"
	"record/tpl"
)

func account(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := tpl.Account.Execute(w, nil); err != nil {
			fmt.Println(err)
			_, _ = fmt.Fprintf(w, "%v", "Error")
		}
		return
	}
}
