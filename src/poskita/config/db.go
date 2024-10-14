package config

import (
	"database/sql"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var e error
// Connect db
func Connect() *gorm.DB {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	// dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := username + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=require"
	dsn := "host=" + host + " port=" + port + " user="+username + " password=" + password + " dbname=" + dbname + " sslmode=require"

	// fmt.Println(dsn)
	// db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })
	sqlDB, err := sql.Open("pgx", dsn)
	if err!=nil {
		log.Fatal("cannot open connection " + e.Error())
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB, 
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err!=nil {
		log.Fatal("cannot connect db" + e.Error())
	}
	return db
}
// Close db connection
func Close(db *gorm.DB) {
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
}