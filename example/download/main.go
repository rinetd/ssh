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
	// download dir
	var local = "/home/ubuntu/go/src/github.com/pytool/ssh/test/download/"
	client.Download(remotedir, local)

	// upload file
	var remotefile = "/root/test/file"

	client.Download(remotefile, local)

}
