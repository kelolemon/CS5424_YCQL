package helper

import (
	"cs5234/dao"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func CopyOrderLineQuantityByOrder() {
	// Step 1. Get Order List
	allOrdersWithDate, err := dao.GetAllOrderIdentifierWithDate()
	if err != nil {
		log.Printf("[warn] Get all orders with date err, err=%v", err)
	}

	// Step 2. For every order, get the order line info
	csvFile, err := os.Create("D:\\GitHub\\CS5424_YCQL\\project_files\\data_files\\order-line-quantity-by-order.csv")
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	for _, owd := range allOrdersWithDate {
		orderLine, err := dao.GetOrderLineByOrder(owd.WarehouseID, owd.DistrictID, owd.OrderID)
		if err != nil {
			log.Printf("[warn] Get orderline info err")
		}

		// Step 2.1 Reassemble the order line info to a map

		var itemIDList []int32
		var quantityList []int32
		var itemNameList []string

		for _, ol := range orderLine {
			itemIDList = append(itemIDList, ol.ItemID)
			itemInfo, err := dao.GetItemInfo(ol.ItemID)
			if err != nil {
				log.Printf("[warn] Item info err, err=%v", err)
			}
			itemNameList = append(itemNameList, itemInfo.Name)
			quantityList = append(quantityList, ol.Quantity)
		}

		var itemQuantityMapString string
		itemQuantityMapString += "{"
		for i := 0; i < len(itemIDList); i++ {
			itemQuantityMapString += strconv.Itoa(int(itemIDList[i]))
			itemQuantityMapString += ": "
			itemQuantityMapString += strconv.Itoa(int(quantityList[i]))
			if i < len(itemIDList)-1 {
				itemQuantityMapString += ", "
			} else {
				itemQuantityMapString += "}"
			}
		}

		var itemIDNameMapString string
		itemIDNameMapString += "{"
		for i := 0; i < len(itemIDList); i++ {
			itemIDNameMapString += strconv.Itoa(int(itemIDList[i]))
			itemIDNameMapString += ": "
			itemIDNameMapString += itemNameList[i]
			if i < len(itemIDList)-1 {
				itemIDNameMapString += ", "
			} else {
				itemIDNameMapString += "}"
			}
		}

		first, middle, last, err := dao.GetCustomerName(owd.CustomerID)
		if err != nil {
			log.Printf("[warn] Get customer name err, err=%v", err)
		}

		line := []string{strconv.Itoa(int(owd.WarehouseID)), strconv.Itoa(int(owd.DistrictID)), strconv.Itoa(int(owd.OrderID)),
			owd.OrderEntryDate.Format("2006-01-02 15:04:05"), itemQuantityMapString, itemIDNameMapString, first, middle, last}

		if err := csvWriter.Write(line); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
