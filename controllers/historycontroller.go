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
	var history models.History
	if err := c.Bind(&history); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to create history data",
			"error":  err.Error(),
		})
	}

	if err := config.DB.Create(&history).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to create history data",
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success create history data",
		"history": history,
	})
}

// GetHistoriesController will get all history data
func GetHistorysController(c echo.Context) error {
	var historys []models.History
	if err := config.DB.Find(&historys).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to get history data",
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success get history data",
		"historys": historys,
	})
}

// GetHistoryController will get a single history data by id
func GetHistoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "invalid parameter",
			"error":  err.Error(),
		})
	}

	var history models.History
	if err := config.DB.First(&history, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to get history data",
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success get history data",
		"history": history,
	})
}

// GetHistoriesByUserController will get all history data by user id
func GetHistoryByUserController(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "invalid parameter",
			"error":  err.Error(),
		})
	}

	var history []models.History
	if err := config.DB.Where("user_id = ?", userID).Find(&history).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to get history data",
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success get history data",
		"history": history,
	})
}

// GetHistoriesByPaymentController will get all history data by payment id
func GetHistoryByPaymentController(c echo.Context) error {
	paymentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "invalid parameter",
			"error":  err.Error(),
		})
	}

	var history []models.History
	if err := config.DB.Where("payment_id = ?", paymentID).Find(&history).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to get history data",
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success get history data",
		"history": history,
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
	var history models.History
	if err := config.DB.First(&history, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to update history data",
			"error":  err.Error(),
		})
	}

	if err := c.Bind(&history); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to update history data",
			"error":  err.Error(),
		})
	}

	if err := config.DB.Save(&history).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to update history data",
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success update history data",
		"history": history,
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

	var history models.History
	if err := config.DB.First(&history, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to delete history data",
			"error":  err.Error(),
		})
	}

	if err := config.DB.Delete(&history).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to delete history data",
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success delete history data",
		"history": history,
	})
}
