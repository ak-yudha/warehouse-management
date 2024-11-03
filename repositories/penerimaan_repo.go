package repositories

import (
	"warehouse-management/config"
	"warehouse-management/models"
)

func CreatePenerimaan(header *models.TransaksiPenerimaanBarangHeader) error {
	return config.DB.Create(header).Error
}

func GetAllPenerimaan() ([]models.TransaksiPenerimaanBarangHeader, error) {
	var penerimaan []models.TransaksiPenerimaanBarangHeader
	err := config.DB.Preload("Details").Find(&penerimaan).Error
	return penerimaan, err
}
