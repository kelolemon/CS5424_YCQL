package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func GetCustomerBalanceInfo(customerID int32) (customerBalanceInfo common.CustomerBalance, err error) {
	if err := client.Session.Query(`SELECT * FROM customerbalance WHERE c_id = ?`, customerID).Scan(&customerBalanceInfo); err != nil {
		log.Printf("[warn] Get customer balance information err, err=%v", err)
		return common.CustomerBalance{}, err
	}

	return customerBalanceInfo, err
}

func SetCustomerBalanceInfo(customerID int32, newBalance float64) (err error) {
	if err := client.Session.Query(`UPDATE customerbalance SET c_balance = ? WHERE c_id = ?`, newBalance, customerID).Exec(); err != nil {
		log.Printf("[warn] Set customer balance information err, err=%v", err)
		return err
	}

	return nil
}

func InsertCustomerBalanceInfo(newCustomerBalanceInfo *common.CustomerBalance) (err error) {
	if err := client.Session.Query(`INSERT INTO customerbalance (c_id, c_balance, c_first, c_middle, c_last, w_name, d_name) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		newCustomerBalanceInfo.ID, newCustomerBalanceInfo.Balance, &newCustomerBalanceInfo.FirstName, newCustomerBalanceInfo.LastName, newCustomerBalanceInfo.WarehouseName, newCustomerBalanceInfo.DistrictName).Exec(); err != nil {
		log.Printf("[warn] Insert customer balance information err, err=%v", err)
		return err
	}
	return nil
}
