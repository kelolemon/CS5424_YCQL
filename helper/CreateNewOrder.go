package helper

import (
	"cs5234/common"
	"cs5234/dao"
	"log"
	"time"
)

func CreateNewOrder(r common.CreateOrderReq) (res common.CreateOrderReq, err error) {
	//1. Let N denote value of the next available order number D NEXT O ID for district (W ID,D ID)
	N, err := dao.GetDNextOID(r.WarehouseID, r.DistrictID)
	if err != nil {
		log.Printf("[warn] GetNextOID error, err=%v", err)
		return common.CreateOrderReq{}, err
	}
	//2. Update the district (W ID, D ID) by incrementing D NEXT O ID by one
	err = dao.SetNewDNextOID(r.WarehouseID, r.DistrictID, N+1)
	if err != nil {
		log.Printf("[warn] SetNextOID error, err=%v", err)
		return common.CreateOrderReq{}, err
	}
	//3.Create a new order
	OrderAllLocal := int32(0)
	for _, WarehouseID := range r.SupplyWarehouse {
		if WarehouseID != r.WarehouseID {
			OrderAllLocal = 1
			break
		}
	}
	OrderEntryDate := time.Now().Unix()
	err = dao.CreateNewOrder(N, r.WarehouseID, r.DistrictID, r.CustomerID, 0, r.NumberItems, OrderAllLocal, OrderEntryDate)
	if err != nil {
		log.Printf("[warn] create new order error, err=%v", err)
		return common.CreateOrderReq{}, err
	}
	//4.Initialize TOTAL AMOUNT = 0
	TotalAmount := int32(0)
	_ = TotalAmount
	//5. handle data，
	return res, nil
}
