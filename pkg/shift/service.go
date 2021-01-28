package shift

import (
	"net/http"

	"github.com/anhthii/staffany-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Route(g *gin.RouterGroup) {
	shifts := g.Group("/shifts")
	shifts.PUT("/:id", s.UpdateShift)
	shifts.DELETE("/:id", s.DeleteShift)
}

type ShiftParams struct {
	ID     uint `json:"id"`
	DateID uint `json:"date_id"`
	UserID uint `json:"user_id"`

	QuarterStart uint `json:"quarter_start"`
	NumQuarter   uint `json:"num_quarter"`

	Title       string `json:"title"`
	Description string `json:"description"`
}

func (s *Service) UpdateShift(c *gin.Context) {
	var params ShiftParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shift := Shift{
		ID:           params.ID,
		DateID:       params.DateID,
		UserID:       params.UserID,
		QuarterStart: params.QuarterStart,
		NumQuarter:   params.NumQuarter,
		Title:        params.Title,
		Description:  params.Description,
	}

	err := s.repo.Update(&shift)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
	return
}

func (s *Service) DeleteShift(c *gin.Context) {
	shiftIDStr := c.Param("id")
	shiftID := utils.StringToUint(shiftIDStr)
	err := s.repo.DeleteByID(shiftID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
	return
}
