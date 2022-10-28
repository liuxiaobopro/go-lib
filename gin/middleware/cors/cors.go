package cors

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			AllowOrigins:  []string{"*"},
			AllowMethods:  []string{"*"},
			AllowHeaders:  []string{"Origin"},
			ExposeHeaders: []string{"Content-Length", "Authorization"},
			// AllowCredentials: true,
			// AllowOriginFunc: func(origin string) bool {
			//	return origin == "https://github.com"
			// },
			MaxAge: 12 * time.Hour,
		})
	}
}
