package main

import (
	"encoding/json"
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
		fmt.Println(err)
		c.JSON(http.StatusBadGateway, gin.H{
			"err": err.Error(),
		})
		return
	}

	var a struct {
		StatusCode        int            `json:"statusCode"`
		Headers           map[string]any `json:"headers"`
		MultiValueHeaders any            `json:"multiValueHeaders"`
		Body              any            `json:"body"`
	}

	fmt.Println(string(res))
	json.Unmarshal(res, &a)

	c.IndentedJSON(http.StatusOK, a)
}

func main() {
	app := gin.Default()

	app.Any("/", proxyController)

	app.Run()
}
