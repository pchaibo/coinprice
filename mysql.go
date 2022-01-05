package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {

	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/admin")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database
	//defer Db.Close() // 注意这行代码要写在上面err判断的下面
}

type Spotprice struct {
	Id     int64  `db:"id"`
	Symbol string `db:"symbol"`
}

func getcoin() {
	var coin []Spotprice
	err := Db.Select(&coin, "select id, symbol from fa_coin_spot_price where status=?", 1)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	//fmt.Println("select succ:", coin)
	var str string
	for _, v := range coin {
		str = str + "/" + v.Symbol + "@miniTicker"
		//fmt.Println(v.Symbol)
	}
	fmt.Println(str)
}
