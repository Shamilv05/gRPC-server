package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"../api"
	"log"
	"net/http"
)

func run(server_addr string, client_addr string) {
	conn, err := grpc.Dial(server_addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := api.NewCombinerClient(conn)

	g := gin.Default()
	g.GET("/string", func(ctx *gin.Context) {
		req := &api.Empty{}

		if response, err := client.Request(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H {
				"result": fmt.Sprint(response.Value),
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
	run("localhost:4048", ":8099")
}
