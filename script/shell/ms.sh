cp -r ./yugabyte-2.15.3.0 /temp/cs4224c/yugabyte && cd /temp/cs4224c/yugabyte
./bin/post_install.sh
./bin/yb-master  \
--fs_data_dirs "/temp/cs4224c/disk" \
--master_addresses xcnd5.comp.nus.edu.sg:7100,xcnd6.comp.nus.edu.sg:7100,xcnd7.comp.nus.edu.sg:7100,xcnd8.comp.nus.edu.sg:7100,xcnd50.comp.nus.edu.sg:7100 \
--rpc_bind_addresses  "$1":7100 \
--webserver_port 7000 \
--replication_factor=5 \
>& /temp/cs4224c/ygdb_master.log &
