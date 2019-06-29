// Created by Hisen at 2019-06-25.
package config

import "sync"

var DbConfig *DBConfig

type DBConfig struct {
	IPAddr       string `json:"ip_addr"`
	Port         int    `json:"port"`
	AuthUser     string `json:"auth_user"`
	AuthPassword string `json:"auth_password"`
	DatabaseName string `json:"database_name"`
	TableName    string `json:"table_name"`
	TimeOut      int    `json:"time_out"`
	sync.RWMutex
}

// 无需这么麻烦,只是测试写法,正常写init函数即可
func (d *DBConfig) InitConfig() (dbconfig *DBConfig) {
	d.Lock()
	defer d.Unlock()
	dbconfig = &DBConfig{
		IPAddr:       "127.0.0.1",
		Port:         3306,
		AuthUser:     "hanxin",
		AuthPassword: "hanxin",
		DatabaseName: "book_store",
		TableName:    "book_info",
		TimeOut:      10,
	}
	return
}

func init() {
	var dbconfig DBConfig
	DbConfig = dbconfig.InitConfig()
}
