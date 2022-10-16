package helper

import (
	"cs5234/common"
	"cs5234/dao"
	"log"
	"strconv"
	"time"
)

func CreateNewOrder(r common.CreateOrderReq) (res common.CreateOrderReq, err error) {
	//1. Let N denote value of the next available order number D NEXT O ID for district (W ID,D ID)
	districtRes, err := dao.GetDistrictInfo(r.WarehouseID, r.DistrictID)
	n := districtRes.DistrictNextOrderID
	if err != nil {
		log.Printf("[warn] GetNextOID error, err=%v", err)
		return common.CreateOrderReq{}, err
	}
	//2. Update the district (W ID, D ID) by incrementing D NEXT O ID by one
	err = dao.SetNewDNextOID(r.WarehouseID, r.DistrictID, n+1)
	if err != nil {
		log.Printf("[warn] SetNextOID error, err=%v", err)
		return common.CreateOrderReq{}, err
	}
	//3.Create a new order
	orderAllLocal := int32(0)
	for _, warehouseID := range r.SupplyWarehouse {
		if warehouseID != r.WarehouseID {
			orderAllLocal = 1
			break
		}
	}
	orderEntryDate := time.Now().Unix()
	err = dao.CreateNewOrder(n, r.WarehouseID, r.DistrictID, r.CustomerID, 0, r.NumberItems, orderAllLocal, orderEntryDate)
	if err != nil {
		log.Printf("[warn] create new order error, err=%v", err)
		return common.CreateOrderReq{}, err
	}
	//4.Initialize TOTAL AMOUNT = 0
	totalAmount := float64(0)
	//5. handle data，
	for i := int32(0); i < r.NumberItems; i++ {
		warehouseID := r.SupplyWarehouse[i]
		itemID := r.ItermNumber[i]
		// (a) Let S QUANTITY denote the stock quantity for item ITEM NUMBER[i] and warehouse SUPPLIER WAREHOUSE[i]
		res, err := dao.GetStockInfo(warehouseID, itemID)
		if err != nil {
			log.Printf("[warn] get stock info error, err=%v", err)
			return common.CreateOrderReq{}, err
		}
		// (b) ADJUSTED QTY = S QUANTITY − QUANTITY [i]
		adjustedQTY := res.StockQuantity - r.Quantity[i]
		// (c) If ADJUSTED QTY < 10, then set ADJUSTED QTY = ADJUSTED QTY + 100
		if adjustedQTY < 10 {
			adjustedQTY += 100
		}
		// (d) Update the stock for (ITEM NUMBER[i], SUPPLIER WAREHOUSE[i])
		res.StockYTD += float64(r.Quantity[i])
		res.StockOrderCnt += 1
		if warehouseID != r.WarehouseID {
			res.StockOrderCnt += 1
		}
		err = dao.UpdateStockInfo(warehouseID, itemID, adjustedQTY, res.StockYTD, res.StockOrderCnt, res.StockRemoteCnt)
		if err != nil {
			log.Printf("[warn] update stock info error, err=%v", err)
			return common.CreateOrderReq{}, err
		}
		// (e) ITEM AMOUNT = QUANTITY[i] × I PRICE, where I PRICE is the price of ITEM NUMBER[i]
		itemPrice, err := dao.GetItermPrice(r.ItermNumber[i])
		if err != nil {
			log.Printf("[warn] get iterm price error, err=%v", err)
			return common.CreateOrderReq{}, err
		}
		itemAmount := float64(r.Quantity[i]) * itemPrice
		// (f) TOTAL AMOUNT = TOTAL AMOUNT + ITEM AMOUNT
		totalAmount += itemAmount
		// (g) Create a new order-line
		err = dao.CreateNewOrderLine(r.WarehouseID, r.DistrictID, n, i, r.SupplyWarehouse[i], 0, r.ItermNumber[i], itemAmount, r.Quantity[i], "S_DIST_"+strconv.FormatInt(int64(r.DistrictID), 10))
		if err != nil {
			log.Printf("[warn] create order line error, err=%v", err)
			return common.CreateOrderReq{}, err
		}
	}
	// 6. TOTAL AMOUNT = TOTAL AMOUNT × (1+D TAX +W TAX) × (1−C DISCOUNT),
	// where W TAX is the tax rate for warehouse W ID, D TAX is the tax rate for district (W ID, D ID),
	//and C DISCOUNT is the discount for customer C ID.
	warehouseRes, err := dao.GetWarehouseInfo(r.WarehouseID)
	if err != nil {
		log.Printf("[warn] get warehouse info error, err=%v", err)
		return common.CreateOrderReq{}, err
	}
	customerRes, err := dao.GetCustomerInfo(r.CustomerID)
	if err != nil {
		log.Printf("[warn] get customer info error, err=%v", err)
		return common.CreateOrderReq{}, err
	}
	totalAmount = totalAmount * (1 + districtRes.DistrictTax) * (1 + warehouseRes.Tax) * (1 - customerRes.Discount)

	return res, nil
}
