// @Description
// @Author  renjingneng
// @CreateTime  2021/12/16 17:37
package core

import (
	"path/filepath"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Context struct {
	src      *gin.Context
	session  sessions.Session
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

func (context *Context) SessionGet(key interface{}) interface{} {
	return context.session.Get(key)
}
func (context *Context) SessionSet(key interface{}, val interface{}) {
	context.session.Set(key, val)
}
func (context *Context) SessionRemove(key interface{}) {
	context.session.Delete(key)
}
func (context *Context) SessionSave() {
	context.session.Save()
}
func (context *Context) DeleteSession() {
	context.session.Options(sessions.Options{
		Path:   "/",
		MaxAge: -1,
	})
	context.session.Clear()
}
