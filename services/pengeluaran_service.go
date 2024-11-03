package services

import (
	"warehouse-management/models"
	"warehouse-management/repositories"
)

func CreatePengeluaran(header *models.PengeluaranBarangHeader) error {
	return repositories.CreatePengeluaran(header)
}

func GetAllPengeluaran() ([]models.PengeluaranBarangHeader, error) {
	return repositories.GetAllPengeluaran()
}
