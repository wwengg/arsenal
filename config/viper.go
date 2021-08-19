// @Title
// @Description
// @Author  Wangwengang  2021/8/18 下午10:01
// @Update  Wangwengang  2021/8/18 下午10:01
package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config string

const (
	ConfigEnv  = "X_CONFIG"
	ConfigFile = "config.yaml"
)

func init() {
	flag.StringVar(&config, "c", "", "choose config file.")
}

func Viper(path ...string) *viper.Viper {
	flag.Parse()
	if len(path) == 0 {
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(ConfigEnv); configEnv == "" {
				config = ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用X_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&config); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&ConfigHub); err != nil {
		fmt.Println(err)
	}
	return v
}
