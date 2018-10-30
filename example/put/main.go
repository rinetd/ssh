package main

import (
	"github.com/pytool/ssh"
)

func main() {

	client, err := ssh.NewClient("root", "localhost", "22", "ubuntu")
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
