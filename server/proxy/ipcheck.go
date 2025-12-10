package proxy

import (
	"strings"
	"sync"
	"time"

	"github.com/acexy/golang-toolkit/sys"
	"github.com/acexy/golang-toolkit/util/coll"
	"github.com/fatedier/frp/acexy/consts"
	"github.com/fatedier/frp/acexy/util"
)

var watchIp sync.Once
var denyIPs map[string]struct{}

func loadConfig() {
	deny, err := util.ReadIfExists(consts.ServerDenyIPsRelativePath)
	if err == nil {
		ips := strings.Split(strings.ReplaceAll(deny, "\r\n", "\n"), "\n")
		denyIPs = coll.SliceFilterToMap(ips, func(v string) (string, struct{}, bool) {
			v = strings.TrimSpace(v)
			if v != "" {
				return v, struct{}{}, true
			}
			return "", struct{}{}, false
		})
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
