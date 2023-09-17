package main

import (
	"github.com/xilepeng/gin-mall/conf"
	"github.com/xilepeng/gin-mall/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
