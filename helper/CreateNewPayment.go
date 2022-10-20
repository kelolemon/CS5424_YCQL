package helper

import (
	"cs5234/common"
	"cs5234/dao"
	"log"
)

func CreateNewPayment(r common.CreateNewPaymentReq) (res common.CreateNewPaymentResp, err error) {
	// step 1. update warehouse (c_w_id) => increment w_ytd by payment
	currentWarehouseInfo, err := dao.GetWarehouseInfo(r.WarehouseID)
	var newWarehouseYTD = currentWarehouseInfo.YTD + r.Payment

	err = dao.SetNewWarehouseYTD(r.WarehouseID, newWarehouseYTD)
	if err != nil {
		return common.CreateNewPaymentResp{}, err
	}

	// step 2. update district (c_w_id, c_d_id) => increment d_ytd by payment
	currentDistrictInfo, err := dao.GetDistrictInfo(r.WarehouseID, r.DistrictID)
	var newDistrictYTD = currentDistrictInfo.YTD + r.Payment

	err = dao.SetNewDistrictYTD(r.WarehouseID, r.DistrictID, newDistrictYTD)
	if err != nil {
		return common.CreateNewPaymentResp{}, err
	}

	// step 3. update customer (c_id) => decrement c_balance by payment; increment c_ytd_payment by payment; increment c_payment_cnt by 1
	currentCustomerInfo, err := dao.GetCustomerInfo(r.CustomerID)
	var newCustomerBalance = currentCustomerInfo.Balance - r.Payment
	var newCustomerYTD = currentCustomerInfo.YtdPayment + r.Payment
	var newCustomerPaymentCnt = currentCustomerInfo.NumPaymentMade + 1

	err = dao.SetNewCustomerPaymentInfo(r.CustomerID, r.WarehouseID, r.DistrictID, newCustomerBalance, newCustomerYTD, newCustomerPaymentCnt)
	if err != nil {
		return common.CreateNewPaymentResp{}, err
	}

	// step 4. update customer balance (c_id) => if customer exists, update, else add a new record
	currentCustomerBalanceInfo, err := dao.GetCustomerBalanceInfo(r.CustomerID)
	if err == nil {
		newCustomerBalanceInfo := common.CustomerBalance{
			ID:            currentCustomerInfo.ID,
			Balance:       newCustomerBalance,
			FirstName:     currentCustomerInfo.FirstName,
			MiddleName:    currentCustomerInfo.MiddleName,
			LastName:      currentCustomerInfo.LastName,
			DistrictName:  currentDistrictInfo.Name,
			WarehouseName: currentWarehouseInfo.Name,
		}

		if err := dao.InsertCustomerBalanceInfo(&newCustomerBalanceInfo); err != nil {
			log.Printf("[warn] insert customer balance info error, err=%v", err)
			return common.CreateNewPaymentResp{}, err
		}
	} else {
		if err := dao.SetCustomerBalanceInfo(r.CustomerID, currentCustomerBalanceInfo.Balance-r.Payment); err != nil {
			log.Printf("[warn] update customer balance info error, err=%v", err)
			return common.CreateNewPaymentResp{}, err
		}
	}

	res = common.CreateNewPaymentResp{
		WarehouseID: r.WarehouseID,
		DistrictID:  r.DistrictID,
		CustomerID:  r.CustomerID,
		FirstName:   currentCustomerInfo.FirstName,
		MiddleName:  currentCustomerInfo.MiddleName,
		LastName:    currentCustomerInfo.LastName,
		Street1:     currentCustomerInfo.Street1,
		Street2:     currentCustomerInfo.Street2,
		City:        currentCustomerInfo.City,
		State:       currentCustomerInfo.State,
		Zip:         currentCustomerInfo.Zip,
		Phone:       currentCustomerInfo.Phone,
		Since:       currentCustomerInfo.CreationTime,
		Credit:      currentCustomerInfo.CreditStatus,
	}

	// return
	return res, nil
}
