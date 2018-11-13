package ssh

import (
	"sync"
	"testing"
)

func GetClient() *Client {
	var (
		once = sync.Once{}
		c    = &Client{}
		err  error
	)
	once.Do(func() {
		c, err = NewClient("localhost", "22", "root", "ubuntu")
	})
	if err != nil {
		panic(err)
	}
	return c
}

// func TestClient_RemoveAll(t *testing.T) {
// 	c := GetClient()
// 	defer c.Close()
// 	var remotedir = "/root/test/"
// 	fmt.Println(c.RemoveAll(remotedir))
// }
func TestClient_Init(t *testing.T) {
	c := GetClient()
	defer c.Close()
	var local = "/home/ubuntu/go/src/github.com/pytool/ssh/test/upload/"

	var remotedir = "/root/test/"
	c.RemoveAll("/root/upload/")

	err := c.Upload(local, remotedir)
	if err != nil {
		println("[Upload]", err.Error())
	}
}
func TestClient_Upload(t *testing.T) {
	c := GetClient()
	defer c.Close()

	var local = "/home/ubuntu/go/src/github.com/pytool/ssh/test/upload/"
	var uploads = map[string][]string{
		local + "null/": []string{"/root/upload/test/null/1", "/root/upload/test/null/2/"},
		local + "null/": []string{"/root/upload/test/null/3", "/root/upload/test/null/4/"},
		local + "file":  []string{"/root/upload/test/file/1", "/root/upload/test/file/2/"},
		local + "file/": []string{"/root/upload/test/file/3", "/root/upload/test/file/4/"},
		local + "dir":   []string{"/root/upload/test/dir/1", "/root/upload/test/dir/2/"},
		local + "dir/":  []string{"/root/upload/test/dir/3", "/root/upload/test/dir/4/"},
	}

	for local, remotes := range uploads {
		for _, remote := range remotes {
			err := c.Upload(local, remote)
			if err != nil {
				println(err.Error())
			}
			// println(remote, "--->", v, "Finish download!")

		}
	}
}

func TestClient_Download(t *testing.T) {
	c := GetClient()
	defer c.Close()

	var local = "/home/ubuntu/go/src/github.com/pytool/ssh/test/download"
	var download = map[string][]string{
		"/root/test/notExist":  []string{local + "/localNotExist/null/1", local + "/localNotExist/null/2/"},
		"/root/test/notExist/": []string{local + "/localNotExist/null/3", local + "/localNotExist/null/4/"},
		"/root/test/file":      []string{local + "/localNotExist/file/1", local + "/localNotExist/file/2/"},
		"/root/test/file/":     []string{local + "/localNotExist/file/3", local + "/localNotExist/file/4/"},
		"/root/test/dir":       []string{local + "/localNotExist/dir/1", local + "/localNotExist/dir/2/"},
		"/root/test/dir/":      []string{local + "/localNotExist/dir/3", local + "/localNotExist/dir/4/"},
	}

	for remote, local := range download {
		for _, v := range local {
			err := c.Download(remote, v)
			if err != nil {
				println(err.Error())
			}
			// println(remote, "--->", v, "Finish download!")

		}
	}

}

