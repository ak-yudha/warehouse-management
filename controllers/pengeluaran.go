package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse-management/models"
	"warehouse-management/services"
)

func CreatePengeluaran(c *gin.Context) {
	var input models.PengeluaranBarangHeader
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreatePengeluaran(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func GetAllPengeluaran(c *gin.Context) {
	data, err := services.GetAllPengeluaran()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
