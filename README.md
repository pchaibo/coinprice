## coinprice
<p align="center">
  <img width="800" src="https://github.com/pchaibo/coinprice/blob/master/images/b.png">
  <img width="800" src="https://github.com/pchaibo/coinprice/blob/master/images/red.png">
  
</p>
使用golang websocket采集币安时实价格

要求安装 redis 采集价格存放redis

## my.ini 为配置文件

coinname 为币类别 每个用/分隔后面加上@miniTicker

proxyid = 1 设置代理(使用本地127.0.0.1:1080端口国内使用)

proxyid = 5 直接连接币安websocket在服务器上使用
## 编译
加载此包：

 go get github.com/garyburd/redigo/redis 

 go get github.com/gorilla/websocket 

 go get github.com/tidwall/gjson

 go get golang.org/x/net/proxy

 go get gopkg.in/ini.v1 

编译文件： go build demo.go

## 运行
windos 直接执行 coinapi.exe

linux 执行后台运行 nohup ./demo >/dev/null 2>&1 &

