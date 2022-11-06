kill -9 $(ps aux | grep "./bin/yb-master" | awk '{print $2}')
kill -9 $(ps aux | grep "./bin/yb-tserver" | awk '{print $2}')
