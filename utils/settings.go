// @Author scy
// @Time 2024/7/23 2:19
// @File settings.go

package utils

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	HttpPort string `mapstructure:"http_port"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

type Config struct {
	AppMode string `mapstructure:"app_mode"`
	Server  ServerConfig
	Mysql   MysqlConfig
}

var Vcf *viper.Viper

func init() {
	// 创建实例
	Vcf = viper.New()
	// 配置文件路径
	Vcf.AddConfigPath("config")
	// 配置文件名
	Vcf.SetConfigName("config")
	// 配置文件类型
	Vcf.SetConfigType("yaml")

	// 读取配置文件
	if err := Vcf.ReadInConfig(); err != nil {
		// 断言
		var err viper.ConfigFileNotFoundError
		if errors.As(err, &err) {
			// 断言成功，配置文件找不到
			panic(err)
		}
	}

	// 配置文件反序列化到结构体中
	var config *Config
	if err := Vcf.Unmarshal(&config); err != nil {
		panic(err)
	}

	// 注册每次配置发生变化后的都会调用的回调函数
	Vcf.OnConfigChange(func(in fsnotify.Event) {
		// 每次配置文件发生变化，需将配置文件重新反序列化到结构体中
		if err := Vcf.Unmarshal(&config); err != nil {
			panic(err)
		}
	})

	// 监视配置文件变化
	Vcf.WatchConfig()
}
