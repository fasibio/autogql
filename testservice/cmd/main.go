package main

import (
	"github.com/fasibio/autogql/testservice"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	dbCon, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbCon = dbCon.Debug()
	testservice.StartServer(dbCon)
}
