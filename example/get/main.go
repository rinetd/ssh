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
	// download dir
	var local = "/home/ubuntu/go/src/github.com/pytool/ssh/test/download/"
	client.Download(remotedir, local)

	// upload file
	var remotefile = "/root/test/file"

	client.Download(remotefile, local)

}
