package models

type MasterWarehouse struct {
	WhsPK   int    `gorm:"primaryKey" json:"whs_pk"`
	WhsName string `json:"whs_name"`
}

// TableName overrides the table name used by MasterWarehouse to `masterwarehouse`
func (MasterWarehouse) TableName() string {
	return "masterwarehouse"
}
