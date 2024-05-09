package initial

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"my-douyin-fighting/glob"
)

func Viper() {
	var err error
	viper.SetConfigName("config")
	viper.SetConfigFile("./config/config.yaml")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&glob.Config)
	if err != nil {
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println(e)
	})
}
