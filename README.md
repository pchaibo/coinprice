## coinprice

使用golang websocket采集币安时实价格

要求安装 redis 采集价格存放redis

demo.go 设置是否代理在45行

id := 1: 1代理(使用本地127.0.0.1:1080端口国内使用) 5：直接连接币安websocket在服务器上使用

编译文件： go build demo.go

linux 执行后台运行 nohup ./demo >/dev/null 2>&1 &

windos 上执行 coinapi.exe