package port

import (
	"fmt"
	"record/dam"
	"record/util"
)

const Menu = `
##########################
#          Port          #
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
# 3. Port                #
# 4. Platform            #
# 5. Comment             #
# 0. Exit                #
##########################
`

func Port() {
	fmt.Print(Menu)
	for {
		L:
		switch util.ReadInt() {
		case 1:
			ports := dam.PortGetAll()
			fmt.Println("##########################################")
			for k, v := range ports {
				if k != 0 {
					fmt.Println()
				}
				fmt.Printf("# ID:[%d] Title:[%s] Platform:[%s] Comment:[%s]\n",
					v.Id, v.Title, v.Platform, v.Comment)
			}
			fmt.Println("##########################################")
		case 2:
			ports := dam.PortGetAll()
			fmt.Println("##########################################")
			for k, v := range ports {
				if k != 0 {
					fmt.Println()
				}
				fmt.Printf("# ID:[%d] Title:[%s] Platform:[%s]\n",
					v.Id, v.Title, v.Platform)
				for _, val := range v.Port {
					fmt.Printf("# Port:[%d] Comment:[%s]\n", val.Port, val.Comment)
				}
				fmt.Printf("# Comment:[%s]\n", v.Comment)
			}
			fmt.Println("##########################################")
		case 3:
			fmt.Println("Please input Title:")
			title := util.ReadLine()
			fmt.Println("Please input Port ('.exit' to exit):")
			port := make([]dam.PortInfo, 0)
			for {
				fmt.Println("Port:")
				p := util.ReadUInt32()
				if p == 0 {
					break
				}
				fmt.Println("Comment:")
				c := util.ReadLine()
				if c == ".exit" {
					break
				}
				port = append(port, dam.PortInfo{
					Port:    p,
					Comment: c,
				})
			}
			fmt.Println("Please input Platform:")
			platform := util.ReadLine()
			fmt.Println("Please input Comment:")
			comment := util.ReadLine()
			fmt.Println("Are you sure to add the record? yes/No")
			opt := "no"
			opt = util.ReadLine()
			if opt != "yes" {
				break
			}
			if id, ok := dam.PortAdd(title, port, platform, comment); ok {
				fmt.Printf("Finished, ID:[%d]\n", id)
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
			if dam.PortDelete(id) {
				fmt.Println("Finished")
			} else {
				fmt.Println("Failure")
			}
		case 5:
			fmt.Println("Please input id:")
			id := util.ReadUInt32()
			port := dam.PortGet(id)
			I:
			for {
				fmt.Println("1. Title")
				fmt.Println("2. Port")
				fmt.Println("3. Platform")
				fmt.Println("4. Comment")
				fmt.Println("7. # Modify")
				fmt.Println("0. # Exit")
				fmt.Println("Please enter the field you want to modify:")
				opt := util.ReadInt()
				switch opt {
				case 1:
					fmt.Println("Please input new title:")
					port.Title = util.ReadLine()
				case 2:
					port.Port = make([]dam.PortInfo, 0)
					fmt.Println("Please input Port ('.exit' to exit):")
					for {
						fmt.Println("Port:")
						p := util.ReadUInt32()
						if p == 0 {
							break
						}
						fmt.Println("Comment:")
						c := util.ReadLine()
						if c == ".exit" {
							break
						}
						port.Port = append(port.Port, dam.PortInfo{
							Port:    p,
							Comment: c,
						})
					}
				case 3:
					fmt.Println("Please input new platform:")
					port.Platform = util.ReadLine()
				case 4:
					fmt.Println("Please input new comment:")
					port.Comment = util.ReadLine()
				case 7:
					fmt.Printf("# ID:[%d] Title:[%s] Platform:[%s]\n",
						port.Id, port.Title, port.Platform)
					for _, val := range port.Port {
						fmt.Printf("# Port:[%d] Comment:[%s]\n", val.Port, val.Comment)
					}
					fmt.Printf("# Comment:[%s]\n", port.Comment)
					fmt.Println("Are you sure to modify the record? yes/No")
					opt := "no"
					opt = util.ReadLine()
					if opt != "yes" {
						break
					}
					dam.PortUpdate(port)
				default:
					break I
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
				field = "port"
			case 4:
				field = "platform"
			case 5:
				field = "comment"
			case 0:
				break L
			default:
				field = "port"
			}
			fmt.Println("Please input regexp:")
			reg := util.ReadLine()
			ports := dam.PortGetByRegexp(field, reg)
			fmt.Println("##########################################")
			for k, v := range ports {
				if k != 0 {
					fmt.Println()
				}
				fmt.Printf("# ID:[%d] Title:[%s] Platform:[%s]\n",
					v.Id, v.Title, v.Platform)
				for _, val := range v.Port {
					fmt.Printf("# Port:[%d] Comment:[%s]\n", val.Port, val.Comment)
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
