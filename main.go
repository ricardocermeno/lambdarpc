package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/blmayer/awslambdarpc/client"
	"github.com/gin-gonic/gin"
)

const LAMBDA_DEST_PORT_ENV = "LAMBDA_DEST_PORT"

func proxyController(c *gin.Context) {

	lambdaport := os.Getenv(LAMBDA_DEST_PORT_ENV)
	if lambdaport == "" {
		fmt.Println("Lambda port dont configured")
		return
	}

	payload, err := io.ReadAll(c.Request.Body)

	if err != nil {
		fmt.Println("payload doesn't provided")
		return
	}

	duration := 15 * time.Second
	res, err := client.Invoke(lambdaport, payload, duration)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(http.StatusOK, string(res))
}

func main() {
	app := gin.Default()

	app.POST("/proxy", proxyController)

	app.Run()
}
