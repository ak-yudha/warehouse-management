package repositories

import (
	"gorm.io/gorm"
	"warehouse-management/models"
)

type WarehouseRepository interface {
	CreatePenerimaan(header *models.PenerimaanBarangHeader, details []models.PenerimaanBarangDetail) error
	CreatePengeluaran(header *models.PengeluaranBarangHeader, details []models.PengeluaranBarangDetail) error
	GetStockReport() ([]models.StockReport, error)
}

type warehouseRepository struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) WarehouseRepository {
	return &warehouseRepository{db}
}

func (r *warehouseRepository) CreatePenerimaan(header *models.PenerimaanBarangHeader, details []models.PenerimaanBarangDetail) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(header).Error; err != nil {
			return err
		}
		for _, detail := range details {
			detail.TrxInIDF = header.TrxInPK
			if err := tx.Create(&detail).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *warehouseRepository) CreatePengeluaran(header *models.PengeluaranBarangHeader, details []models.PengeluaranBarangDetail) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(header).Error; err != nil {
			return err
		}
		for _, detail := range details {
			detail.TrxOutIDF = header.TrxOutPK
			if err := tx.Create(&detail).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *warehouseRepository) GetStockReport() ([]models.StockReport, error) {
	var stockReport []models.StockReport
	err := r.db.Raw(`
        SELECT 
            w.WhsName AS warehouse, 
            p.ProductName AS product, 
            COALESCE(SUM(pd.TrxInDQtyDus), 0) - COALESCE(SUM(od.TrxOutDQtyDus), 0) AS qty_dus,
            COALESCE(SUM(pd.TrxInDQtyPcs), 0) - COALESCE(SUM(od.TrxOutDQtyPcs), 0) AS qty_pcs
        FROM MasterWarehouse w
        JOIN MasterProduct p
        LEFT JOIN PenerimaanBarangDetail pd ON pd.TrxInDProductIdf = p.ProductPK
        LEFT JOIN PengeluaranBarangDetail od ON od.TrxOutDProductIdf = p.ProductPK
        GROUP BY w.WhsName, p.ProductName
    `).Scan(&stockReport).Error
	return stockReport, err
}
