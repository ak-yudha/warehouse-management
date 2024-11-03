package routes

import (
	"github.com/gin-gonic/gin"
	"warehouse-management/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Endpoint untuk Penerimaan Barang
	router.POST("/penerimaan", controllers.CreatePenerimaan)
	router.GET("/penerimaan", controllers.GetAllPenerimaan)

	// Endpoint untuk Pengeluaran Barang
	router.POST("/pengeluaran", controllers.CreatePengeluaran)
	router.GET("/pengeluaran", controllers.GetAllPengeluaran)

	// Endpoint untuk Laporan Stok Barang
	router.GET("/stok", controllers.GetStockReport)

	return router
}
