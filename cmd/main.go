package main

import (
	"log"
	"os"

	"github.com/anhthii/staffany-backend/pkg/date"
	"github.com/anhthii/staffany-backend/pkg/shift"
	"github.com/anhthii/staffany-backend/pkg/user"
	"github.com/anhthii/staffany-backend/pkg/week"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// bad, never do this in production
	db.AutoMigrate(&user.User{}, &week.Week{}, &date.Date{}, &shift.Shift{})
	r := gin.Default()
	r.Use(cors.Default())

	userRepository := user.NewRepository(db)
	weekRepository := week.NewRepository(db)
	dateRepository := date.NewRepository(db)
	shiftRepository := shift.NewRepository(db)

	userservice := user.NewService(userRepository)
	weekService := week.NewService(weekRepository, dateRepository, shiftRepository)
	shiftService := shift.NewService(shiftRepository)

	v1 := r.Group("/api/v1")

	userservice.Route(v1)
	weekService.Route(v1)
	shiftService.Route(v1)

	r.Run(":8080")
}
