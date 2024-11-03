package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse-management/models"
	"warehouse-management/services"
)

func CreatePenerimaan(c *gin.Context) {
	var input models.TransaksiPenerimaanBarangHeader
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreatePenerimaan(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func GetAllPenerimaan(c *gin.Context) {
	data, err := services.GetAllPenerimaan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
