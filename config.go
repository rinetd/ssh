package ssh

import (
	"time"
)

type Config struct {
	User     string
	Host     string
	Port     int
	Password string
	KeyFiles []string

	// DisableAgentForwarding, if true, will not forward the SSH agent.
	DisableAgentForwarding bool

	// HandshakeTimeout limits the amount of time we'll wait to handshake before
	// saying the connection failed.
	HandshakeTimeout time.Duration

	// KeepAliveInterval sets how often we send a channel request to the
	// server. A value < 0 disables.
	KeepAliveInterval time.Duration

	// Timeout is how long to wait for a read or write to succeed.
	Timeout time.Duration
}

var DefaultConfig = &Config{
	User:     "root",
	Port:     22,
	KeyFiles: []string{"~/.ssh/id_rsa"},
}

//
func (c *Client) WithUser(user string) *Client {
	if user == "" {
		user = "root"
	}
	c.User = user
	return c
}

//
func (c *Client) WithHost(host string) *Client {
	if host == "" {
		host = "localhost"
	}
	c.Host = host
	return c
}
func (c *Client) WithPassword(password string) *Client {
	c.Password = password
	return c
}

//
func (c *Client) SetKeys(keyfiles []string) *Client {
	if keyfiles == nil {
		return c
	}
	t := make([]string, len(keyfiles))
	copy(t, keyfiles)
	c.KeyFiles = t
	return c
}

//
func (c *Client) WithKey(keyfile string) *Client {
	if keyfile == "" {
		keyfile = "~/.ssh/id_rsa"
	}
	for _, s := range c.KeyFiles {
		if s == keyfile {
			return c
		}
	}
	c.KeyFiles = append(c.KeyFiles, keyfile)
	return c
}
