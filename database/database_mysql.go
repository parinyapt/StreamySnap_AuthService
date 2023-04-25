package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	originalmysql "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

func initializeConnectMySQL() {

	dsn := originalmysql.Config{
		User:      os.Getenv("DATABASE_MYSQL_USERNAME"),
		Passwd:    os.Getenv("DATABASE_MYSQL_PASSWORD"),
		Net:       "tcp",
		Addr:      os.Getenv("DATABASE_MYSQL_HOST"),
		DBName:    os.Getenv("DATABASE_MYSQL_DBNAME"),
		AllowNativePasswords: true,
		ParseTime: true,
		Loc:       time.Local,
	}
	database, err := gorm.Open(mysql.Open(dsn.FormatDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("[Error]->Failed to connect database : %s", err)
	}

	// database.AutoMigrate(&models.data1{})
	// err = database.AutoMigrate(&models.Data1{}, &models.Data2{}, &models.Data3{})
	// if err != nil {
	// 		log.Fatalf("[Error]->Failed to AutoMigrate database : %s", err)
	// }

	DB = database

}