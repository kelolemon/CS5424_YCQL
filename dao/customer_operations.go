package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func SetNewCustomerPaymentInfo(customerID int32, customerWarehouseID int32, customerDistrictID int32, newCustomerBalance float64, newCustomerYTD float64, newCustomerPaymentCnt int32) (err error) {
	if err := client.Session.Query(`UPDATE customer SET c_balance = ?, c_ytd_payment = ?, c_payment_cnt = ? WHERE c_id = ? AND c_w_id = ? AND c_d_id = ?`, newCustomerBalance, newCustomerYTD, newCustomerPaymentCnt, customerID, customerWarehouseID, customerDistrictID).Exec(); err != nil {
		log.Printf("[warn] Set new customer payment information err, err=%v", err)
		return err
	}

	return nil
}

func GetCustomerInfo(customerID int32, warehouseID int32, districtID int32) (customer common.Customer, err error) {
	rawMap := make(map[string]interface{})
	if err := client.Session.Query(`SELECT * FROM customer WHERE c_id = ? AND c_w_id = ? AND c_d_id = ?`, customerID, warehouseID, districtID).MapScan(rawMap); err != nil {
		log.Printf("[warn] Get customer information error, err=%v", err)
		return common.Customer{}, err
	}

	err = common.ToCqlStruct(rawMap, &customer)
	if err != nil {
		log.Printf("[warn] To cql struct error, err=%v", err)
	}
	return customer, nil
}
