package getdata

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func connect(user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}

var FORMAT = "2006-01-02"
var (
	err        error
	sftpClient *sftp.Client
)
var dbnames = []string{"tower", "mengyin", "pingyi", "shizhi", "tancheng", "yinan", "yishui", "feixian", "gaoxinqu", "hedong", "jingkaiqu", "junan", "luozhuang", "lanling", "lanshan", "lingang", "linshu"}

func main() {

	// 这里换成实际的 SSH 连接的 用户名，密码，主机名或IP，SSH端口
	// sftpClient, err = connect("root", "sdlylshl871016", "111.235.181.127", 443)
	sftpClient, err = connect("root", "HR2018!!", "15.14.12.150", 22)
	if err != nil {
		log.Println(err)
	}
	defer sftpClient.Close()
	// 用来测试的远程文件路径 和 本地文件夹
	// fmt.Println(shizhi)
	// var localDir = "."
	date_dir := "db_" + time.Now().Format(FORMAT)
	os.Mkdir(date_dir, 0755)
	var lzkpbi = "/docker/backup/" + time.Now().Format(FORMAT) + "_lzkp_bi_inner.zip"
	Down(lzkpbi, date_dir)
	for _, n := range dbnames {
		p := "/docker/backup/" + time.Now().Format(FORMAT) + "_" + n + "_inner.zip"
		// fmt.Println(p)

		Down(p, date_dir)
	}
	// fmt.Scanln()
}

func Down(src, dst string) {
	fmt.Println(src, "数据正在复制中，请耐心等待...")
	srcFile, err := sftpClient.Open(src)
	if err != nil {
		log.Println(err)
		return
	}
	defer srcFile.Close()

	var localFileName = path.Base(src)
	dstFile, err := os.Create(path.Join(dst, localFileName))
	if err != nil {
		log.Println(err)
		return
	}
	defer dstFile.Close()

	if _, err = srcFile.WriteTo(dstFile); err != nil {
		log.Println(err)
		return
	}

	fmt.Println(src, "数据复制完成!")

}
