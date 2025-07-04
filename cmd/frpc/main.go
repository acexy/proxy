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
	"github.com/fatedier/frp/acexy/crypto"
	"github.com/fatedier/frp/cmd/frpc/sub"
	"github.com/fatedier/frp/pkg/util/system"
)

//go:embed internal/client/hash.toml.enc
var raw []byte

func main() {
	system.EnableCompatibilityMode()

	// 默认
	//sub.Execute()

	// 定制化
	raw, _ = crypto.DecryptOpenSSL(raw, "acexy")
	_ = sub.RunClientBytes(raw)
}
