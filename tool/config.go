package tool

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	App App `json:"app"`
	Db  Db  `json:"db"`
}
type App struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
type Db struct {
	Name     string `json:"name"`
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
}

var AllConfig *Config = nil

func ReadConfigFile(file string) (*Config, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	v := viper.New()
	v.AddConfigPath(path + "/config/")
	v.SetConfigName(file)
	v.SetConfigType("yml")
	err = v.ReadInConfig()
	if err != nil {
		
		panic("读取配置文件失败\n")
		return nil, err
	}
	err = v.Unmarshal(&AllConfig)
	if err != nil {
		panic("配置转换结构体失败\n")
		return nil, err
	}
	return AllConfig, nil
}
