package models

type MasterProduct struct {
	ProductPK   int    `gorm:"primaryKey" json:"product_pk"`
	ProductName string `json:"product_name"`
}

// TableName overrides the table name used by MasterProduct to `masterproduct`
func (MasterProduct) TableName() string {
	return "masterproduct"
}
