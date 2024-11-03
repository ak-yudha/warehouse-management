package models

type PengeluaranBarangHeader struct {
	TrxOutPK      int    `gorm:"primaryKey" json:"trx_out_pk"`
	TrxOutNo      string `json:"trx_out_no"`
	WhsIdf        int    `json:"whs_idf"`
	TrxOutDate    string `json:"trx_out_date"`
	TrxOutSuppIdf int    `json:"trx_out_supp_idf"`
	TrxOutNotes   string `json:"trx_out_notes"`
}

func (PengeluaranBarangHeader) TableName() string {
	return "PengeluaranBarangHeader"
}

type PengeluaranBarangDetail struct {
	TrxOutDPK         int `gorm:"primaryKey" json:"trx_out_d_pk"`
	TrxOutIDF         int `json:"trx_out_idf"`
	TrxOutDProductIdf int `json:"trx_out_d_product_idf"`
	TrxOutDQtyDus     int `json:"trx_out_d_qty_dus"`
	TrxOutDQtyPcs     int `json:"trx_out_d_qty_pcs"`
}

func (PengeluaranBarangDetail) TableName() string {
	return "PengeluaranBarangDetail"
}
