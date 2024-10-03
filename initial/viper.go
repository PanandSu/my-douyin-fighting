package initial

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	gb "my-douyin-fighting/glob"
)

func Viper() {
	var err error
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config/config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Panic("读取配置文件失败")
	}
	err = viper.Unmarshal(&gb.Cfg)
	if err != nil {
		log.Panic("viper反序列化失败")
	}
	fmt.Println(gb.Cfg)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("配置文件被修改:", e.Name)
	})
}
