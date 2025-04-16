package sql

import (
	"fmt"
	"log"
	"mikhael-project-go/config"
	"mikhael-project-go/pkg/drivers/common"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenDbConnection() *gorm.DB {

	v := common.NewDbInfo(
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)

	// log.Println("Hasil : -> ", v)
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", v.User, v.Password, v.Host, v.Port, v.Name)
	log.Println("Connection Db : ", connectionString)

	// dsn := "root:adm1234@tcp(127.0.0.1:3306)/go_cashier_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect Databases : " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get database object: " + err.Error())
	}

	// Ping untuk cek koneksi
	if err := sqlDB.Ping(); err != nil {
		panic("Database not reachable: " + err.Error())
	}
	return db
}
