package week

import (
	"net/http"
	"time"

	"github.com/anhthii/staffany-backend/pkg/date"
	"github.com/anhthii/staffany-backend/pkg/shift"
	"github.com/anhthii/staffany-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Service struct {
	weekRepo  Repository
	dateRepo  date.Repository
	shiftRepo shift.Repository
}

func NewService(weekRepo Repository, dateRepo date.Repository, shiftRepo shift.Repository) *Service {
	return &Service{
		weekRepo:  weekRepo,
		dateRepo:  dateRepo,
		shiftRepo: shiftRepo}
}

func (s *Service) Route(g *gin.RouterGroup) {
	weeks := g.Group("/weeks")
	weeks.POST("/", s.CreateWeek)
	weeks.GET("/current_week/:user_id", s.GetCurrentWeek)
	weeks.POST("/:id/shifts", s.CreateShift)
}

type WeekParams struct {
	UserID uint `json:"user_id"`
}

func (s *Service) GetCurrentWeek(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID := utils.StringToUint(userIDStr)
	week, err := s.weekRepo.GetCurrentWeek(userID)

	if err == RecordNotFound {
		now := utils.GetDateString(time.Now())
		// create current week if not exist
		_, err := s.weekRepo.Create(now, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		week, err := s.weekRepo.GetCurrentWeek(userID)
		c.JSON(http.StatusOK, week)
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, week)
	return
}

type ShiftParams struct {
	DateID       uint   `json:"date_id"`
	Date         string `json:"date"`
	UserID       uint   `json:"user_id"`
	QuarterStart uint   `json:"quarter_start"`
	NumQuarter   uint   `json:"num_quarter"`

	Title       string `json:"title"`
	Description string `json:"description"`
}

func (s *Service) CreateShift(c *gin.Context) {
	weekIDStr := c.Param("id")
	weekID := utils.StringToUint(weekIDStr)

	var params ShiftParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dateID uint
	var retDate *date.Date

	// the date contains the shift does not exist
	// so we have to create the date first
	if params.DateID == 0 {
		date := date.Date{
			Date:   utils.DateStringToInt(params.Date),
			WeekID: weekID,
			UserID: params.UserID,
		}

		d, err := s.dateRepo.Create(&date)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		retDate = d

	} else {
		dateID = params.DateID
	}

	// if date already exists

	shift := shift.Shift{
		DateID:       dateID,
		UserID:       params.UserID,
		QuarterStart: params.QuarterStart,
		NumQuarter:   params.NumQuarter,
		Title:        params.Title,
		Description:  params.Description,
	}

	shiftID, err := s.shiftRepo.Create(&shift)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"shift_id": shiftID, "date": retDate})
	return
}

func (s *Service) CreateWeek(c *gin.Context) {
	// userIDStr := c.Param("userID")
	// userID := utils.StringToUint(userIDStr)

	return
}
