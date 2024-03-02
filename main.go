package main

import (
	"errors"
	"github.com/Woringsuhang/ServiceLayer/server"
	"github.com/Woringsuhang/mess/user"
	"github.com/Woringsuhang/user/common"
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

	//nacos.InitViper()
	//nacos.Consul()
	//nacos.InitMysql()
	common.Consul()
	common.InitZap()

	err := model.InitMysql()
	if err != nil {
		panic(err)
	}
	model.Tx(Tx)
	err = grpcs.Registration(func(g *grpc.Server) {
		user.RegisterUserServer(g, &server.Servers{})
	}, "./document/server_cert.pem", "./document/server_key.pem")
	if err != nil {
		zap.S().Info(err)
		return
	}
}
