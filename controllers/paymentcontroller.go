package controllers

import (
	"MiniProject/config"
	"MiniProject/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreatePaymentController will create a new payment data
func CreatePaymentController(c echo.Context) error {
	var payment models.Payment
	if err := c.Bind(&payment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to create payment data",
			"error":  err.Error(),
		})
	}

	if err := config.DB.Create(&payment).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to create payment data",
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success create payment data",
		"payment": payment,
	})
}

// GetPaymentsController will get all payment data
func GetPaymentsController(c echo.Context) error {
	var payments []models.Payment
	if err := config.DB.Find(&payments).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to get payment data",
			"error":  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success get payment data",
		"payments": payments,
	})
}

func GetPaymentController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "invalid parameter",
			"error":  err.Error(),
		})
	}
	var payment models.Payment
	if err := config.DB.Where("id = ?", id).First(&payment).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to get payment data",
			"error":  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success get payment data",
		"payment": payment,
	})
}

// UpdatePaymentController will update payment data by id
func UpdatePaymentController(c echo.Context) error {
	body := new(models.Payment)

	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "parameter salah",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id parameter",
		})
	}

	var payment models.Payment
	if err := config.DB.Where("id_payment = ?", id).First(&payment).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "payment tidak tersedia",
		})
	}
	payment.BillId = body.BillId
	payment.UserId = body.UserId
	payment.Amount = body.Amount
	payment.Status = body.Status

	if err := config.DB.Save(&payment).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "gagal memperbarui data payment",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "berhasil memperbarui data payment",
		"payment": payment,
	})
}

// delete payment by id
func DeletePaymentController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "invalid parameter",
		})
	}
	var payment models.Payment
	if err := config.DB.Where("id_payment = ?", id).First(&payment).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "payment not available",
		})
	}

	if err := config.DB.Delete(&payment).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to delete payment",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success delete payment",
		"payment": payment,
	})
}
