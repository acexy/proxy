// Copyright 2018 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/fatedier/frp/acexy/crypto"
	_ "github.com/fatedier/frp/assets/frps"
	_ "github.com/fatedier/frp/pkg/metrics"
	"github.com/fatedier/frp/pkg/util/system"
	"github.com/fatedier/frp/pkg/util/version"
)

//go:embed internal/server/acexy.toml.enc
var raw []byte

func readIfExists(path string) (string, error) {
	// 判断文件是否存在
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
func main() {
	system.EnableCompatibilityMode()

	// 默认
	//Execute()

	// acexy定制化

	// 定制化
	for _, arg := range os.Args[1:] {
		if arg == "--version" || arg == "-v" {
			fmt.Println(version.Full())
			return
		}
	}
	raw, _ = crypto.DecryptOpenSSL(raw, "acexy")
	configContent, err := readIfExists("./proxys.toml")
	if err == nil {
		raw = append(raw, []byte("\n"+configContent)...)
	}
	_ = RunServerBytes(raw)
}
