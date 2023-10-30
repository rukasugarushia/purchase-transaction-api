package routes

import (
	"test/wex/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/transactions/by-country-currency/:country-currency", controllers.ShowAllTransactions)
	r.POST("/transactions", controllers.CreateNewTransaction)
	r.Run()
}
