package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToMySQL() (*gorm.DB, error) {

	// var databaseName string = os.Getenv("DATABASE_NAME")
	// var clusterName string = os.Getenv("DATABASE_HOST")

	var user string = "root"
	var password string = "root"
	var databaseName string = "dhp"
	var clusterName string = "127.0.0.1"
	var port string = "3306"

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=UTC&timeout=30s", user, password, clusterName, port, databaseName)

	fmt.Printf("\nlog:DSN: %v \n", DSN)

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database!")
	}

	fmt.Println("")
	fmt.Printf("Database: %s:%s => %s\n", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), databaseName)
	fmt.Println("")

	return db, err
}
