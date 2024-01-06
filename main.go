package main

import (
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func returnRequestBody(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, "error while retreiving body frorm request")
	}
	waitTime := time.Duration(rand.Intn(2)+7) * time.Second
	time.Sleep(waitTime)
	ctx.String(http.StatusAccepted, string(body))
}

func main() {

	server := gin.Default()
	server.POST("/helloworld", returnRequestBody)

	server.Run()

}
