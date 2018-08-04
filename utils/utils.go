//// Copyright 2015 Spring Rain Software Compnay LTD. All Rights Reserved.
//// Licensed under the MIT (MIT-LICENSE.txt) license.

package utils

import (
	"net"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"unsafe"

	"github.com/wfxiang08/cyutils/utils/config"
	"github.com/wfxiang08/cyutils/utils/errors"
	"github.com/wfxiang08/cyutils/utils/log"
)

func InitConfigFromFile(filename string) (*config.Cfg, error) {
	ret := config.NewCfg(filename)
	if err := ret.Load(); err != nil {
		return nil, errors.Trace(err)
	}
	return ret, nil
}

//
// 获取带有指定Prefix的Ip
//
func GetIpWithPrefix(prefix string) string {

	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			ipAddr := ip.String()
			// fmt.Println("ipAddr: ", ipAddr)
			if strings.HasPrefix(ipAddr, prefix) {
				return ipAddr
			}

		}
	}
	return ""
}

func GetExecutorPath() string {
	filedirectory := filepath.Dir(os.Args[0])
	execPath, err := filepath.Abs(filedirectory)
	if err != nil {
		log.PanicErrorf(err, "get executor path failed")
	}
	return execPath
}

type Strings []string

func (s1 Strings) Eq(s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

const (
	EMPTY_MSG = ""
)

//
// <head> "", tail... ----> head, tail...
// 将msgs拆分成为两部分, 第一部分为: head(包含路由信息);第二部分为: tails包含信息部分
//
func Unwrap(msgs []string) (head string, tails []string) {
	head = msgs[0]
	if len(msgs) > 1 && msgs[1] == EMPTY_MSG {
		tails = msgs[2:]
	} else {
		tails = msgs[1:]
	}
	return
}


func Copy(s string) string {
	var b []byte
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	h.Data = (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	h.Len = len(s)
	h.Cap = len(s)

	return string(b)
}

// 判断给定的文件是否存在
func FileExist(file string) bool {
	var err error
	_, err = os.Stat(file)
	return !os.IsNotExist(err)
}

