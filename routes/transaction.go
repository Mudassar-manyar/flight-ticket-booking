package routes

import (
	"flight-ticket-booking/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(r *gin.Engine) {
	txn := r.Group("/transactions")
	{
		txn.POST("/", controllers.CreateTransaction)
		txn.GET("/", controllers.GetAllTransactions)
	}
}
