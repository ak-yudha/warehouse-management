package services

import (
	"warehouse-management/models"
	"warehouse-management/repositories"
)

type WarehouseService interface {
	RecordPenerimaan(header *models.PenerimaanBarangHeader, details []models.PenerimaanBarangDetail) error
	RecordPengeluaran(header *models.PengeluaranBarangHeader, details []models.PengeluaranBarangDetail) error
	GetStockReport() ([]models.StockReport, error)
}

type warehouseService struct {
	repository repositories.WarehouseRepository
}

func NewWarehouseService(repository repositories.WarehouseRepository) WarehouseService {
	return &warehouseService{repository}
}

func (s *warehouseService) RecordPenerimaan(header *models.PenerimaanBarangHeader, details []models.PenerimaanBarangDetail) error {
	return s.repository.CreatePenerimaan(header, details)
}

func (s *warehouseService) RecordPengeluaran(header *models.PengeluaranBarangHeader, details []models.PengeluaranBarangDetail) error {
	return s.repository.CreatePengeluaran(header, details)
}

func (s *warehouseService) GetStockReport() ([]models.StockReport, error) {
	return s.repository.GetStockReport()
}
