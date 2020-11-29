package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	_ "nano-world-backend-golang/docs"
	"nano-world-backend-golang/routes"
	"os"
	"path/filepath"
	"time"
)

// @title Swagger nano backend API
// @version 1.0
// @description Nano backend developed in golang.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081

func initServer() *gin.Engine {

	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create(filepath.Join("logs", "gin.log"))
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.New()

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	r.GET("/getBalance/:address", routes.GetBalance())
	r.GET("/getPeers", routes.GetPeers())
	r.GET("/getGeoLocations", routes.GetGeoLocations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func main() {

	fmt.Print("listen and serve on 0.0.0.0:8081")
	initServer().Run(":8081") // listen and serve on 0.0.0.0:8081
}
