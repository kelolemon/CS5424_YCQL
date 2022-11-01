package helper

import (
	"cs5234/common"
	"cs5234/dao"
	"fmt"
)

func GetOrderPopularItems(req common.GetPopularItemReq) (common.GetPopularItemResp, error) {
	orderQuantities, err := dao.GetLastLOrdersQuantity(req.WarehouseID, req.DistrictID, req.NumLastOrders)
	if err != nil {
		return common.GetPopularItemResp{}, err
	}

	orderPopularItems := make([]*common.OrderPopularItem, 0)
	popularItemPercentages := make([]*common.PopularItemPercentage, 0)

	itemFrequency := make(map[int32]int32)
	allPopularItems := make(map[int32]string)

	// Get popular items for each order
	for _, oq := range orderQuantities {
		popularItems := make([]*common.PopularItem, 0)
		largest := int32(0)
		for _, quantity := range oq.OrderLineQuantitiesMap {
			if quantity > largest {
				largest = quantity
			}
		}

		for itemId, quantity := range oq.OrderLineQuantitiesMap {
			if quantity == largest {
				popularItems = append(popularItems, &common.PopularItem{
					ItemName: oq.OrderItemsIDNameMap[itemId],
					Quantity: quantity,
				})
				allPopularItems[itemId] = oq.OrderItemsIDNameMap[itemId]
			}
			itemFrequency[itemId] += 1
		}

		orderPopularItem := &common.OrderPopularItem{
			OrderID:            oq.OrderID,
			OrderEntryTime:     oq.OrderEntryTime,
			CustomerFirstName:  oq.CustomerFirstName,
			CustomerMiddleName: oq.CustomerMiddleName,
			CustomerLastName:   oq.CustomerLastName,
			PopularItems:       popularItems,
		}
		orderPopularItems = append(orderPopularItems, orderPopularItem)
	}

	// Calculate popular items percentage
	for itemId, itemName := range allPopularItems {
		percentage := float32(itemFrequency[itemId]) / float32(req.NumLastOrders)
		popularItemPercentages = append(popularItemPercentages, &common.PopularItemPercentage{
			ItemName:                itemName,
			PercentageOrderWithItem: fmt.Sprintf("%f%", percentage),
		})
	}

	popularItemResp := common.GetPopularItemResp{
		WarehouseID:            req.WarehouseID,
		DistrictID:             req.DistrictID,
		NumLastOrders:          req.NumLastOrders,
		OrderPopularItems:      orderPopularItems,
		PopularItemPercentages: popularItemPercentages,
	}

	return popularItemResp, nil
}
