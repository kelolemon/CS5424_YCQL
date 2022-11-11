# CS5424_YCQL

## config 
the config file is in config.yaml, we can config cluster ips and cql ports in it.
here is a sample 
```yaml
port: 9142
ips:
  - "xcnd5.comp.nus.edu.sg"
  - "xcnd6.comp.nus.edu.sg"
  - "xcnd7.comp.nus.edu.sg"
  - "xcnd8.comp.nus.edu.sg"
  - "xcnd50.comp.nus.edu.sg"
```

## API
+ POST /api/cql/order corresponds to the trans1
```go
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
```
+ POST /api/cql/payment corresponds to the trans2 
```go
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
```
+ POST /api/cql/delivery corresponds to the trans3 
```go
type CreateNewDeliveryReq struct {
	WarehouseID int32 `json:"w_id"`
	CarrierID   int32 `json:"carrier_id"`
}

type CreateNewDeliveryResp struct {
}
```
+ GET /api/cql/status corresponds to the trans4 
```go
type GetLastOrderStatusReq struct {
	WarehouseID int32 `form:"c_w_id"`
	DistrictID  int32 `form:"c_d_id"`
	CustomerID  int32 `form:"c_id"`
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
```
+ GET /api/cql/stock corresponds to the trans5 
```go
type GetStockLevelLowItemNumberReq struct {
	WarehouseID    int32 `form:"w_id"`
	DistrictID     int32 `form:"d_id"`
	StockThreshold int32 `form:"t"`
	LastOrders     int32 `form:"l"`
}

type GetStockLevelLowItemNumberResp struct {
	StockLevelLowItemNumber int32 `json:"stock_level_low_item_number"`
}
```
+ GET /api/cql/item corresponds to the trans6 
```go
type GetPopularItemReq struct {
	WarehouseID   int32 `form:"w_id"`
	DistrictID    int32 `form:"d_id"`
	NumLastOrders int32 `form:"num_last_orders"`
}

type GetPopularItemResp struct {
	WarehouseID            int32
	DistrictID             int32
	NumLastOrders          int32
	OrderPopularItems      []*OrderPopularItem
	PopularItemPercentages []*PopularItemPercentage
}
```
+ GET /api/cql/transaction corresponds to the trans7 
```go
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
```
+ GET /api/cql/customer corresponds to the trans8
```go
type GetRelatedCustomerReq struct {
	WarehouseID int32 `form:"c_w_id"`
	DistrictID  int32 `form:"c_d_id"`
	CustomerID  int32 `form:"c_id"`
}

type CustomerList struct {
	WarehouseID int32 `json:"c_w_id"`
	DistrictID  int32 `json:"c_d_id"`
	CustomerID  int32 `json:"c_id"`
}

type GetRelatedCustomerResp struct {
	CustomerList []CustomerList `json:"customer_list"`
}
```

## deployment
+ server start
```shell
# check the config/config.yaml before run server
go get -u 
go mod tidy
go build main.go
./main
```
+ client start
```shell
cd client
pip3 install -r req.txt
python3 client.py
# then input transaction, and stop input by using ctr+d
```
+ benchmark start
```shell
cd client
chmod +x benchmark.sh
./benchmark.sh
```
+ cluster start 
```shell
cd shell
./sacluster.sh
```
+ load_data
```shell
# copy warehouse from '../../data_files/warehouse.csv' WITH delimiter=',' AND HEADER=FALSE;
# copy district from '../../data_files/district.csv' WITH delimiter=',' AND HEADER=FALSE;
# copy item from '../../data_files/item.csv' WITH delimiter=',' AND HEADER=FALSE;
# copy stock from '../../data_files/stock.csv' WITH delimiter=',' AND HEADER=FALSE;
# copy customer from '../../data_files/customer.csv' WITH delimiter=',' AND HEADER=FALSE;
# copy order from '../../data_files/order.csv' WITH delimiter=',' AND HEADER=FALSE;
# copy orderline from '../../data_files/order-line.csv' WITH delimiter=',' AND HEADER=FALSE;
# using the github cassandra-loader
wget https://github.com/brianmhess/cassandra-loader/releases/download/v0.0.27/cassandra-loader

./cassandra-loader -f data_files/warehouse.csv -host xcnd5.comp.nus.edu.sg -port 9142 -schema "cs5424.warehouse(w_id, w_zip, w_name, w_street_1, w_street_2, w_city, w_state, w_tax, w_ytd)"
./cassandra-loader -f data_files/district.csv -host xcnd5.comp.nus.edu.sg -port 9142 -schema "cs5424.district(d_w_id, d_id, d_zip, d_name, d_street_1, d_street_2, d_city, d_state, d_tax, d_ytd, d_next_o_id)"
./cassandra-loader -f data_files/item.csv -host xcnd5.comp.nus.edu.sg -port 9142 -schema "cs5424.item(i_id, i_name, i_price, i_im_id, i_data)"
./cassandra-loader -f data_files/order.csv -host xcnd5.comp.nus.edu.sg -port 9142 -schema "cs5424."order"(o_w_id, o_d_id, o_id, o_c_id, o_carrier_id, o_ol_cnt, o_all_local, o_entry_d)" -dateFormat "YYYY-MM-DD hh:mm:ss"
./cassandra-loader -f data_files/order-line.csv -host xcnd5.comp.nus.edu.sg -port 9142 -schema "cs5424.orderline(ol_w_id, ol_d_id, ol_o_id, ol_number, ol_i_id, ol_delivery_d, ol_amount, ol_supply_w_id, ol_quantity, ol_dist_info)" -dateFormat "YYYY-MM-DD hh:mm:ss"
./cassandra-loader -f data_files/customer.csv -host xcnd5.comp.nus.edu.sg -port 9142 -schema "cs5424.customer(c_w_id, c_d_id, c_id, c_zip, c_first, c_middle, c_last, c_street_1, c_street_2, c_city, c_state, c_phone, c_since, c_credit, c_credit_lim, c_discount, c_balance, c_ytd_payment, c_payment_cnt, c_delivery_cnt, c_data)" -dateFormat "YYYY-MM-DD hh:mm:ss"
./cassandra-loader -f data_files/stock.csv -host xcnd5.comp.nus.edu.sg -port 9142 -schema "cs5424.stock(s_w_id, s_i_id, s_quantity, s_ytd, s_order_cnt, s_remote_cnt, s_dist_01, s_dist_02, s_dist_03, s_dist_04, s_dist_05, s_dist_06, s_dist_07, s_dist_08, s_dist_09, s_dist_10, s_data)"
./cassandra-loader -f data_files/customer-balance.csv -host xcnd5.comp.nus.edu.sg -port 9142 -schema "cs5424.customerbalance(c_w_id, c_d_id, c_id, c_balance, c_first, c_middle, c_last, w_name, d_name)"
./cassandra-loader -f data_files/order-by-customer.csv -host xcnd5.comp.nus.edu.sg -port 9142 -schema "cs5424.orderbycustomer(c_w_id, c_d_id, c_id, o_entry_d, c_first, c_middle, c_last, c_balance, c_last_o_id, o_carrier_id)" -dateFormat "YYYY-MM-DD hh:mm:ss"
./cassandra-loader -f data_files/order-line-quantity-by-order.csv -host xcnd5.comp.nus.edu.sg -port 9142 -schema "cs5424.orderlinequantitybyorder(w_id, d_id, o_id, ol_quantity_map, items_id_name_map, c_id, o_entry_d, c_first, c_middle, c_last)" -dateFormat "YYYY-MM-DD hh:mm:ss"
```
