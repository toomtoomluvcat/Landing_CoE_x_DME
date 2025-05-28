package connect

import (
	"fmt"
	"log"
	"os"
	// "scavenger_hunt/app/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Pg() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Fail to load env: %v", err)
	}
	Host := os.Getenv("Host")
	Database := os.Getenv("DATABASE")
	User := os.Getenv("USER")
	Password := os.Getenv("PASSWORD")
	Port := os.Getenv("PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", Host, User, Password, Database, Port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Fail to connect database: %v", err)
	}

	// DB.AutoMigrate(&model.Group{}, &model.Participant{}, &model.Join{})
}
