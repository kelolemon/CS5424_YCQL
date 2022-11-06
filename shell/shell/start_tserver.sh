rm -rf /temp/cs4224c/yugabyte
cp -r ./yugabyte-2.15.3.0 /temp/cs4224c/yugabyte
cd /temp/cs4224c/yugabyte
rm -rf /temp/cs4224c/hdd
mkdir -p /temp/cs4224c/hdd
./bin/post_install.sh
./bin/yb-tserver   \
	--tserver_master_addrs xcnd5.comp.nus.edu.sg:7100,xcnd6.comp.nus.edu.sg:7100,xcnd8.comp.nus.edu.sg:7100 \
	--rpc_bind_addresses  "${1}":9100  \
	--enable_ysql   \
	--pgsql_proxy_bind_address "${1}":5433  \
	--cql_proxy_bind_address "${1}":9142   \
	--fs_data_dirs "/temp/cs4224c/hdd" \
	--redis_proxy_bind_address "${1}":6479 \
	--redis_proxy_webserver_port 11100 \
	--client_read_write_timeout_ms=10000000 \
	>& /temp/cs4224c/ygdb_tserver.log &
