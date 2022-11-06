rm -rf /temp/cs4224c/yugabyte
cp -r ./yugabyte-2.15.3.0 /temp/cs4224c/yugabyte
cd /temp/cs4224c/yugabyte
rm -rf /mnt/ramdisk/cs4224c
mkdir -p /mnt/ramdisk/cs4224c/disk
mkdir -p /mnt/ramdisk/cs4224c/disk2
./bin/post_install.sh
./bin/yb-master  \
--fs_data_dirs "/mnt/ramdisk/cs4224c/disk,/mnt/ramdisk/cs4224c/disk2" \
--master_addresses xcnd5.comp.nus.edu.sg:7100,xcnd6.comp.nus.edu.sg:7100,xcnd8.comp.nus.edu.sg:7100 \
--rpc_bind_addresses  "$1":7100 \
--webserver_port 7000 \
--replication_factor=3 \
>& /temp/cs4224c/ygdb_master.log &
./bin/yb-tserver   \
--tserver_master_addrs xcnd5.comp.nus.edu.sg:7100,xcnd6.comp.nus.edu.sg:7100,xcnd8.comp.nus.edu.sg:7100 \
--rpc_bind_addresses  "${1}":9100  \
--enable_ysql   \
--pgsql_proxy_bind_address "${1}":5533  \
--cql_proxy_bind_address "${1}":9142   \
--fs_data_dirs "/mnt/ramdisk/cs4224c/disk,/mnt/ramdisk/cs4224c/disk2" \
--redis_proxy_bind_address "${1}":6479 \
--redis_proxy_webserver_port 11100 \
--client_read_write_timeout_ms=10000000 \
>& /temp/cs4224c/ygdb_tserver.log &