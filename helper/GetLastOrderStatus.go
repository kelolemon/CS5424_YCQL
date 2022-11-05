package helper

import (
	"cs5234/common"
	"cs5234/dao"
	"log"
)

func GetLastOrderStatus(r common.GetLastOrderStatusReq) (res common.GetLastOrderStatusResp, err error) {
	// step 1. get orderByCustomer info. from table `orderByCustomer`
	orderByCustomer, err := dao.GetOrderByCustomerInfo(r.CustomerID, r.WarehouseID, r.DistrictID)
	if err != nil {
		return common.GetLastOrderStatusResp{}, err
	}

	// step 2. use orderByCustomer.c_last_o_id to get the orderLine information in table `orderLine`
	orderLines, err := dao.GetOrderLineByOrder(orderByCustomer.WarehouseID, orderByCustomer.DistrictID, orderByCustomer.LastOrderID)
	log.Printf("[error] get orderline by order err=%v, r=%v", err, r)
	// step 3. pack the output data
	res = common.GetLastOrderStatusResp{
		FirstName:      orderByCustomer.FirstName,
		MiddleName:     orderByCustomer.MiddleName,
		LastName:       orderByCustomer.LastName,
		Balance:        orderByCustomer.Balance,
		OrderID:        orderByCustomer.LastOrderID,
		OrderCarrierID: orderByCustomer.CarrierID,
		OrderEntryDate: orderByCustomer.OrderEntryTime,
	}

	var lastOrderStatusItem common.LastOrderStatusItemList
	for i := 0; i < len(orderLines); i++ {
		lastOrderStatusItem.ItemID = orderLines[i].ItemID
		lastOrderStatusItem.SupplyWarehouseID = orderLines[i].SupplyWarehouseID
		lastOrderStatusItem.Quantity = orderLines[i].Quantity
		lastOrderStatusItem.Amount = orderLines[i].Amount
		lastOrderStatusItem.DeliveryDate = orderLines[i].DeliveryTime
		res.Items = append(res.Items, lastOrderStatusItem)
	}

	return res, nil
}
