package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"../api"
	"log"
	"net/http"
	"strconv"
)

func run(server_addr string, client_addr string) {
	conn, err := grpc.Dial(server_addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := api.NewUUIDSenderClient(conn)

	g := gin.Default()
	g.GET("/id", func(ctx *gin.Context) {
		req := &api.Empty{}

		if response, err := client.IdSend(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H {
				"result": fmt.Sprint(response.Uuid),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H {
				"error": err.Error(),
			})
		}
	})

	if err := g.Run(client_addr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}


func main() {
	for i := 0; i < 4; i++ {
		go run("localhost:404" + strconv.Itoa(i), ":808" + strconv.Itoa(i))
	}
	fmt.Scanf("%d")
}
