package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"record/def"
	"record/handler"
	"record/maintain"
	"record/mod/account"
	"record/mod/port"
	"record/tpl"
	"record/util"
	"time"
)

const Menu = `
##########################
# 1. Account             #
# 2. Port                #
# 3. HTTP Server         #
# 0. Exit                #
##########################
`

var httpStarted = flag.Bool("http", false, "HTTP Server")
var initDatabase = flag.Bool("init", false, "Init Database")

func main() {
	flag.Parse()

	if *initDatabase {
		maintenance.InitDatabase()
	}

	if *httpStarted {
		*httpStarted = false
		HTTPServer()
	}

	go maintenance.ShutDownListener()

	if !def.CheckEncryptKey() {
		os.Exit(def.WrongEncryptKeyLength)
	}

	if util.GetPassword() != def.Password {
		os.Exit(def.WrongPassword)
	}

	fmt.Print(Menu)
	for {
		switch util.ReadInt() {
		case 1:
			account.Account()
		case 2:
			port.Port()
		case 3:
			HTTPServer()
		case -1:
			fmt.Printf("Encrypt Key: [%s]", def.EncryptKey)
		case 0:
			os.Exit(0)
		}
		fmt.Print(Menu)
	}
}

func HTTPServer() {
	if *httpStarted {
		return
	}
	*httpStarted = true

	tpl.Parse()
	handler.ParsePrefix()
	addr := fmt.Sprintf(":%d", def.Port)
	server := http.Server{
		Addr:              addr,
		Handler:           &handler.MyHandler{},
		ReadTimeout:       20 * time.Minute,
	}
	log.Printf("http://127.0.0.1%s\n", addr)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
}
