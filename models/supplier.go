package models

type MasterSupplier struct {
	SupplierPK   int    `gorm:"primaryKey" json:"supplier_pk"`
	SupplierName string `json:"supplier_name"`
}

// TableName overrides the table name used by MasterSupplier to `mastersupplier`
func (MasterSupplier) TableName() string {
	return "mastersupplier"
}
