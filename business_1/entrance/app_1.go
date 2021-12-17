package main

import (
	"github.com/renjingneng/goapp/business_1/controller/admin"
	"github.com/renjingneng/goapp/business_1/controller/api"
	"github.com/renjingneng/goapp/core"
	"github.com/renjingneng/goapp/core/config"
)

// main app_1应用入口
//
// @Author  renjingneng
//
// @CreateTime  2020/9/24 11:08
func main() {
	//读取配置
	config.LoadConfig()
	//绑定控制器
	core.RegisterController(&api.IndexController{})
	core.RegisterController(&admin.IndexController{})
	//开始
	core.Start("0.0.0.0:" + config.Get("Port"))
}
