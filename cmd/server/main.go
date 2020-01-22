package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Kit-Bunrong/sample-rest/cmd/server/apis"
	"github.com/Kit-Bunrong/sample-rest/cmd/server/config"
	"github.com/Kit-Bunrong/sample-rest/cmd/server/httputil"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "github.com/Kit-Bunrong/sample-rest/cmd/server/docs"
)

// @title Server Swagger API
// @version 1.0
// @description Swagger API for Golang Project Server
// @termOfService http://swagger.io/terms/

// @contact.name API Support 
// @contect.email Kit-Bunrong@gmail.com

// @license.name MIT

// @BasePath /api/v1

func main() {
	// Create a router without any middleware by default
	r := gin.Default()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.Use(auth())
		v1.GET("/users/:id", apis.GetUser)
	}

	config.Config.DB, config.Config.DBErr = gorm.Open("postgres", config.Config.DNS)
	// config.Config.DB, config.Config.DBErr = gorm.Open("postgres", "host=localhost dbname=postgres sslmode=disable user=postgres password=new-sqlpostgre")
	if config.Config.DBErr != nil {
		panic(config.Config.DBErr)
	}

	// config.Config.DB.AutoMigrate(&models.User{}) // This is needed for generation of schema for postgres images

	defer config.Config.DB.Close()

	log.Println("Successfully connected to database")
	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))

}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		if authHeader != config.Config.ApiKey {
			httputil.NewError(c, http.StatusUnauthorized, fmt.Errorf("This user isn't authorized to this operation: api_key=%s", authHeader))
			c.Abort()
		}
		c.Next()
	}
}
