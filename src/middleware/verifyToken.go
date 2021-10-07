package middleware

import (
	"log"
	"net/http"

	"github.com/gopla/maro/src/helper"
	"github.com/gopla/maro/src/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ExtractToken(c *gin.Context) string {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	
	if authHeader == "" {
		response := helper.BuildErrorResponse("No token found", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	tokenString := authHeader[len(BEARER_SCHEMA):]
	return tokenString
}

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := ExtractToken(c)		

		token, err := jwtService.ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse( err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

	}
}