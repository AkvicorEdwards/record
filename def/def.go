package def

const Password = "12345678"

const EncryptKey = "1234567812345678"

const DatabaseFileName = "record.db"

func CheckEncryptKey() bool {
	return len(EncryptKey) == 16 || len(EncryptKey) == 24 || len(EncryptKey) == 32
}
