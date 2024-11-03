package models

type TransaksiPenerimaanBarangHeader struct {
	TrxInPK      int    `gorm:"primaryKey" json:"trx_in_pk"`
	TrxInNo      string `json:"trx_in_no"`
	WhsIdf       int    `json:"whs_idf"`
	TrxInDate    string `json:"trx_in_date"`
	TrxInSuppIdf int    `json:"trx_in_supp_idf"`
	TrxInNotes   string `json:"trx_in_notes"`
}

// TableName overrides the table name used by TransaksiPenerimaanBarangHeader to `transaksipenerimaanbarangheader`
func (TransaksiPenerimaanBarangHeader) TableName() string {
	return "transaksipenerimaanbarangheader"
}

type TransaksiPenerimaanBarangDetail struct {
	TrxInDPK         int `gorm:"primaryKey" json:"trx_in_d_pk"`
	TrxInIDF         int `json:"trx_in_idf"`
	TrxInDProductIdf int `json:"trx_in_d_product_idf"`
	TrxInDQtyDus     int `json:"trx_in_d_qty_dus"`
	TrxInDQtyPcs     int `json:"trx_in_d_qty_pcs"`
}

// TableName overrides the table name used by TransaksiPenerimaanBarangDetail to `transaksipenerimaanbarangdetail`
func (TransaksiPenerimaanBarangDetail) TableName() string {
	return "transaksipenerimaanbarangdetail"
}
