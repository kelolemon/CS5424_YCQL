package helper

import (
	"cs5234/common"
	"cs5234/dao"
)

func CreateNewPayment(r common.CreateNewPaymentReq) (res common.CreateNewPaymentReq, err error) {
	// step 1. update warehouse (c_w_id) => increment w_ytd by payment
	err = dao.SetNewWYTD(r.WarehouseID, r.Payment)
	if err != nil {
		return common.CreateNewPaymentReq{}, err
	}

	// step 2. update district (c_w_id, c_d_id) => increment d_ytd by payment
	err = dao.SetNewDYTD(r.WarehouseID, r.DistrictID, r.Payment)
	if err != nil {
		return common.CreateNewPaymentReq{}, err
	}

	// step 3. update customer (c_id) => decrement c_balance by payment; increment c_ytd_payment by payment; increment c_payment_cnt by 1
	err = dao.SetNewCPaymentSet(r.CustomerID, r.Payment)
	if err != nil {
		return common.CreateNewPaymentReq{}, err
	}

	// return
	return res, nil
}
