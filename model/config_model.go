package model

type Config struct {
	Nacos struct {
		ServerIp    string `yaml:"serverIp"`
		ServerPort  int    `yaml:"serverPort"`
		ServiceName string `yaml:"serviceName"`
		ClusterName string `yaml:"clusterName"`
		Port        int    `yaml:"port"`
	} `yaml:"nacos"`
}
