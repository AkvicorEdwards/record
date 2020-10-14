package maintenance

import (
	"log"
	"os"
	"record/dam"
	"record/def"
)

const (
	sql = `
create table inc
(
    name text,
    val  integer
);

` + `
create table account
(
    id              integer
        constraint account_pk
            primary key,
    title           text,
    account         text,
    password        text,
    secret_question text,
    two_factor      text,
    comment         text
);
` + `
create table port
(
    id       integer
        constraint port_pk
            primary key,
    title    text,
    port     text,
    platform text,
    comment  text
);
` + `
create table heartbeat
(
    id   integer
        constraint heartbeat_pk
            primary key,
    date text,
    data text
);

create unique index heartbeat_date_uindex
    on heartbeat (date);
` + `
create table key
(
    id      integer
        constraint key_pk
            primary key,
    title   text,
    key     text,
    comment text
);

create unique index key_title_uindex
    on key (title);
` + `

INSERT INTO inc (name, val) VALUES ('account', 0);
INSERT INTO inc (name, val) VALUES ('port', 0);
INSERT INTO inc (name, val) VALUES ('heartbeat', 0);
INSERT INTO inc (name, val) VALUES ('key', 0);
`
)

func InitDatabase() {
	if !IsFile(def.DatabaseFileName) {
		log.Println("record.db do not exist")
		os.Exit(-1)
	}

	err := dam.Exec(sql).Error
	if err != nil {
		log.Println(err)
		os.Exit(-2)
	}

	log.Println("Finished")

	os.Exit(0)
}

func Exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	if !Exists(path) {
		return false
	}
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}