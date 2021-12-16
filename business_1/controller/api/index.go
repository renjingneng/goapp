package api

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/renjingneng/goapp/business_1/database/mysql"
	"github.com/renjingneng/goapp/business_1/database/redis"
	"github.com/renjingneng/goapp/core"
	"time"
)

type IndexController struct {
	Tablename    string
	OpenDatabase *mysql.Open
	OpenRedis    *redis.Open
}

// Test1Action 测试请求返回
func (ctl *IndexController) Test1Action(context *core.Context) map[string]interface{} {
	temp := make(map[string]interface{})
	temp1 := map[string]interface{}{"key1": []int{1, 2, 3, 4}}
	temp["key1"] = temp1
	temp["key2"] = ctl.Tablename
	temp["key3"] = context.DefaultQuery("hello", "key3")
	return temp
}

// Test2Action 测试mysql
func (ctl *IndexController) Test2Action(context *core.Context) map[string]interface{} {
	temp := make(map[string]interface{})
	ctl.OpenDatabase.SetTablename(ctl.Tablename)
	temp["key1"] = ctl.OpenDatabase.FetchRowShort(map[string]string{"id": "1"})
	temp["key2"] = context.DefaultQuery("hello", "key3")
	return temp
}

// Test3Action 测试mysql
func (ctl *IndexController) Test3Action(context *core.Context) map[string]interface{} {
	temp := make(map[string]interface{})
	ctl.OpenRedis.Set("dddfff", "fff", 500*time.Second)
	temp["key1"], _ = ctl.OpenRedis.Get("dddfff")
	return temp
}

// Test4Action spew.Dump深度打印值
func (ctl *IndexController) Test4Action(context *core.Context) map[string]interface{} {
	temp := make(map[string]interface{})
	temp["key1"] = "ddd"
	spew.Dump(ctl)
	return temp
}
