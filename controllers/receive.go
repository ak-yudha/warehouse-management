package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse-management/models"
	"warehouse-management/services"
)

type ReceiveController struct {
	receiveService services.ReceiveService
}

func NewReceiveController(service services.ReceiveService) *ReceiveController {
	return &ReceiveController{receiveService: service}
}

func (c *ReceiveController) AddReceive(ctx *gin.Context) {
	var receive models.TransaksiPenerimaanBarangHeader
	if err := ctx.ShouldBindJSON(&receive); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.receiveService.AddReceive(&receive); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Penerimaan barang berhasil dicatat"})
}
