## Initial

go mod init github.com/akhil/go-bookstore

go get "github.com/jinzhu/gorm"
go get "github.com/jinzhu/gorm/dialects/mysql"
go get "github.com/gorilla/mux"

## Build

cd go-bookstore/cmd/main
go build
go run main.go
