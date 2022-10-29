package helper

import (
	"cs5234/common"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func createOrder(data [][]string) (orderList []common.Order) {
	for _, line := range data {
		var order common.Order
		for i, field := range line {
			if i <= 6 {
				if field == "null" {
					field = "0"
				}

				temp, err := strconv.ParseInt(field, 10, 32)
				if err != nil {
					log.Fatal(err)
				}
				switch i {
				case 0:
					order.WarehouseID = int32(temp)
				case 1:
					order.DistrictID = int32(temp)
				case 2:
					order.ID = int32(temp)
				case 3:
					order.CustomerID = int32(temp)
				case 4:
					order.CarrierID = int32(temp)
				case 5:
					order.NumItemOrdered = int32(temp)
				case 6:
					order.OrderAllLocal = int32(temp)
				default:
					log.Printf("[warn] Order transform err")
				}
			} else {
				temp, err := time.Parse("2006-02-01 15:04:05", field)
				if err != nil {
					log.Fatal(err)
				}
				order.OrderEntryTime = temp
			}
		}
		orderList = append(orderList, order)
	}
	return orderList
}

func CopyOrder() {
	f, err := os.Open("D:\\GitHub\\CS5424_YCQL\\project_files\\data_files\\test.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	orderList := createOrder(data)
	fmt.Printf("%v\t%v\n", len(orderList), orderList[0])
}
