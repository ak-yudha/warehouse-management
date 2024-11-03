package services

import (
	"warehouse-management/config"
)

func GetStockReport() ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := config.DB.Table("penerimaan_barang_details as pbd").
		Select("pbd.trx_in_d_product_idf AS product_id", "pbd.trx_in_d_qty_dus AS qty_dus", "pbd.trx_in_d_qty_pcs AS qty_pcs").
		Scan(&result).Error
	return result, err
}
