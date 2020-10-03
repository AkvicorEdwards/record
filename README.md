# Record

- [x] Manage Password
- [x] Manage Port Allocation
- [x] Visualize Heartbeat Data

## Development environment

- Windows 10
- Go 1.14
- Sqlite3

## Usage

```shell script
vim ./def/def.go
# change Password and EncryptKey
## Password: Enter the program
## EncryptKey: Encrypted data
```

```shell script
go build
# init database
./record init
./record
```
