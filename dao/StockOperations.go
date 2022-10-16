package dao

import (
	"cs5234/client"
	"log"
)

func UpdateOrderStockInfo(WareHouseID int32, ItemID int32, StockQuantity int32, StockYTD float64, StockOrderCnt int32, StockRemoteCnt int32) (err error) {
	session, err := client.DBCluster.CreateSession()
	if err != nil {
		log.Printf("[warn] Get DB session err, err=%v", err)
		return err
	}
	defer session.Close()

	if err := session.Query(`UPDATE Stock SET S_QUANTITY = ?, S_YTD = ?, S_ORDER_CNT = ?, S_REMOTE_CNT = ? FROM WHERE S_W_ID = ? and S_I_ID = ?`,
		StockQuantity, StockYTD, StockOrderCnt, StockRemoteCnt, WareHouseID, ItemID).Exec(); err != nil {
		log.Printf("[warn] Querry err, err=%v", err)
		return err
	}

	return nil
}
