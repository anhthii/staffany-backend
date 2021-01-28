package main

import (
	"github.com/anhthii/anystaff-backend/pkg/date"
	"github.com/anhthii/anystaff-backend/pkg/shift"
	"github.com/anhthii/anystaff-backend/pkg/user"
	"github.com/anhthii/anystaff-backend/pkg/week"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/anystaff?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&user.User{}, &week.Week{}, &date.Date{}, &shift.Shift{})
	// timeStr := "2021-01-28"
	r := gin.Default()

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

	r.Run()
	// t, _ := time.Parse(layoutISO, timeStr)
	// wyear, week := isoweek.FromDate(t.Year(), t.Month(), t.Day())
	// fmt.Println(wyear, week)
}
