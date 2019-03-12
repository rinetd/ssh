package main

import (
	"fmt"

	"github.com/pytool/ssh"
)

func main() {
	config := ssh.Default.WithHost("15.14.12.153").WithPassword("HR2018!!")
	// config.Host = "15.14.12.153"
	client, err := ssh.New(config)
	// client, err := ssh.NewClient("localhost", "22", "root", "ubuntu")
	if err != nil {
		// panic(err)
		fmt.Println("连接失败,按Enter键退出!")
		fmt.Scanln()
	}
	defer client.Close()

	err = client.Exec("sh /root/shetl/etl.sh")
	if err != nil {
		fmt.Println(err)
		// panic(err)
		fmt.Println("执行失败,按Enter键退出!")
		fmt.Scanln()
	}

	// fmt.Printf("Uptime: %s\n", output)

}
