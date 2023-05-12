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

func TestGetBillsController(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/bills", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetBillsController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "success get all bills")
	}

	// Test negative path with invalid request
	req = httptest.NewRequest(http.MethodPost, "/bills", nil) // Menggunakan MethodPost sebagai contoh request yang tidak valid
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	err := GetBillsController(c).(*echo.HTTPError)
	if assert.Error(t, err) {
		assert.Equal(t, http.StatusMethodNotAllowed, err.Code)
		assert.Contains(t, err.Message.(string), "method not allowed")
	}
}

func TestGetBillController(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/bills/3", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	if assert.NoError(t, GetBillController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "success get Bill by id")
	}

	// Test negative path with invalid request
	req = httptest.NewRequest(http.MethodGet, "/bills/invalid", nil) // Menggunakan nilai yang tidak valid sebagai contoh parameter id yang tidak valid
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	err := GetBillController(c).(*echo.HTTPError)
	if assert.Error(t, err) {
		assert.Equal(t, http.StatusBadRequest, err.Code)           // Mengharapkan kode status 400 (Bad Request)
		assert.Contains(t, err.Message.(string), "invalid syntax") // Mengharapkan pesan kesalahan "invalid syntax"
	}
}
func TestCreateBillController(t *testing.T) {
	// setup
	e := echo.New()
	reqBody := `{
        "user_id": 20,
        "description": "pembayaran bulan mei",
        "amount": 500000,
        "status": "lunas"
    }`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// call controller
	err := CreateBillController(c)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response struct {
		Message string      `json:"message"`
		Bill    models.Bill `json:"bill"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success create new bill", response.Message)
	assert.Equal(t, 20, response.Bill.UserId)
	assert.Equal(t, "pembayaran bulan mei", response.Bill.Description)
	assert.Equal(t, 500000, response.Bill.Amount)
	assert.Equal(t, "lunas", response.Bill.Status)
}

func TestUpdateBillController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/bills/6", strings.NewReader(`{"user_id": 20, "description": "pembayaran bulan mei", "amount": 500000, "status": "success"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("6")
	// test UpdateBillController
	if assert.NoError(t, UpdateBillController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// check response
		var response map[string]interface{}
		err := json.Unmarshal([]byte(rec.Body.String()), &response)
		if assert.NoError(t, err) {
			assert.Equal(t, "berhasil update data", response["status"])
			assert.NotNil(t, response["bill"])
		}
		// check updated bill
		var bill models.Bill
		if err := config.DB.Where("id = ?", 6).First(&bill).Error; err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 20, bill.UserId)
		assert.Equal(t, "pembayaran bulan juni", bill.Description)
		assert.Equal(t, 500000, bill.Amount)
		assert.Equal(t, "success", bill.Status)
	}
}

func TestDeleteBillController(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test case 1: Delete bill with valid ID
	// Create a bill to delete
	bill := models.Bill{UserId: 20, Description: "mpembayaran bulan mei", Amount: 500000, Status: "lunas"}
	config.DB.Create(&bill)
	// Set the ID of the bill to delete
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(bill.ID)))
	// Call the DeleteBillController function
	err := DeleteBillController(c)
	assert.NoError(t, err)
	// Check if the response code is 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)
	// Check if the response body contains "berhasil menghapus data"
	assert.Contains(t, rec.Body.String(), "berhasil menghapus data")
	// Check if the bill is deleted from the database
	var deletedBill models.Bill
	err = config.DB.Where("id = ?", bill.ID).First(&deletedBill).Error
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	// Test case 2: Delete bill with invalid ID
	// Set an invalid bill ID (ID of a non-existent bill)
	c.SetParamValues("999")
	// Call the DeleteBillController function
	err = DeleteBillController(c)
	assert.NoError(t, err)
	// Check if the response code is 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)
	// Check if the response body contains "bill not available"
	assert.Contains(t, rec.Body.String(), "bill not available")
}
