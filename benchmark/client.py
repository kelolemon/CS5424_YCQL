import sys
import threading
import time
import requests
import json
import argparse
import numpy as np


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


def handler(s, nub):
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
            w_id = int(s[i][2])
            d_id = int(s[i][3])
            c_id = int(s[i][1])
            num_items = int(s[i][4])
            item_ids = []
            supply_warehouse = []
            quantity = []
            for cnt in range(0, num_items):
                line = i + cnt + 1
                item_ids.append(int(s[line][0]))
                supply_warehouse.append(int(s[line][1]))
                quantity.append(int(s[line][2]))
            request_str = {
                'w_id': w_id,
                'd_id': d_id,
                'c_id': c_id,
                'number_items': num_items,
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
            w_id = int(s[i][1])
            carrier_id = int(s[i][2])
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
            c_w_id = int(s[i][1])
            c_d_id = int(s[i][2])
            c_id = int(s[i][3])
            request_str = {
                'c_w_id': c_w_id,
                'c_d_id': c_d_id,
                'c_id': c_id
            }
            response, latency = send_request('GET', header, request_str, api_url + '/status')
            time_counter += latency
            requests_counter += 1
            latency_list.append(latency)
            # print response content at stdout
            print(response)
        elif trans_type == 'S':
            w_id = int(s[i][1])
            d_id = int(s[i][2])
            t = int(s[i][3])
            l = int(s[i][4])
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
            w_id = int(s[i][1])
            d_id = int(s[i][2])
            l = int(s[i][3])
            request_str = {
                'w_id': w_id,
                'd_id': d_id,
                'num_last_orders': l
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
            c_w_id = int(s[i][1])
            c_d_id = int(s[i][2])
            c_id = int(s[i][3])
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
        else:
            pass
    print(str(nub) + "," + str(requests_counter) + "," + str(time_counter) + "," + str(requests_counter / time_counter )
          + "," + str(np.mean(latency_list * 1000)) + "," + str(np.median(latency_list) * 1000) + "," + str(np.percentile(latency_list, 95) * 1000) + ","
          + str(np.percentile(latency_list, 99) * 1000), file=sys.stderr)


def read_from_file(nub):
    dirname = "../project_files/xact_files/"
    filename = str(nub) + ".txt"
    with open(dirname + filename, 'r+') as f:
        s = [line[:-1].split(',') for line in f.readlines()]
        return s


def read_from_stdin():
    s = [line[:-1].split(',') for line in sys.stdin.readlines()]
    return s


def benchmark():
    threads = []
    for i in range(0, 5):
        threads.append(threading.Thread(target=handler, args=(read_from_file(i), i,)))
    for thread in threads:
        thread.start()
    for thread in threads:
        if thread.is_alive():
            thread.join()


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--benchmark', type=bool, default=None)
    args = parser.parse_args()
    if args.benchmark is not None:
        benchmark()
    else:
        s = read_from_stdin()
        handler(s, -1)


if __name__ == "__main__":
    main()
