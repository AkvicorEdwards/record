package dam

import (
	"fmt"
	"log"
	"regexp"
)

func HeartbeatAddRecord(date, data string) (uint32, bool) {
	if !Connected {
		Connect()
	}
	lockHeartbeat.Lock()
	defer lockHeartbeat.Unlock()

	h := Heartbeat{
		Id:   GetInc("heartbeat") + 1,
		Date: date,
		Data: data,
	}

	h.Encrypt()

	if err := db.Table("heartbeat").Create(&h).Error; err != nil {
		log.Println(err)
		return 0, false
	}

	UpdateInc("heartbeat", h.Id)

	return h.Id, true
}

func HeartbeatDelete(id uint32) bool {
	if !Connected {
		Connect()
	}
	lockHeartbeat.Lock()
	defer lockHeartbeat.Unlock()

	res := db.Table("heartbeat").Delete(&Heartbeat{}, id)
	if res.Error != nil {
		log.Println(res.Error)
		return false
	}
	if res.RowsAffected == 0 {
		return false
	}
	return true
}

func HeartbeatUpdate(heartbeat Heartbeat) bool {
	if !Connected {
		Connect()
	}
	lockHeartbeat.Lock()
	defer lockHeartbeat.Unlock()

	heartbeat.Encrypt()

	res := db.Table("heartbeat").Where("id=?", heartbeat.Id).Updates(map[string]interface{}{
		"date":     heartbeat.Date,
		"data":     heartbeat.Data,
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

func HeartbeatGetByID(id uint32) Heartbeat {
	if !Connected {
		Connect()
	}

	heartbeat := Heartbeat{}

	db.Table("heartbeat").Where("id=?", id).First(&heartbeat)

	heartbeat.Decrypt()

	return heartbeat
}

func HeartbeatGetByDate(date string) Heartbeat {
	if !Connected {
		Connect()
	}

	heartbeat := Heartbeat{}

	db.Table("heartbeat").Where("date=?", date).First(&heartbeat)

	heartbeat.Decrypt()

	return heartbeat
}

func HeartbeatGetAll() []Heartbeat {
	if !Connected {
		Connect()
	}

	heartbeat := make([]Heartbeat, 0)

	db.Table("heartbeat").Find(&heartbeat)

	for k, _ := range heartbeat {
		heartbeat[k].Decrypt()
	}

	return heartbeat
}

func HeartbeatGetByRegexp(field, reg string) []Heartbeat {
	if !Connected {
		Connect()
	}
	ports := HeartbeatGetAll()
	res := make([]Heartbeat, 0)
	var matched bool
	var err error
	for _, v := range ports {
		switch field {
		case "id":
			matched, err = regexp.MatchString(reg, fmt.Sprint(v.Id))
		case "date":
			matched, err = regexp.MatchString(reg, v.Date)
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

	for k, _ := range res {
		res[k].Decrypt()
	}

	return res
}

