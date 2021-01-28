package shift

import (
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

func (s *Service) UpdateShift(c *gin.Context) {
	// userIDStr := c.Param("userID")
	// userID := utils.StringToUint(userIDStr)

	return
}

func (s *Service) DeleteShift(c *gin.Context) {
	// userIDStr := c.Param("userID")
	// userID := utils.StringToUint(userIDStr)

	return
}
