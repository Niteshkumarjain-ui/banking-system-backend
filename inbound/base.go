package inbound

import (
	"banking-system-backend/util"
	"fmt"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func HttpService() {
	defer util.GlobalWaitGroup.Done()

	logger := util.GetLogger()

	gin.SetMode(gin.ReleaseMode)
	appRouter := gin.Default()

	appRouter.Use(gzip.Gzip(gzip.DefaultCompression))

	router := appRouter.Group(fmt.Sprintf("/api/%s", util.Configuration.Meta.Version))

	healthGroup := router.Group("/health")
	healthGroup.GET("", healthGet)

	userAuth := router.Group("/auth")
	userAuth.POST("/register", register)
	userAuth.POST("/login", login)
	userAuth.GET("/session", session)

	accountGroup := router.Group("/account")
	accountGroup.POST("", util.AuthorizeRole("customer or employee"), createAccount)
	accountGroup.GET("/:id", util.AuthorizeRole("customer or employee"), getAccount)
	accountGroup.GET("", util.AuthorizeRole("employee"), getAllAccount)
	accountGroup.PUT("/:id", util.AuthorizeRole("customer or employee"), updateAccount)
	accountGroup.DELETE("/:id", util.AuthorizeRole("customer or employee"), deleteAccount)

	transactionGroup := router.Group("/transaction")
	transactionGroup.POST("/deposit", util.AuthorizeRole("customer or employee"), depositFunds)
	transactionGroup.POST("/withdrawl", util.AuthorizeRole("customer or employee"), withdrawlFunds)

	logger.Infof("HTTP server staring...")
	listenAddress := fmt.Sprintf("%s:%s", util.Configuration.HTTPServer.Host, util.Configuration.HTTPServer.Port)
	if err := appRouter.Run(listenAddress); err != nil {
		logger.Errorf("HTTP server couldn't be started: %v", err)
	}
}
