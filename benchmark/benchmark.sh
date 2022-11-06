pip3 install -r req.txt
python3 client.py --benchmark=true 2>metrics.out 1>response.out
echo "client_number,measurement_a,measurement_b,measurement_c,measurement_d,measurement_e,measurement_f,measurement_g" >> clients.csv
cat metrics.out >> clients.csv
echo "min_throughput, max_throughput, avg_throughput" >> throughput.csv
awk -F " " 'BEGIN {max = 0} {if ($3 > max) max=$3 } END {print max}' metrics.outpu > max_throughput
awk -F " " 'BEGIN {max = 0} {if ($3 > max) max=$3 } END {print max}' metrics.outpu > max_throughput
awk -F " " 'BEGIN {max = 0} {if ($3 > max) max=$3 } END {print max}' metrics.outpu > max_throughput