package main

import (
	"errors"
	"github.com/Woringsuhang/ServiceLayer/server"
	"github.com/Woringsuhang/mess/user"
	"github.com/Woringsuhang/user/common"
	"github.com/Woringsuhang/user/global"
	"github.com/Woringsuhang/user/grpcs"
	"github.com/Woringsuhang/user/model"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var Tx = func(tx *gorm.DB) error {
	if tx.Error != nil {
		return errors.New("数据库连接失败")
	}
	return nil
}

func main() {
	err := common.ConsulClient()
	if err != nil {
		zap.S().Info(err)
		return
	}
	err = common.AgentService("10.2.171.94", 8080)
	if err != nil {
		zap.S().Info(err)
		return
	}
	common.InitViper("./configs.yaml")
	common.Nacos(global.NacosConfig.DataId, global.NacosConfig.Group)
	common.InitZap()
	model.InitMysql()
	model.Tx(Tx)
	err = grpcs.Registration(func(g *grpc.Server) {
		user.RegisterUserServer(g, &server.Servers{})
	}, "./document/server_cert.pem", "./document/server_key.pem")
	if err != nil {
		zap.S().Info(err)
		return
	}
}
