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

	"github.com/acexy/proxy/acexy/consts"
	"github.com/acexy/proxy/acexy/crypto"
	"github.com/acexy/proxy/acexy/util"
	_ "github.com/acexy/proxy/pkg/metrics"
	"github.com/acexy/proxy/pkg/util/system"
	"github.com/acexy/proxy/pkg/util/version"
	_ "github.com/acexy/proxy/web/frps"
)

//go:embed internal/server/acexy.toml.enc
var bytes []byte

func main() {
	system.EnableCompatibilityMode()

	// 默认
	//Execute()

	// ---------------------------

	// 定制化
	for _, arg := range os.Args[1:] {
		if arg == "--version" || arg == "-v" {
			fmt.Println(version.Full())
			return
		}
	}

	bytes, _ = crypto.DecryptOpenSSL(bytes, consts.ConfigEncPassword)
	configContent, err := util.ReadIfExists(consts.ServerConfigRelativePath)
	if err == nil {
		bytes = append(bytes, []byte("\n"+configContent)...)
	}
	err = runServerWithConfigBytes(bytes)
	if err != nil {
		panic(err)
	}
}
