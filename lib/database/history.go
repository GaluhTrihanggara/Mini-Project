package database

import (
	"MiniProject/config"
	"MiniProject/models"
)

// CreateHistory function to create new history record
func CreateHistory(history *models.Histories) error {
	err := config.DB.Create(&history).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAllHistories function to get all history records
func GetHistories() ([]models.Histories, error) {
	var historys []models.Histories
	err := config.DB.Find(&historys).Error
	if err != nil {
		return nil, err
	}
	return historys, nil
}

// GetHistoriesByPaymentID function to get all history records by payment id
func GetHistoriesByPaymentID(paymentID int) ([]models.Histories, error) {
	var history []models.Histories
	err := config.DB.Where("payment_id = ?", paymentID).Find(&history).Error
	if err != nil {
		return nil, err
	}
	return history, nil
}

// DeleteHistoryByID function to delete a single history record by id
func DeleteHistoryByID(id int) error {
	err := config.DB.Where("id = ?", id).Delete(&models.Histories{}).Error
	if err != nil {
		return err
	}
	return nil
}
