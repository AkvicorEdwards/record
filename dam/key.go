package dam

import (
	"fmt"
	"log"
	"regexp"
)

func KeyAdd(title string, key []KeyInfo, comment string) (uint32, bool) {
	if !Connected {
		Connect()
	}
	lockKey.Lock()
	defer lockKey.Unlock()
	p := Key{
		Id:       GetInc("key") + 1,
		Title:    title,
		Key:     	key,
		Comment:  comment,
	}
	pdb := p.Transfer()
	if err := db.Table("key").Create(&pdb).Error; err != nil {
		log.Println(err)
		return 0, false
	}

	UpdateInc("key", pdb.Id)

	return pdb.Id, true
}

func KeyDelete(id uint32) bool {
	if !Connected {
		Connect()
	}
	lockKey.Lock()
	defer lockKey.Unlock()

	res := db.Table("key").Delete(&Key{}, id)
	if res.Error != nil {
		log.Println(res.Error)
		return false
	}
	if res.RowsAffected == 0 {
		return false
	}
	return true
}

func KeyUpdate(key Key) bool {
	if !Connected {
		Connect()
	}
	lockKey.Lock()
	defer lockKey.Unlock()

	p := key.Transfer()

	res := db.Table("key").Where("id=?", p.Id).Updates(map[string]interface{}{
		"title":     p.Title,
		"key":     p.Key,
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

func KeyGet(id uint32) Key {
	if !Connected {
		Connect()
	}

	key := KeyDB{}

	db.Table("key").Where("id=?", id).First(&key)

	return key.Transfer()
}

func KeyGetByRegexp(field, reg string) []Key {
	if !Connected {
		Connect()
	}
	ports := KeyGetAll()
	res := make([]Key, 0)
	var matched bool
	var err error
	for _, v := range ports {
		switch field {
		case "id":
			matched, err = regexp.MatchString(reg, fmt.Sprint(v.Id))
		case "title":
			matched, err = regexp.MatchString(reg, v.Title)
		case "key":
			for _, val := range v.Key {
				matched, err = regexp.MatchString(reg, fmt.Sprint(val.Key))
				if err != nil || matched {
					break
				}
			}
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

func KeyGetAll() []Key {
	if !Connected {
		Connect()
	}

	ports := make([]KeyDB, 0)

	db.Table("key").Find(&ports)

	res := make([]Key, len(ports))

	for k, v := range ports {
		res[k] = v.Transfer()
	}

	return res
}

