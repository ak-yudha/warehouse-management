package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse-management/services"
)

func GetStockReport(c *gin.Context) {
	data, err := services.GetStockReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
