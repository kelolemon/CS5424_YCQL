package helper

import (
	"cs5234/common"
	"cs5234/dao"
	"fmt"
	"log"
	"time"
)

func CreateNewDelivery(r common.CreateNewDeliveryReq) (res common.CreateNewDeliveryResp, err error) {
	start := time.Now()
	errChan := make(chan error, 10)
	for districtID := 1; districtID <= 10; districtID++ {
		go func(districtID int) {
			// (a) Let N denote the value of the smallest order number O ID for district (W ID,DISTRICT NO)
			// with O CARRIER ID = null; i.e.,
			if districtID == 1 {
				elapsed := time.Since(start)
				fmt.Println("step 1：", elapsed)
			}
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
			if districtID == 1 {
				elapsed := time.Since(start)
				fmt.Println("step 1.5：", elapsed)
			}
			var lastOrderNotDelivery common.Order
			for _, o := range orderRes {
				if o.ID >= lastOrderNotDelivery.ID {
					lastOrderNotDelivery = o
				}
			}
			if districtID == 1 {
				elapsed := time.Since(start)
				fmt.Println("step 2：", elapsed)
			}
			// (b) Update the order X by setting O CARRIER ID to CARRIER ID
			err = dao.SetCarrierInfo(r.WarehouseID, int32(districtID), lastOrderNotDelivery, r.CarrierID)
			if err != nil {
				log.Printf("[warn] update delivery error err=%v", err)
				errChan <- err
				return
			}
			if districtID == 1 {
				elapsed := time.Since(start)
				fmt.Println("step 3：", elapsed)
			}
			// (c) Update all the order-lines in X by setting OL DELIVERY D to the current date and time
			orderDeliveryDate := time.Unix(time.Now().Unix(), 0)
			err = dao.SetOrderLinesDeliveryDate(orderDeliveryDate, r.WarehouseID, int32(districtID), lastOrderNotDelivery.ID)
			if err != nil {
				log.Printf("[warn] update delivery date error err=%v", err)
				errChan <- err
				return
			}
			if districtID == 1 {
				elapsed := time.Since(start)
				fmt.Println("step 4：", elapsed)
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
			if districtID == 1 {
				elapsed := time.Since(start)
				fmt.Println("step 5：", elapsed)
			}
			customerRes, err := dao.GetCustomerInfo(lastOrderNotDelivery.CustomerID, r.WarehouseID, int32(districtID))

			if err != nil {
				log.Printf("[warn] get customer info error, err=%v", err)
				errChan <- err
				return
			}
			if districtID == 1 {
				elapsed := time.Since(start)
				fmt.Println("step 6：", elapsed)
			}
			err = dao.SetCustomerBalance(customerRes.ID, r.WarehouseID, int32(districtID), customerRes.Balance+b, customerRes.NumDeliveryMade+1)
			if err != nil {
				log.Printf("[warn] update order delivery balance error, err=%v", err)
				errChan <- err
				return
			}
			if districtID == 1 {
				elapsed := time.Since(start)
				fmt.Println("step 7：", elapsed)
			}
			// update self order by customer table
			err = dao.SetOrderByCustomerBalanceInfo(customerRes.Balance+b, r.CarrierID, r.WarehouseID, int32(districtID), customerRes.ID, lastOrderNotDelivery.OrderEntryTime)
			if err != nil {
				log.Printf("[warn] update order by customer delivery balance error, err=%v", err)
				errChan <- err
				return
			}
			// update select customer balance table
			if districtID == 1 {
				elapsed := time.Since(start)
				fmt.Println("step 8：", elapsed)
			}
			customerBalanceRes, err := dao.GetCustomerBalanceInfo(customerRes.ID, r.WarehouseID, int32(districtID))
			if err != nil {
				log.Printf("[warn] get customer balance info error, err=%v, costomer=%v", err, customerRes)
				errChan <- err
				return
			}
			customerBalanceRes.Balance = customerRes.Balance + b
			if districtID == 1 {
				elapsed := time.Since(start)
				fmt.Println("step 9：", elapsed)
			}
			err = dao.DeleteOrderByCustomerInfo(customerRes.ID, r.WarehouseID, int32(districtID))
			if err != nil {
				log.Printf("[warn] delete customer balance info error, err=%v", err)
				errChan <- err
				return
			}
			if districtID == 1 {
				elapsed := time.Since(start)
				fmt.Println("step 10：", elapsed)
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
	elapsed := time.Since(start)
	fmt.Println("step 11：", elapsed)
	return common.CreateNewDeliveryResp{}, nil
}
