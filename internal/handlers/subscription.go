package handlers

import (
	"net/http"
	"time"

	"github.com/JayBoba/EM_test/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubsInput struct {
	ServiceName string  `json:"service_name" binding:"required"`
	Price       int     `json:"price" binding:"required,min=0"`
	UserID      string  `json:"user_id" binding:"required,uuid"`
	StartDate   string  `json:"start_date" binding:"required"`
	EndDate     *string `json:"end_date,omitempty"`
}

func parseMonthYear(s string) (time.Time, error) {
	return time.Parse("01-2001", s)
}

func CreateSubscription(c *gin.Context) {
	var inputJSON SubsInput
	if err := c.ShouldBindJSON(&inputJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := parseMonthYear(inputJSON.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format, use MM-YYYY"})
		return
	}

	var endDate *time.Time
	if inputJSON.EndDate != nil {
		parsed, err := parseMonthYear(*inputJSON.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format, use MM-YYYY"})
			return
		}
		endDate = &parsed
	}

	sub := models.Subscription{
		ServiceName: inputJSON.ServiceName,
		Price:       inputJSON.Price,
		UserID:      uuid.MustParse(inputJSON.UserID),
		StartDate:   startDate,
		EndDate:     endDate,
	}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&sub).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sub)
}
