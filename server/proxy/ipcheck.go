package proxy

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/acexy/golang-toolkit/crypto/hashing"
	"github.com/acexy/golang-toolkit/sys"
	"github.com/acexy/golang-toolkit/util/coll"
	"github.com/acexy/golang-toolkit/util/net"
	"github.com/fatedier/frp/acexy/consts"
)

var watchOnce sync.Once
var ipChecker *net.IpChecker
var currentIpRules []string
var lastFileMd5 string

func loadConfig() {
	_, err := os.Stat(consts.ServerDenyIPsRelativePath)
	if err != nil {
		return
	}
	currentMd5, _ := hashing.Md5FileHex(consts.ServerDenyIPsRelativePath)
	if lastFileMd5 != currentMd5 {
		data, err := os.ReadFile(consts.ServerDenyIPsRelativePath)
		if err == nil {
			deny := string(data)
			ips := strings.Split(strings.ReplaceAll(deny, "\r\n", "\n"), "\n")
			fileIpRules := coll.SliceDistinct(coll.SliceCollect(ips, func(v string) string {
				v = strings.TrimSpace(v)
				if v != "" {
					return v
				}
				return ""
			}))
			added, removed := coll.SliceDiff(currentIpRules, fileIpRules)
			if len(added) > 0 {
				ipChecker.AddRuleIp(added...)
			}
			if len(removed) > 0 {
				ipChecker.RemoveRuleIp(removed...)
			}
			currentIpRules = fileIpRules
		}
		lastFileMd5 = currentMd5
	}
}

// IsDenyIP 判断ip是否被禁止
func IsDenyIP(ip string) bool {
	watchOnce.Do(func() {
		ipChecker = net.NewIpChecker()
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
	ok, _ := ipChecker.Match(ip)
	return ok
}
