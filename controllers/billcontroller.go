package controllers

import (
	"MiniProject/config"
	"MiniProject/lib/database"
	"MiniProject/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// GetBillsController will allow admin to get all bills
func GetBillsController(c echo.Context) error {
	// get all bills from database
	var bills []models.Bill
	if err := config.DB.Find(&bills).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "failed to get bills",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"bills":  bills,
	})
}

// GetBillController gets a bill by its ID
func GetBillController(c echo.Context) error {
	// get bill ID from request parameter
	billID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "invalid parameter",
			"error":  err.Error(),
		})
	}
	// get the bill with given ID
	bill, err := database.GetBillByID(billID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status": "bill not found",
			})
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to get bill data",
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   bill,
	})
}

// CreateBillController creates a new bill
func CreateBillController(c echo.Context) error {
	// get request body
	bill := new(models.Bill)
	if err := c.Bind(bill); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "invalid request body",
			"error":  err.Error(),
		})
	}

	// create new bill
	if err := database.CreateBill(bill); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to create bill",
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"data":   bill,
	})
}

func UpdateBillController(c echo.Context) error {
	body := new(models.Bill)
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

	bill, err := database.GetBillByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status": "bill not found",
			})
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to get bill data",
			"error":  err.Error(),
		})
	}

	bill.UserId = body.UserId
	bill.Description = body.Description
	bill.Amount = body.Amount
	bill.Status = body.Status

	if err := database.UpdateBill(bill); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to update bill data",
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "successfully updated bill data",
		"bill":   bill,
	})
}

func DeleteBillController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "invalid id parameter",
		})
	}

	bill, err := database.GetBillByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status": "bill not found",
		})
	}

	if err := database.DeleteBill(bill); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed to delete bill",
			"error":  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "bill deleted successfully",
		"bill":   bill,
	})
}
