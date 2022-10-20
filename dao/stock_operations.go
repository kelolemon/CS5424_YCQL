package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func UpdateStockInfo(WareHouseID int32, ItemID int32, StockQuantity int32, StockYTD float64, StockOrderCnt int32, StockRemoteCnt int32) (err error) {
	if err := client.Session.Query(`UPDATE Stock SET S_QUANTITY = ?, S_YTD = ?, S_ORDER_CNT = ?, S_REMOTE_CNT = ? WHERE S_W_ID = ? and S_I_ID = ?`,
		StockQuantity, StockYTD, StockOrderCnt, StockRemoteCnt, WareHouseID, ItemID).Exec(); err != nil {
		log.Printf("[warn] Query err, err=%v", err)
		return err
	}

	return nil
}

func GetStockInfo(WareHouseID int32, ItemID int32) (res common.Stock, err error) {
	if err := client.Session.Query("select * from Stock where S_W_ID = ? and S_I_ID = ?", WareHouseID, ItemID).Scan(&res); err != nil {
		log.Printf("[warn] Query err, err=%v", err)
		return common.Stock{}, err
	}
	return res, nil
}
