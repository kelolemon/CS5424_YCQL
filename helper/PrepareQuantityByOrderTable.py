import pandas as pd

# CREATE TABLE CS5424.OrderLineQuantityByOrder (
#   W_ID int,
#   D_ID int,
#   O_ID int,
#   O_ENTRY_D timestamp,
#   OL_QUANTITY_MAP map<int, int>,
#   ITEMS_ID_NAME_MAP map<int, text>,
#   C_ID int,
#   C_FIRST text,
#   C_MIDDLE text,
#   C_LAST text,
#   PRIMARY KEY ((W_ID, D_ID), O_ID)
# )
# WITH CLUSTERING ORDER By (O_ID DESC);

# Command to import data into Cassandra table
# COPY orderlinequantitybyorder(w_id, d_id, o_id, ol_quantity_map, items_id_name_map, c_id, o_entry_d, c_first, c_middle, c_last)
# FROM '/Users/yiyangliu/Desktop/NUS Master/CS5424/CS5424_YCQL/data_files/orderline-quantity-by-order.csv' WITH header = true;

# input csv
customer_csv = '../data_files/customer.csv'
item_csv = '../data_files/item.csv'
order_csv = '../data_files/order.csv'
order_line_csv = '../data_files/order-line.csv'
# output csv
order_line_quantity_by_order_csv = '../data_files/orderline-quantity-by-order.csv'

customer_df = pd.read_csv(customer_csv, header=None, usecols=[0, 1, 2, 3, 4, 5],
                          names=['w_id', 'd_id', 'c_id', 'c_first', 'c_middle', 'c_last'])
# print(customer_df)

item_df = pd.read_csv(item_csv, header=None, usecols=[0, 1], names=['i_id', 'i_name'])
# print(item_df)

order_df = pd.read_csv(order_csv, header=None, usecols=[0, 1, 2, 3, 7],
                       names=['w_id', 'd_id', 'o_id', 'c_id', 'o_entry_d'])
# print(order_df)

order_line_df = pd.read_csv(order_line_csv, header=None, usecols=[0, 1, 2, 3, 4, 8],
                            names=['w_id', 'd_id', 'o_id', 'ol_id', 'i_id', 'quantity'])
# print(order_line_df)

order_line_item_df = order_line_df.join(item_df.set_index('i_id'), on='i_id')
# print(order_line_item_df)

order_line_item_df['ol_quantity_map'] = order_line_item_df['i_id'].astype(str) + ":" + order_line_item_df['quantity'].astype(str)

order_line_item_df['items_id_name_map'] = order_line_item_df['i_id'].astype(str) + ":" + order_line_item_df['i_name']

order_line_quantity_by_order = order_line_item_df.groupby(['w_id', 'd_id', 'o_id'], as_index=False) \
    .agg({'ol_quantity_map': ', '.join, 'items_id_name_map': ', '.join}) \
    .join(order_df.set_index(['w_id', 'd_id', 'o_id']), on=['w_id', 'd_id', 'o_id']) \
    .join(customer_df.set_index(['w_id', 'd_id', 'c_id']), on=['w_id', 'd_id', 'c_id'])
order_line_quantity_by_order['ol_quantity_map'] = '{' + order_line_quantity_by_order['ol_quantity_map'] + '}'
order_line_quantity_by_order['items_id_name_map'] = '{' + order_line_quantity_by_order['items_id_name_map'] + '}'

print(order_line_quantity_by_order)

order_line_quantity_by_order.to_csv(order_line_quantity_by_order_csv, index=False)
