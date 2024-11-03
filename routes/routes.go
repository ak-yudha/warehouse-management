package routes

import (
	"github.com/gin-gonic/gin"
	"warehouse-management/controllers"
)

func RegisterRoutes(router *gin.Engine, warehouseController *controllers.WarehouseController) {
	api := router.Group("/api")
	{
		api.POST("/penerimaan", warehouseController.RecordPenerimaan)
		api.POST("/pengeluaran", warehouseController.RecordPengeluaran)
		api.GET("/stock", warehouseController.GetStockReport)
	}
}
