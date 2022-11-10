package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
	"time"
)

func GetOrderByCustomerInfo(customerID int32, warehouseID int32, districtID int32) (orderByCustomer common.OrderByCustomer, err error) {
	rawMap := make(map[string]interface{})
	err = client.Session.Query(`SELECT * FROM orderbycustomer WHERE c_id = ? AND c_w_id = ? AND c_d_id = ? order by o_entry_d DESC LIMIT 1`, customerID, warehouseID, districtID).MapScan(rawMap)
	if err != nil {
		log.Printf("[warn] Get order by customer information error, err=%v", err)
		return common.OrderByCustomer{}, err
	}

	err = common.ToCqlStruct(rawMap, &orderByCustomer)
	if err != nil {
		log.Printf("[warn] To cql struct error, err=%v", err)
		return common.OrderByCustomer{}, err
	}

	return orderByCustomer, nil
}

func InsertOrderByCustomerInfo(orderByCustomer *common.OrderByCustomer) (err error) {
	err = client.Session.Query(`INSERT INTO orderbycustomer (c_w_id, c_d_id, c_id, o_entry_d, c_first, c_middle, 
                             c_last, c_balance, c_last_o_id, o_carrier_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		orderByCustomer.WarehouseID, orderByCustomer.DistrictID, orderByCustomer.CustomerID,
		orderByCustomer.OrderEntryTime, orderByCustomer.FirstName, orderByCustomer.MiddleName,
		orderByCustomer.LastName, orderByCustomer.Balance, orderByCustomer.LastOrderID,
		orderByCustomer.CarrierID).Exec()
	if err != nil {
		log.Printf("[warn] Insert new order by customer information err, err=%v", err)
		return err
	}

	return nil
}

func DeleteOrderByCustomerInfo(customerID int32, warehouseID int32, districtID int32) (err error) {
	err = client.Session.Query(`DELETE FROM orderbycustomer WHERE c_w_id = ? AND c_d_id = ? AND c_id = ?`, warehouseID, districtID, customerID).Exec()
	if err != nil {
		log.Printf("[warn] Delete order by customer information err, err=%v", err)
		return err
	}

	return nil
}

func SetOrderByCustomerBalanceInfo(balance float64, carrierID int32, warehouseID int32, districtID int32, customerID int32, entryDate time.Time) (err error) {
	stmt := `UPDATE orderbycustomer SET c_balance = ?, o_carrier_id = ? 
                       WHERE c_w_id = ? and c_d_id = ? and c_id = ? AND o_entry_d = ?`
	if err = client.Session.Query(
		stmt,
		balance, carrierID, warehouseID, districtID, customerID, entryDate).Exec(); err != nil {
		log.Printf("[warn] Set new customer balance information err, err=%v", err)
		return err
	}

	return nil
}
