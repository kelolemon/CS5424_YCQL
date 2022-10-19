package common

type CreateOrderReq struct {
	WarehouseID     int32   `json:"w_id"`
	DistrictID      int32   `json:"d_id"`
	CustomerID      int32   `json:"c_id"`
	NumberItems     int32   `json:"number_items"`
	ItermNumber     []int32 `json:"iterm_number"`
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
	EntryDate    int64      `json:"o_entry_d"`
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
	WarehouseID int32  `json:"c_w_id"`
	DistrictID  int32  `json:"c_d_id"`
	CustomerID  int32  `json:"c_id"`
	FirstName   string `json:"c_first"`
	MiddleName  string `json:"c_middle"`
	LastName    string `json:"c_last"`
	Street1     string `json:"c_street_1"`
	Street2     string `json:"c_street_2"`
	City        string `json:"c_city"`
	State       string `json:"c_state"`
	Zip         int32  `json:"c_zip"`
	Phone       int32  `json:"c_phone"`
	Since       int64  `json:"c_since"`
	Credit      string `json:"c_credit"`
}
