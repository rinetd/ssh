package main

import (
	"github.com/pytool/ssh"
)

func main() {

	client, err := ssh.NewClient("localhost", "22", "root", "ubuntu")
	if err != nil {
		panic(err)
	}
	defer client.Close()
	var remotedir = "/root/test/"
	// upload dir
	var local = "/home/ubuntu/go/src/github.com/pytool/ssh/test/upload/"
	client.Upload(local, remotedir)
	// upload file
	local = "/home/ubuntu/go/src/github.com/pytool/ssh/test/upload/file"
	client.Upload(local, remotedir)

}
