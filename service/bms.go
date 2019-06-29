package main

import (
	_ "code.hanx.xin/bms/config"
	"code.hanx.xin/bms/database"
	"code.hanx.xin/bms/web/controllers"
	"code.hanx.xin/bms/web/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := database.InitDatabaseConnect(); err != nil {
		panic(fmt.Sprintf("初始化数据库失败,%s", err.Error()))
	}
	bms := gin.Default()
	bms.Use(middleware.LogHandleStatus())
	bmsV1 := bms.Group("/v1")
	{
		bmsV1.GET("/book/:id", middleware.MarkTraceID(), controllers.GetBookByIDHandlers)
		bmsV1.GET("/book", controllers.GetBooksListHandlers)
		bmsV1.POST("/book", controllers.PostBooksListHandlers)
		bmsV1.PUT("/book", controllers.PutBooksListHandlers)
		bmsV1.DELETE("/book/:id", controllers.DeleteBooksListHandlers)
	}
	srv := &http.Server{
		Addr:    ":8899",
		Handler: bms,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("listen error,%v", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("关闭server")
	// 延时五秒等待,等待请求处理完,5秒后强制停止
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("server shutdown error %v", err)
	} else {
		fmt.Println("server shutdown success")
	}

}
