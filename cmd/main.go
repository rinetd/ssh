package main

import (
	"fmt"

	"github.com/pytool/ssh"
)

func main() {
	config := ssh.Default.WithPassword("ubuntu")
	client, err := ssh.New(config)
	// client, err := ssh.NewClient("localhost", "22", "root", "ubuntu")
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
