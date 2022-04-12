package server

import (
	"we-share-go/router"

	"github.com/gin-gonic/gin"
)

func ServerStart(port string) {
	r := gin.Default()
	router.RegRouter(r)
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
