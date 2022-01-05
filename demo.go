package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/websocket" //这里使用的是 gorilla 的 websocket 库
	"github.com/tidwall/gjson"     //处理json
	"golang.org/x/net/proxy"
	"gopkg.in/ini.v1"
)

type Symbol struct {
	Symbol string
	C      float64
	H      float64
	I      float64
}

var ch = make(chan int, 1) //通道
var pool *redis.Pool       //创建redis连接池
var coinname string

func init() {
	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	cfg, _ := ini.Load("my.ini")
	coinname = cfg.Section("").Key("coinname").String() //取配置文件
	//go Limit()                                          //定时处理
	id := 5 //1代理 5直连接
	go apisocket(id)
	for {
		<-ch                        //通道取数据
		time.Sleep(time.Second * 2) //休眠2秒
		apisocket(id)

	}

}

func apisocket(s int) {

	var id int = s
	dialer := websocket.Dialer{}
	if id < 2 {
		//代理
		netDialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, nil)
		if err != nil {
			log.Println(err)
		}
		dialer = websocket.Dialer{NetDial: netDialer.Dial}
	}

	//dialer := websocket.Dialer{}
	//向服务器发送连接请求，websocket 统一使用 ws://，默认端口和http一样都是80
	//coinstring = "btcusdt@miniTicker"
	wss := "wss://stream.binance.com:9443/stream?streams="
	connect, _, err := dialer.Dial(wss+coinname, nil)
	if nil != err {
		ch <- 10
		fmt.Print("connecterro")
		//log.Println(err)
		return
	}
	//离开作用域关闭连接，go 的常规操作
	defer connect.Close()

	//定时向客户端发送数据
	//go tickWriter(connect)

	//启动数据读取循环，读取客户端发送来的数据
	for {
		//从 websocket 中读取数据
		//messageType 消息类型，websocket 标准
		//messageData 消息数据
		messageType, messageData, err := connect.ReadMessage()
		if nil != err {
			log.Println(err)
			break
		}
		switch messageType {
		case websocket.TextMessage: //文本数据
			Redis_set(string(messageData))
			//fmt.Println(string(messageData))
		case websocket.BinaryMessage: //二进制数据
			//fmt.Println(messageData)
		case websocket.CloseMessage: //关闭
			//fmt.Println("close")
		case websocket.PingMessage: //Ping
			fmt.Println("ping")
		case websocket.PongMessage: //Pong
		default:

		}

	}
	ch <- 10 //退出
}

func Redis_set(str string) {
	//处理json
	coinname := gjson.Get(str, "data.s")
	cointype := strings.ToLower(coinname.String()) //转小写

	var symbol Symbol
	symbol.Symbol = cointype
	symbol.C, _ = strconv.ParseFloat(gjson.Get(str, "data.c").String(), 64)
	symbol.H, _ = strconv.ParseFloat(gjson.Get(str, "data.h").String(), 64)
	symbol.I, _ = strconv.ParseFloat(gjson.Get(str, "data.l").String(), 64)

	jsonstr, _ := json.Marshal(symbol) //返回json

	//fmt.Println(symbol)
	c := pool.Get() //从连接池，取一个链接
	defer c.Close() //函数运行结束 ，把连接放回连接池

	_, err := c.Do("HSet", "coinprice", cointype, jsonstr) //放入hash
	if err != nil {
		fmt.Println(err)
		return
	}
	//pool.Close() //关闭连接池

}

func tickWriter(connect *websocket.Conn) {
	student := make(map[string]interface{})
	student["pong"] = 1492420473027
	b, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	for {
		//向客户端发送类型为文本的数据
		err := connect.WriteMessage(websocket.PongMessage, []byte(b))
		if nil != err {
			log.Println(err)
			break
		}
		//休息一秒
		time.Sleep(time.Second * 5)
	}
}
