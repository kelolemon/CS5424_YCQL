package common

type Stock struct {
	StockWarehouseID int32   `cql:"s_w_id"`
	StockItemID      int32   `cql:"s_i_id"`
	StockQuantity    int32   `cql:"s_quantity"`
	StockYTD         float64 `cql:"s_ytd"`
	StockOrderCnt    int32   `cql:"s_order_cnt"`
	StockRemoteCnt   int32   `cql:"s_remote_cnt"`
	StockDist1       string  `cql:"s_dist_01"`
	StockDist2       string  `cql:"s_dist_02"`
	StockDist3       string  `cql:"s_dist_03"`
	StockDist4       string  `cql:"s_dist_04"`
	StockDist5       string  `cql:"s_dist_05"`
	StockDist6       string  `cql:"s_dist_06"`
	StockDist7       string  `cql:"s_dist_07"`
	StockDist8       string  `cql:"s_dist_08"`
	StockDist9       string  `cql:"s_dist_09"`
	StockDist10      string  `cql:"s_dist_10"`
	StockData        string  `cql:"s_date"`
}

type District struct {
	DistrictWarehouseID int32   `cql:"d_w_id"`
	DistrictID          int32   `cql:"d_id"`
	DistrictName        string  `cql:"d_name"`
	DistrictStreet1     string  `cql:"d_street_1"`
	DistrictStreet2     string  `cql:"d_street_2"`
	DistrictCity        string  `cql:"d_city"`
	DistrictState       string  `cql:"d_state"`
	DistrictZip         string  `cql:"d_zip"`
	DistrictTax         float64 `cql:"d_tax"`
	DistrictYTD         float64 `cql:"d_ytd"`
	DistrictNextOrderID int32   `cql:"d_next_o_id"`
}

type Customer struct {
	WarehouseID     int32   `cql:"c_w_id"`
	DistrictID      int32   `cql:"c_d_id"`
	ID              int32   `cql:"c_id"`
	FirstName       string  `cql:"c_first"`
	MiddleName      string  `cql:"c_middle"`
	LastName        string  `cql:"c_last"`
	Street1         string  `cql:"c_street_1"`
	Street2         string  `cql:"c_street_2"`
	City            string  `cql:"c_city"`
	State           string  `cql:"c_state"`
	Zip             string  `cql:"c_zip"`
	Phone           string  `cql:"c_phone"`
	CreationTime    int64   `cql:"c_since"`
	CreditStatus    string  `cql:"c_credit"`
	CreditLimit     float64 `cql:"c_credit_lim"`
	Discount        float64 `cql:"c_discount"`
	Balance         float64 `cql:"c_balance"`
	YtdPayment      float64 `cql:"c_ytd_payment"`
	NumPaymentMade  int32   `cql:"c_payment_cnt"`
	NumDeliveryMade int32   `cql:"c_delivery_cnt"`
	Data            string  `cql:"c_data"`
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
	Ytd     float64 `cql:"w_ytd"`
}

type Item struct {
	ID      int32  `cql:"i_id"`
	Name    string `cql:"i_name"`
	Price   string `cql:"i_price"`
	ImageID int32  `cql:"i_im_id"`
	Data    string `cql:"i_data"`
}
