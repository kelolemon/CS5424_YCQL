package helper

import (
	"cs5234/common"
	"cs5234/dao"
	"log"
	"time"
)

func CreateNewDelivery(r common.CreateNewDeliveryReq) (res common.CreateNewDeliveryResp, err error) {
	for districtID := 0; districtID <= 10; districtID++ {
		// (a) Let N denote the value of the smallest order number O ID for district (W ID,DISTRICT NO)
		// with O CARRIER ID = null; i.e.,
		orderRes, err := dao.GetOldestNotDelivery(r.WarehouseID, int32(districtID))
		if err != nil {
			log.Printf("[warn] get oldest not delivery error err=%v", err)
			return common.CreateNewDeliveryResp{}, err
		}
		// (b) Update the order X by setting O CARRIER ID to CARRIER ID
		err = dao.SetCarrierInfo(r.WarehouseID, int32(districtID), orderRes.ID, r.CarrierID)
		if err != nil {
			log.Printf("[warn] update delivery error err=%v", err)
			return common.CreateNewDeliveryResp{}, err
		}
		// (c) Update all the order-lines in X by setting OL DELIVERY D to the current date and time
		orderDeliveryDate := time.Unix(time.Now().Unix(), 0)
		err = dao.SetOrderLineDeliveryDate(orderDeliveryDate, r.WarehouseID, int32(districtID), orderRes.ID)
		if err != nil {
			log.Printf("[warn] update delivery date error err=%v", err)
			return common.CreateNewDeliveryResp{}, err
		}
		// Update customer C as follows:
		//• Increment C BALANCE by B, where B denote the sum of OL AMOUNT for all the items placed in order X
		//• Increment C DELIVERY CNT by 1
		b, err := dao.GetOrderAmount(r.WarehouseID, int32(districtID), orderRes.ID)
		if err != nil {
			log.Printf("[warn] get tot amount error err=%v", err)
			return common.CreateNewDeliveryResp{}, err
		}
		customerRes, err := dao.GetCustomerInfo(orderRes.CustomerID, r.WarehouseID, int32(districtID))
		if err != nil {
			log.Printf("[warn] get customer info error, err=%v", err)
			return common.CreateNewDeliveryResp{}, err
		}
		err = dao.SetCustomerBalance(customerRes.ID, r.WarehouseID, int32(districtID), customerRes.Balance+b, customerRes.NumDeliveryMade+1)
		if err != nil {
			log.Printf("[warn] update order delivery balance error, err=%v", err)
			return common.CreateNewDeliveryResp{}, err
		}
		// update self order by customer table
		err = dao.SetOrderByCustomerBalanceINfo(customerRes.Balance+b, r.CarrierID, customerRes.ID, orderRes.OrderEntryTime)
		if err != nil {
			log.Printf("[warn] update order by customer delivery balance error, err=%v", err)
			return common.CreateNewDeliveryResp{}, err
		}
		// update select customer balance table
		customerBalanceRes, err := dao.GetCustomerBalanceInfo(customerRes.ID, r.WarehouseID, int32(districtID))
		if err != nil {
			log.Printf("[warn] get customer balance info error, err=%v", err)
			return common.CreateNewDeliveryResp{}, err
		}
		customerBalanceRes.Balance = customerRes.Balance + b
		err = dao.DeleteOrderByCustomerInfo(customerRes.ID, r.WarehouseID, int32(districtID))
		if err != nil {
			log.Printf("[warn] delete customer balance info error, err=%v", err)
			return common.CreateNewDeliveryResp{}, err
		}
		err = dao.InsertCustomerBalanceInfo(&customerBalanceRes)
		if err != nil {
			log.Printf("[warn] create customer balance info error, err=%v", err)
			return common.CreateNewDeliveryResp{}, err
		}
	}
	return common.CreateNewDeliveryResp{}, err
}
