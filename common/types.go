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

type CreateOrderResp struct {
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
	//Since       timestamppb.Timestamp `json:"c_since"`
	Credit string `json:"c_credit"`
}
