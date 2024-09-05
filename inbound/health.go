package inbound

import (
	"banking-system-backend/application"
	"banking-system-backend/domain"
	"banking-system-backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthGet(ctx *gin.Context) {

	logger := util.GetLogger()

	logger.Debugf("Health up!")

	response, err := application.HealthGet()
	if err != nil {
		ctx.JSON(util.ERROR_GLOSSARY["ERR101"].HTTPStatusCode, &domain.HTTPError{
			ErrorCode:    util.ERROR_GLOSSARY["ERR101"].ErrorCode,
			ErrorMessage: util.ERROR_GLOSSARY["ERR101"].ErrorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
