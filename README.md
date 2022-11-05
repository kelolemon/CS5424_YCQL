# CS5424_YCQL







import data from project_files/data_files/*.csv:

```shell
# deprecated
# copy warehouse from '../../project_files/data_files/warehouse.csv' WITH delimiter=',' AND HEADER=FALSE;
# copy district from '../../project_files/data_files/district.csv' WITH delimiter=',' AND HEADER=FALSE;
# copy item from '../../project_files/data_files/item.csv' WITH delimiter=',' AND HEADER=FALSE;
# copy stock from '../../project_files/data_files/stock.csv' WITH delimiter=',' AND HEADER=FALSE;
# copy customer from '../../project_files/data_files/customer.csv' WITH delimiter=',' AND HEADER=FALSE;
# copy order from '../../project_files/data_files/order.csv' WITH delimiter=',' AND HEADER=FALSE;
# copy orderline from '../../project_files/data_files/order-line.csv' WITH delimiter=',' AND HEADER=FALSE;

./cassandra-loader -f project_files/data_files/warehouse.csv -host localhost -schema "cs5424.warehouse(w_id, w_zip, w_name, w_street_1, w_street_2, w_city, w_state, w_tax, w_ytd)"
./cassandra-loader -f project_files/data_files/district.csv -host localhost -schema "cs5424.district(d_w_id, d_id, d_zip, d_name, d_street_1, d_street_2, d_city, d_state, d_tax, d_ytd, d_next_o_id)"
./cassandra-loader -f project_files/data_files/item.csv -host localhost -schema "cs5424.item(i_id, i_name, i_price, i_im_id, i_data)"
./cassandra-loader -f project_files/data_files/order.csv -host localhost -schema "cs5424."order"(o_w_id, o_d_id, o_id, o_c_id, o_carrier_id, o_ol_cnt, o_all_local, o_entry_d)" -dateFormat "YYYY-MM-DD hh:mm:ss"
./cassandra-loader -f project_files/data_files/order-line.csv -host localhost -schema "cs5424.orderline(ol_w_id, ol_d_id, ol_o_id, ol_number, ol_i_id, ol_delivery_d, ol_amount, ol_supply_w_id, ol_quantity, ol_dist_info)" -dateFormat "YYYY-MM-DD hh:mm:ss"
./cassandra-loader -f project_files/data_files/customer.csv -host localhost -schema "cs5424.customer(c_w_id, c_d_id, c_id, c_zip, c_first, c_middle, c_last, c_street_1, c_street_2, c_city, c_state, c_phone, c_since, c_credit, c_credit_lim, c_discount, c_balance, c_ytd_payment, c_payment_cnt, c_delivery_cnt, c_data)" -dateFormat "YYYY-MM-DD hh:mm:ss"
./cassandra-loader -f project_files/data_files/stock.csv -host localhost -schema "cs5424.stock(s_w_id, s_i_id, s_quantity, s_ytd, s_order_cnt, s_remote_cnt, s_dist_01, s_dist_02, s_dist_03, s_dist_04, s_dist_05, s_dist_06, s_dist_07, s_dist_08, s_dist_09, s_dist_10, s_data)"
./cassandra-loader -f project_files/data_files/order-line-quantity-by-order.csv -host localhost -schema "cs5424.orderlinequantitybyorder(w_id, d_id, o_id, ol_quantity_map, items_id_name_map, c_id, o_entry_d, c_first, c_middle, c_last)" -dateFormat "YYYY-MM-DD hh:mm:ss"
./cassandra-loader -f project_files/data_files/customer-balance.csv -host localhost -schema "cs5424.customerbalance(c_id, c_w_id, c_d_id, c_balance, c_first, c_middle, c_last, w_name, d_name)"
./cassandra-loader -f project_files/data_files/order-by-customer.csv -host localhost -schema "cs5424.orderbycustomer(c_w_id, c_d_id, c_id, o_entry_d, c_first, c_middle, c_last, c_balance, c_last_o_id, o_carrier_id)" -dateFormat "YYYY-MM-DD hh:mm:ss"
```

