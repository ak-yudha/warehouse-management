package services

import (
	"warehouse-management/models"
	"warehouse-management/repositories"
)

type ReceiveService interface {
	AddReceive(receive *models.TransaksiPenerimaanBarangHeader) error
}

type receiveService struct {
	receiveRepo repositories.ReceiveRepository
}

func NewReceiveService(repo repositories.ReceiveRepository) ReceiveService {
	return &receiveService{receiveRepo: repo}
}

func (s *receiveService) AddReceive(receive *models.TransaksiPenerimaanBarangHeader) error {
	// Business logic can go here if needed
	return s.receiveRepo.CreateReceive(receive)
}
