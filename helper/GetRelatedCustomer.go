package helper

import (
	"cs5234/common"
	"cs5234/dao"
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

func GetRelatedCustomer(r common.GetRelatedCustomerReq) (res common.GetRelatedCustomerResp, err error) {
	// step 1. get all orders of the target customer
	targetOrderQuantities, err := dao.GetAllOrderLineQuantitiesByCustomer(r.WarehouseID, r.DistrictID, r.CustomerID)

	var targetItemLists [][]int32
	for _, targetOrderQuantity := range targetOrderQuantities {
		targetItemLists = append(targetItemLists, mapKeysToList(targetOrderQuantity.OrderLineQuantitiesMap))
	}

	// step 2. for every target order line, get the item related order identifiers (o_w_id, o_d_id, o_id)
	errChan := make(chan error, len(targetItemLists))
	var relatedOrderIdentifiers []common.OrderIdentifierList
	for _, targetItemList := range targetItemLists {
		go func(targetItemList []int32) {
			subErrChan := make(chan error, len(targetItemList))
			var orderIdentifiersOnce []common.OrderIdentifierList
			for _, targetItem := range targetItemList {
				go func(targetItem int32) {
					orderIdentifiers, err := dao.GetOrderIdentifiersByItemID(targetItem)
					if err != nil {
						log.Printf("[warn] Get order identifiers err, err=%v", orderIdentifiers)
						subErrChan <- err
						return
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
					subErrChan <- nil
				}(targetItem)
			}
			for i := 0; i < len(targetItemList); i++ {
				if err := <-subErrChan; err != nil {
					errChan <- err
					return
				}
			}
			errChan <- nil
		}(targetItemList)
	}
	for i := 0; i < len(targetItemLists); i++ {
		if err := <-errChan; err != nil {
			return common.GetRelatedCustomerResp{}, err
		}
	}
	// step 3. find the customer identifiers wrt order identifiers in related orders
	errChan = make(chan error, len(relatedOrderIdentifiers))
	for _, relatedOrderIdentifier := range relatedOrderIdentifiers {
		go func(relatedOrderIdentifier common.OrderIdentifierList) {
			relatedCustomerID, err := dao.GetCustomerIDByOrderIdentifier(relatedOrderIdentifier.WarehouseID,
				relatedOrderIdentifier.DistrictID, relatedOrderIdentifier.OrderID)
			if err != nil {
				log.Printf("[warn] Get customer id err, err=%v", err)
				errChan <- err
			}
			customerIdentifierList := common.CustomerList{
				WarehouseID: relatedOrderIdentifier.WarehouseID,
				DistrictID:  relatedOrderIdentifier.DistrictID,
				CustomerID:  relatedCustomerID,
			}

			if Contains(res.CustomerList, customerIdentifierList) == false {
				res.CustomerList = append(res.CustomerList, customerIdentifierList)
			}
			errChan <- nil
		}(relatedOrderIdentifier)
	}
	for i := 0; i < len(relatedOrderIdentifiers); i++ {
		if err := <-errChan; err != nil {
			return common.GetRelatedCustomerResp{}, err
		}
	}

	return res, nil
}
