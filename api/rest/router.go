package rest

import (
	"context"
	"ddd/api/rest/handler"
	"ddd/conf"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Router struct {
	setting *conf.RestSetting
	server  *http.Server
	handler *handler.Handler
}

var once = &sync.Once{}
var router *Router

func Init() {
	once.Do(func() {
		var restSetting = conf.GetSetting().Rest
		var restHandler = handler.GetHandler()
		gin.SetMode(strings.ToLower(restSetting.Mode))
		router = &Router{handler: restHandler, setting: restSetting}
		router.buildRouter()
	})
}

func GetRouter() *Router {
	if router == nil {
		panic("get router failed, please init router first")
	}
	return router
}

func (r *Router) buildRouter() {
	engine := gin.Default()
	engine.Use(gin.Recovery())

	if gin.Mode() != gin.ReleaseMode {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	engine.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, "service is alive") })
	{
		apiV1 := engine.Group("/svc-name/api/v1")
		apiV1.GET("/demo/hello", r.handler.Demo.Demo)
	}
	r.server = &http.Server{Addr: r.setting.Port, Handler: engine}
}

func (r *Router) Run() {
	go func() {
		if err := r.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic("run http server failed: " + err.Error())
		}
	}()
	fmt.Println("run http server succeed")
}

func (r *Router) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := r.server.Shutdown(ctx); err != nil {
		fmt.Println("stop http server failed: ", err.Error())
	}
	fmt.Println("stop http server succeed")
}
