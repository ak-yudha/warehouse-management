package models

type MasterCustomer struct {
	CustomerPK   int    `gorm:"primaryKey" json:"customer_pk"`
	CustomerName string `json:"customer_name"`
}

// TableName overrides the table name used by MasterCustomer to `mastercustomer`
func (MasterCustomer) TableName() string {
	return "mastercustomer"
}
