package controllers

import (
	"MiniProject/config"
	"MiniProject/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateHistoryController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	reqBody := `{
	"payment_id": 2,
	"status": "success",
	"description": "Pembayaran berhasil diproses"
}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// call controller
	err := CreateHistoryController(c)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response struct {
		Message string           `json:"message"`
		History models.Histories `json:"history"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success create new history", response.Message)
	assert.Equal(t, 2, response.History.PaymentID)
	assert.Equal(t, "success", response.History.Status)
	assert.Equal(t, "Pembayaran berhasil diproses", response.History.Description)
}

func TestGetHistoriesController(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/histories", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetHistoriesController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "success get all histories")
	}

	// Test negative path with invalid request
	req = httptest.NewRequest(http.MethodPost, "/histories", nil) // Menggunakan MethodPost sebagai contoh request yang tidak valid
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	err := GetHistoriesController(c).(*echo.HTTPError)
	if assert.Error(t, err) {
		assert.Equal(t, http.StatusMethodNotAllowed, err.Code)
		assert.Contains(t, err.Message.(string), "method not allowed")
	}
}

func TestGetHistoriesByPaymentController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/payments/1/histories", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// create dummy histories data
	paymentID := 1
	histories := []models.Histories{
		{PaymentID: paymentID, Status: "success", Description: "Pembayaran berhasil diproses"},
	}
	config.DB.Create(&histories)

	// call GetHistoriesByPaymentController
	if assert.NoError(t, GetHistoriesByPaymentController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// check response
		var response map[string]interface{}
		err := json.Unmarshal([]byte(rec.Body.String()), &response)
		if assert.NoError(t, err) {
			assert.Equal(t, "success get history data", response["status"])
			assert.NotNil(t, response["history"])
		}

		// check histories data
		var historiesFromDB []models.Histories
		if err := config.DB.Where("payment_id = ?", paymentID).Find(&historiesFromDB).Error; err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, len(histories), len(historiesFromDB))
	}
}

func TestUpdateHistoryController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	reqBody := `{
		"status": "success",
		"description": "Pembayaran berhasil diproses"
	}`
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// call controller
	err := UpdateHistoryController(c)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response struct {
		Status  string           `json:"status"`
		History models.Histories `json:"history"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success update history data", response.Status)
	assert.Equal(t, 1, int(response.History.ID))
	assert.Equal(t, 1, response.History.PaymentID)
	assert.Equal(t, "paid", response.History.Status)
	assert.Equal(t, "payment has been received", response.History.Description)
}

func TestDeleteHistoryController(t *testing.T) {
	// setup
	e := echo.New()
	history := models.Histories{
		PaymentID:   1,
		Status:      "paid",
		Description: "payment received",
	}
	config.DB.Create(&history)
	url := fmt.Sprintf("/histories/%d", history.ID)
	req := httptest.NewRequest(http.MethodDelete, url, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// call controller
	err := DeleteHistoryController(c)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response struct {
		Status  string           `json:"status"`
		History models.Histories `json:"history"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success delete history data", response.Status)
	assert.Equal(t, history.PaymentID, response.History.PaymentID)
	assert.Equal(t, history.Status, response.History.Status)
	assert.Equal(t, history.Description, response.History.Description)

	// check if history has been deleted from database
	var deletedHistory models.Histories
	config.DB.First(&deletedHistory, history.ID)
	assert.True(t, config.DB.RecordNotFound())
}