func TestClient_DownloadFile(t *testing.T) {
	c := GetClient()
	defer c.Close()

	var local = "/home/ubuntu/go/src/github.com/pytool/ssh/test/downloadfile"
	var download = map[string][]string{
		"/root/test/notExist":  []string{local + "/localNotExist/null/1", local + "/localNotExist/null/2/"},
		"/root/test/notExist/": []string{local + "/localNotExist/null/3", local + "/localNotExist/null/4/"},
		"/root/test/file":      []string{local + "/localNotExist/file/1", local + "/localNotExist/file/2/"},
		"/root/test/file/":     []string{local + "/localNotExist/file/3", local + "/localNotExist/file/4/"},
		"/root/test/dir":       []string{local + "/localNotExist/dir/1", local + "/localNotExist/dir/2/"},
		"/root/test/dir/":      []string{local + "/localNotExist/dir/3", local + "/localNotExist/dir/4/"},
	}

	for remote, local := range download {
		for _, v := range local {
			err := c.downloadFile(remote, v)
			if err != nil {
				println(err.Error())
			}
			// println(remote, "--->", v, "Finish download!")

		}
	}
}
func TestClient_DownloadDir(t *testing.T) {
	c := GetClient()
	defer c.Close()

	var local = "/home/ubuntu/go/src/github.com/pytool/ssh/test/downloaddir"
	var download = map[string][]string{
		"/root/test/notExist":  []string{local + "/localNotExist/null/1", local + "/localNotExist/null/2/"},
		"/root/test/notExist/": []string{local + "/localNotExist/null/3", local + "/localNotExist/null/4/"},
		"/root/test/file":      []string{local + "/localNotExist/file/1", local + "/localNotExist/file/2/"},
		"/root/test/file/":     []string{local + "/localNotExist/file/3", local + "/localNotExist/file/4/"},
		"/root/test/dir":       []string{local + "/localNotExist/dir/1", local + "/localNotExist/dir/2/"},
		"/root/test/dir/":      []string{local + "/localNotExist/dir/3", local + "/localNotExist/dir/4/"},
	}

	for remote, local := range download {
		for _, v := range local {
			err := c.downloadDir(remote, v)
			if err != nil {
				println(err.Error())
			}
			// println(remote, "--->", v, "Finish download!")

		}
	}
}
func TestClient_UploadFile(t *testing.T) {
	c := GetClient()
	defer c.Close()
	var local = "/home/ubuntu/go/src/github.com/pytool/ssh/test/upload/"
	var uploads = map[string][]string{
		local + "null":  []string{"/root/upload/file_test/null/1", "/root/upload/file_test/null/2/"},
		local + "null/": []string{"/root/upload/file_test/null/3", "/root/upload/file_test/null/4/"},
		local + "file":  []string{"/root/upload/file_test/file/1", "/root/upload/file_test/file/2/"},
		local + "file/": []string{"/root/upload/file_test/file/3", "/root/upload/file_test/file/4/"},
		local + "dir":   []string{"/root/upload/file_test/dir/1", "/root/upload/file_test/dir/2/"},
		local + "dir/":  []string{"/root/upload/file_test/dir/3", "/root/upload/file_test/dir/4/"},
	}

	for local, remotes := range uploads {
		for _, remote := range remotes {
			err := c.UploadFile(local, remote)
			if err != nil {
				println(err.Error())
			}
			// println(remote, "--->", v, "Finish download!")

		}
	}
}

func TestClient_UploadDir(t *testing.T) {
	c := GetClient()
	defer c.Close()
	var local = "/home/ubuntu/go/src/github.com/pytool/ssh/test/upload/"
	var uploads = map[string][]string{
		local + "null/": []string{"/root/upload/dir_test/null/1", "/root/upload/dir_test/null/2/"},
		local + "null/": []string{"/root/upload/dir_test/null/3", "/root/upload/dir_test/null/4/"},
		local + "file":  []string{"/root/upload/dir_test/file/1", "/root/upload/dir_test/file/2/"},
		local + "file/": []string{"/root/upload/dir_test/file/3", "/root/upload/dir_test/file/4/"},
		local + "dir":   []string{"/root/upload/dir_test/dir/1", "/root/upload/dir_test/dir/2/"},
		local + "dir/":  []string{"/root/upload/dir_test/dir/3", "/root/upload/dir_test/dir/4/"},
	}

	for local, remotes := range uploads {
		for _, remote := range remotes {
			err := c.UploadDir(local, remote)
			if err != nil {
				println(err.Error())
			}
			// println(remote, "--->", v, "Finish download!")

		}
	}
}
