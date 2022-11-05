package common

import "time"

type DBTest struct {
	TestID  int32  `cql:"test_id"`
	TestMsg string `cql:"test_msg"`
	TestNub int32  `cql:"test_nub"`
}

type Stock struct {
	WarehouseID int32   `cql:"s_w_id"`
	ItemID      int32   `cql:"s_i_id"`
	Quantity    int32   `cql:"s_quantity"`
	YTD         float64 `cql:"s_ytd"`
	OrderCnt    int32   `cql:"s_order_cnt"`
	RemoteCnt   int32   `cql:"s_remote_cnt"`
	Dist1       string  `cql:"s_dist_01"`
	Dist2       string  `cql:"s_dist_02"`
	Dist3       string  `cql:"s_dist_03"`
	Dist4       string  `cql:"s_dist_04"`
	Dist5       string  `cql:"s_dist_05"`
	Dist6       string  `cql:"s_dist_06"`
	Dist7       string  `cql:"s_dist_07"`
	Dist8       string  `cql:"s_dist_08"`
	Dist9       string  `cql:"s_dist_09"`
	Dist10      string  `cql:"s_dist_10"`
	Data        string  `cql:"s_data"`
}

type District struct {
	WarehouseID int32   `cql:"d_w_id"`
	ID          int32   `cql:"d_id"`
	Name        string  `cql:"d_name"`
	Street1     string  `cql:"d_street_1"`
	Street2     string  `cql:"d_street_2"`
	City        string  `cql:"d_city"`
	State       string  `cql:"d_state"`
	Zip         string  `cql:"d_zip"`
	Tax         float64 `cql:"d_tax"`
	YTD         float64 `cql:"d_ytd"`
	NextOrderID int32   `cql:"d_next_o_id"`
}

type Customer struct {
	WarehouseID     int32     `cql:"c_w_id"`
	DistrictID      int32     `cql:"c_d_id"`
	ID              int32     `cql:"c_id"`
	FirstName       string    `cql:"c_first"`
	MiddleName      string    `cql:"c_middle"`
	LastName        string    `cql:"c_last"`
	Street1         string    `cql:"c_street_1"`
	Street2         string    `cql:"c_street_2"`
	City            string    `cql:"c_city"`
	State           string    `cql:"c_state"`
	Zip             string    `cql:"c_zip"`
	Phone           string    `cql:"c_phone"`
	CreationTime    time.Time `cql:"c_since"`
	CreditStatus    string    `cql:"c_credit"`
	CreditLimit     float64   `cql:"c_credit_lim"`
	Discount        float64   `cql:"c_discount"`
	Balance         float64   `cql:"c_balance"`
	YtdPayment      float64   `cql:"c_ytd_payment"`
	NumPaymentMade  int32     `cql:"c_payment_cnt"`
	NumDeliveryMade int32     `cql:"c_delivery_cnt"`
	Data            string    `cql:"c_data"`
}

type Order struct {
	WarehouseID    int32     `cql:"o_w_id"`
	DistrictID     int32     `cql:"o_d_id"`
	ID             int32     `cql:"o_id"`
	CustomerID     int32     `cql:"o_c_id"`
	CarrierID      int32     `cql:"o_carrier_id"`
	NumItemOrdered int32     `cql:"o_ol_cnt"`
	OrderAllLocal  int32     `cql:"o_all_local"`
	OrderEntryTime time.Time `cql:"o_entry_d"`
}

type OrderLine struct {
	WarehouseID       int32     `cql:"ol_w_id"`
	DistrictID        int32     `cql:"ol_d_id"`
	OrderID           int32     `cql:"ol_o_id"`
	ID                int32     `cql:"ol_number"`
	ItemID            int32     `cql:"ol_i_id"`
	DeliveryTime      time.Time `cql:"ol_delivery_d"`
	Amount            float64   `cql:"ol_amount"`
	SupplyWarehouseID int32     `cql:"ol_supply_w_id"`
	Quantity          int32     `cql:"ol_quantity"`
	Info              string    `cql:"ol_dist_info"`
}

