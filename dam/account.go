package dam

import (
	"fmt"
	"log"
	"regexp"
)

func AccountAdd(title, account, password string, question []SecretQuestion,
	factor []string, comment string) (uint32, bool) {
	if !Connected {
		Connect()
	}
	lockAccount.Lock()
	defer lockAccount.Unlock()
	a := Account{
		Id:             GetInc("account") + 1,
		Title:          title,
		Account:        account,
		Password:       password,
		SecretQuestion: question,
		TwoFactor:      factor,
		Comment:        comment,
	}

	aDB := a.Transfer()

	if err := db.Table("account").Create(&aDB).Error; err != nil {
		log.Println(err)
		return 0, false
	}

	UpdateInc("account", aDB.Id)

	return aDB.Id, true
}

func AccountDelete(id uint32) bool {
	if !Connected {
		Connect()
	}
	lockAccount.Lock()
	defer lockAccount.Unlock()

	res := db.Table("account").Delete(&AccountDB{}, id)
	if res.Error != nil {
		log.Println(res.Error)
		return false
	}
	if res.RowsAffected == 0 {
		return false
	}
	return true
}

func AccountUpdate(account Account) bool {
	if !Connected {
		Connect()
	}
	lockAccount.Lock()
	defer lockAccount.Unlock()

	ac := account.Transfer()

	res := db.Table("account").Where("id=?", account.Id).Updates(map[string]interface{}{
		"title":            ac.Title,
		"account":         ac.Account,
		"password":        ac.Password,
		"secret_question": ac.SecretQuestion,
		"two_factor":      ac.TwoFactor,
		"comment":         ac.Comment,
	})

	if res.Error != nil {
		log.Println(res.Error)
		return false
	}

	if res.RowsAffected == 0 {
		return false
	}

	return true
}

func AccountGet(id uint32) Account {
	if !Connected {
		Connect()
	}

	port := AccountDB{}

	db.Table("account").Where("id=?", id).First(&port)

	return port.Transfer()
}

func AccountGetByRegexp(field, reg string) []Account {
	if !Connected {
		Connect()
	}
	accounts := AccountGetAll()
	res := make([]Account, 0)
	var matched bool
	var err error
	for _, v := range accounts {
		switch field {
		case "id":
			matched, err = regexp.MatchString(reg, fmt.Sprint(v.Id))
		case "title":
			matched, err = regexp.MatchString(reg, v.Title)
		case "account":
			matched, err = regexp.MatchString(reg, v.Account)
		case "password":
			matched, err = regexp.MatchString(reg, v.Password)
		case "comment":
			matched, err = regexp.MatchString(reg, v.Comment)
		default:
			continue
		}
		if err != nil {
			break
		}
		if matched {
			res = append(res, v)
		}
	}
	return res
}

func AccountGetAll() []Account {
	if !Connected {
		Connect()
	}

	accounts := make([]AccountDB, 0)

	db.Table("account").Find(&accounts)

	res := make([]Account, len(accounts))

	for k, v := range accounts {
		res[k] = v.Transfer()
	}

	return res
}