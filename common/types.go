package common

import "time"

type CreateOrderReq struct {
	WarehouseID     int32   `json:"w_id"`
	DistrictID      int32   `json:"d_id"`
	CustomerID      int32   `json:"c_id"`
	NumberItems     int32   `json:"number_items"`
	ItemNumber      []int32 `json:"item_number"`
	SupplyWarehouse []int32 `json:"supply_warehouse"`
	Quantity        []int32 `json:"quantity"`
}

type ItemList struct {
	ItemNumber        int32   `json:"item_number"`
	ItemName          string  `json:"item_name"`
	SupplyWarehouseID int32   `json:"item_supply_warehouse_id"`
	Quantity          int32   `json:"quantity"`
	OrderAmount       float64 `json:"order_amount"`
	StockQuantity     int32   `json:"stock_quantity"`
}

type CreateOrderResp struct {
	WarehouseID  int32      `json:"w_id"`
	DistrictID   int32      `json:"d_id"`
	CustomerID   int32      `json:"c_id"`
	LastName     string     `json:"c_last"`
	CreditStatus string     `json:"c_credit"`
	Discount     float64    `json:"c_discount"`
	WarehouseTax float64    `json:"w_tax"`
	DistrictTax  float64    `json:"d_tax"`
	OrderID      int32      `json:"o_id"`
	EntryDate    time.Time  `json:"o_entry_d"`
	NumberItems  int32      `json:"number_items"`
	TotalAmount  float64    `json:"total_amount"`
	Items        []ItemList `json:"items"`
}

type CreateNewPaymentReq struct {
	WarehouseID int32   `json:"c_w_id"`
	DistrictID  int32   `json:"c_d_id"`
	CustomerID  int32   `json:"c_id"`
	Payment     float64 `json:"payment"`
}

type CreateNewPaymentResp struct {
	WarehouseID int32     `json:"c_w_id"`
	DistrictID  int32     `json:"c_d_id"`
	CustomerID  int32     `json:"c_id"`
	FirstName   string    `json:"c_first"`
	MiddleName  string    `json:"c_middle"`
	LastName    string    `json:"c_last"`
	Street1     string    `json:"c_street_1"`
	Street2     string    `json:"c_street_2"`
	City        string    `json:"c_city"`
	State       string    `json:"c_state"`
	Zip         string    `json:"c_zip"`
	Phone       string    `json:"c_phone"`
	Since       time.Time `json:"c_since"`
	Credit      string    `json:"c_credit"`
}

type GetLastOrderStatusReq struct {
	WarehouseID int32 `json:"c_w_id"`
	DistrictID  int32 `json:"c_d_id"`
	CustomerID  int32 `json:"c_id"`
}

type LastOrderStatusItemList struct {
	ItemID            int32     `json:"ol_i_id"`
	SupplyWarehouseID int32     `json:"ol_supply_w_id"`
	Quantity          int32     `json:"ol_quantity"`
	Amount            float64   `json:"ol_amount"`
	DeliveryDate      time.Time `json:"ol_delivery_d"`
}

type GetLastOrderStatusResp struct {
	FirstName      string                    `json:"c_first"`
	MiddleName     string                    `json:"c_middle"`
	LastName       string                    `json:"c_last"`
	Balance        float64                   `json:"c_balance"`
	OrderID        int32                     `json:"o_id"`
	OrderEntryDate time.Time                 `json:"o_entry_d"`
	OrderCarrierID int32                     `json:"o_carrier_id"`
	Items          []LastOrderStatusItemList `json:"last_o_items"`
}

type GetStockLevelLowItemNumberReq struct {
	WarehouseID    int32 `json:"w_id"`
	DistrictID     int32 `json:"d_id"`
	StockThreshold int32 `json:"stock_threshold"`
	LastOrders     int32 `json:"last_orders"`
}

type GetStockLevelLowItemNumberResp struct {
	StockLevelLowItemNumber int32 `json:"stock_level_low_item_number"`
}

type GetTopBalanceCustomerReq struct {
}

type CustomerBalanceInfo struct {
	FirstName     string  `json:"c_first"`
	MiddleName    string  `json:"c_middle"`
	LastName      string  `json:"c_last"`
	Balance       float64 `json:"c_balance"`
	WarehouseName string  `json:"w_name"`
	DistrictName  string  `json:"d_name"`
}

type GetTopBalanceCustomerResp struct {
	CustomerBalanceInfoList []CustomerBalanceInfo `json:"customer_balance_info_list"`
}
