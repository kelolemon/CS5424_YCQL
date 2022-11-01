package helper

import (
	"cs5234/common"
	"cs5234/dao"
	"fmt"
	"log"
)

func mapKeysToList(rawMap map[int32]int32) (res []int32) {
	for k := range rawMap {
		res = append(res, k)
	}
	return res
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func GetRelativeCustomer(r common.GetRelatedCustomerReq) (res common.GetRelatedCustomerResp, err error) {
	// step 1. get all orders of the target customer
	targetOrderQuantities, err := dao.GetAllOrderLineQuantitiesByCustomer(r.WarehouseID, r.DistrictID, r.CustomerID)

	var targetItemLists [][]int32
	for _, targetOrderQuantity := range targetOrderQuantities {
		targetItemLists = append(targetItemLists, mapKeysToList(targetOrderQuantity.OrderLineQuantitiesMap))
		fmt.Print(mapKeysToList(targetOrderQuantity.OrderLineQuantitiesMap))
	}
	fmt.Print(targetItemLists)
	fmt.Println()

	// step 2. for every target order line, get the item related order identifiers (o_w_id, o_d_id, o_id)
	var relatedOrderIdentifiers []common.OrderIdentifierList
	for _, targetItemList := range targetItemLists {
		var orderIdentifiersOnce []common.OrderIdentifierList

		for _, targetItem := range targetItemList {
			orderIdentifiers, err := dao.GetOrderIdentifiersByItemID(targetItem)
			if err != nil {
				log.Printf("[warn] Get order identifiers err, err=%v", orderIdentifiers)
			}
			for _, orderIdentifier := range orderIdentifiers {
				if orderIdentifier.WarehouseID != r.WarehouseID {
					if Contains(orderIdentifiersOnce, orderIdentifier) {
						relatedOrderIdentifiers = append(relatedOrderIdentifiers, orderIdentifier)
					} else {
						orderIdentifiersOnce = append(orderIdentifiersOnce, orderIdentifier)
					}
				}
			}
		}
	}

	// step 3. find the customer identifiers wrt order identifiers in related orders
	for _, relatedOrderIdentifier := range relatedOrderIdentifiers {
		relatedCustomerID, err := dao.GetCustomerIDByOrderIdentifier(relatedOrderIdentifier.WarehouseID,
			relatedOrderIdentifier.DistrictID, relatedOrderIdentifier.OrderID)
		if err != nil {
			log.Printf("[warn] Get customer id err, err=%v", err)
		}
		customerIdentifierList := common.CustomerList{
			WarehouseID: relatedOrderIdentifier.WarehouseID,
			DistrictID:  relatedOrderIdentifier.DistrictID,
			CustomerID:  relatedCustomerID,
		}

		if Contains(res.CustomerList, customerIdentifierList) == false {
			res.CustomerList = append(res.CustomerList, customerIdentifierList)
		}
	}

	return res, nil
}
