package middleware

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

// middleware de autenticación, validará el token recibido en el header (debe coincidir con: testAdminToken321), 
// con el token almacenado en el archivo .env
// go get github.com/joho/godotenv 

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		adminToken := os.Getenv("adminToken")
		tokenHeader := ctx.GetHeader("token")
		if tokenHeader != adminToken {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"mensaje": "Usuario No autenticado como Administrador.",
			})
		} 
		ctx.Next()
	}
}