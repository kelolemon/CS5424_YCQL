package helper

import (
	"cs5234/dao"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func CopyCustomerBalance() {
	// step 1. Get customer list
	allCustomers, err := dao.GetAllCustomerList()
	if err != nil {
		log.Printf("[warn] Get all customer list err, err=%v", err)
	}

	// step 2. Get warehouse name & district name
	// step 2.1. Get warehouse list
	allWarehouses, err := dao.GetAllWarehouseList()
	if err != nil {
		log.Printf("[warn] Get all warehouse list err, err=%v", err)
	}

	// step 2.2. Make map of warehouseID : warehouseName
	warehouseIDNameMap := make(map[int32]string)
	for _, w := range allWarehouses {
		warehouseIDNameMap[w.WarehouseID] = w.WarehouseName
	}

	// step 2.3. Get district list
	allDistricts, err := dao.GetAllDistrictList()
	if err != nil {
		log.Printf("[warn] Get all district list err, err=%v", err)
	}

	// step 2.4. Make map of (districtWarehouseID, districtID) : districtName
	type districtKey struct {
		WarehouseID int32
		DistrictID  int32
	}

	districtIDNameMap := make(map[districtKey]string)
	for _, d := range allDistricts {
		districtIDNameMap[districtKey{d.WarehouseID, d.DistrictID}] = d.DistrictName
	}

	// step 3. assemble customer balance data
	csvFile, err := os.Create("D:\\GitHub\\CS5424_YCQL\\project_files\\data_files\\customer-balance.csv")
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	for _, row := range allCustomers {
		line := []string{strconv.Itoa(int(row.WarehouseID)), strconv.Itoa(int(row.DistrictID)), strconv.Itoa(int(row.CustomerID)),
			fmt.Sprintf("%f", row.Balance), row.FirstName, row.MiddleName, row.LastName, warehouseIDNameMap[row.WarehouseID],
			districtIDNameMap[districtKey{row.WarehouseID, row.DistrictID}]}
		if err := csvWriter.Write(line); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
