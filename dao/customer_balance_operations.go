package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func GetCustomerBalanceInfo(customerID int32, warehouseID int32, districtID int32) (customerBalance common.CustomerBalance, err error) {
	rawMap := make(map[string]interface{})
	if err := client.Session.Query(`SELECT * FROM customerbalance WHERE c_id = ? AND c_w_id = ? AND c_d_id = ?`, customerID, warehouseID, districtID).MapScan(rawMap); err != nil {
		log.Printf("[warn] Get customer balance information error, err=%v", err)
		return common.CustomerBalance{}, err
	}

	err = common.ToCqlStruct(rawMap, &customerBalance)
	if err != nil {
		log.Printf("[warn] To cql struct error, err=%v", err)
		return common.CustomerBalance{}, err
	}
	return customerBalance, nil
}

func GetTopCustomerBalanceInfo() (customerBalances []common.CustomerBalance, err error) {
	scanner := client.Session.Query(`SELECT * FROM customerbalance order by c_balance desc limit 10`).Iter().Scanner()
	customerBalance := common.CustomerBalance{}
	for scanner.Next() {
		err = scanner.Scan(&customerBalance.WarehouseID, &customerBalance.DistrictID, &customerBalance.ID, &customerBalance.Balance, &customerBalance.FirstName, &customerBalance.MiddleName, &customerBalance.LastName, &customerBalance.WarehouseName, &customerBalance.DistrictName)
		if err != nil {
			log.Printf("[warn] read customer balance err, err=%v", err)
			return nil, err
		}
		customerBalances = append(customerBalances, customerBalance)
	}
	return customerBalances, nil
}

func DeleteCustomerBalance(customerID int32, warehouseID int32, districtID int32) (err error) {
	if err := client.Session.Query(`DELETE FROM customerbalance WHERE c_id = ? AND c_w_id = ? AND c_d_id = ?`, customerID, warehouseID, districtID, customerID).Exec(); err != nil {
		log.Printf("[warn] Delete customer balance information err, err=%v", err)
		return err
	}

	return nil
}

func InsertCustomerBalanceInfo(newCustomerBalanceInfo *common.CustomerBalance) (err error) {
	if err := client.Session.Query(`INSERT INTO customerbalance (c_id, c_w_id, c_d_id, c_balance, c_first, c_middle, c_last, w_name, d_name) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		newCustomerBalanceInfo.ID, newCustomerBalanceInfo.WarehouseID, newCustomerBalanceInfo.DistrictID, newCustomerBalanceInfo.Balance, newCustomerBalanceInfo.FirstName, newCustomerBalanceInfo.MiddleName, newCustomerBalanceInfo.LastName, newCustomerBalanceInfo.WarehouseName, newCustomerBalanceInfo.DistrictName).Exec(); err != nil {
		log.Printf("[warn] Insert customer balance information err, err=%v", err)
		return err
	}
	return nil
}
