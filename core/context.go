// @Description
// @Author  renjingneng
// @CreateTime  2021/12/16 17:37
package core

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
)

type Context struct {
	src      *gin.Context
	resType  string                 //json/html
	json     map[string]interface{} //can be used as response json or params in html
	htmlName string
}

func (context *Context) Query(key string) string {
	return context.src.Query(key)
}

func (context *Context) DefaultQuery(key, defaultValue string) string {
	return context.src.DefaultQuery(key, defaultValue)
}

func (context *Context) PostForm(key string) string {
	return context.src.PostForm(key)
}
func (context *Context) DefaultPostForm(key, defaultValue string) string {
	return context.src.DefaultPostForm(key, defaultValue)
}
func (context *Context) Json(val map[string]interface{}) {
	context.json = val
	context.resType = "json"
}

func (context *Context) Html(path string, val map[string]interface{}) {
	context.json = val
	context.resType = "html"
	path = "template/" + path
	context.htmlName = filepath.Base(path)
	router.LoadHTMLFiles(path)
}
