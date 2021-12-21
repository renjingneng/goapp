package config

import (
	"flag"
	"github.com/renjingneng/goapp/core/log"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var viperInstance = viper.New()

// LoadConfig 载入配置文件
//
// @Author  renjingneng
//
// @CreateTime  2020/8/19 11:20
func LoadConfig() {
	var filename string
	var envFlag = flag.String("env", "local", "请输入env参数,默认值为local!")
	flag.Parse()

	if *envFlag == "prod" {
		filename = "config-prod"
	} else if *envFlag == "dev" {
		filename = "config-dev"
	} else {
		*envFlag = "local"
		filename = "config-local"
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Error("config error")
	}
	viperInstance.SetConfigName(filename)
	viperInstance.SetConfigType("yaml")
	viperInstance.AddConfigPath(dir + string(os.PathSeparator) + "config")
	if err := viperInstance.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Error(err)
		} else {
			// Config file was found but another error was produced
			log.Error(err)
		}
	}
	//设置环境变量值
	viperInstance.Set("Env", *envFlag)
}

// Get 获取配置信息,示例Get("Port")、Get("Hostlist.Open")
//
// @Author  renjingneng
//
// @CreateTime  2020/8/19 11:02
func Get(key string) string {
	return viperInstance.GetString(key)
}

// GetMap 获取map[string]string
//
// @Author  renjingneng
//
// @CreateTime  2020/8/19 11:06
func GetMap(key string) map[string]string {
	return viperInstance.GetStringMapString(key)
}

// GetList 获取[]string
//
// @Author  renjingneng
//
// @CreateTime  2020/8/19 11:06
func GetList(key string) []string {
	return viperInstance.GetStringSlice(key)
}
