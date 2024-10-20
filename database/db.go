package database

//import packages
import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connect to the Database
func ConnectDB() {
	//dsn := "root:root@tcp(127.0.0.1:8889)/gonotes?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:root@tcp(127.0.0.1:8889)/gonotes?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	//attempt to establish DB connection
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//if there is an error
	if err != nil {
		fmt.Println("Failed to connect to database")
		return
	}
	fmt.Println("Database connected successfully")
}
