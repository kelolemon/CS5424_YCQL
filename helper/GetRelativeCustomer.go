package helper

import (
	"cs5234/common"
	"cs5234/dao"
	"fmt"
	"log"
)

func compareItems(items1 []int32, items2 []int32) (res bool) {
	refs := make(map[int32]int)
	intersect := make([]int32, 0)

	for _, v := range items1 {
		refs[v]++
	}

	for _, v := range items2 {
		times, _ := refs[v]
		if times == 1 {
			intersect = append(intersect, v)
		}
	}

	if len(intersect) >= 2 {
		return true
	}

	return false
}

func GetRelativeCustomer(r common.GetRelatedCustomerReq) (res common.GetRelatedCustomerResp, err error) {
	// step 1. get candidate customers
	diffWarehouseCustomerList, err := dao.GetDiffWarehouseCustomerList(r.CustomerID, r.WarehouseID)

	// step 2. get orders of target customer
	targetOrderIdentifiers, err := dao.GetOrderIdentifier(r.WarehouseID, r.DistrictID, r.CustomerID)
	fmt.Printf("targetOrderIdentifiers: %v", targetOrderIdentifiers)

	// step 3. get order items of target customer
	targetItems := make([][]int32, len(targetOrderIdentifiers))
	curr := 0
	for _, v := range targetOrderIdentifiers {
		fmt.Printf("v: %v", v)
		orderLineQuantity, err := dao.GetOrderLineQuantity(v.WarehouseID, v.DistrictID, v.OrderID)
		fmt.Printf("orderLineQuantity: %v\n", orderLineQuantity)
		if err != nil {
			log.Printf("[warn] Get order line quantity err, err=%v", err)
			return common.GetRelatedCustomerResp{}, err
		}

		// item list of current order
		targetItem := make([]int32, len(orderLineQuantity.OrderLineQuantitiesMap))

		i := 0
		for k := range orderLineQuantity.OrderLineQuantitiesMap {
			targetItem[i] = k
			i++
		}

		// append to the targetItems
		targetItems[curr] = targetItem
		fmt.Printf("targetItem: %v\n", targetItem)
		curr++
	}

	fmt.Printf("targetItems: %v\n", targetItems)
	// step 3. get orders of candidate customers
	// for every candidate customer, init the flag, get the item list of order line, and compare the items
	flag := false
	for _, candidateIdentifier := range diffWarehouseCustomerList {
		flag = false
		candidateOrderIdentifiers, err := dao.GetOrderIdentifier(candidateIdentifier.WarehouseID, candidateIdentifier.DistrictID, candidateIdentifier.CustomerID)
		if err != nil {
			log.Printf("Get candidate order identifiers err, err=%v", err)
			return common.GetRelatedCustomerResp{}, err
		}

		// for every order, get the order line and transform to item list
		for _, candidateOrderIdentifier := range candidateOrderIdentifiers {
			candidateOrderLineQuantity, err := dao.GetOrderLineQuantity(candidateOrderIdentifier.WarehouseID, candidateOrderIdentifier.DistrictID, candidateOrderIdentifier.OrderID)
			if err != nil {
				log.Printf("Get candidate order line quantity err, err=%v", err)
				return common.GetRelatedCustomerResp{}, err
			}

			// item list of current order
			candidateItem := make([]int32, len(candidateOrderLineQuantity.OrderLineQuantitiesMap))

			i := 0
			for k := range candidateOrderLineQuantity.OrderLineQuantitiesMap {
				candidateItem[i] = k
				i++
			}

			for _, targetItem := range targetItems {
				if compareRes := compareItems(targetItem, candidateItem); compareRes == true {
					res.CustomerList = append(res.CustomerList, candidateIdentifier)
					flag = true
					break
				}
			}

			if flag == true {
				break
			}
		}
	}

	return res, nil
}
