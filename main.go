package main

import (
	"cs5234/client"
	"cs5234/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	client.InitDB()
	r := gin.Default()
	router.InitRouters(r)
	err := r.Run()
	if err != nil {
		log.Fatalf("init gin error, error=%v", err)
	}
}
