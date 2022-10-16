package common

type Stock struct {
	StockWarehouseID int32   `cql:"S_W_ID"`
	StockItemID      int32   `cql:"S_I_ID"`
	StockQuantity    int32   `cql:"S_QUANTITY"`
	StockYTD         float64 `cql:"S_YTD"`
	StockOrderCnt    int32   `cql:"S_ORDER_CNT"`
	StockRemoteCnt   int32   `cql:"S_REMOTE_CNT"`
	StockDist1       string  `cql:"S_DIST_01"`
	StockDist2       string  `cql:"S_DIST_02"`
	StockDist3       string  `cql:"S_DIST_03"`
	StockDist4       string  `cql:"S_DIST_04"`
	StockDist5       string  `cql:"S_DIST_05"`
	StockDist6       string  `cql:"S_DIST_06"`
	StockDist7       string  `cql:"S_DIST_07"`
	StockDist8       string  `cql:"S_DIST_08"`
	StockDist9       string  `cql:"S_DIST_09"`
	StockDist10      string  `cql:"S_DIST_10"`
	StockData        string  `cql:"S_DATA"`
}

type District struct {
	DistrictWarehouseID int32   `cql:"D_W_ID"`
	DistrictID          int32   `cql:"D_ID"`
	DistrictName        string  `cql:"D_NAME"`
	DistrictStreet1     string  `cql:"D_STREET_1"`
	DistrictStreet2     string  `cql:"D_STREET_2"`
	DistrictCity        string  `cql:"D_CITY"`
	DistrictState       string  `cql:"D_STATE"`
	DistrictZip         string  `cql:"D_ZIP"`
	DistrictTax         float64 `cql:"D_TAX"`
	DistrictYTD         float64 `cql:"D_YTD"`
	DistrictNextOrderID int32   `cql:"D_NEXT_O_ID"`
}

type Customer struct {
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
	ID    int32  `cql:"i_id"`
	Name  string `cql:"i_name"`
	Price string `cql:"i_price"`
	IMID  int32  `cql:"i_im_id"`
	Data  string `cql:"i_daya"`
}
