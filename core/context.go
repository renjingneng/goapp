// @Description
// @Author  renjingneng
// @CreateTime  2021/12/16 17:37
package core

import "github.com/gin-gonic/gin"

type Context struct {
	src *gin.Context
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
