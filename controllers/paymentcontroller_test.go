package controllers

import (
	"MiniProject/config"
	"MiniProject/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreatePaymentController(t *testing.T) {
	// setup
	e := echo.New()
	reqBody := `{
		"bill_id": 6,
		"user_id": 20,
		"amount": 500000,
		"status": "success"
	}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// call controller
	err := CreatePaymentController(c)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response struct {
		Message string         `json:"message"`
		Payment models.Payment `json:"payment"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success create new payment", response.Message)
	assert.Equal(t, 6, response.Payment.BillId)
	assert.Equal(t, 20, response.Payment.UserId)
	assert.Equal(t, 500000, response.Payment.Amount)
	assert.Equal(t, "success", response.Payment.Status)
}

func TestGetPaymentsController(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/payments", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetPaymentsController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "success get all payments")
	}

	// Test negative path with invalid request
	req = httptest.NewRequest(http.MethodPost, "/payments", nil) // Menggunakan MethodPost sebagai contoh request yang tidak valid
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	err := GetPaymentsController(c).(*echo.HTTPError)
	if assert.Error(t, err) {
		assert.Equal(t, http.StatusMethodNotAllowed, err.Code)
		assert.Contains(t, err.Message.(string), "method not allowed")
	}
}

func TestGetPaymentController(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/Payments/4", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("4")

	if assert.NoError(t, GetPaymentController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "success get payment by id")
	}

	// Test negative path with invalid request
	req = httptest.NewRequest(http.MethodGet, "/payments/invalid", nil) // Menggunakan nilai yang tidak valid sebagai contoh parameter id yang tidak valid
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	err := GetPaymentController(c).(*echo.HTTPError)
	if assert.Error(t, err) {
		assert.Equal(t, http.StatusBadRequest, err.Code)           // Mengharapkan kode status 400 (Bad Request)
		assert.Contains(t, err.Message.(string), "invalid syntax") // Mengharapkan pesan kesalahan "invalid syntax"
	}
}

func TestUpdatePaymentController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/payments/6", strings.NewReader(`{"bill_id": 6, "user_id": 20, "amount": 500000, "status": "success"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("6")

	// test UpdatePaymentController
	if assert.NoError(t, UpdatePaymentController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// check response
		var response map[string]interface{}
		err := json.Unmarshal([]byte(rec.Body.String()), &response)
		if assert.NoError(t, err) {
			assert.Equal(t, "berhasil update data", response["status"])
			assert.NotNil(t, response["payment"])
		}

		// check updated payment
		var payment models.Payment
		if err := config.DB.Where("id = ?", 6).First(&payment).Error; err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 6, payment.BillId)
		assert.Equal(t, 20, payment.UserId)
		assert.Equal(t, 50000, payment.Amount)
		assert.Equal(t, "success", payment.Status)
	}
}

func TestDeletePaymentController(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test case 1: Delete payment with valid ID
	// Create a payment to delete
	payment := models.Payment{BillId: 6, UserId: 20, Amount: 500000, Status: "succes"}
	config.DB.Create(&payment)
	// Set the ID of the payment to delete
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(payment.ID)))
	// Call the DeletePaymentController function
	err := DeletePaymentController(c)
	assert.NoError(t, err)
	// Check if the response code is 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)
	// Check if the response body contains "berhasil menghapus data"
	assert.Contains(t, rec.Body.String(), "berhasil menghapus data")
	// Check if the payment is deleted from the database
	var deletedPayment models.Payment
	err = config.DB.Where("id = ?", payment.ID).First(&deletedPayment).Error
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	// Test case 2: Delete payment with invalid ID
	// Set an invalid payment ID (ID of a non-existent payment)
	c.SetParamValues("999")
	// Call the DeletePaymentController function
	err = DeletePaymentController(c)
	assert.NoError(t, err)
	// Check if the response code is 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)
	// Check if the response body contains "payment not available"
	assert.Contains(t, rec.Body.String(), "payment not available")
}
