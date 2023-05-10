package controllers

import (
	"MiniProject/config"
	"MiniProject/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateHistoryController will create a new history data
func CreateHistoryController(c echo.Context) error {
	var histories models.Histories
	if err := c.Bind(&histories); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to create history data",
			"error":  err.Error(),
		})
	}

	if err := config.DB.Create(&histories).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to create history data",
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success create history data",
		"history": histories,
	})
}

// GetHistoriesController will get all history data
func GetHistoriesController(c echo.Context) error {
	var histories []models.Histories
	if err := config.DB.Find(&histories).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to get history data",
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success get history data",
		"historys": histories,
	})
}

// GetHistoriesByPaymentController will get all history data by payment id
func GetHistoriesByPaymentController(c echo.Context) error {
	paymentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "invalid parameter",
			"error":  err.Error(),
		})
	}

	var histories []models.Histories
	if err := config.DB.Where("payment_id = ?", paymentID).Find(&histories).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to get history data",
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success get history data",
		"history": histories,
	})
}

// UpdateHistoryController will update a history data by id
func UpdateHistoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "invalid parameter",
			"error":  err.Error(),
		})
	}
	var histories models.Histories
	if err := config.DB.First(&histories, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to update history data",
			"error":  err.Error(),
		})
	}

	if err := c.Bind(&histories); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to update history data",
			"error":  err.Error(),
		})
	}

	if err := config.DB.Save(&histories).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to update history data",
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success update history data",
		"history": histories,
	})
}

// DeleteHistoryController will delete a history data by id
func DeleteHistoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "invalid parameter",
			"error":  err.Error(),
		})
	}

	var histories models.Histories
	if err := config.DB.First(&histories, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to delete history data",
			"error":  err.Error(),
		})
	}

	if err := config.DB.Delete(&histories).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to delete history data",
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success delete history data",
		"history": histories,
	})
}
