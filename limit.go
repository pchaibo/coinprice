package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Limit() {
	go getorder()
	go coinspot() //现货
	getdate()     //合约

}

//合约
func getdate() {
	for true {
		httpget("task") //合约任务
		time.Sleep(time.Duration(2) * time.Second)
	}

}

//合约订单处理
func getorder() {
	var id int
	for true {
		id = id + 1
		ss := strconv.Itoa(id)
		fmt.Printf(ss + " \n")
		time.Sleep(time.Duration(10) * time.Second)
		httpget("getorder")

	}

}

//现货
func coinspot() {
	fmt.Printf("getspot：\n")
	id := 0
	for true {
		id++
		fmt.Println(id)
		//time.Sleep(100000000) //0.1秒
		httpgetcoinspot("task")
		time.Sleep(time.Duration(2) * time.Second)
	}
}
func httpgetcoinspot(add string) {
	// 使用Get方法获取服务器响应包数据
	resp, err := http.Get("http://c.55youtao.com/api/coinspot/" + add)
	if err != nil {
		fmt.Println("Get err:", err)
		return
	}
	defer resp.Body.Close()
	//
	f, err := os.Create(add)
	if err != nil {
		panic(err)
	}
	io.Copy(f, resp.Body)

}

func httpget(add string) {
	// 使用Get方法获取服务器响应包数据
	resp, err := http.Get("http://c.55youtao.com/api/coinheyue/" + add)
	if err != nil {
		fmt.Println("Get err:", err)
		return
	}
	defer resp.Body.Close()
	//
	f, err := os.Create(add)
	if err != nil {
		panic(err)
	}
	io.Copy(f, resp.Body)

}
