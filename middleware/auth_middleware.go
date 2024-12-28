package middleware

import (
	"log"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

var secretKey = os.Getenv("SECRET_KEY")

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el valor del encabezado Authorization
		authHeader := c.GetHeader("Authorization")

		// Verificar que el encabezado no esté vacío
		if authHeader == "" {
			log.Println("Authorization header is missing")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Eliminar el prefijo "Bearer " si está presente
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Loguear el token recibido (aunque no sea una buena práctica loguear tokens reales en producción)
		log.Printf("Received token: %s", tokenString)

		// Verificar el formato del token
		if len(strings.Split(tokenString, ".")) != 3 {
			log.Println("Invalid JWT token format")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT token format"})
			c.Abort()
			return
		}

		// Parsear el token y verificar su validez
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		// Validar el token y extraer las reclamaciones
		if err != nil {
			log.Printf("Error parsing token: %s", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Println("Token is valid, extracting claims...")
			// Loguear las reclamaciones si es necesario
			log.Printf("User claims: %v", claims)
			c.Set("user", claims)
		} else {
			log.Println("Invalid token claims")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
