package dao

import (
	"cs5234/client"
	"cs5234/common"
	"log"
)

func GetOrderByCustomerInfo(customerID int32) (orderByCustomer common.OrderByCustomer, err error) {
	rawMap := make(map[string]interface{})
	err = client.Session.Query(`SELECT * FROM orderbycustomer WHERE c_id = ? order by o_entry_d DESC LIMIT 1`, customerID).MapScan(rawMap)
	if err != nil {
		log.Printf("[warn] Get warehouse information error, err=%v", err)
		return common.OrderByCustomer{}, err
	}

	err = common.ToCqlStruct(rawMap, &orderByCustomer)
	if err != nil {
		log.Printf("[warn] To cql struct error, err=%v", err)
	}

	return orderByCustomer, nil
}
