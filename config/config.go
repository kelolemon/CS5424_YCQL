package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ClusterConfigStr struct {
	Ips  []string `yaml:"ips"`
	Port int      `yaml:"port"`
}

var ClusterConfig ClusterConfigStr

func init() {
	config, _ := ioutil.ReadFile("config/config.yaml")
	_ = yaml.Unmarshal(config, &ClusterConfig)
	fmt.Println(ClusterConfig)
}
