// Copyright 2016 fatedier, fatedier@gmail.com
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
	"os"

	"github.com/fatedier/frp/acexy/consts"
	"github.com/fatedier/frp/acexy/crypto"
	"github.com/fatedier/frp/cmd/frpc/sub"
	"github.com/fatedier/frp/pkg/util/system"
	_ "github.com/fatedier/frp/web/frpc"
)

//go:embed internal/client/hx-mini-mac.toml.enc
var bytes []byte

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
	//sub.Execute()

	// ----------------------

	// 定制化
	bytes, _ = crypto.DecryptOpenSSL(bytes, consts.ConfigEncPassword)
	configContent, err := readIfExists(consts.ClientConfigRelativePath)
	if err == nil {
		bytes = append(bytes, []byte("\n"+configContent)...)
	}
	err = sub.RunClientBytes(bytes)
	if err != nil {
		panic(err)
	}
}
