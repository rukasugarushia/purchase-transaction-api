package controllers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"test/wex/database"
	"test/wex/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	dateLayout = "2006-01-02"
	baseURL    = "https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange"
)

func ShowAllTransactions(c *gin.Context) {
	currencyParam := c.Params.ByName("country-currency")
	transactions := getPurchaseTransactions()

	client := models.NewHTTPClient(baseURL)

	var purchaseTransactionOutputs []models.PurchaseTransactionOutput

	for _, transaction := range transactions {
		sixMonthsAgo := transaction.TransactionDate.AddDate(0, -6, 0)

		exchangeRates := fetchExchangeRates(currencyParam, sixMonthsAgo, client)
		if len(exchangeRates.Data) > 0 {
			parsedExchangeRate, err := strconv.ParseFloat(exchangeRates.Data[0].ExchangeRate, 64)
			if err != nil {
				fmt.Println("Error when parsing exchangeRate to float64:", err)
				return
			}

			convertedAmount := roundToNearestCent(transaction.PurchaseAmount * parsedExchangeRate)

			purchaseTransactionOutputs = append(purchaseTransactionOutputs, models.PurchaseTransactionOutput{
				ID:              transaction.ID,
				Description:     transaction.Description,
				TransactionDate: transaction.TransactionDate,
				PurchaseAmount:  transaction.PurchaseAmount,
				ExchangeRate:    parsedExchangeRate,
				ConvertedAmount: convertedAmount,
			})
		} else {
			fmt.Println("No records found for transaction", transaction.ID)
		}
	}

	c.JSON(http.StatusOK, purchaseTransactionOutputs)
}

func CreateNewTransaction(c *gin.Context) {
	var transaction models.PurchaseTransaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validatePurchaseTransaction(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction.RoundToNearestCent()
	transaction.TransactionDate = time.Now()
	savePurchaseTransaction(&transaction)

	c.JSON(http.StatusOK, transaction)
}

func getPurchaseTransactions() []models.PurchaseTransaction {
	var transactions []models.PurchaseTransaction
	database.DB.Find(&transactions)
	return transactions
}

func fetchExchangeRates(currencyParam string, recordDateGteFilter time.Time, client *models.HTTPClient) models.ExchangeRateData {
	params := url.Values{}
	params.Set("fields", "country_currency_desc,exchange_rate,record_date")
	params.Set("page[size]", "1")
	params.Set("sort", "-record_date")
	params.Set("filter", fmt.Sprintf("country_currency_desc:eq:%s,record_date:gte:%s", currencyParam, recordDateGteFilter.Format(dateLayout)))

	response, err := client.MakeGETRequest(params.Encode())
	if err != nil {
		fmt.Println("Error:", err)
	}

	var exchangeRates models.ExchangeRateData
	errJson := json.Unmarshal([]byte(response), &exchangeRates)
	if errJson != nil {
		fmt.Println("Error:", errJson)
	}

	return exchangeRates
}

func roundToNearestCent(amount float64) float64 {
	return math.Round(amount*100) / 100
}

func validatePurchaseTransaction(transaction *models.PurchaseTransaction) error {
	validate := validator.New()
	return validate.Struct(transaction)
}

func savePurchaseTransaction(transaction *models.PurchaseTransaction) {
	database.DB.Create(transaction)
}
