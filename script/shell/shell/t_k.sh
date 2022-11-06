kill -9 $(ps aux | grep "./bin/yb-tserver" | awk '{print $2}')