type Warehouse struct {
	ID      int32   `cql:"w_id"`
	Name    string  `cql:"w_name"`
	Street1 string  `cql:"w_street_1"`
	Street2 string  `cql:"w_street_2"`
	City    string  `cql:"w_city"`
	State   string  `cql:"w_state"`
	Zip     string  `cql:"w_zip"`
	Tax     float64 `cql:"w_tax"`
	YTD     float64 `cql:"w_ytd"`
}

type Item struct {
	ID      int32   `cql:"i_id"`
	Name    string  `cql:"i_name"`
	Price   float64 `cql:"i_price"`
	ImageID int32   `cql:"i_im_id"`
	Data    string  `cql:"i_data"`
}

type OrderByCustomer struct {
	CustomerID     int32     `cql:"c_id"`
	WarehouseID    int32     `cql:"c_w_id"`
	DistrictID     int32     `cql:"c_d_id"`
	OrderEntryTime time.Time `cql:"o_entry_d"`
	FirstName      string    `cql:"c_first"`
	MiddleName     string    `cql:"c_middle"`
	LastName       string    `cql:"c_last"`
	Balance        float64   `cql:"c_balance"`
	LastOrderID    int32     `cql:"c_last_o_id"`
	CarrierID      int32     `cql:"o_carrier_id"`
}

type OrderLineQuantityByOrder struct {
	WarehouseID            int32            `cql:"w_id"`
	DistrictID             int32            `cql:"d_id"`
	OrderID                int32            `cql:"o_id"`
	OrderEntryTime         time.Time        `cql:"o_entry_d"`
	OrderLineQuantitiesMap map[int32]int32  `cql:"ol_quantity_map"`
	OrderItemsIDNameMap    map[int32]string `cql:"items_id_name_map"`
	CustomerID             int32            `cql:"c_id"`
	CustomerFirstName      string           `cql:"c_first"`
	CustomerMiddleName     string           `cql:"c_middle"`
	CustomerLastName       string           `cql:"c_last"`
}

type CustomerBalance struct {
	ID            int32   `cql:"c_id"`
	WarehouseID   int32   `cql:"c_w_id"`
	DistrictID    int32   `cql:"c_d_id"`
	Balance       float64 `cq1:"c_balance"`
	FirstName     string  `cql:"c_first"`
	MiddleName    string  `cql:"c_middle"`
	LastName      string  `cql:"c_last"`
	WarehouseName string  `cql:"w_name"`
	DistrictName  string  `cql:"d_name"`
}

type CustomerBasicInfoList struct {
	WarehouseID int32   `json:"c_w_id"`
	DistrictID  int32   `json:"c_d_id"`
	CustomerID  int32   `json:"c_id"`
	Balance     float64 `json:"c_balance"`
	FirstName   string  `json:"c_first"`
	MiddleName  string  `json:"c_middle"`
	LastName    string  `json:"c_last"`
}

type WarehouseWithNameList struct {
	WarehouseID   int32  `json:"w_id"`
	WarehouseName string `json:"w_name"`
}

type DistrictWithNameList struct {
	WarehouseID  int32  `json:"d_w_id"`
	DistrictID   int32  `json:"d_id"`
	DistrictName string `json:"d_name"`
}

type OrderWithDateList struct {
	WarehouseID    int32     `json:"o_w_id"`
	DistrictID     int32     `json:"o_d_id"`
	OrderID        int32     `json:"o_id"`
	CustomerID     int32     `json:"o_c_id"`
	OrderEntryDate time.Time `json:"o_entry_d"`
}

type ItemWithStockQtyList struct {
	ItemID   int32 `json:"s_i_id"`
	StockQty int32 `json:"s_quantity"`
}
