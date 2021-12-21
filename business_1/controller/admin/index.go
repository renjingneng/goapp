package admin

import (
	"github.com/renjingneng/goapp/core"
)

type IndexController struct {
}

// Init 初始化
func (ctl *IndexController) Init() {
}

// Test1Action 渲染html
func (ctl *IndexController) Test1Action(context *core.Context) *core.Context {
	context.Html("house/index.html", map[string]interface{}{"title": "Main website"})
	return context
}

// Test2Action session set
func (ctl *IndexController) Test2Action(context *core.Context) *core.Context {
	var count int
	v := context.SessionGet("count")
	if v == nil {
		count = 1
	} else {
		count = v.(int)
		count++
	}
	context.SessionSet("count", count)
	context.SessionSave()
	temp := make(map[string]interface{})
	temp["count"] = count
	context.Json(temp)
	return context
}

// Test3Action  delete session
func (ctl *IndexController) Test3Action(context *core.Context) *core.Context {
	context.DeleteSession()
	context.SessionSave()
	temp := make(map[string]interface{})
	temp["count"] = 1
	context.Json(temp)
	return context
}
