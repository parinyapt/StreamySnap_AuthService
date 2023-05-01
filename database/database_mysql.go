package database

import (
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	originalmysql "github.com/go-sql-driver/mysql"

	"github.com/parinyapt/StreamySnap_AuthService/logger"
	"github.com/parinyapt/StreamySnap_AuthService/model/database"
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
		logger.Fatal("Failed to connect mysql database", logger.Field("error", err))
	}

	// database.AutoMigrate(&modelDatabase.data1{})
	err = database.AutoMigrate(
		&modelDatabase.Client{}, 
		&modelDatabase.ClientService{}, 
		&modelDatabase.Account{}, 
		&modelDatabase.AccountOAuth{}, 
		&modelDatabase.OAuthProvider{}, 
		&modelDatabase.AuthPageSession{}, 
		&modelDatabase.AuthTemporaryToken{}, 
		&modelDatabase.AuthHistory{},
	)
	if err != nil {
			logger.Error("Failed to AutoMigrate database", logger.Field("error", err))
	}

	DB = database

	logger.Info("Initialize MySQL Database Success")
}