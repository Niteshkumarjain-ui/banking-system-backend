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
	accountGroup.POST("", util.AuthorizeRole("customer or employee or admin"), createAccount)
	accountGroup.GET("/:id", util.AuthorizeRole("customer or employee or admin"), getAccount)
	accountGroup.GET("", util.AuthorizeRole("employee or admin"), getAllAccount)
	accountGroup.PUT("/:id", util.AuthorizeRole("customer or employee or admin"), updateAccount)
	accountGroup.DELETE("/:id", util.AuthorizeRole("customer or employee or admin"), deleteAccount)

	transactionGroup := router.Group("/transaction")
	transactionGroup.POST("/deposit", util.AuthorizeRole("customer or employee or admin"), depositFunds)
	transactionGroup.POST("/withdrawl", util.AuthorizeRole("customer or employee or admin"), withdrawlFunds)
	transactionGroup.POST("/transfer", util.AuthorizeRole("customer or employee or admin"), transferFunds)
	transactionGroup.GET("/history/:account_id", util.AuthorizeRole("customer or employee or admin"), getAccountStatement)
	transactionGroup.GET("/:transaction_id", util.AuthorizeRole("customer or employee or admin"), getTransaction)

	userGroup := router.Group("/user")
	userGroup.GET("/:id", util.AuthorizeRole("customer or employee or admin"), getUser)
	userGroup.GET("", util.AuthorizeRole("employee or admin"), getAllUser)
	userGroup.PUT("/:id", util.AuthorizeRole("customer or employee or admin"), updateUser)
	userGroup.DELETE("/:id", util.AuthorizeRole("employee or admin"), deleteUser)
	userGroup.PUT("/role/:id", util.AuthorizeRole("admin"), giveUserRole)

	logger.Infof("HTTP server staring...")
	listenAddress := fmt.Sprintf("%s:%s", util.Configuration.HTTPServer.Host, util.Configuration.HTTPServer.Port)
	if err := appRouter.Run(listenAddress); err != nil {
		logger.Errorf("HTTP server couldn't be started: %v", err)
	}
}
