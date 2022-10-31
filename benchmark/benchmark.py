import time

import requests
import json


def main():
    api_url = "http://127.0.0.1:8080/api/cql"
    header = {'Content-Type': 'application/json'}
    for nub in range(0, 20):
        dirname = "../project_files/xact_files/"
        filename = str(nub) + ".txt"
        with open(dirname + filename, 'r+') as f:
            s = [line[:-1].split(',') for line in f.readlines()]
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
                    begin_time = time.time()
                    request = requests.post(api_url + '/payment', headers=header, data=json.dumps(request_str))
                    end_time = time.time()
                    response = request.content


if __name__ == "__main__":
    main()
