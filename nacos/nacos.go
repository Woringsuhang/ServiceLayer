package nacos

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	_ "github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	_ "github.com/nacos-group/nacos-sdk-go/common/logger"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	Db *gorm.DB
)

func Consul() {
	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         NacosConfig.NamespaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      NacosConfig.IpAddr,
			ContextPath: "/nacos",
			Port:        uint64(NacosConfig.Port),
			Scheme:      "http",
		},
	}

	// Create config client for dynamic configuration
	configs, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	config, err := configs.GetConfig(vo.ConfigParam{
		DataId: "configuration",
		Group:  "dev",
	})
	if err != nil {
		log.Printf("Error getting configuration")
		return
	}
	fmt.Println(config)
	err = json.Unmarshal([]byte(config), &NacosT)
	fmt.Println(NacosT)
	//监听

	err = configs.ListenConfig(vo.ConfigParam{
		DataId: "configuration",
		Group:  "dev",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ",dataId:" + dataId + ",data:" + data)
			err = json.Unmarshal([]byte(data), &NacosT)

			dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				NacosT.Mysql.Username, NacosT.Mysql.Password, NacosT.Mysql.Host,
				NacosT.Mysql.Port, NacosT.Mysql.Library)

			updateDbConnection(dsn)
		},
	})

	if err != nil {
		panic(err)
	}
}

func updateDbConnection(config string) {
	// 关闭现有连接池（如果存在）
	Dbs, _ := Db.DB()
	if Dbs != nil {
		_ = Dbs.Close()
	}

	// 使用新的配置信息创建数据库连接
	var err error
	Db, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	// 假设 config 是有效的数据库 DSN
	if err != nil {
		log.Fatalf("Failed to create database connection: %v", err)
	}

	// 可能需要对 db 进行额外配置，如设置连接池大小等

	fmt.Println("Database connection updated successfully.")
}

func InitMysql() {
	var err error
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		NacosT.Mysql.Username, NacosT.Mysql.Password, NacosT.Mysql.Host,
		NacosT.Mysql.Port, NacosT.Mysql.Library)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}
