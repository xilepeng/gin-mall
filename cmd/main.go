package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xilepeng/gin-mall/conf"
)

func main() {
	r := gin.Default()
	conf.Init()
	r.Run()
}
