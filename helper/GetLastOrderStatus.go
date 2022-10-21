package helper

import (
	"cs5234/common"
	"cs5234/dao"
)

func GetLastOrderStatus(r common.GetLastOrderStatusReq) (res common.GetLastOrderStatusResp, err error) {
	// step 1. get orderByCustomer info. from table `orderByCustomer`
	orderByCustomer, err := dao.GetOrderByCustomerInfo(r.CustomerID)
	if err != nil {
		return common.GetLastOrderStatusResp{}, err
	}

	// step 2. use orderByCustomer.c_last_o_id to get the orderLine information in table `orderLine`
	orderLines, err := dao.GetOrderLineByOrder(orderByCustomer.CustomerLastOrderID)

	// step 3. pack the output data
	res = common.GetLastOrderStatusResp{
		FirstName:      orderByCustomer.FirstName,
		MiddleName:     orderByCustomer.MiddleName,
		LastName:       orderByCustomer.LastName,
		Balance:        orderByCustomer.CustomerBalance,
		OrderID:        orderByCustomer.CustomerLastOrderID,
		OrderCarrierID: orderByCustomer.OrderCarrierID,
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
