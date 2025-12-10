package proxy

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/acexy/golang-toolkit/crypto/hashing"
	"github.com/acexy/golang-toolkit/sys"
	"github.com/acexy/golang-toolkit/util/coll"
	"github.com/fatedier/frp/acexy/consts"
)

var watchIp sync.Once
var denyIPs map[string]struct{}
var lastMd5 string

func loadConfig() {
	_, err := os.Stat(consts.ServerDenyIPsRelativePath)
	if err != nil {
		return
	}
	currentMd5, _ := hashing.Md5FileHex(consts.ServerDenyIPsRelativePath)
	if lastMd5 != currentMd5 {
		data, err := os.ReadFile(consts.ServerDenyIPsRelativePath)
		if err == nil {
			deny := string(data)
			ips := strings.Split(strings.ReplaceAll(deny, "\r\n", "\n"), "\n")
			denyIPs = coll.SliceFilterToMap(ips, func(v string) (string, struct{}, bool) {
				v = strings.TrimSpace(v)
				if v != "" {
					return v, struct{}{}, true
				}
				return "", struct{}{}, false
			})
		}
		lastMd5 = currentMd5
	}
}

// IsDenyIP 判断ip是否被禁止
func IsDenyIP(ip string) bool {
	watchIp.Do(func() {
		loadConfig()
		go func() {
			ticker := time.NewTicker(time.Second * 3)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					loadConfig()
				case <-sys.ShutdownSignal():
					return
				}
			}
		}()
	})
	if len(denyIPs) == 0 {
		return false
	}
	_, ok := denyIPs[ip]
	return ok
}
