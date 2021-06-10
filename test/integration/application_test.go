package integration

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Server_cmd(t *testing.T) {
	command := "logger --udp -n 0.0.0.0 --port 1514 --rfc5424 User 'username' Executed the 'Configure Term' Command."
	expectedStr := "time=\"2021-05-17T18:20:17Z\" level=info facility=1 fields.msg=\"User username Executed the Configure Term Command.\" fields.time=\"2021-05-17 21:20:17.162689 +0300 +0300\" severity=5\n"
	assert := assert.New(t)
	cmd := exec.
		Command("/bin/sh", "-c", command)
	cmd.Run()
	actualByte, _ := exec.Command("/bin/sh", "-c", "docker logs event-server").CombinedOutput()
	actualStr := string(actualByte)
	assert.Equal(expectedStr, actualStr)
}

func Test_Server_cmdNotEqual(t *testing.T) {
	command := "logger --udp -n 0.0.0.0 --port 1514 --rfc5424 User"
	expectedStr := "time=\"2021-05-17T18:20:17Z\" level=info facility=1 fields.msg=\"User username Executed the Configure Term Command.\" fields.time=\"2021-05-17 21:20:17.162689 +0300 +0300\" severity=5\n"
	assert := assert.New(t)
	cmd := exec.
		Command("/bin/sh", "-c", command)
	cmd.Run()
	actualByte, _ := exec.Command("/bin/sh", "-c", "docker logs event-server").CombinedOutput()
	actualStr := string(actualByte)
	assert.NotEqual(expectedStr, actualStr)
}
