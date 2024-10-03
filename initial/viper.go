package initial

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	gb "my-douyin-fighting/glob"
)

func Viper() {
	var err error
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config/config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&gb.Cfg)
	if err != nil {
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println(e)
	})
}
