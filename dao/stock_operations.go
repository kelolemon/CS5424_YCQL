package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func UpdateStockInfo(WareHouseID int32, ItemID int32, StockQuantity int32, StockYTD float64, StockOrderCnt int32, StockRemoteCnt int32) (err error) {
	if err := client.Session.Query(`UPDATE Stock SET S_QUANTITY = ?, S_YTD = ?, S_ORDER_CNT = ?, S_REMOTE_CNT = ? WHERE S_W_ID = ? and S_I_ID = ?`,
		StockQuantity, StockYTD, StockOrderCnt, StockRemoteCnt, WareHouseID, ItemID).Exec(); err != nil {
		log.Printf("[warn] Update stock information err, err=%v", err)
		return err
	}

	return nil
}

func GetStockInfo(warehouseID int32, itemID int32) (stock common.Stock, err error) {
	rawMap := make(map[string]interface{})
	if err := client.Session.Query(`SELECT * FROM stock WHERE s_w_id = ? AND s_i_id = ?`, warehouseID, itemID).MapScan(rawMap); err != nil {
		log.Printf("[warn] Get stock information error, err=%v", err)
		return common.Stock{}, err
	}

	err = common.ToCqlStruct(rawMap, &stock)
	if err != nil {
		log.Printf("[warn] To cql struct error, err=%v", err)
		return common.Stock{}, err
	}

	return stock, nil
}

func InsertNewStockInfo(newStock *common.Stock) (err error) {
	err = client.Session.Query(`INSERT INTO stock (s_w_id, s_i_id, s_quantity, s_ytd, s_order_cnt, s_remote_cnt,
s_dist_01, s_dist_02, s_dist_03, s_dist_04, s_dist_05, s_dist_06, s_dist_07, s_dist_08, s_dist_09, s_dist_10, s_data) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, newStock.WarehouseID, newStock.ItemID, newStock.Quantity,
		newStock.YTD, newStock.OrderCnt, newStock.RemoteCnt, newStock.Dist1, newStock.Dist2, newStock.Dist3, newStock.Dist4,
		newStock.Dist5, newStock.Dist6, newStock.Dist7, newStock.Dist8, newStock.Dist9, newStock.Dist10, newStock.Data).Exec()
	if err != nil {
		log.Printf("[warn] Insert new stock information err, err=%v", err)
		return err
	}

	return nil
}
