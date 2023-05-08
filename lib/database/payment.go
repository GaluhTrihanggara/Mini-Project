package database

import (
	"MiniProject/config"
	"MiniProject/models"
)

// GetPaymentByID returns a payment by its ID
func GetPaymentByID(id int) (*models.Payment, error) {
	var payment models.Payment
	if err := config.DB.First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func GetPayments() ([]models.Payment, error) {
	var payments []models.Payment
	if err := config.DB.Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

func CreatePayment(payment *models.Payment) error {
	if err := config.DB.Create(&payment).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePayment(payment *models.Payment) error {
	if err := config.DB.Save(&payment).Error; err != nil {
		return err
	}
	return nil
}

func DeletePayment(payment *models.Payment) error {
	if err := config.DB.Delete(&payment).Error; err != nil {
		return err
	}
	return nil
}
