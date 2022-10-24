package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func SetNewCustomerPaymentInfo(customerID int32, customerWarehouseID int32, customerDistrictID int32, newCustomerBalance float64, newCustomerYTD float64, newCustomerPaymentCnt int32) (err error) {
	if err = client.Session.Query(`UPDATE customer SET c_balance = ?, c_ytd_payment = ?, c_payment_cnt = ? WHERE c_id = ? AND c_w_id = ? AND c_d_id = ?`, newCustomerBalance, newCustomerYTD, newCustomerPaymentCnt, customerID, customerWarehouseID, customerDistrictID).Exec(); err != nil {
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
		return common.Customer{}, err
	}
	return customer, nil
}

func InsertNewCustomerInfo(newCustomer *common.Customer) (err error) {
	err = client.Session.Query(`INSERT INTO customer (c_w_id, c_d_id, c_id, c_zip, c_first, c_middle, c_last, c_street_1, 
c_street_2, c_city, c_state, c_phone, c_since, c_credit, c_credit_lim, c_discount, c_balance, c_ytd_payment, c_payment_cnt, 
c_delivery_cnt, c_data) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, newCustomer.WarehouseID,
		newCustomer.DistrictID, newCustomer.ID, newCustomer.Zip, newCustomer.FirstName, newCustomer.MiddleName, newCustomer.LastName,
		newCustomer.Street1, newCustomer.Street2, newCustomer.City, newCustomer.State, newCustomer.Phone, newCustomer.CreationTime,
		newCustomer.CreditStatus, newCustomer.CreditLimit, newCustomer.Discount, newCustomer.Balance, newCustomer.YtdPayment,
		newCustomer.NumPaymentMade, newCustomer.NumDeliveryMade, newCustomer.Data).Exec()
	if err != nil {
		log.Printf("[warn] Insert new item information err, err=%v", err)
		return err
	}

	return nil
}

func SetCustomerBalance(customerID int32, warehouseID int32, districtID int32, balance float64, deliveryCnt int32) (err error) {
	if err = client.Session.Query(`UPDATE customer SET c_balance = ?, c_delivery_cnt = ? WHERE c_id = ? AND c_w_id = ? AND c_d_id = ?`, balance, deliveryCnt, customerID, warehouseID, districtID).Exec(); err != nil {
		log.Printf("[warn] Set new customer balance information err, err=%v", err)
		return err
	}

	return nil
}
