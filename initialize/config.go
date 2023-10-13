package initialize

import (
	"go-api-base/pkg/global"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func NewConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("configs/")

	// 读取配置信息
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("init.ReadInConfig err: %v", err)
	}

	// 热更新配置
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(global.Conf); err != nil {
			log.Fatalf("viper.Unmarshal err: %s", err)
		}
	})

	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(global.Conf); err != nil {
		log.Fatalf("viper.Unmarshal err: %s", err)
	}
}
