package helper

import (
	"cs5234/common"
	"cs5234/dao"
	"log"
	"time"
)

func CreateNewDelivery(r common.CreateNewDeliveryReq) (res common.CreateNewDeliveryResp, err error) {
	errChan := make(chan error, 10)
	for districtID := 1; districtID <= 10; districtID++ {
		go func(districtID int) {
			// (a) Let N denote the value of the smallest order number O ID for district (W ID,DISTRICT NO)
			// with O CARRIER ID = null; i.e.,
			orderRes, err := dao.GetALlOrdersNotDelivery(r.WarehouseID, int32(districtID))
			if err != nil {
				log.Printf("[warn] get oldest not delivery error err=%v", err)
				errChan <- err
				return
			}

			if len(orderRes) == 0 {
				errChan <- nil
				return
			}

			var lastOrderNotDelivery common.Order
			for _, o := range orderRes {
				if o.ID >= lastOrderNotDelivery.ID {
					lastOrderNotDelivery = o
				}
			}

			// (b) Update the order X by setting O CARRIER ID to CARRIER ID
			err = dao.SetCarrierInfo(r.WarehouseID, int32(districtID), lastOrderNotDelivery.ID, r.CarrierID)
			if err != nil {
				log.Printf("[warn] update delivery error err=%v", err)
				errChan <- err
				return
			}
			// (c) Update all the order-lines in X by setting OL DELIVERY D to the current date and time
			orderDeliveryDate := time.Unix(time.Now().Unix(), 0)
			err = dao.SetOrderLinesDeliveryDate(orderDeliveryDate, r.WarehouseID, int32(districtID), lastOrderNotDelivery.ID)
			if err != nil {
				log.Printf("[warn] update delivery date error err=%v", err)
				errChan <- err
				return
			}
			// Update customer C as follows:
			//• Increment C BALANCE by B, where B denote the sum of OL AMOUNT for all the items placed in order X
			//• Increment C DELIVERY CNT by 1
			b, err := dao.GetOrderAmount(r.WarehouseID, int32(districtID), lastOrderNotDelivery.ID)
			if err != nil {
				log.Printf("[warn] get tot amount error err=%v", err)
				errChan <- err
				return
			}
			customerRes, err := dao.GetCustomerInfo(lastOrderNotDelivery.CustomerID, r.WarehouseID, int32(districtID))
			if err != nil {
				log.Printf("[warn] get customer info error, err=%v", err)
				errChan <- err
				return
			}
			err = dao.SetCustomerBalance(customerRes.ID, r.WarehouseID, int32(districtID), customerRes.Balance+b, customerRes.NumDeliveryMade+1)
			if err != nil {
				log.Printf("[warn] update order delivery balance error, err=%v", err)
				errChan <- err
				return
			}
			// update self order by customer table
			err = dao.SetOrderByCustomerBalanceINfo(customerRes.Balance+b, r.CarrierID, r.WarehouseID, int32(districtID), customerRes.ID, lastOrderNotDelivery.OrderEntryTime)
			if err != nil {
				log.Printf("[warn] update order by customer delivery balance error, err=%v", err)
				errChan <- err
				return
			}
			// update select customer balance table
			customerBalanceRes, err := dao.GetCustomerBalanceInfo(customerRes.ID, r.WarehouseID, int32(districtID))
			if err != nil {
				log.Printf("[warn] get customer balance info error, err=%v, costomer=%v", err, customerRes)
				errChan <- err
				return
			}
			customerBalanceRes.Balance = customerRes.Balance + b
			err = dao.DeleteOrderByCustomerInfo(customerRes.ID, r.WarehouseID, int32(districtID))
			if err != nil {
				log.Printf("[warn] delete customer balance info error, err=%v", err)
				errChan <- err
				return
			}
			err = dao.InsertCustomerBalanceInfo(&customerBalanceRes)
			if err != nil {
				log.Printf("[warn] create customer balance info error, err=%v", err)
				errChan <- err
				return
			}
			errChan <- nil
		}(districtID)
	}
	for i := 1; i <= 10; i++ {
		if err := <-errChan; err != nil {
			log.Printf("[error] create new delivery error err=%v", err)
			return common.CreateNewDeliveryResp{}, err
		}
	}
	return common.CreateNewDeliveryResp{}, nil
}
