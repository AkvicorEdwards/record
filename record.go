package main

import (
	"fmt"
	"os"
	"record/def"
	"record/maintain"
	"record/mod/account"
	"record/mod/port"
	"record/util"
)

const Menu = `
##########################
# 1. Account             #
# 2. Port                #
# 0. Exit                #
##########################
`

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "init" {
		maintenance.InitDatabase()
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
		case -1:
			fmt.Printf("Encrypt Key: [%s]", def.EncryptKey)
		case 0:
			os.Exit(0)
		}
		fmt.Print(Menu)
	}
}
