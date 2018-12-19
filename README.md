
## 项目简介
本项目是基于golang标准库 ssh 和 sftp 开发

本项目是对标准库进行一个简单的高层封装,使得可以在在 Windows Linux Mac 上非常容易的执行 ssh 命令,
以及文件,文件夹的上传,下载等操作.
1. 当src 为目录时
文件上传下载模仿rsync: 只和源有关.  
// rsync -av src/ dst     ./src/* --> /root/dst/*  
// rsync -av src/ dst/    ./src/* --> /root/dst/*  
// rsync -av src  dst     ./src/* --> /root/dst/src/*  
// rsync -av src  dst/    ./src/* --> /root/dst/src/*  
2. 当src 为文件时
当dst为目录，以"/"结尾，则自动拼接上文件名
当dst为文件，不以“/”结尾时，则重命名文件
## Install
`go get github.com/pytool/ssh`
## Example

### 在远程执行ssh命令
提供3个方法: Run() Exec() Output() 
1. Run() : 程序执行后,不再受执行者控制. 适用于启动服务端进程.
2. Exec() : 在控制台同步实时输出程序的执行结果.
3. Output() : 会等待程序执行完成后,输出执行结果,在需要对执行的结果进行操作时使用.
```go
package main
import (
	"fmt"
	"github.com/pytool/ssh"
)
func main() {

	c, err := ssh.NewClient("localhost", "22", "root", "ubuntu")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	output, err := c.Output("uptime")
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

	client, err := ssh.NewClient( "localhost", "22", "root", "ubuntu")
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

	client, err := ssh.NewClient( "localhost", "22", "root", "ubuntu")
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


