package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var MySQLDB *sql.DB
var MysqlDbErr error

const (
	UserName = ""
	PassWord = ""
	HOST     = "localhost"
	PORT     = "3306"
	DATABASE = "qukeduo"
	CHARSET  = "utf8"
)

// automatic run when runtime
func init() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", UserName, PassWord, HOST, PORT, DATABASE, CHARSET)

	// 打开连接失败
	MySQLDB, MysqlDbErr = sql.Open("mysql", dbDSN)
	//defer MySQLDB.Close();
	if MysqlDbErr != nil {
		log.Println("dbDSN: " + dbDSN)
		panic("数据源配置不正确: " + MysqlDbErr.Error())
	}

	// 最大连接数
	MySQLDB.SetMaxOpenConns(100)
	// 闲置连接数
	MySQLDB.SetMaxIdleConns(20)
	// 最大连接周期
	MySQLDB.SetConnMaxLifetime(100 * time.Second)

	if MysqlDbErr = MySQLDB.Ping(); nil != MysqlDbErr {
		panic("数据库链接失败: " + MysqlDbErr.Error())
	}
}
