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
	WarehouseID int32 `json:"c_w_id"`
	DistrictID  int32 `json:"c_d_id"`
	CustomerID  int32 `json:"c_id"`
	Payment     int32 `json:"payment"`
}

type CreateNewPaymentResp struct {
}
