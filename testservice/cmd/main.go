package main

import (
	"github.com/fasibio/autogql/testservice"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	d := postgres.New(postgres.Config{
		DSN: "host=localhost user=postgres password=postgres dbname=autogql port=5432 sslmode=disable",
	})
	// m := mysql.Open("root:password@tcp(127.0.0.1:3306)/autogql")
	dbCon, err := gorm.Open(d, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	dbCon = dbCon.Debug()
	testservice.StartServer(dbCon)
}
