package ssh

import (
	"testing"
)

func TestClient_IsCheck(t *testing.T) {
	c := GetClient()
	defer c.Close()
	var remotes = []string{
		"/root/test/notExist",
		"/root/test/notExist/",
		"/root/test/file",
		"/root/test/file/", // 不存在
		"/root/test/dir",
		"/root/test/dir/",
	}

	// /root/test/file 		存在
	// /root/test/file/ 	不存在
	// /root/test/dir 		存在
	// /root/test/dir/ 		存在
	for _, v := range remotes {
		is := c.IsExist(v)
		if is {
			println(v, "\t存在")
		} else {
			println(v, "\t不存在")
		}
	}

	// /root/test/file 		不是一个目录
	// /root/test/file/ 	不是一个目录
	// /root/test/dir 		是一个目录
	// /root/test/dir/ 		是一个目录
	println()
	for _, v := range remotes {
		isdir := c.IsDir(v)
		if isdir {
			println(v, "\t是一个目录")
		} else {
			println(v, "\t不是一个目录")
		}
	}

	// /root/test/file 		是一个文件
	// /root/test/file/ 	不是一个文件
	// /root/test/dir 		不是一个文件
	// /root/test/dir/ 		不是一个文件
	println()
	for _, v := range remotes {
		isfile := c.IsFile(v)
		if isfile {
			println(v, "\t是一个文件")
		} else {
			println(v, "\t不是一个文件")
		}

	}

}
