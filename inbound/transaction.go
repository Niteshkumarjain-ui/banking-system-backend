package inbound

import (
	"banking-system-backend/application"
	"banking-system-backend/domain"
	"banking-system-backend/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func depositFunds(ctx *gin.Context) {

	logger := util.GetLogger()

	logger.Debugf("deposit Funds Called!!")
	var request domain.DepositFundsRequest
	var err error
	var jwtClaims domain.JwtValidate
	if ctx.GetHeader("Content-Type") != "application/json" {
		logger.Warnf("%v", ctx.GetHeader("Content-Type"))
		ctx.JSON(util.ERROR_GLOSSARY["ERR102"].HTTPStatusCode, &domain.HTTPError{
			ErrorCode:    util.ERROR_GLOSSARY["ERR102"].ErrorCode,
			ErrorMessage: util.ERROR_GLOSSARY["ERR102"].ErrorMessage,
		})
		return
	}

	if err = ctx.BindJSON(&request); err != nil {
		logger.Warnf("Bad request %v", err)
		ctx.JSON(util.ERROR_GLOSSARY["ERR103"].HTTPStatusCode, &domain.HTTPError{
			ErrorCode:    util.ERROR_GLOSSARY["ERR103"].ErrorCode,
			ErrorMessage: util.ERROR_GLOSSARY["ERR103"].ErrorMessage,
		})
		return
	}
	jwtClaims, _ = util.ValidateJWT(ctx.GetHeader("Authorization"))

	var response domain.DepositFundsResponse
	response, err = application.DepositFunds(request, jwtClaims)
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

func withdrawlFunds(ctx *gin.Context) {

	logger := util.GetLogger()

	logger.Debugf("Withdrawl Funds Called!!")
	var request domain.DepositFundsRequest
	var err error
	var jwtClaims domain.JwtValidate
	if ctx.GetHeader("Content-Type") != "application/json" {
		logger.Warnf("%v", ctx.GetHeader("Content-Type"))
		ctx.JSON(util.ERROR_GLOSSARY["ERR102"].HTTPStatusCode, &domain.HTTPError{
			ErrorCode:    util.ERROR_GLOSSARY["ERR102"].ErrorCode,
			ErrorMessage: util.ERROR_GLOSSARY["ERR102"].ErrorMessage,
		})
		return
	}

	if err = ctx.BindJSON(&request); err != nil {
		logger.Warnf("Bad request %v", err)
		ctx.JSON(util.ERROR_GLOSSARY["ERR103"].HTTPStatusCode, &domain.HTTPError{
			ErrorCode:    util.ERROR_GLOSSARY["ERR103"].ErrorCode,
			ErrorMessage: util.ERROR_GLOSSARY["ERR103"].ErrorMessage,
		})
		return
	}
	jwtClaims, _ = util.ValidateJWT(ctx.GetHeader("Authorization"))

	var response domain.DepositFundsResponse
	response, err = application.WithdrawlFunds(request, jwtClaims)
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
		if strings.Contains(err.Error(), "Insufficent Balance") {
			ctx.JSON(util.ERROR_GLOSSARY["ERR113"].HTTPStatusCode, &domain.HTTPError{
				ErrorCode:    util.ERROR_GLOSSARY["ERR113"].ErrorCode,
				ErrorMessage: util.ERROR_GLOSSARY["ERR113"].ErrorMessage,
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
