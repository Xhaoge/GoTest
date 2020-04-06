package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = "3306"
	DATABASE = "myblog"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int //0 正常状态， 1 删除
	Createtime int64
}

// 查询单行；
func QueryOne(db *sql.DB) {
	fmt.Println("query times:", 1)
	user := new(User)
	row := db.QueryRow("select * from users where id=?", 1)
	// row.scan中的字段必须是按照数据库存入字段的顺序，否则报错；
	if err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Status, &user.Createtime); err != nil {
		fmt.Printf("scan failed,err:%v", err)
		return
	}
	fmt.Println(*user)
}

func main() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open mysql failed ,err:%v\n", err)
		return
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close；
}
