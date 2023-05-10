package database

import (
	"MiniProject/config"
	"MiniProject/models"
)

// GetBillByID returns a bill by its ID
func GetBillByID(id int) (*models.Bill, error) {
	var bill models.Bill
	if err := config.DB.First(&bill, id).Error; err != nil {
		return nil, err
	}
	return &bill, nil
}

// GetBills returns all bills
func GetBills() ([]models.Bill, error) {
	var bills []models.Bill
	if err := config.DB.Find(&bills).Error; err != nil {
		return nil, err
	}
	return bills, nil
}

// CreateBill creates a new bill
func CreateBill(bill *models.Bill) error {
	if err := config.DB.Create(&bill).Error; err != nil {
		return err
	}
	return nil
}

// UpdateBill updates an existing bill
func UpdateBill(bill *models.Bill) error {
	if err := config.DB.Save(&bill).Error; err != nil {
		return err
	}
	return nil
}

// DeleteBill deletes a bill
func DeleteBill(bill *models.Bill) error {
	if err := config.DB.Delete(&bill).Error; err != nil {
		return err
	}
	return nil
}
