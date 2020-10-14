package account

import (
	"fmt"
	"record/dam"
	"record/util"
)

const Menu = `
##########################
#        Account         #
# 1. Index               #
# 2. List                #
# 3. Add record          #
# 4. Delete record       #
# 5. Update record       #
# 6. Search by field     #
# 0. Exit                #
##########################
`

const Field = `
##########################
# 1. ID                  #
# 2. Title               #
# 3. Account             #
# 4. Password            #
# 5. Comment             #
# 0. Exit                #
##########################
`

func Account() {
	fmt.Print(Menu)
	for {
		L:
		switch util.ReadInt() {
		case 1:
			accounts := dam.AccountGetAll()
			fmt.Println("##########################################")
			for k, v := range accounts {
				if k != 0 {
					fmt.Println()
				}
				fmt.Printf("# ID:[%d] Title:[%s] Comment:[%s]\n", v.Id, v.Title, v.Comment)
			}
			fmt.Println("##########################################")
		case 2:
			accounts := dam.AccountGetAll()
			fmt.Println("##########################################")
			for k, v := range accounts {
				if k != 0 {
					fmt.Println()
				}
				fmt.Printf("# ID:[%d] Title:[%s]\n", v.Id, v.Title)
				fmt.Printf("# Account:[%s] Password:[%s]\n", v.Account, v.Password)
				fmt.Println("# Secret Question:")
				for _, val := range v.SecretQuestion {
					fmt.Printf("#   Quest:[%s] Answer:[%s]\n", val.Quest, val.Answer)
				}
				fmt.Println("# Two Factor:")
				for _, val := range v.TwoFactor {
					fmt.Printf("#   [%s]\n", val)
				}
				fmt.Printf("# Comment:[%s]\n", v.Comment)
			}
			fmt.Println("##########################################")
		case 3:
			fmt.Println("Please input Title:")
			title := util.ReadLine()
			fmt.Println("Please input Account:")
			account := util.ReadLine()
			fmt.Println("Please input Password:")
			password := util.ReadLine()
			question := make([]dam.SecretQuestion, 0)
			fmt.Println("Please input Secret Question ('.exit' to exit):")
			for {
				fmt.Println("Question:")
				quest := util.ReadLine()
				if quest == ".exit" {
					break
				}
				fmt.Println("Answer:")
				answer := util.ReadLine()
				if answer == ".exit" {
					break
				}
				if len(quest) == 0 && len(answer) == 0 {
					continue
				}
				question = append(question, dam.SecretQuestion{
					Quest:  quest,
					Answer: answer,
				})
			}
			factor := make([]string, 0)
			fmt.Println("Please input Two Factor ('.exit' to exit):")
			for {
				fmt.Println("Recovery code:")
				code := util.ReadLine()
				if code == ".exit" {
					break
				}
				if len(code) == 0 {
					continue
				}
				factor = append(factor, code)
			}

			fmt.Println("Please input Comment:")
			comment := util.ReadLine()
			fmt.Println("Are you sure to add the record? yes/No")
			opt := "no"
			opt = util.ReadLine()
			if opt != "yes" {
				break
			}
			if id, ok := dam.AccountAdd(title, account, password, question, factor, comment); ok {
				fmt.Println()
				fmt.Printf("# ID:[%d] Title:[%s]\n", id, title)
				fmt.Printf("# Account:[%s] Password:[%s]\n", account, password)
				fmt.Println("# Secret Question:")
				for _, val := range question {
					fmt.Printf("#   Quest:[%s] Answer:[%s]\n", val.Quest, val.Answer)
				}
				fmt.Println("# Two Factor:")
				for _, val := range factor {
					fmt.Printf("#   [%s]\n", val)
				}
				fmt.Printf("# Comment:[%s]\n", comment)
				fmt.Println("\nFinished")
			} else {
				fmt.Println("Failure")
			}
		case 4:
			fmt.Println("Please input id:")
			id := util.ReadUInt32()
			fmt.Printf("Are you sure to delete the [%d] record? yes/No\n", id)
			opt := "no"
			opt = util.ReadLine()
			if opt != "yes" {
				break
			}
			if dam.AccountDelete(id) {
				fmt.Println("Finished")
			} else {
				fmt.Println("Failure")
			}
		case 5:
			fmt.Println("Please input id:")
			id := util.ReadUInt32()
			account := dam.AccountGet(id)
			I:
			for {
				fmt.Println("1. Title")
				fmt.Println("2. Account")
				fmt.Println("3. Password")
				fmt.Println("4. Secret Question")
				fmt.Println("5. Two Factor")
				fmt.Println("6. Comment")
				fmt.Println("7. # Modify")
				fmt.Println("0. # Exit")
				fmt.Println("Please enter the field you want to modify:")
				opt := util.ReadInt()
				switch opt {
				case 1:
					fmt.Println("Please enter new Title:")
					account.Title = util.ReadLine()
				case 2:
					fmt.Println("Please enter new Account:")
					account.Account = util.ReadLine()
				case 3:
					fmt.Println("Please enter new Password:")
					account.Password = util.ReadLine()
				case 4:
					question := make([]dam.SecretQuestion, 0)
					fmt.Println("Please input new Secret Question ('.exit' to exit):")
					for {
						fmt.Println("Question:")
						quest := util.ReadLine()
						if quest == ".exit" {
							break
						}
						fmt.Println("Answer:")
						answer := util.ReadLine()
						if answer == ".exit" {
							break
						}
						if len(quest) == 0 && len(answer) == 0 {
							continue
						}
						question = append(question, dam.SecretQuestion{
							Quest:  quest,
							Answer: answer,
						})
					}
					account.SecretQuestion = question
				case 5:
					factor := make([]string, 0)
					fmt.Println("Please input new Two Factor ('.exit' to exit):")
					for {
						fmt.Println("Recovery code:")
						code := util.ReadLine()
						if code == ".exit" {
							break
						}
						if len(code) == 0 {
							continue
						}
						factor = append(factor, code)
					}
					account.TwoFactor = factor
				case 6:
					fmt.Println("Please enter new Comment:")
					account.Comment = util.ReadLine()
				case 7:
					fmt.Printf("# ID:[%d] Title:[%s]\n", account.Id, account.Title)
					fmt.Printf("# Account:[%s] Password:[%s]\n", account.Account, account.Password)
					fmt.Println("# Secret Question:")
					for _, val := range account.SecretQuestion {
						fmt.Printf("#   Quest:[%s] Answer:[%s]\n", val.Quest, val.Answer)
					}
					fmt.Println("# Two Factor:")
					for _, val := range account.TwoFactor {
						fmt.Printf("#   [%s]\n", val)
					}
					fmt.Printf("# Comment:[%s]\n", account.Comment)
					fmt.Println("Are you sure to modify the record? yes/No")
					opt := "no"
					opt = util.ReadLine()
					if opt != "yes" {
						break
					}
					dam.AccountUpdate(account)
				case 0:
					break I
				default:
					break
				}
			}
		case 6:
			fmt.Print(Field)
			fmt.Println("Please input field id:")
			id := util.ReadUInt32()
			field := ""
			switch id {
			case 1:
				field = "id"
			case 2:
				field = "title"
			case 3:
				field = "account"
			case 4:
				field = "password"
			case 5:
				field = "comment"
			case 0:
				break L
			default:
				field = "title"
			}
			fmt.Println("Please input regexp:")
			reg := util.ReadLine()
			accounts := dam.AccountGetByRegexp(field, reg)
			fmt.Println("##########################################")
			for k, v := range accounts {
				if k != 0 {
					fmt.Println()
				}
				fmt.Printf("# ID:[%d] Title:[%s]\n", v.Id, v.Title)
				fmt.Printf("# Account:[%s] Password:[%s]\n", v.Account, v.Password)
				fmt.Println("# Secret Question:")
				for _, val := range v.SecretQuestion {
					fmt.Printf("#   Quest:[%s] Answer:[%s]\n", val.Quest, val.Answer)
				}
				fmt.Println("# Two Factor:")
				for _, val := range v.TwoFactor {
					fmt.Printf("#   [%s]\n", val)
				}
				fmt.Printf("# Comment:[%s]\n", v.Comment)
			}
			fmt.Println("##########################################")
		case 0:
			return
		default:
			break
		}
		fmt.Print(Menu)
	}
}
