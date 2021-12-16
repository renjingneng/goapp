// Package core
// @Description
// @Author  renjingneng
// @CreateTime  2021/12/15 17:06
package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

var contlContainer map[string]interface{}
var router *gin.Engine

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
func init() {
	contlContainer = make(map[string]interface{})
	gin.SetMode(gin.ReleaseMode)
	router = gin.New()
}

func RegisterController(contlPtr interface{}) {
	contlPtrVal := reflect.ValueOf(contlPtr)
	contlVal := reflect.Indirect(contlPtrVal)
	contlType := contlVal.Type()
	pkgPathList := strings.Split(contlType.PkgPath(), "/")
	pkgName := pkgPathList[(len(pkgPathList) - 1)]
	contlShortName := contlType.Name()[:len(contlType.Name())-10]
	contlShortName = strings.ToLower(contlShortName)
	contlContainer[pkgName+"_"+contlShortName] = contlPtr
}

func Start(addr string) {
	router.Any("/:pkg/:contl/:method", controllerProcess)
	router.Run(addr)
}

func controllerProcess(c *gin.Context) {
	contlPtr, ok := contlContainer[c.Param("pkg")+"_"+c.Param("contl")]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"status": 0})
		return
	}
	methodUp := strings.ToUpper(c.Param("method"))
	methodLow := strings.ToLower(c.Param("method"))
	methodName := string(methodUp[0]) + string(methodLow[1:]) + "Action"
	contlPtrVal := reflect.ValueOf(contlPtr)
	method := contlPtrVal.MethodByName(methodName)
	if !method.IsValid() {
		c.JSON(http.StatusNotFound, gin.H{"status": 0})
		return
	}
	context := &Context{src: c}
	res := method.Call([]reflect.Value{reflect.ValueOf(context)})
	c.JSON(http.StatusOK, gin.H{"status": 1, "res": res[0].Interface()})
}
