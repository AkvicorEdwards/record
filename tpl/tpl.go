package tpl

import (
	"html/template"
	"log"
)

var err error
var errCnt = 0

var Login *template.Template
var Heartbeat *template.Template

func Parse() {
	err = nil
	errCnt = 0

	Login = addFromFile("./tpl/login.html")
	Heartbeat = addFromFile("./tpl/heartbeat.html")

	log.Printf("Parsing the html template was completed with %d errors\n", errCnt)
}

func add(name, tpl string) (t *template.Template) {
	t, err = template.New(name).Parse(tpl)
	if err != nil {
		errCnt++
		log.Println(err)
	}
	return
}

func addFromFile(file string) (t *template.Template) {
	t, err = template.ParseFiles(file)
	if err != nil {
		errCnt++
		log.Println(err)
	}
	return
}