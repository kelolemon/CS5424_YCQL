package helper

import (
	"cs5234/common"
	"cs5234/dao"
)

func CreateNewPayment(r common.CreateNewPaymentReq) (res common.CreateNewPaymentReq, err error) {
	// step 1. update warehouse (c_w_id) => increment w_ytd by payment
	currentWarehouseInfo, err := dao.GetWarehouseInfo(r.WarehouseID)
	var newWarehouseYTD = currentWarehouseInfo.Ytd + r.Payment

	err = dao.SetNewWarehouseYTD(r.WarehouseID, newWarehouseYTD)
	if err != nil {
		return common.CreateNewPaymentReq{}, err
	}

	// step 2. update district (c_w_id, c_d_id) => increment d_ytd by payment
	currentDistrictInfo, err := dao.GetDistrictInfo(r.WarehouseID, r.DistrictID)
	var newDistrictYTD = currentDistrictInfo.YTD + r.Payment

	err = dao.SetNewDistrictYTD(r.WarehouseID, r.DistrictID, newDistrictYTD)
	if err != nil {
		return common.CreateNewPaymentReq{}, err
	}

	// step 3. update customer (c_id) => decrement c_balance by payment; increment c_ytd_payment by payment; increment c_payment_cnt by 1
	currentCustomerInfo, err := dao.GetCustomerInfo(r.CustomerID)
	var newCustomerBalance = currentCustomerInfo.Balance - r.Payment
	var newCustomerYTD = currentCustomerInfo.YtdPayment + r.Payment
	var newCustomerPaymentCnt = currentCustomerInfo.NumPaymentMade + 1

	err = dao.SetNewCustomerPaymentInfo(r.CustomerID, newCustomerBalance, newCustomerYTD, newCustomerPaymentCnt)
	if err != nil {
		return common.CreateNewPaymentReq{}, err
	}

	// return
	return res, nil
}
