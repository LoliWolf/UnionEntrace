package caller

import (
	"UnionEntrace/model"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gopkg.in/yaml.v3"
	"net"
	"os"
)

var (
	GlobalConfig          model.Config // 新增：全局配置实例
	ConfigClient          config_client.IConfigClient
	ServiceDiscoverClient naming_client.INamingClient
	ServiceName           string
	ClusterName           string
	Port                  int
	NacosServerIp         string
	NacosServerPort       int
)

func LoadConfig() {
	//cwd, err := os.Getwd()
	//if err != nil {
	//	panic("获取工作目录失败: " + err.Error())
	//}
	//// 拼接配置文件路径（假设 config.yaml 在工作目录根目录）
	//configPath := filepath.Join(cwd, "config.yaml")
	//fmt.Print(configPath)
	file, err := os.Open("config.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&GlobalConfig); err != nil {
		panic(err)
	}

	// 从配置文件赋值到全局变量
	ServiceName = GlobalConfig.Nacos.ServiceName
	ClusterName = GlobalConfig.Nacos.ClusterName
	Port = GlobalConfig.Nacos.Port

	NacosServerIp = GlobalConfig.Nacos.ServerIp
	NacosServerPort = GlobalConfig.Nacos.ServerPort
}

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
		IpAddr: NacosServerIp,
		Port:   uint64(NacosServerPort),
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

	// 创建服务发现客户端
	ServiceDiscoverClient, err = clients.CreateNamingClient(map[string]interface{}{
		"clientConfig": clientConfig,
		"serverConfigs": []constant.ServerConfig{
			serverConfig,
		},
	})
	if err != nil {
		panic(err)
	}

	// 注册实例
	err = RegisterInstance()
	if err != nil {
		panic(err)
	}
	return
}

func RegisterInstance() error {
	// 获取本机ip 指定端口
	ip, err := getLocalIP()
	if err != nil {
		return err
	}
	// 注册实例
	_, err = ServiceDiscoverClient.RegisterInstance(vo.RegisterInstanceParam{
		ServiceName: ServiceName,
		ClusterName: ClusterName,
		Ip:          ip,
		Port:        uint64(Port),
		Enable:      true,
		Healthy:     true,
	})
	return err
}

func getLocalIP() (string, error) {
	// 获取本机ip
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	// 选择一个192.168开头的ipv4
	for _, ifi := range interfaces {
		addrs, err := ifi.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && ipNet.IP.To4() != nil {
				if ipNet.IP.String()[:7] == "192.168" {
					return ipNet.IP.String(), nil
				}
			}
		}
	}
	return "", nil
}
