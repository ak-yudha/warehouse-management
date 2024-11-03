package repositories

import (
	"warehouse-management/config"
	"warehouse-management/models"
)

func CreatePengeluaran(header *models.PengeluaranBarangHeader) error {
	return config.DB.Create(header).Error
}

func GetAllPengeluaran() ([]models.PengeluaranBarangHeader, error) {
	var pengeluaran []models.PengeluaranBarangHeader
	err := config.DB.Preload("Details").Find(&pengeluaran).Error
	return pengeluaran, err
}
