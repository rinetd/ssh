package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/pytool/ssh"

	"github.com/go-sql-driver/mysql"
)

var dsn = `lzkp:yqhtfjzm@tcp(192.168.5.100:3306)/?parseTime=true&loc=Local`
var DBNAME = "shizhi"
var db *sql.DB

func Prepare() {
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		// return FAIL, fmt.Errorf("Unable to open connection to database server: %s", err.Error())
		fmt.Print("")
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		// return FAIL, fmt.Errorf("Unable to ping database server: %s", err.Error())
		fmt.Print("")
	}
	// _, err = db.Exec("CREATE DATABASE IF NOT EXISTS" + DBNAME)
	// if err != nil {
	// 	// return FAIL, fmt.Errorf("Unable to create database %s: %s", DBNAME, err.Error())
	// 	fmt.Print("")
	// }
	// defer db.Exec("DROP DATABASE dbgrep")
	_, err = db.Exec("use " + DBNAME)
	if err != nil {
		fmt.Errorf("Unable to select database %s: %s", DBNAME, err.Error())
	}
	// return m.Run(), nil
}

func main() {
	// SSH的连接参数:
	config := ssh.Default.WithPassword("HR2018!!").WithHost("192.168.5.157")
	client, err := ssh.New(config)
	// client, err := ssh.NewClient("localhost", "22", "root", "ubuntu")
	if err != nil {
		panic(err)
	}
	defer client.Close()
	fmt.Println(client.Output("id"))

	// 1. 注册自定义的 Dial 命名为:mysql+ssh
	// Now we register the ViaSSHDialer with the ssh connection as a parameter
	mysql.RegisterDial("mysql+ssh", func(addr string) (net.Conn, error) {
		return client.SSHClient.Dial("tcp", addr)
	})

	// DB数据库的连接参数:
	dbUser := "root"           // DB username
	dbPass := ""               // DB Password
	dbHost := "localhost:3306" // DB Hostname/IP
	dbName := "shizhi"         // Database name
	// 2. 使用自定义命名为:mysql+ssh的 Dial 进行mysql连接
	// And now we can use our new driver with the regular mysql connection string tunneled through the SSH connection
	if db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@mysql+ssh(%s)/%s", dbUser, dbPass, dbHost, dbName)); err == nil {
		
		if rows, err := db.Query("SELECT user, host FROM mysql.user "); err == nil {
			for rows.Next() {
				var id string
				var name string
				rows.Scan(&id, &name)
				fmt.Printf("ID: %s\tName: %s\n", id, name)
			}
			rows.Close()
		} else {
			fmt.Printf("Failure: %s", err.Error())
		}

		db.Close()
		fmt.Printf("Successfully connected to the db\n")
	}
}
