package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func GetCustomerBalanceInfo(customerID int32) (customerBalanceInfo common.CustomerBalance, err error) {
	if err := client.Session.Query(`SELECT * FROM CustomerBalance WHERE c_id = ?`, customerID).Scan(&customerBalanceInfo); err != nil {
		log.Printf("[warn] Query err, err=%v", err)
		return common.CustomerBalance{}, err
	}

	return customerBalanceInfo, err
}

func SetCustomerBalanceInfo(customerID int32, newBalance float64) (err error) {
	if err := client.Session.Query(`UPDATE CustomerBalance SET balance = ?`, newBalance).Exec(); err != nil {
		log.Printf("[warn] Query err, err=%v", err)
		return err
	}

	return nil
}

func InsertCustomerBalanceInfo(newCustomerBalanceInfo common.CustomerBalance) (err error) {
	if err := client.Session.Query(`INSERT INTO CustomerBalance (ID, Balance, FirstName, MiddleName, LastName, WarehouseName, DistrictName) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		newCustomerBalanceInfo.ID, newCustomerBalanceInfo.Balance, newCustomerBalanceInfo.FirstName, newCustomerBalanceInfo.LastName, newCustomerBalanceInfo.WarehouseName, newCustomerBalanceInfo.DistrictName).Exec(); err != nil {
		log.Printf("[warn] Query err, err=%v", err)
		return err
	}
	return nil
}
