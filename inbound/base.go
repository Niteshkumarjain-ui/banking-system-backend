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
	accountGroup.POST("", util.AuthorizeRole(util.ALL_USER_ALLOWED), createAccount)
	accountGroup.GET("/:id", util.AuthorizeRole(util.ALL_USER_ALLOWED), getAccount)
	accountGroup.GET("", util.AuthorizeRole(util.EMP_ADMIN_ALLOWED), getAllAccount)
	accountGroup.PUT("/:id", util.AuthorizeRole(util.ALL_USER_ALLOWED), updateAccount)
	accountGroup.DELETE("/:id", util.AuthorizeRole(util.ALL_USER_ALLOWED), deleteAccount)

	transactionGroup := router.Group("/transaction")
	transactionGroup.POST("/deposit", util.AuthorizeRole(util.ALL_USER_ALLOWED), depositFunds)
	transactionGroup.POST("/withdrawl", util.AuthorizeRole(util.ALL_USER_ALLOWED), withdrawlFunds)
	transactionGroup.POST("/transfer", util.AuthorizeRole(util.ALL_USER_ALLOWED), transferFunds)
	transactionGroup.GET("/history/:account_id", util.AuthorizeRole(util.ALL_USER_ALLOWED), getAccountStatement)
	transactionGroup.GET("/:transaction_id", util.AuthorizeRole(util.ALL_USER_ALLOWED), getTransaction)

	userGroup := router.Group("/user")
	userGroup.GET("/:id", util.AuthorizeRole(util.ALL_USER_ALLOWED), getUser)
	userGroup.GET("", util.AuthorizeRole(util.EMP_ADMIN_ALLOWED), getAllUser)
	userGroup.PUT("/:id", util.AuthorizeRole(util.ALL_USER_ALLOWED), updateUser)
	userGroup.DELETE("/:id", util.AuthorizeRole(util.EMP_ADMIN_ALLOWED), deleteUser)
	userGroup.PUT("/role/:id", util.AuthorizeRole(util.ADMIN_ALLOWED), giveUserRole)

	reportGroup := router.Group("/report")
	reportGroup.GET("/account-balance/:id", util.AuthorizeRole(util.ALL_USER_ALLOWED), getAccountBalance)
	reportGroup.GET("/daily-transaction", util.AuthorizeRole(util.EMP_ADMIN_ALLOWED), getDailyTransaction)
	reportGroup.GET("/user-financial/:id", util.AuthorizeRole(util.ALL_USER_ALLOWED), getUserFinancial)

	logger.Infof("HTTP server staring...")
	listenAddress := fmt.Sprintf("%s:%s", util.Configuration.HTTPServer.Host, util.Configuration.HTTPServer.Port)
	if err := appRouter.Run(listenAddress); err != nil {
		logger.Errorf("HTTP server couldn't be started: %v", err)
	}
}
