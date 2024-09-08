package util

import (
	"banking-system-backend/domain"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthorizeRole(requiredRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		var jwtClaims domain.JwtValidate
		tokenString := ctx.GetHeader("Authorization")
		jwtClaims, err = ValidateJWT(tokenString)
		if err != nil {
			ctx.JSON(ERROR_GLOSSARY["ERR109"].HTTPStatusCode, &domain.HTTPError{
				ErrorCode:    ERROR_GLOSSARY["ERR109"].ErrorCode,
				ErrorMessage: ERROR_GLOSSARY["ERR109"].ErrorMessage,
			})
			ctx.Abort()
			return
		}

		role := jwtClaims.Claims["role"].(string)
		fmt.Println(role)
		fmt.Println(requiredRole)
		if !strings.Contains(requiredRole, role) {
			ctx.JSON(ERROR_GLOSSARY["ERR110"].HTTPStatusCode, &domain.HTTPError{
				ErrorCode:    ERROR_GLOSSARY["ERR110"].ErrorCode,
				ErrorMessage: ERROR_GLOSSARY["ERR110"].ErrorMessage,
			})
			ctx.Abort()
			return
		}

		ctx.Next() // Proceed if authorized
	}
}
