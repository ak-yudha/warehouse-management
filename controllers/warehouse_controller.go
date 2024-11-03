package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse-management/models"
	"warehouse-management/services"
)

type WarehouseController struct {
	Service services.WarehouseService
}

func NewWarehouseController(service services.WarehouseService) *WarehouseController {
	return &WarehouseController{Service: service}
}

func (ctrl *WarehouseController) RecordPenerimaan(ctx *gin.Context) {
	var header models.PenerimaanBarangHeader
	var details []models.PenerimaanBarangDetail

	if err := ctx.ShouldBindJSON(&header); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.Service.RecordPenerimaan(&header, details); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Penerimaan barang berhasil dicatat"})
}

func (ctrl *WarehouseController) RecordPengeluaran(ctx *gin.Context) {
	var header models.PengeluaranBarangHeader
	var details []models.PengeluaranBarangDetail

	if err := ctx.ShouldBindJSON(&header); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.Service.RecordPengeluaran(&header, details); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Pengeluaran barang berhasil dicatat"})
}

func (ctrl *WarehouseController) GetStockReport(ctx *gin.Context) {
	stockReport, err := ctrl.Service.GetStockReport()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, stockReport)
}
