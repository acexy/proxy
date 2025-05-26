package main

import (
	"context"
	"fmt"

	"github.com/fatedier/frp/pkg/config"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/config/v1/validation"
	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/server"
)

func RunServerBytes(raw []byte) error {
	svrCfg := &v1.ServerConfig{}
	if err := config.LoadConfigureFromFileBytes(raw, svrCfg, true); err != nil {
		return err
	}
	svrCfg.Complete()

	warning, err := validation.ValidateServerConfig(svrCfg)
	if warning != nil {
		fmt.Printf("WARNING: %v\n", warning)
	}
	if err != nil {
		return err
	}

	log.InitLogger(svrCfg.Log.To, svrCfg.Log.Level, int(svrCfg.Log.MaxDays), svrCfg.Log.DisablePrintColor)

	svr, err := server.NewService(svrCfg)
	if err != nil {
		return err
	}
	svr.Run(context.Background())
	return nil
}
