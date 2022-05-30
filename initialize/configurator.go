package initialize

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/svcodestore/sv-resource-gin/config"
	"github.com/svcodestore/sv-resource-gin/global"
)

func InitConfigurator(path ...string) config.Configurator {
	var conf string
	if len(path) == 0 {
		flag.StringVar(&conf, "c", "", "choose config file.")
		flag.Parse()
		if conf == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(config.ConfigEnvFile); configEnv == "" {
				conf = config.ConfigFile
				fmt.Printf("系统配置文件是 %v\n", config.ConfigFile)
			} else {
				conf = configEnv
				fmt.Printf("系统配置使用SSG_CONFIG环境变量，路径为 %v\n", conf)
			}
		} else {
			fmt.Printf("系统配置使用命令行的 -c 参数传递的值，路径为 %v\n", conf)
		}
	} else {
		conf = path[0]
		fmt.Printf("系统配置使用func Viper()传递的值,config的路径为 %v\n", conf)
	}

	v := viper.New()
	v.SetConfigFile(conf)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
