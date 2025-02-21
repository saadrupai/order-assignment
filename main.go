package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"

	"github.com/saadrupai/order-assignment/app/config"
	"github.com/saadrupai/order-assignment/app/container"
)

func main() {
	g := gin.Default()
	config.SetConfig()
	container.Serve(g)
	fmt.Println("Server starting..., pid: ", strconv.Itoa(os.Getpid()))

}
