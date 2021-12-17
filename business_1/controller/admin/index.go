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
