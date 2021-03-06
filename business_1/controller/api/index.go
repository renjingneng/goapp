package api

import (
	"github.com/davecgh/go-spew/spew"

	"github.com/renjingneng/goapp/business_1/service"
	"github.com/renjingneng/goapp/core"
)

type IndexController struct {
	account *service.AccountService
}

// Init 初始化
func (ctl *IndexController) Init() {
	ctl.account = service.NewAccountService()
}

// Test1Action 测试服务
func (ctl *IndexController) Test1Action(context *core.Context) *core.Context {
	temp := ctl.account.GetInfoById(context.Query("id"))
	context.Json(temp)
	return context
}

// Test2Action spew.Dump深度打印值
func (ctl *IndexController) Test2Action(context *core.Context) *core.Context {
	temp := make(map[string]interface{})
	temp["key1"] = "ddd"
	spew.Dump(ctl)
	context.Json(temp)
	return context
}
