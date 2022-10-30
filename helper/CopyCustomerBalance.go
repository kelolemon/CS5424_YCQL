package helper

import (
	"cs5234/client"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CustomerBasicInfoList struct {
	WarehouseID int32   `json:"c_w_id"`
	DistrictID  int32   `json:"c_d_id"`
	CustomerID  int32   `json:"c_id"`
	Balance     float64 `json:"c_balance"`
	FirstName   string  `json:"c_first"`
	MiddleName  string  `json:"c_middle"`
	LastName    string  `json:"c_last"`
}

type WarehouseBasicInfoList struct {
	WarehouseID   int32  `json:"w_id"`
	WarehouseName string `json:"w_name"`
}

type DistrictBasicInfoList struct {
	WarehouseID  int32  `json:"d_w_id"`
	DistrictID   int32  `json:"d_id"`
	DistrictName string `json:"d_name"`
}

func GetAllCustomerList() (customerBasicInfoLists []CustomerBasicInfoList, err error) {
	scanner := client.Session.Query(`SELECT c_w_id, c_d_id, c_id, c_balance, c_first, c_middle, c_last FROM customer`).Iter().Scanner()
	customerBasicInfoList := CustomerBasicInfoList{}
	for scanner.Next() {
		err = scanner.Scan(&customerBasicInfoList.WarehouseID, &customerBasicInfoList.DistrictID, &customerBasicInfoList.CustomerID,
			&customerBasicInfoList.Balance, &customerBasicInfoList.FirstName, &customerBasicInfoList.MiddleName, &customerBasicInfoList.LastName)
		if err != nil {
			log.Printf("[warn] Read customer basic info list err, err=%v", err)
			return nil, err
		}
		customerBasicInfoLists = append(customerBasicInfoLists, customerBasicInfoList)
	}

	if err = scanner.Err(); err != nil {
		log.Printf("[warn] Scanner err, err=%v", err)
		return []CustomerBasicInfoList{}, err
	}

	return customerBasicInfoLists, nil
}

func GetAllWarehouseList() (warehouseBasicInfoLists []WarehouseBasicInfoList, err error) {
	scanner := client.Session.Query(`SELECT w_id, w_name FROM warehouse`).Iter().Scanner()
	warehouseBasicInfoList := WarehouseBasicInfoList{}
	for scanner.Next() {
		err = scanner.Scan(&warehouseBasicInfoList.WarehouseID, &warehouseBasicInfoList.WarehouseName)
		if err != nil {
			log.Printf("[warn] Read warehouse basic info list err, err=%v", err)
			return nil, err
		}
		warehouseBasicInfoLists = append(warehouseBasicInfoLists, warehouseBasicInfoList)
	}

	if err = scanner.Err(); err != nil {
		log.Printf("[warn] Scanner err, err=%v", err)
		return []WarehouseBasicInfoList{}, err
	}

	return warehouseBasicInfoLists, nil
}

func GetAllDistrictList() (districtBasicInfoLists []DistrictBasicInfoList, err error) {
	scanner := client.Session.Query(`SELECT d_w_id, d_id, d_name FROM district`).Iter().Scanner()
	districtBasicInfoList := DistrictBasicInfoList{}
	for scanner.Next() {
		err = scanner.Scan(&districtBasicInfoList.WarehouseID, &districtBasicInfoList.DistrictID, &districtBasicInfoList.DistrictName)
		if err != nil {
			log.Printf("[warn] Read district basic info list err, err=%v", err)
			return nil, err
		}
		districtBasicInfoLists = append(districtBasicInfoLists, districtBasicInfoList)
	}

	if err = scanner.Err(); err != nil {
		log.Printf("[warn] Scanner err, err=%v", err)
		return []DistrictBasicInfoList{}, err
	}

	return districtBasicInfoLists, nil
}

func CopyCustomerBalance() {
	// step 1. Get customer list
	allCustomers, err := GetAllCustomerList()
	if err != nil {
		log.Printf("[warn] Get all customer list err, err=%v", err)
	}

	// step 2. Get warehouse name & district name
	// step 2.1. Get warehouse list
	allWarehouses, err := GetAllWarehouseList()
	if err != nil {
		log.Printf("[warn] Get all warehouse list err, err=%v", err)
	}

	// step 2.2. Make map of warehouseID : warehouseName
	warehouseIDNameMap := make(map[int32]string)
	for _, w := range allWarehouses {
		warehouseIDNameMap[w.WarehouseID] = w.WarehouseName
	}

	// step 2.3. Get district list
	allDistricts, err := GetAllDistrictList()
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
