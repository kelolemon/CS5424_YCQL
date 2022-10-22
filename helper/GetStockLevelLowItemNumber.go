package helper

import (
	"cs5234/common"
	"cs5234/dao"
	"log"
)

func GetStockLevelLowItemNumber(r common.GetStockLevelLowItemNumberReq) (res common.GetStockLevelLowItemNumberResp, err error) {
	// 1. Let N denote the value of the next available order number D NEXT O ID for district (W ID,D ID)
	districtRes, err := dao.GetDistrictInfo(r.WarehouseID, r.DistrictID)
	n := districtRes.NextOrderID
	if err != nil {
		log.Printf("[warn] GetNextOID error, err=%v", err)
		return common.GetStockLevelLowItemNumberResp{}, err
	}
	//Let S denote the set of items from the last L orders for district (W ID,D ID); i.e.,
	//S = {t.OL I ID | t ∈ Order-Line, t.OL D ID = D ID, t.OL W ID = W ID, t.OL O ID ∈ [N−L, N)}
	totalNumberItems := int32(0)
	for i := n - r.LastOrders; i < n; i++ {
		stockByOrderLineRes, err := dao.GetLastStockByOrderLineInfo(r.WarehouseID, r.DistrictID, i)
		if err != nil {
			log.Printf("[warn] get last sotck by order line info error, err=%v", err)
			return common.GetStockLevelLowItemNumberResp{}, err
		}
		for _, v := range stockByOrderLineRes.StockQuantitiesMap {
			if v < r.StockThreshold {
				totalNumberItems++
			}
		}
	}

	return common.GetStockLevelLowItemNumberResp{
		StockLevelLowItemNumber: totalNumberItems,
	}, nil
}
