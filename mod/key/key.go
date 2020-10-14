package key

import (
	"fmt"
	"record/dam"
	"record/util"
)

const Menu = `
##########################
#          Key           #
# 1. List                #
# 2. Add record          #
# 3. Delete record       #
# 4. Update record       #
# 5. Search by field     #
# 0. Exit                #
##########################
`

const Field = `
##########################
# 1. ID                  #
# 2. Title               #
# 3. Key                 #
# 4. Comment             #
# 0. Exit                #
##########################
`

func Key() {
	fmt.Print(Menu)
	for {
	L:
		switch util.ReadInt() {
		case 1:
			keys := dam.KeyGetAll()
			fmt.Println("##########################################")
			for k, v := range keys {
				if k != 0 {
					fmt.Println()
				}
				fmt.Printf("# ID:[%d] Title:[%s]\n",
					v.Id, v.Title)
				for _, val := range v.Key {
					fmt.Printf("# Key:[\n%s\n] Comment:[%s]\n", val.Key, val.Comment)
				}
				fmt.Printf("# Comment:[%s]\n", v.Comment)
			}
			fmt.Println("##########################################")
		case 2:
			fmt.Println("Please input Title:")
			title := util.ReadLine()
			fmt.Println("Please input Key ('.exit' to exit):")
			fmt.Println("'EOF' to finish input")
			key := make([]dam.KeyInfo, 0)
			K:
			for {
				fmt.Println("Key:")
				p := ""
				for {
					pt := util.ReadLine()
					if pt == ".exit" {
						break K
					} else if pt == "EOF" {
						break
					}
					if len(p) != 0 {
						p += "\n"
					}
					p += pt
				}
				fmt.Println("Comment:")
				c := util.ReadLine()
				if c == ".exit" {
					break
				}
				key = append(key, dam.KeyInfo{
					Key:    p,
					Comment: c,
				})
			}
			fmt.Println("Please input Comment:")
			comment := util.ReadLine()
			fmt.Println("Are you sure to add the record? yes/No")
			opt := "no"
			opt = util.ReadLine()
			if opt != "yes" {
				break
			}
			if id, ok := dam.KeyAdd(title, key, comment); ok {
				fmt.Printf("Finished, ID:[%d]\n", id)
			} else {
				fmt.Println("Failure")
			}
		case 3:
			fmt.Println("Please input id:")
			id := util.ReadUInt32()
			fmt.Printf("Are you sure to delete the [%d] record? yes/No\n", id)
			opt := "no"
			opt = util.ReadLine()
			if opt != "yes" {
				break
			}
			if dam.KeyDelete(id) {
				fmt.Println("Finished")
			} else {
				fmt.Println("Failure")
			}
		case 4:
			fmt.Println("Please input id:")
			id := util.ReadUInt32()
			key := dam.KeyGet(id)
		I:
			for {
				fmt.Println("1. Title")
				fmt.Println("2. Key")
				fmt.Println("3. Comment")
				fmt.Println("7. # Modify")
				fmt.Println("0. # Exit")
				fmt.Println("Please enter the field you want to modify:")
				opt := util.ReadInt()
				switch opt {
				case 1:
					fmt.Println("Please input new title:")
					key.Title = util.ReadLine()
				case 2:
					key.Key = make([]dam.KeyInfo, 0)
					fmt.Println("Please input Key ('.exit' to exit):")
					fmt.Println("'EOF' to finish input")
					K2:
					for {
						fmt.Println("Key:")
						p := ""
						for {
							pt := util.ReadLine()
							if pt == ".exit" {
								break K2
							} else if pt == "EOF" {
								break
							}
							if len(p) != 0 {
								p += "\n"
							}
							p += pt
						}
						fmt.Println("Comment:")
						c := util.ReadLine()
						if c == ".exit" {
							break
						}
						key.Key = append(key.Key, dam.KeyInfo{
							Key:    p,
							Comment: c,
						})
					}
				case 3:
					fmt.Println("Please input new comment:")
					key.Comment = util.ReadLine()
				case 7:
					fmt.Printf("# ID:[%d] Title:[%s]\n",
						key.Id, key.Title)
					for _, val := range key.Key {
						fmt.Printf("# Key:[\n%s\n] Comment:[%s]\n", val.Key, val.Comment)
					}
					fmt.Printf("# Comment:[%s]\n", key.Comment)
					fmt.Println("Are you sure to modify the record? yes/No")
					opt := "no"
					opt = util.ReadLine()
					if opt != "yes" {
						break
					}
					dam.KeyUpdate(key)
				default:
					break I
				}
			}
		case 5:
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
				field = "key"
			case 4:
				field = "comment"
			case 0:
				break L
			default:
				field = "key"
			}
			fmt.Println("Please input regexp:")
			reg := util.ReadLine()
			keys := dam.KeyGetByRegexp(field, reg)
			fmt.Println("##########################################")
			for k, v := range keys {
				if k != 0 {
					fmt.Println()
				}
				fmt.Printf("# ID:[%d] Title:[%s]\n",
					v.Id, v.Title)
				for _, val := range v.Key {
					fmt.Printf("# Key:[\n%s\n] Comment:[%s]\n", val.Key, val.Comment)
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
