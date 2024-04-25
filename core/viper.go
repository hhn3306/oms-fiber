package core

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"oms-fiber/core/internal"
	"oms-fiber/global"
	"os"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
func Viper(path ...string) {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			switch os.Getenv(internal.ConfigEnv) {
			case "prod":
				config = internal.ConfigProdFile
				fmt.Printf("您正在使用 %s 环境,config的文件名为%s\n", "prod", internal.ConfigProdFile)
			case "dev":
				config = internal.ConfigDevFile
				fmt.Printf("您正在使用 %s 环境,config的文件名为%s\n", "dev", internal.ConfigDevFile)
			default:
				config = internal.ConfigDefaultFile
				fmt.Printf("您正在使用 %s 环境,config的文件名为%s\n", "默认", internal.ConfigDefaultFile)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的文件名为%s\n", config)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的文件名为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	//v.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("config file changed:", e.Name)
	//	if err = v.Unmarshal(&global.Config); err != nil {
	//		fmt.Println(err)
	//	}
	//})
	if err = v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	//global.Config.AutoCode.Root, _ = filepath.Abs("..")
}
