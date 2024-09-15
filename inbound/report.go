package inbound

import (
	"banking-system-backend/application"
	"banking-system-backend/domain"
	"banking-system-backend/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getAccountBalance(ctx *gin.Context) {
	span_ctx, span := util.InboudGetSpan(ctx, "getAccountBalance")
	defer span.End()

	logger := util.GetLogger()

	logger.Debugf("Get Account Balance Called !")
	var err error
	var jwtClaims domain.JwtValidate
	accountID, _ := strconv.Atoi(ctx.Param("id"))

	jwtClaims, _ = util.ValidateJWT(ctx.GetHeader("Authorization"))

	var response domain.GetAccountBalanceResponse
	response, err = application.GetAccountBalance(span_ctx, accountID, jwtClaims)
	if err != nil {
		logger.Warnf("Bad request %v", err)
		if strings.Contains(err.Error(), "You are not authorized to access this account") {
			ctx.JSON(util.ERROR_GLOSSARY["ERR110"].HTTPStatusCode, &domain.HTTPError{
				ErrorCode:    util.ERROR_GLOSSARY["ERR110"].ErrorCode,
				ErrorMessage: util.ERROR_GLOSSARY["ERR110"].ErrorMessage,
			})
			return
		}
		if strings.Contains(err.Error(), "Account Not Found") {
			ctx.JSON(util.ERROR_GLOSSARY["ERR112"].HTTPStatusCode, &domain.HTTPError{
				ErrorCode:    util.ERROR_GLOSSARY["ERR112"].ErrorCode,
				ErrorMessage: util.ERROR_GLOSSARY["ERR112"].ErrorMessage,
			})
			return
		}
		ctx.JSON(util.ERROR_GLOSSARY["ERR105"].HTTPStatusCode, &domain.HTTPError{
			ErrorCode:    util.ERROR_GLOSSARY["ERR105"].ErrorCode,
			ErrorMessage: util.ERROR_GLOSSARY["ERR105"].ErrorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func getUserFinancial(ctx *gin.Context) {
	span_ctx, span := util.InboudGetSpan(ctx, "getUserFinancial")
	defer span.End()

	logger := util.GetLogger()

	logger.Debugf("Get User Financial Called !")
	var err error
	var jwtClaims domain.JwtValidate
	accountID, _ := strconv.Atoi(ctx.Param("id"))

	jwtClaims, _ = util.ValidateJWT(ctx.GetHeader("Authorization"))

	var response domain.GetFinancialReportResponse
	response, err = application.GetFinancialReport(span_ctx, accountID, jwtClaims)
	if err != nil {
		logger.Warnf("Bad request %v", err)
		if strings.Contains(err.Error(), "You are not authorized to access this account") {
			ctx.JSON(util.ERROR_GLOSSARY["ERR110"].HTTPStatusCode, &domain.HTTPError{
				ErrorCode:    util.ERROR_GLOSSARY["ERR110"].ErrorCode,
				ErrorMessage: util.ERROR_GLOSSARY["ERR110"].ErrorMessage,
			})
			return
		}
		if strings.Contains(err.Error(), "Account Not Found") {
			ctx.JSON(util.ERROR_GLOSSARY["ERR112"].HTTPStatusCode, &domain.HTTPError{
				ErrorCode:    util.ERROR_GLOSSARY["ERR112"].ErrorCode,
				ErrorMessage: util.ERROR_GLOSSARY["ERR112"].ErrorMessage,
			})
			return
		}
		ctx.JSON(util.ERROR_GLOSSARY["ERR105"].HTTPStatusCode, &domain.HTTPError{
			ErrorCode:    util.ERROR_GLOSSARY["ERR105"].ErrorCode,
			ErrorMessage: util.ERROR_GLOSSARY["ERR105"].ErrorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func getDailyTransaction(ctx *gin.Context) {
	span_ctx, span := util.InboudGetSpan(ctx, "getDailyTransaction")
	defer span.End()

	logger := util.GetLogger()

	logger.Debugf("Get Daily Transaction Called !")
	var err error
	var response []domain.GetDailyTransactionReportResponse
	response, err = application.GetDailyTransactionReport(span_ctx)
	if err != nil {
		logger.Warnf("Bad request %v", err)
		ctx.JSON(util.ERROR_GLOSSARY["ERR105"].HTTPStatusCode, &domain.HTTPError{
			ErrorCode:    util.ERROR_GLOSSARY["ERR105"].ErrorCode,
			ErrorMessage: util.ERROR_GLOSSARY["ERR105"].ErrorMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
