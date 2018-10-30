
## 项目简介
本项目是基于golang标准库 ssh 和 sftp 开发

本项目是对标准库进行一个简单的高层封装,使得可以在在 Windows Linux Mac 上非常容易的执行 ssh 命令,
以及文件,文件夹的上传,下载等操作.

文件上传下载模仿rsync方式: 只和源有关.
// rsync -av src/ dst     ./src/* --> /root/dst/*
// rsync -av src/ dst/    ./src/* --> /root/dst/*
// rsync -av src  dst     ./src/* --> /root/dst/src/*
// rsync -av src  dst/    ./src/* --> /root/dst/src/*

## Example

### 在远程执行ssh命令
```go
package main
import (
	"fmt"
	"github.com/pytool/ssh"
)
func main() {

	c, err := ssh.NewClient("root", "localhost", "22", "ubuntu")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	output, err := c.Exec("uptime")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Uptime: %s\n", output)
}

```
### 文件下载
```go
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

```

### 文件上传
```go
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

```


