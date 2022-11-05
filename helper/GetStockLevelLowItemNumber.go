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
	ItemAndQtyList, err := dao.GetStockInfoByStock(r.WarehouseID)
	ItemQtyMap := make(map[int32]int32)
	for _, v := range ItemAndQtyList {
		ItemQtyMap[v.ItemID] = v.StockQty
	}

	errChan := make(chan error, r.LastOrders)
	for i := n - r.LastOrders; i < n; i++ {
		go func(i int32) {
			stockByOrderLineRes, err := dao.GetOrderLineQuantity(r.WarehouseID, r.DistrictID, i)
			if err != nil {
				log.Printf("[warn] get last stock by order line info error, err=%v", err)
				errChan <- err
				return
			}
			for v := range stockByOrderLineRes.OrderItemsIDNameMap {
				if ItemQtyMap[v] < r.StockThreshold {
					totalNumberItems++
				}
			}
			errChan <- nil
		}(i)
	}
	for i := n - r.LastOrders; i < n; i++ {
		if err := <-errChan; err != nil {
			return common.GetStockLevelLowItemNumberResp{}, err
		}
	}

	return common.GetStockLevelLowItemNumberResp{
		StockLevelLowItemNumber: totalNumberItems,
	}, nil
}
