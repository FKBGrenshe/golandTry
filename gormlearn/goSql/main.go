package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := initialConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	if db != nil {
		defer db.Close()
	}

}

// 初始化db连接
func initialConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gormlearn?charset=utf8mb4&parseTime=True&loc=Local")
	db.SetMaxOpenConns(100)          // 最大打开连接数
	db.SetMaxIdleConns(10)           // 最大空闲连接数
	db.SetConnMaxLifetime(time.Hour) // 连接最大存活时间

	if err = db.Ping(); err != nil {
		fmt.Println("initialConnection failed")
		return nil, err
	}
	fmt.Println("initialConnection success")

	return db, nil
}
