package def

// Login Password
const Password = "12345678"

// Use for encrypt data
const EncryptKey = "123456781234567812345678"

const DatabaseFileName = "record.db"

func CheckEncryptKey() bool {
	return len(EncryptKey) == 16 || len(EncryptKey) == 24 || len(EncryptKey) == 32
}

const Protocol = "http"
const Port = 8080
