package database

import (
	"MiniProject/config"
	"MiniProject/models"
)

// CreateHistory function to create new history record
func CreateHistory(history *models.History) error {
	err := config.DB.Create(&history).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAllHistories function to get all history records
func GetHistorys() ([]models.History, error) {
	var historys []models.History
	err := config.DB.Find(&historys).Error
	if err != nil {
		return nil, err
	}
	return historys, nil
}

// GetHistoryByID function to get a single history record by id
func GetHistoryByID(id int) (*models.History, error) {
	history := new(models.History)
	err := config.DB.Where("id = ?", id).First(&history).Error
	if err != nil {
		return nil, err
	}
	return history, nil
}

// GetHistoriesByUserID function to get all history records by user id
func GetHistoryByUserID(userID int) ([]models.History, error) {
	var history []models.History
	err := config.DB.Where("user_id = ?", userID).Find(&history).Error
	if err != nil {
		return nil, err
	}
	return history, nil
}

// GetHistoriesByPaymentID function to get all history records by payment id
func GetHistoryByPaymentID(paymentID int) ([]models.History, error) {
	var history []models.History
	err := config.DB.Where("payment_id = ?", paymentID).Find(&history).Error
	if err != nil {
		return nil, err
	}
	return history, nil
}

// DeleteHistoryByID function to delete a single history record by id
func DeleteHistoryByID(id int) error {
	err := config.DB.Where("id = ?", id).Delete(&models.History{}).Error
	if err != nil {
		return err
	}
	return nil
}
