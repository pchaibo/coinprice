## coinprice

使用golang websocket采集币安时实价格

使用包:gjson gorilla/websocket ini.v1 proxy redis 编译前go get * 加载此包

要求安装 redis 采集价格存放redis

## my.ini 为配置文件

coinname 为币类别 每个用/分隔后面加上@miniTicker

proxyid = 1 设置代理(使用本地127.0.0.1:1080端口国内使用)

proxyid = 5 直接连接币安websocket在服务器上使用

编译文件： go build demo.go

linux 执行后台运行 nohup ./demo >/dev/null 2>&1 &

windos 上执行 coinapi.exe