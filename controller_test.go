package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"test/wex/controllers"
	"test/wex/database"
	"test/wex/models"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) MakeGETRequest(url string) (string, error) {
	args := m.Called(url)
	return args.String(0), args.Error(1)
}

func TestShowAllTransactions(t *testing.T) {
	initTestDB()
	prepareTestData()

	mockClient := new(MockHTTPClient)

	exchangeRateResponse := `{
		"data": [
			{
				"country_currency_desc": "Mexico-Peso",
				"exchange_rate": "17.471",
				"record_date": "2023-09-30"
			}
		]
	}`

	mockClient.On("MakeGETRequest", mock.Anything).Return(exchangeRateResponse, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	controllers.ShowAllTransactions(c)

	assert.Equal(t, http.StatusOK, w.Code)

	closeTestDB()
}

func TestCreateNewTransaction(t *testing.T) {
	initTestDB()
	prepareTestData()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	requestBody := `{
		"description": "Test Transaction"
		"purchase_amount": 97.9876
	}`

	c.Request, _ = http.NewRequest("POST", "/", strings.NewReader(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	controllers.CreateNewTransaction(c)

	assert.Equal(t, http.StatusOK, w.Code)

	closeTestDB()
}

func prepareTestData() {
	transaction1 := models.PurchaseTransaction{
		Description:     "Test Transaction 1",
		TransactionDate: time.Now(),
		PurchaseAmount:  200.0,
	}
	database.DB.Create(&transaction1)
}

func initTestDB() {
	testDBUser := "root"
	testDBPassword := "root"
	testDBName := "root"

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", testDBUser, testDBPassword, testDBName)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	database.DB = db

	db.AutoMigrate(&models.PurchaseTransaction{
		Description:     "My testing purchase",
		TransactionDate: time.Now(),
		PurchaseAmount:  197.98})
}

func closeTestDB() {
	db, err := database.DB.DB()
	if err != nil {
		panic("Failed to get the underlying database connection: " + err.Error())
	}

	db.Close()
}
