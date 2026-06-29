package handlers

import (
	"errors"
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
	return time.Parse("01-2006", s)
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

func GetSubscription(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
	}

	db := c.MustGet("db").(*gorm.DB)
	var sub models.Subscription
	if err := db.First(&sub, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sub)
}

func UpdateSubscription(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
	}
	var input SubsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	startDate, err := parseMonthYear(input.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format, use MM-YYYY"})
		return
	}

	var endDate *time.Time
	if input.EndDate != nil {
		parsed, err := parseMonthYear(*input.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format, use MM-YYYY"})
			return
		}
		endDate = &parsed
	}
	db := c.MustGet("db").(*gorm.DB)
	var sub models.Subscription
	if err := db.First(&sub, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	sub.ServiceName = input.ServiceName
	sub.Price = input.Price
	sub.UserID = uuid.MustParse(input.UserID)
	sub.StartDate = startDate
	sub.EndDate = endDate

	if err := db.Save(&sub).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sub)
}

func DeleteSubscription(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Delete(&models.Subscription{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func ListSubscriptions(c *gin.Context) {
	var query struct {
		UserID      string `form:"user_id"`
		ServiceName string `form:"service_name"`
	}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.MustGet("db").(*gorm.DB).Model(&models.Subscription{})
	if query.UserID != "" {
		db = db.Where("user_id = ?", query.UserID)
	}
	if query.ServiceName != "" {
		db = db.Where("service_name = ?", query.ServiceName)
	}
	var subs []models.Subscription
	if err := db.Find(&subs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subs)
}

func AggregateCost(c *gin.Context) {
	var query struct {
		UserID      string `form:"user_id"`
		ServiceName string `form:"service_name"`
		StartDate   string `form:"start_date" binding:"required"`
		EndDate     string `form:"end_date" binding:"required"`
	}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	start, err := parseMonthYear(query.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format"})
		return
	}
	end, err := parseMonthYear(query.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format"})
		return
	}
	end = end.AddDate(0, 1, -1)
	db := c.MustGet("db").(*gorm.DB).Model(&models.Subscription{})
	db = db.Where("start_date <= ? AND (end_date IS NULL OR end_date >= ?)", end, start)

	if query.UserID != "" {
		db = db.Where("user_id = ?", query.UserID)
	}
	if query.ServiceName != "" {
		db = db.Where("service_name = ?", query.ServiceName)
	}
	var total int64
	db.Select("COALESCE(SUM(price), 0)").Scan(&total)
	c.JSON(http.StatusOK, gin.H{"total_cost": total})

}
