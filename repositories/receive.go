package repositories

import (
	"gorm.io/gorm"
	"warehouse-management/models"
)

type ReceiveRepository interface {
	CreateReceive(receive *models.TransaksiPenerimaanBarangHeader) error
}

type receiveRepository struct {
	db *gorm.DB
}

func NewReceiveRepository(db *gorm.DB) ReceiveRepository {
	return &receiveRepository{db: db}
}

func (r *receiveRepository) CreateReceive(receive *models.TransaksiPenerimaanBarangHeader) error {
	return r.db.Create(receive).Error
}
