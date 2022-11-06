pip3 install -r req.txt
python3 client.py --benchmark=true 2>metrics.out 1>response.out
echo "client_number,measurement_a,measurement_b,measurement_c,measurement_d,measurement_e,measurement_f,measurement_g" >> clients.csv
cat metrics.out >> clients.csv
echo "min_throughput, max_throughput, avg_throughput" >> throughput.csv
awk -F " " 'BEGIN {max = 0} {if ($3 > max) max=$3 } END {print max}' metrics.out > max_throughput
awk -F " " 'BEGIN {min = 0} {if ($3 < min) min=$3 } END {print min}' metrics.out > min_throughput
awk -F " " 'BEGIN {sum = 0} {sum += $3} END {print sum/NR}' metrics.out > abg_throughput
paste -d, max_throughput min_throughput abg_throughput >> throughput.csv
