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

func TestLoginUserController(t *testing.T) {
	e := echo.New()

	userJSON := `{"email": "galuhhanggara@gmail.com", "password": "12345"}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, LoginUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "success login")
	}
}

func TestGetUsersController(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "success get all users")
	}

	// Test negative path with invalid request
	req = httptest.NewRequest(http.MethodPost, "/users", nil) // Menggunakan MethodPost sebagai contoh request yang tidak valid
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	err := GetUsersController(c).(*echo.HTTPError)
	if assert.Error(t, err) {
		assert.Equal(t, http.StatusMethodNotAllowed, err.Code)
		assert.Contains(t, err.Message.(string), "method not allowed")
	}
}

func TestGetUserController(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users/4", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("4")

	if assert.NoError(t, GetUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "success get user by id")
	}

	// Test negative path with invalid request
	req = httptest.NewRequest(http.MethodGet, "/users/invalid", nil) // Menggunakan nilai yang tidak valid sebagai contoh parameter id yang tidak valid
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("invalid")

	err := GetUserController(c).(*echo.HTTPError)
	if assert.Error(t, err) {
		assert.Equal(t, http.StatusBadRequest, err.Code)           // Mengharapkan kode status 400 (Bad Request)
		assert.Contains(t, err.Message.(string), "invalid syntax") // Mengharapkan pesan kesalahan "invalid syntax"
	}
}

func TestCreateUserController(t *testing.T) {
	// setup
	e := echo.New()
	reqBody := `{
			"nis": 12345,
			"name": "Davy Hazami",
			"username": "davy",
			"password": "30123",
			"email": "hazami@gmail.com",
			"jenis_kelamin": "male",
			"tahun_ajaran": "2022/2023"
		}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// call controller
	err := CreateUserController(c)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response struct {
		Message string      `json:"message"`
		User    models.User `json:"user"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success create new user", response.Message)
	assert.Equal(t, 12345, response.User.Nis)
	assert.Equal(t, "Davy Hazami", response.User.Name)
	assert.Equal(t, "davy", response.User.Username)
	assert.Equal(t, "30123", response.User.Password)
	assert.Equal(t, "hazami@gmail.com", response.User.Email)
	assert.Equal(t, "male", response.User.JenisKelamin)
	assert.Equal(t, "2022/2023", response.User.TahunAjaran)
}

func TestDeleteUserController(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Test case 1: Delete user with valid ID
	// Create a user to delete
	user := models.User{Name: "Hanggara", Email: "Hanggara123@gmail.com", Password: "23145"}
	config.DB.Create(&user)
	// Set the ID of the user to delete
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(user.ID)))
	// Call the DeleteUserController function
	err := DeleteUserController(c)
	assert.NoError(t, err)
	// Check if the response code is 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)
	// Check if the response body contains "berhasil menghapus data"
	assert.Contains(t, rec.Body.String(), "berhasil menghapus data")
	// Check if the user is deleted from the database
	var deletedUser models.User
	err = config.DB.Where("id = ?", user.ID).First(&deletedUser).Error
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	// Test case 2: Delete user with invalid ID
	// Set an invalid user ID (ID of a non-existent user)
	c.SetParamValues("999")
	// Call the DeleteUserController function
	err = DeleteUserController(c)
	assert.NoError(t, err)
	// Check if the response code is 200 OK
	assert.Equal(t, http.StatusOK, rec.Code)
	// Check if the response body contains "user not available"
	assert.Contains(t, rec.Body.String(), "user not available")
}

func TestUpdateUserController(t *testing.T) {
	// setup
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/users/9", strings.NewReader(`{"name":"Hanggara Gilang", "email":"Hanggara123@gmail.com", "password":"34567"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("6")

	// test UpdateUserController
	if assert.NoError(t, UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// check response
		var response map[string]interface{}
		err := json.Unmarshal([]byte(rec.Body.String()), &response)
		if assert.NoError(t, err) {
			assert.Equal(t, "berhasil update data", response["status"])
			assert.NotNil(t, response["user"])
		}

		// check updated user
		var user models.User
		if err := config.DB.Where("id = ?", 6).First(&user).Error; err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "Hanggara Gilang", user.Name)
		assert.Equal(t, "Hanggara123@gmail.com", user.Email)
		assert.Equal(t, "34567", user.Password)
	}
}
