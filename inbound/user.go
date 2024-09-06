package inbound

import (
	"banking-system-backend/application"
	"banking-system-backend/domain"
	"banking-system-backend/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func register(ctx *gin.Context) {

	logger := util.GetLogger()

	logger.Debugf("User Registry Called !")
	var request domain.UserRegisterRequest
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

	var response domain.UserRegisterResponse
	response, err = application.Register(request)
	if err != nil {
		logger.Warnf("Bad request %v", err)
		if err == bcrypt.ErrMismatchedHashAndPassword || strings.Contains(err.Error(), "bcrypt") {
			ctx.JSON(util.ERROR_GLOSSARY["ERR106"].HTTPStatusCode, &domain.HTTPError{
				ErrorCode:    util.ERROR_GLOSSARY["ERR106"].ErrorCode,
				ErrorMessage: util.ERROR_GLOSSARY["ERR106"].ErrorMessage,
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

func login(ctx *gin.Context) {

	logger := util.GetLogger()

	logger.Debugf("User Login Called !")
	var request domain.UserLoginRequest
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

	if request.Username == "" && request.Email == "" {
		ctx.JSON(util.ERROR_GLOSSARY["ERR101"].HTTPStatusCode, &domain.HTTPError{
			ErrorCode:    util.ERROR_GLOSSARY["ERR101"].ErrorCode,
			ErrorMessage: util.ERROR_GLOSSARY["ERR101"].ErrorMessage,
		})
		return
	}

	var response domain.UserLoginResponse
	response, err = application.Login(request)
	if err != nil {
		logger.Warnf("Bad request %v", err)
		if err == bcrypt.ErrMismatchedHashAndPassword || strings.Contains(err.Error(), "bcrypt") {
			ctx.JSON(util.ERROR_GLOSSARY["ERR106"].HTTPStatusCode, &domain.HTTPError{
				ErrorCode:    util.ERROR_GLOSSARY["ERR106"].ErrorCode,
				ErrorMessage: util.ERROR_GLOSSARY["ERR106"].ErrorMessage,
			})
			return
		}
		if strings.Contains(err.Error(), "record not found") {
			ctx.JSON(util.ERROR_GLOSSARY["ERR104"].HTTPStatusCode, &domain.HTTPError{
				ErrorCode:    util.ERROR_GLOSSARY["ERR104"].ErrorCode,
				ErrorMessage: util.ERROR_GLOSSARY["ERR104"].ErrorMessage,
			})
			return
		}
		if strings.Contains(err.Error(), "Invalid") {
			ctx.JSON(util.ERROR_GLOSSARY["ERR107"].HTTPStatusCode, &domain.HTTPError{
				ErrorCode:    util.ERROR_GLOSSARY["ERR107"].ErrorCode,
				ErrorMessage: util.ERROR_GLOSSARY["ERR107"].ErrorMessage,
			})
			return
		}
		if strings.Contains(err.Error(), "jwt") {
			ctx.JSON(util.ERROR_GLOSSARY["ERR108"].HTTPStatusCode, &domain.HTTPError{
				ErrorCode:    util.ERROR_GLOSSARY["ERR108"].ErrorCode,
				ErrorMessage: util.ERROR_GLOSSARY["ERR108"].ErrorMessage,
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

func session(ctx *gin.Context) {
	var err error
	var jwt domain.JwtValidate
	var response domain.UserSessionResponse
	logger := util.GetLogger()

	logger.Debugf("Get Session Called!")

	token := ctx.GetHeader("Authorization")

	jwt, err = util.ValidateJWT(token)

	if err != nil {
		ctx.JSON(util.ERROR_GLOSSARY["ERR109"].HTTPStatusCode, &domain.HTTPError{
			ErrorCode:    util.ERROR_GLOSSARY["ERR109"].ErrorCode,
			ErrorMessage: util.ERROR_GLOSSARY["ERR109"].ErrorMessage,
		})
		return
	}
	response.UserId = jwt.Claims["user_id"].(float64)
	response.Email = jwt.Claims["email"].(string)
	response.Role = jwt.Claims["role"].(string)

	ctx.JSON(http.StatusOK, response)
}
