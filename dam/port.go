package dam

import (
	"fmt"
	"log"
	"regexp"
)

func PortAdd(title string, port []PortInfo, platform, comment string) (uint32, bool) {
	if !Connected {
		Connect()
	}
	lockPort.Lock()
	defer lockPort.Unlock()
	p := Port{
		Id:       GetInc("port") + 1,
		Title:    title,
		Port:     port,
		Platform: platform,
		Comment:  comment,
	}
	pdb := p.Transfer()
	if err := db.Table("port").Create(&pdb).Error; err != nil {
		log.Println(err)
		return 0, false
	}

	UpdateInc("port", pdb.Id)

	return pdb.Id, true
}

func PortDelete(id uint32) bool {
	if !Connected {
		Connect()
	}
	lockPort.Lock()
	defer lockPort.Unlock()

	res := db.Table("port").Delete(&Port{}, id)
	if res.Error != nil {
		log.Println(res.Error)
		return false
	}
	if res.RowsAffected == 0 {
		return false
	}
	return true
}

func PortUpdate(port Port) bool {
	if !Connected {
		Connect()
	}
	lockPort.Lock()
	defer lockPort.Unlock()

	p := port.Transfer()

	res := db.Table("port").Where("id=?", p.Id).Updates(map[string]interface{}{
		"title":     p.Title,
		"port":     p.Port,
		"platform": p.Platform,
		"comment":  p.Comment,
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

func PortGet(id uint32) Port {
	if !Connected {
		Connect()
	}

	port := PortDB{}

	db.Table("port").Where("id=?", id).First(&port)

	return port.Transfer()
}

func PortGetByRegexp(field, reg string) []Port {
	if !Connected {
		Connect()
	}
	ports := PortGetAll()
	res := make([]Port, 0)
	var matched bool
	var err error
	for _, v := range ports {
		switch field {
		case "id":
			matched, err = regexp.MatchString(reg, fmt.Sprint(v.Id))
		case "title":
			matched, err = regexp.MatchString(reg, v.Title)
		case "port":
			for _, val := range v.Port {
				matched, err = regexp.MatchString(reg, fmt.Sprint(val.Port))
				if err != nil || matched {
					break
				}
			}
		case "platform":
			matched, err = regexp.MatchString(reg, v.Platform)
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

func PortGetAll() []Port {
	if !Connected {
		Connect()
	}

	ports := make([]PortDB, 0)

	db.Table("port").Find(&ports)

	res := make([]Port, len(ports))

	for k, v := range ports {
		res[k] = v.Transfer()
	}

	return res
}

