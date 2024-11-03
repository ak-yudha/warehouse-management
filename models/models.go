package models

import (
	"time"
)

// Model untuk tabel MasterSupplier
type MasterSupplier struct {
	SupplierPK   int    `gorm:"primaryKey;autoIncrement"`
	SupplierName string `gorm:"type:nvarchar(255);not null"`
}

func (MasterSupplier) TableName() string {
	return "MasterSupplier"
}

// Model untuk tabel MasterCustomer
type MasterCustomer struct {
	CustomerPK   int    `gorm:"primaryKey;autoIncrement"`
	CustomerName string `gorm:"type:nvarchar(255);not null"`
}

func (MasterCustomer) TableName() string {
	return "MasterCustomer"
}

// Model untuk tabel MasterProduct
type MasterProduct struct {
	ProductPK   int    `gorm:"primaryKey;autoIncrement"`
	ProductName string `gorm:"type:nvarchar(255);not null"`
}

func (MasterProduct) TableName() string {
	return "MasterProduct"
}

// Model untuk tabel MasterWarehouse
type MasterWarehouse struct {
	WhsPK   int    `gorm:"primaryKey;autoIncrement"`
	WhsName string `gorm:"type:nvarchar(255);not null"`
}

func (MasterWarehouse) TableName() string {
	return "MasterWarehouse"
}

// Model untuk tabel PenerimaanBarangHeader
type PenerimaanBarangHeader struct {
	TrxInPK      int       `gorm:"primaryKey;autoIncrement"`
	TrxInNo      string    `gorm:"type:nvarchar(50);not null"`
	WhsIdf       int       `gorm:"not null"`
	TrxInDate    time.Time `gorm:"not null"`
	TrxInSuppIdf int       `gorm:"not null"`
	TrxInNotes   string    `gorm:"type:nvarchar(255);"`

	Warehouse MasterWarehouse          `gorm:"foreignKey:WhsIdf;references:WhsPK"`
	Supplier  MasterSupplier           `gorm:"foreignKey:TrxInSuppIdf;references:SupplierPK"`
	Details   []PenerimaanBarangDetail `gorm:"foreignKey:TrxInIDF"`
}

func (PenerimaanBarangHeader) TableName() string {
	return "PenerimaanBarangHeader"
}

// Model untuk tabel PenerimaanBarangDetail
type PenerimaanBarangDetail struct {
	TrxInDPK         int `gorm:"primaryKey;autoIncrement"`
	TrxInIDF         int `gorm:"not null"`
	TrxInDProductIdf int `gorm:"not null"`
	TrxInDQtyDus     int `gorm:"not null"`
	TrxInDQtyPcs     int `gorm:"not null"`

	Header  PenerimaanBarangHeader `gorm:"foreignKey:TrxInIDF;references:TrxInPK"`
	Product MasterProduct          `gorm:"foreignKey:TrxInDProductIdf;references:ProductPK"`
}

func (PenerimaanBarangDetail) TableName() string {
	return "PenerimaanBarangDetail"
}

// Model untuk tabel PengeluaranBarangHeader
type PengeluaranBarangHeader struct {
	TrxOutPK      int       `gorm:"primaryKey;autoIncrement"`
	TrxOutNo      string    `gorm:"type:nvarchar(50);not null"`
	WhsIdf        int       `gorm:"not null"`
	TrxOutDate    time.Time `gorm:"not null"`
	TrxOutSuppIdf int       `gorm:"not null"`
	TrxOutNotes   string    `gorm:"type:nvarchar(255);"`

	Warehouse MasterWarehouse           `gorm:"foreignKey:WhsIdf;references:WhsPK"`
	Supplier  MasterSupplier            `gorm:"foreignKey:TrxOutSuppIdf;references:SupplierPK"`
	Details   []PengeluaranBarangDetail `gorm:"foreignKey:TrxOutIDF"`
}

func (PengeluaranBarangHeader) TableName() string {
	return "PengeluaranBarangHeader"
}

// Model untuk tabel PengeluaranBarangDetail
type PengeluaranBarangDetail struct {
	TrxOutDPK         int `gorm:"primaryKey;autoIncrement"`
	TrxOutIDF         int `gorm:"not null"`
	TrxOutDProductIdf int `gorm:"not null"`
	TrxOutDQtyDus     int `gorm:"not null"`
	TrxOutDQtyPcs     int `gorm:"not null"`

	Header  PengeluaranBarangHeader `gorm:"foreignKey:TrxOutIDF;references:TrxOutPK"`
	Product MasterProduct           `gorm:"foreignKey:TrxOutDProductIdf;references:ProductPK"`
}

func (PengeluaranBarangDetail) TableName() string {
	return "PengeluaranBarangDetail"
}

// StockReport digunakan untuk menampilkan laporan stok
type StockReport struct {
	Warehouse string `json:"warehouse"`
	Product   string `json:"product"`
	QtyDus    int    `json:"qty_dus"`
	QtyPcs    int    `json:"qty_pcs"`
}
