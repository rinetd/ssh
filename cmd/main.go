package main

import (
	"fmt"

	"github.com/pytool/ssh"
)

func main() {

	client, err := ssh.NewClient("root", "localhost", "22", "ubuntu")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	output, err := client.Exec("uptime")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Uptime: %s\n", output)

}
