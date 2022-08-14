package main

import (
	"fmt"
	"log"
	"time"

	handler "go-grpc/client/handler"
	userPB "go-grpc/client/proto"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {

	r := gin.Default()

	logger, _ := zap.NewProduction()

	/*  Add a ginzap middleware, which:
	    - Logs all requests, like a combined access and error log.
	    - Logs to stdout.
		- RFC3339 with UTC time format.
	*/
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	/* Logs all panic to error log - stack means whether output the stack info. */
	r.Use(ginzap.RecoveryWithZap(logger, true))

	tokenConn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}

	client := userPB.NewUserServiceClient(tokenConn)

	/* added default port address*/
	port := ":5000"

	/* User handler */
	handler.NewUserHandler(r, client)

	r.Run(port)

	fmt.Println("Server is running on port", port)
}
