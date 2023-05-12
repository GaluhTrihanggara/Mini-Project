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

func TestLoginAdminController(t *testing.T) {
	e := echo.New()

	adminJSON := `{"username": "Yanto", "password": "12345"}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(adminJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, LoginAdminController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "success login")
	}
}

func TestGetAdminsController(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/admins", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetAdminsController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "success get all admins")
	}

	// Test negative path with invalid request
	req = httptest.NewRequest(http.MethodPost, "/admins", nil) // Menggunakan MethodPost sebagai contoh request yang tidak valid
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	err := GetAdminsController(c).(*echo.HTTPError)
	if assert.Error(t, err) {
		assert.Equal(t, http.StatusMethodNotAllowed, err.Code)
		assert.Contains(t, err.Message.(string), "method not allowed")
	}
}

func TestGetAdminController(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/admins/3", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	if assert.NoError(t, GetAdminController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "success get admin by id")
	}

	// Test negative path with invalid request
	req = httptest.NewRequest(http.MethodGet, "/admins/invalid", nil) // Menggunakan nilai yang tidak valid sebagai contoh parameter id yang tidak valid
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	err := GetAdminController(c).(*echo.HTTPError)
	if assert.Error(t, err) {
		assert.Equal(t, http.StatusBadRequest, err.Code)           // Mengharapkan kode status 400 (Bad Request)
		assert.Contains(t, err.Message.(string), "invalid syntax") // Mengharapkan pesan kesalahan "invalid syntax"
	}
}

func TestCreateAdminController(t *testing.T) {
	// setup
	e := echo.New()
	reqBody := `{
		"name": "Kusmanto",
		"username": "Kusmanto",
		"password": "31823"
	}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// call controller
	err := CreateAdminController(c)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response struct {
		Message string       `json:"message"`
		Admin   models.Admin `json:"admin"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success create new admin", response.Message)
	assert.Equal(t, "Kusmanto", response.Admin.Name)
	assert.Equal(t, "Kusmanto", response.Admin.Username)
	assert.Equal(t, "31823", response.Admin.Password)
}

func TestDeleteAdminController(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test case 1: Delete admin with valid ID
	// Create a admin to delete
	admin := models.Admin{Name: "Kusmanto", Username: "Kusmanto", Password: "31823"}
	config.DB.Create(&admin)
	// Set the ID of the admin to delete
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(admin.ID)))
	// Call the DeleteAdminController function
	err := DeleteAdminController(c)
	assert.NoError(t, err)
	// Check if the response code is 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)
	// Check if the response body contains "berhasil menghapus data"
	assert.Contains(t, rec.Body.String(), "berhasil menghapus data")
	// Check if the admin is deleted from the database
	var deletedAdmin models.Admin
	err = config.DB.Where("id = ?", admin.ID).First(&deletedAdmin).Error
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	// Test case 2: Delete Admin with invalid ID
	// Set an invalid user ID (ID of a non-existent user)
	c.SetParamValues("999")
	// Call the DeleteAdminController function
	err = DeleteAdminController(c)
	assert.NoError(t, err)
	// Check if the response code is 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)
	// Check if the response body contains "admin not available"
	assert.Contains(t, rec.Body.String(), "admin not available")
}

func TestUpdateAdminController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/admins/8", strings.NewReader(`{"name":"Kusmanto", "username":"Kusmanto", "password":"31823"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("4")

	// test UpdateUserController
	if assert.NoError(t, UpdateAdminController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// check response
		var response map[string]interface{}
		err := json.Unmarshal([]byte(rec.Body.String()), &response)
		if assert.NoError(t, err) {
			assert.Equal(t, "berhasil memperbarui data admin", response["status"])
			assert.NotNil(t, response["admin"])
		}

		// check updated user
		var admin models.Admin
		if err := config.DB.Where("id = ?", 8).First(&admin).Error; err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "Kusmanto", admin.Name)
		assert.Equal(t, "Kusmanto", admin.Username)
		assert.Equal(t, "31823", admin.Password)
	}
}
