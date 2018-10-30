package ssh

import (
	"errors"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

const DefaultTimeout = 30 * time.Second

type Client struct {
	*Config
	SSHClient  *ssh.Client
	SFTPClient *sftp.Client
}

// NewClient 根据配置
func NewClient(user, host, port, password string) (client *Client, err error) {
	p, err := strconv.Atoi(port)
	if err == nil || p == 0 {
		p = 22
	}
	if user == "" {
		user = "root"
	}
	var config = &Config{
		User:     user,
		Host:     host,
		Port:     p,
		Password: password,
		// KeyFiles: []string{"~/.ssh/id_rsa"},
	}
	return New(config)
}

// New 创建SSH client
func New(config *Config) (client *Client, err error) {
	clientConfig := &ssh.ClientConfig{
		User:            config.User,
		Timeout:         DefaultTimeout,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 2. 密码方式
	if config.Password != "" {
		clientConfig.Auth = append(clientConfig.Auth, ssh.Password(config.Password))
	}

	// 3. privite key file
	if len(config.KeyFiles) == 0 {
		keyPath := os.Getenv("HOME") + "/.ssh/id_rsa"
		if auth, err := WithPrivateKey(keyPath, config.Password); err != nil {
			clientConfig.Auth = append(clientConfig.Auth, auth)
		}
	} else {
		if auth, err := WithPrivateKeys(config.KeyFiles, config.Password); err != nil {
			clientConfig.Auth = append(clientConfig.Auth, auth)
		}
	}
	// 1. agent
	if auth, err := WithAgent(); err != nil {
		clientConfig.Auth = append(clientConfig.Auth, auth)
	}
	if config.Port == 0 {
		config.Port = 22
	}
	// hostPort := config.Host + ":" + strconv.Itoa(config.Port)
	sshClient, err := ssh.Dial("tcp", net.JoinHostPort(config.Host, strconv.Itoa(config.Port)), clientConfig)

	if err != nil {
		return client, errors.New("Failed to dial ssh: " + err.Error())
	}

	// create sftp client
	var sftpClient *sftp.Client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return client, errors.New("Failed to conn sftp: " + err.Error())
	}

	return &Client{SSHClient: sshClient, SFTPClient: sftpClient}, nil
}

// Execute cmd on the remote host and return stderr and stdout
func (c *Client) Exec(cmd string) ([]byte, error) {
	session, err := c.SSHClient.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	return session.CombinedOutput(cmd)
}

// Close the underlying SSH connection
func (c *Client) Close() {
	c.SFTPClient.Close()
	c.SSHClient.Close()
}

func addPortToHost(host string) string {
	_, _, err := net.SplitHostPort(host)

	// We got an error so blindly try to add a port number
	if err != nil {
		return net.JoinHostPort(host, "22")
	}

	return host
}
