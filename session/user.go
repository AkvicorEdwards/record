package session

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

// 获取全局session
func GetUser(r *http.Request) (*sessions.Session, error) {
	return Get(r, "user")
}

// 设置用户的Session
func SetPer(w http.ResponseWriter, r *http.Request, per int) {
	// Session
	ses, _ := GetUser(r)
	ses.Values["per"] = per
	ses.Options.MaxAge = 60 * 60 * 24
	err := ses.Save(r, w)
	if err != nil {
		_, _ = fmt.Fprintln(w, "ERROR session SetUserInfo")
		return
	}
}

func GetPer(r *http.Request) int {
	ses, err := Get(r, "user")
	if err != nil {
		return 0
	}
	per, ok := ses.Values["per"].(int)
	if !ok {
		return 0
	}
	return per
}
