package services

import (
	"warehouse-management/models"
	"warehouse-management/repositories"
)

func CreatePenerimaan(header *models.TransaksiPenerimaanBarangHeader) error {
	return repositories.CreatePenerimaan(header)
}

func GetAllPenerimaan() ([]models.TransaksiPenerimaanBarangHeader, error) {
	return repositories.GetAllPenerimaan()
}
