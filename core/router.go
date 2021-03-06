// Package core
// @Description
// @Author  renjingneng
// @CreateTime  2021/12/15 17:06
package core

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"

	"github.com/renjingneng/goapp/core/config"
)

var contlContainer map[string]interface{}
var router *gin.Engine

func Init() {
	contlContainer = make(map[string]interface{})
	//gin.SetMode(gin.ReleaseMode)
	router = gin.New()
	//Session
	store, _ := redis.NewStore(10, "tcp", config.Get("SessionRedisAddress"), config.Get("SessionRedisPassword"), []byte(config.Get("SessionEncryptionKey")))
	router.Use(sessions.Sessions(config.Get("SessionName"), store))
}

func RegisterController(contlPtr interface{}) {
	contlPtrVal := reflect.ValueOf(contlPtr)
	contlVal := reflect.Indirect(contlPtrVal)
	contlType := contlVal.Type()
	pkgPathList := strings.Split(contlType.PkgPath(), "/")
	pkgName := pkgPathList[(len(pkgPathList) - 1)]
	contlShortName := contlType.Name()[:len(contlType.Name())-10]
	contlShortName = strings.ToLower(contlShortName)
	//判断容器里面是否已经有了
	if _, ok := contlContainer[pkgName+"_"+contlShortName]; !ok {
		contlContainer[pkgName+"_"+contlShortName] = contlPtr
		//调用控制器初始化方法
		method := contlPtrVal.MethodByName("Init")
		if !method.IsValid() {
			panic("contl init error")
		}
		method.Call([]reflect.Value{})
	}
}

func Start(addr string) {
	router.Any("/:pkg/:contl/:method", controllerProcess)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	//step1
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatal("listen error ", err)
		}
	}()

	//step2
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
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
	context.session = sessions.Default(context.src)
	res := method.Call([]reflect.Value{reflect.ValueOf(context)})
	context, _ = res[0].Interface().(*Context)
	if context.resType == "json" {
		c.JSON(http.StatusOK, gin.H{"status": 1, "res": context.json})
	} else if context.resType == "html" {
		c.HTML(http.StatusOK, context.htmlName, gin.H(context.json))
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0})
	}
}
