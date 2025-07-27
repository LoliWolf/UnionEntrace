package caller

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
)

var (
	ConfigClient config_client.IConfigClient
)

func initNacos() {
	clientConfig := constant.ClientConfig{
		NamespaceId:         "3cbc1c69-938f-4b35-a37d-9e28c2f55338",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
		LogRollingConfig: &constant.ClientLogRollingConfig{
			MaxSize:   100,
			MaxAge:    30,
			LocalTime: true,
			Compress:  false,
		},
	}
	serverConfig := constant.ServerConfig{
		IpAddr: "192.168.31.234",
		Port:   8848,
	}

	// 创建动态配置客户端
	var err error
	ConfigClient, err = clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
		"serverConfigs": []constant.ServerConfig{
			serverConfig,
		},
	})
	if err != nil {
		panic(err)
	}
	return
}
