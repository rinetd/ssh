package ssh

import (
	"fmt"
	"path/filepath"
)

// Run Execute cmd on the remote host and return stderr and stdout
func (c *Client) Run(cmd string) ([]byte, error) {
	session, err := c.SSHClient.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	return session.CombinedOutput(cmd)
}

//Exec Execute cmd on the remote host and return stderr and stdout
func (c *Client) Exec(cmd string) ([]byte, error) {
	session, err := c.SSHClient.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	return session.CombinedOutput(cmd)
}

// RunScript Executes a shell script file on the remote machine.
// It is copied in the tmp folder and ran in a single session.
// chmod +x is applied before running.
// Returns an SshResponse and an error if any has occured
func (c *Client) RunScript(scriptPath string) ([]byte, error) {
	session, err := c.SSHClient.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	// 1. 上传 script
	remotePath := fmt.Sprintf("/tmp/script/%s", filepath.Base(scriptPath))
	if err := c.UploadFile(scriptPath, remotePath); err != nil {
		return nil, err
	}
	// 2. 执行script
	rCmd := fmt.Sprintf("chmod +x %s ; %s", remotePath, remotePath)
	return session.CombinedOutput(rCmd)
}
