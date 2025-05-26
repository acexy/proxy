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
	_ "github.com/fatedier/frp/assets/frps"

	"fmt"
	_ "github.com/fatedier/frp/pkg/metrics"
	"github.com/fatedier/frp/pkg/util/system"
	"github.com/fatedier/frp/pkg/util/version"
	"os"
)

//go:embed internal/server/qing-prd-peer1.toml
var raw []byte

func main() {
	for _, arg := range os.Args[1:] {
		if arg == "--version" || arg == "-v" {
			fmt.Println(version.Full())
			return
		}
	}
	system.EnableCompatibilityMode()
	//Execute()
	_ = RunServerBytes(raw)
}
