package main

import (
	"fmt"

	conf "github.com/xilepeng/gin-mall/config"
	"github.com/xilepeng/gin-mall/routes"
)

func main() {
	loading() // 加载配置
	r := routes.NewRouter()
	_ = r.Run(conf.Config.System.HttpPort)
	fmt.Println("启动配成功...")
}

func loading() {
	conf.InitConfig()

}
