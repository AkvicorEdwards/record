package dam

import (
	"encoding/json"
	"record/util"
)

type Inc struct {
	Name string // 表名
	Val  uint32 // 自增值
}

type SecretQuestion struct {
	Quest string `json:"quest"`
	Answer string `json:"answer"`
}

type Account struct {
	Id uint32 `json:"id"`
	Title string `json:"title"`
	Account string `json:"account"`
	Password string `json:"password"`
	SecretQuestion []SecretQuestion `json:"secret_question"`
	TwoFactor []string `json:"two_factor"`
	Comment string `json:"comment"`
}

func (a *Account) Transfer() AccountDB {
	res := AccountDB{
		Id:             a.Id,
		Title:          util.Encrypt(a.Title),
		Account:        util.Encrypt(a.Account),
		Password:       util.Encrypt(a.Password),
		SecretQuestion: "",
		TwoFactor:      "",
		Comment:        util.Encrypt(a.Comment),
	}
	data, _ := json.Marshal(a.SecretQuestion)
	res.SecretQuestion = util.Encrypt(string(data))
	data, _ = json.Marshal(a.TwoFactor)
	res.TwoFactor = util.Encrypt(string(data))
	return res
}

type AccountDB struct {
	Id uint32
	Title string
	Account string
	Password string
	SecretQuestion string
	TwoFactor string
	Comment string
}

func (a *AccountDB) Transfer() Account {
	res := Account{
		Id:             a.Id,
		Title:          util.Decrypt(a.Title),
		Account:        util.Decrypt(a.Account),
		Password:       util.Decrypt(a.Password),
		SecretQuestion: make([]SecretQuestion, 0),
		TwoFactor:      make([]string, 0),
		Comment:        util.Decrypt(a.Comment),
	}
	_ = json.Unmarshal([]byte(util.Decrypt(a.SecretQuestion)), &res.SecretQuestion)
	_ = json.Unmarshal([]byte(util.Decrypt(a.TwoFactor)), &res.TwoFactor)
	return res
}

type PortInfo struct {
	Port uint32 `json:"port"`
	Comment string `json:"comment"`
}

type Port struct {
	Id uint32 `json:"id"`
	Title string `json:"title"`
	Port []PortInfo `json:"port"`
	Platform string `json:"platform"`
	Comment string `json:"comment"`
}

func (p *Port) Transfer() PortDB {
	port := PortDB{
		Id:       p.Id,
		Title:    util.Encrypt(p.Title),
		Port:     "",
		Platform: util.Encrypt(p.Platform),
		Comment:  util.Encrypt(p.Comment),
	}
	data, _ := json.Marshal(p.Port)
	port.Port = util.Encrypt(string(data))
	return port
}

type PortDB struct {
	Id uint32
	Title string
	Port string
	Platform string
	Comment string
}

func (p *PortDB) Transfer() Port {
	port := Port{
		Id:       p.Id,
		Title:    util.Decrypt(p.Title),
		Port:     make([]PortInfo, 0),
		Platform: util.Decrypt(p.Platform),
		Comment:  util.Decrypt(p.Comment),
	}
	_ = json.Unmarshal([]byte(util.Decrypt(p.Port)), &port.Port)
	return port
}