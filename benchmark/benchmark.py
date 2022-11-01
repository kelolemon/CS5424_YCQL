import time
import requests
import json


def send_request(request_type, request_header, request_data, request_url):
    if request_type == 'POST':
        begin_time = time.time()
        request = requests.post(request_url, headers=request_header, data=request_data)
        end_time = time.time()
        response = request.content
        latency = end_time - begin_time
        return response, latency
    if request_type == 'GET':
        begin_time = time.time()
        request = requests.get(request_url, params=request_data)
        end_time = time.time()
        response = request.content
        latency = end_time - begin_time
        return response, latency


def handler(s):
    time_counter = 0
    requests_counter = 0
    latency_list = []
    api_url = "http://127.0.0.1:8080/api/cql"
    header = {'Content-Type': 'application/json'}
    for i in range(0, len(s)):
        trans_type = s[i][0]
        if trans_type == 'P':
            c_w_id = int(s[i][1])
            c_d_id = int(s[i][2])
            c_id = int(s[i][3])
            payment = float(s[i][4])
            request_str = {
                'c_w_id': c_w_id,
                'c_d_id': c_d_id,
                'c_id': c_id,
                'payment': payment
            }
            response, latency = send_request('POST', header, json.dumps(request_str), api_url + '/payment')
            time_counter += latency
            requests_counter += 1
            latency_list.append(latency)
            # print response content at stdout
            print(response)
        elif trans_type == 'N':
            w_id = s[i][1]
            d_id = s[i][2]
            c_id = s[i][3]
            num_items = s[i][4]
            item_ids = []
            supply_warehouse = []
            quantity = []
            for cnt in range(0, num_items):
                i += 1
                item_ids.append(s[i][0])
                supply_warehouse.append(s[i][1])
                quantity.append(s[i][2])
            request_str = {
                'w_id': w_id,
                'd_id': d_id,
                'c_id': c_id,
                'num_items': num_items,
                'item_number': item_ids,
                'supply_warehouse': supply_warehouse,
                'quantity': quantity
            }
            response, latency = send_request('POST', header, json.dumps(request_str), api_url + '/order')
            time_counter += latency
            requests_counter += 1
            latency_list.append(latency)
            # print response content at stdout
            print(response)
        elif trans_type == 'D':
            w_id = s[i][1]
            carrier_id = s[i][2]
            request_str = {
                'w_id': w_id,
                'carrier_id': carrier_id
            }
            response, latency = send_request('POST', header, json.dumps(request_str), api_url + '/delivery')
            time_counter += latency
            requests_counter += 1
            latency_list.append(latency)
            # print response content at stdout
            print(response)
        elif trans_type == 'O':
            c_w_id = s[i][1]
            c_d_id = s[i][2]
            c_id = s[i][3]
            request_str = {
                'c_w_id': c_w_id,
                'c_d_id': c_d_id,
                'c_id': c_id
            }
            response, latency = send_request('GET', header, request_str, api_url + '/order_status')
            time_counter += latency
            requests_counter += 1
            latency_list.append(latency)
            # print response content at stdout
            print(response)
        elif trans_type == 'S':
            w_id = s[i][1]
            d_id = s[i][2]
            t = s[i][3]
            l = s[i][4]
            request_str = {
                'w_id': w_id,
                'd_id': d_id,
                't': t,
                'l': l
            }
            response, latency = send_request('GET', header, request_str, api_url + '/stock')
            time_counter += latency
            requests_counter += 1
            latency_list.append(latency)
            # print response content at stdout
            print(response)
        elif trans_type == 'I':
            w_id = s[i][1]
            d_id = s[i][2]
            l = s[i][3]
            request_str = {
                'w_id': w_id,
                'd_id': d_id,
                'l': l
            }
            response, latency = send_request('GET', header, request_str, api_url + '/item')
            time_counter += latency
            requests_counter += 1
            latency_list.append(latency)
            # print response content at stdout
            print(response)
        elif trans_type == 'T':
            response, latency = send_request('GET', header, {}, api_url + '/transaction')
            time_counter += latency
            requests_counter += 1
            latency_list.append(latency)
            # print response content at stdout
            print(response)
        elif trans_type == 'R':
            c_w_id = s[i][1]
            c_d_id = s[i][2]
            c_id = s[i][3]
            request_str = {
                'c_w_id': c_w_id,
                'c_d_id': c_d_id,
                'c_id': c_id
            }
            response, latency = send_request('GET', header, request_str, api_url + '/customer')
            time_counter += latency
            requests_counter += 1
            latency_list.append(latency)
            # print response content at stdout
            print(response)


def benchmark():
    handler([])


if __name__ == "__main__":
    benchmark()
