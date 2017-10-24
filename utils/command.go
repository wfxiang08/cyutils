package utils

import (
	log "github.com/wfxiang08/cyutils/utils/rolling_log"
	"os/exec"
	"strings"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
)

// 执行指定的Command
func ExecCommand(cmd string, dir string) ([]byte, error) {
	if len(dir) > 0 {
		log.Printf("Execute command: %s AT: %s", cmd, dir)
	}

	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:]

	//start := time.Now()
	command := exec.Command(head, parts...)

	if len(dir) > 0 {
		command.Dir = dir
	}

	// 需要同时看到: stdout 和 stderr
	result, err := command.CombinedOutput()
	if err != nil {
		log.Printf("Command succeed: %s", cmd)
	}
	return result, err
}

func GetFileMD5(filePath string) string {
	data, _ := ioutil.ReadFile(filePath)
	sum := md5.Sum(data)
	return hex.EncodeToString(sum[:])
}