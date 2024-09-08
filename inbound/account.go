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

func createAccount(ctx *gin.Context) {

	logger := util.GetLogger()

	logger.Debugf("Create Account Called !")
	var request domain.AccountRequest
	var err error

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

	var response domain.AccountResponse
	response, err = application.CreateAccount(request)
	if err != nil {
		logger.Warnf("Bad request %v", err)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			ctx.JSON(util.ERROR_GLOSSARY["ERR111"].HTTPStatusCode, &domain.HTTPError{
				ErrorCode:    util.ERROR_GLOSSARY["ERR111"].ErrorCode,
				ErrorMessage: util.ERROR_GLOSSARY["ERR111"].ErrorMessage,
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

func getAllAccount(ctx *gin.Context) {

	logger := util.GetLogger()

	logger.Debugf("Get All Account Called !")
	var err error

	var response []domain.GetAccountResponse
	response, err = application.GetAllAccount()
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

func getAccount(ctx *gin.Context) {

	logger := util.GetLogger()

	logger.Debugf("Get Account Called !")
	var err error
	var jwtClaims domain.JwtValidate
	accountID, _ := strconv.Atoi(ctx.Param("id"))

	jwtClaims, _ = util.ValidateJWT(ctx.GetHeader("Authorization"))

	var response domain.GetAccountResponse
	response, err = application.GetAccount(accountID, jwtClaims)
	if err != nil {
		logger.Warnf("Bad request %v", err)
		if strings.Contains(err.Error(), "You are not authorized to access this account") {
			ctx.JSON(util.ERROR_GLOSSARY["ERR110"].HTTPStatusCode, &domain.HTTPError{
				ErrorCode:    util.ERROR_GLOSSARY["ERR110"].ErrorCode,
				ErrorMessage: util.ERROR_GLOSSARY["ERR110"].ErrorMessage,
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

func updateAccount(ctx *gin.Context) {

	logger := util.GetLogger()

	logger.Debugf("Update Account Called !")
	var request domain.UpdateAccountRequest
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

	accountID, _ := strconv.Atoi(ctx.Param("id"))
	jwtClaims, _ = util.ValidateJWT(ctx.GetHeader("Authorization"))
	request.ID = accountID
	var response domain.AccountResponse
	response, err = application.UpdateAccount(request, jwtClaims)
	if err != nil {
		logger.Warnf("Bad request %v", err)
		if strings.Contains(err.Error(), "You are not authorized to access this account") {
			ctx.JSON(util.ERROR_GLOSSARY["ERR110"].HTTPStatusCode, &domain.HTTPError{
				ErrorCode:    util.ERROR_GLOSSARY["ERR110"].ErrorCode,
				ErrorMessage: util.ERROR_GLOSSARY["ERR110"].ErrorMessage,
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

func deleteAccount(ctx *gin.Context) {

	logger := util.GetLogger()

	logger.Debugf("Delete Account Called !")
	var err error
	var jwtClaims domain.JwtValidate
	accountID, _ := strconv.Atoi(ctx.Param("id"))

	jwtClaims, _ = util.ValidateJWT(ctx.GetHeader("Authorization"))

	var response domain.AccountResponse
	response, err = application.DeleteAccount(accountID, jwtClaims)
	if err != nil {
		logger.Warnf("Bad request %v", err)
		if strings.Contains(err.Error(), "You are not authorized to access this account") {
			ctx.JSON(util.ERROR_GLOSSARY["ERR110"].HTTPStatusCode, &domain.HTTPError{
				ErrorCode:    util.ERROR_GLOSSARY["ERR110"].ErrorCode,
				ErrorMessage: util.ERROR_GLOSSARY["ERR110"].ErrorMessage,
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
