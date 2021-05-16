package rest

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"ddd/api/rest/handler"
	"ddd/application"
	"ddd/conf"
)

type Router struct {
	setting *conf.RestSetting
	server  *http.Server
	handler *handler.Factory
}

func NewRouter(setting *conf.RestSetting, controller *application.Factory) *Router {

	var router = &Router{}
	gin.SetMode(strings.ToLower(setting.Mode))
	router.setting = setting

	var demoHandler = handler.NewUserHandler(controller.UserSvc)
	router.handler = handler.NewHandler(demoHandler)

	router.buildRouter()
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
		apiV1 := engine.Group("/ddd/api/v1")
		apiV1.POST("/user", r.handler.UserHdl.AddUser)
		apiV1.GET("/users", r.handler.UserHdl.GetUsers)
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
