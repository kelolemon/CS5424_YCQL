package helper

import (
	"cs5234/dao"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func CopyOrderByCustomer() {
	// Step 1. Get all customer list
	allCustomers, err := dao.GetAllCustomerList()
	if err != nil {
		log.Printf("[warn] Get all customer list err, err=%v", err)
	}

	// Step 2. For every customer, get the last order info (o_id, o_entry_d, o_carrier_id)
	csvFile, err := os.Create("D:\\GitHub\\CS5424_YCQL\\project_files\\data_files\\order-by-customer.csv")
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	for _, c := range allCustomers {
		lastOrderInfo, err := dao.GetLastOrderInfo(c.WarehouseID, c.DistrictID, c.CustomerID)

		if err != nil {
			log.Printf("[warn] Get last order info err, err=%v", err)
		}

		line := []string{strconv.Itoa(int(c.WarehouseID)), strconv.Itoa(int(c.DistrictID)), strconv.Itoa(int(c.CustomerID)),
			lastOrderInfo.OrderEntryTime.Format("2006-01-02 15:04:05"), c.FirstName, c.MiddleName, c.LastName,
			fmt.Sprintf("%f", c.Balance), strconv.Itoa(int(lastOrderInfo.ID)),
			strconv.Itoa(int(lastOrderInfo.CarrierID))}

		if err := csvWriter.Write(line); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
