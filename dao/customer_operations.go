package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func SetNewCustomerPaymentInfo(CustomerID int32, newCustomerBalance float64, newCustomerYTD float64, newCustomerPaymentCnt int32) (err error) {
	if err := client.Session.Query(`UPDATE Customer SET C_BALANCE = ?, C_YTD_PAYMENT = ?, C_PAYMENT_CNT = newCustomerPaymentCnt WHERE C_ID = ?`, newCustomerBalance, newCustomerYTD, newCustomerPaymentCnt, CustomerID).Exec(); err != nil {
		log.Printf("[warn] Query err, err=%v", err)
		return err
	}

	return nil
}

func GetCustomerInfo(customerID int32) (customerInfo common.Customer, err error) {
	if err := client.Session.Query(`SELECT * FROM Customer WHERE C_ID = ?`, customerID).Scan(&customerInfo); err != nil {
		log.Printf("[warn] Querry err, err=%v", err)
		return common.Customer{}, err
	}

	return customerInfo, nil
}
