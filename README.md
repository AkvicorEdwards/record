# Record

- [x] Manage Password
- [x] Manage Port Allocation
- [x] Visualize Heartbeat Data
- [x] Manage Key

## Development environment

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
./record -init
./record
```
