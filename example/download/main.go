package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/yeka/zip"

	"github.com/pytool/ssh"
)

var (
	err error
	// sftpClient *sftp.Client
)
var FORMAT = "2006-01-02"
var dbnames = []string{"tower", "mengyin", "pingyi", "shizhi", "tancheng", "yinan", "yishui", "feixian", "gaoxinqu", "hedong", "jingkaiqu", "junan", "luozhuang", "lanling", "lanshan", "lingang", "linshu"}

func main() {

	// client, err := ssh.NewClient("localhost", "22", "root", "ubuntu")
	// if err != nil {
	// 	panic(err)
	// }
	// defer client.Close()

	tmp := os.TempDir()

	fmt.Println(tmp)
	// var remotedir = "/root/test/"
	// // download dir
	// var local = "/home/ubuntu/go/src/github.com/pytool/ssh/test/download/"
	// client.Download(remotedir, local)

	// // download file
	// var remotefile = "/root/test/file"
	// client.Download(remotefile, local)

}
func Down() {
	// tmp:=os.TempDir()
	tmp_dir := "db_" + time.Now().Format(FORMAT)
	os.Mkdir(tmp_dir, 0755)

}
func NewZipWriter(name string) *zip.Writer {
	zipname, err := os.Create(name)
	if err != nil {
		log.Fatalln(err)
	}
	return zip.NewWriter(zipname)
}
func DownLoadZip(client *ssh.Client, zw *zip.Writer, src string) {
	// fmt.Println(src, "数据正在复制中，请耐心等待...")
	srcFile, err := client.SFTPClient.Open(src)
	if err != nil {
		log.Println(err)
		return
	}
	defer srcFile.Close()

	var localFileName = path.Base(src)
	// dstFile, err := os.Create(path.Join(dst, localFileName))
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// defer dstFile.Close()
	w, err := zw.Encrypt(localFileName, `hangruan2017`, zip.AES256Encryption)
	if err != nil {
		return
	}
	if _, err = srcFile.WriteTo(w); err != nil {
		log.Println(err)
		return
	}
}
