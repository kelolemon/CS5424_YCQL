kill -9 $(ps aux | grep "./bin/yb-tserver" | awk '{print $2}')
cd /temp/cs4224c/yugabyte
./bin/yb-tserver   \
--tserver_master_addrs xcnd5.comp.nus.edu.sg:7100,xcnd6.comp.nus.edu.sg:7100,xcnd8.comp.nus.edu.sg:7100 \
--rpc_bind_addresses  "${1}":9100  \
--enable_ysql   \
--pgsql_proxy_bind_address "${1}":5533  \
--cql_proxy_bind_address "${1}":9142   \
--fs_data_dirs "/mnt/ramdisk/cs4224c/disk,/mnt/ramdisk/cs4224c/disk2" \
--redis_proxy_bind_address "${1}":6479 \
--redis_proxy_webserver_port 11100 \
--max_stale_read_bound_time_ms=10000000 \
--client_read_write_timeout_ms=10000000 \
>& /temp/cs4224c/ygdb_tserver.log &
