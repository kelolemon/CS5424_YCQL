package main

import (
	"cs5234/dao"
	"log"

	"cs5234/client"
	"cs5234/router"

	"github.com/gin-gonic/gin"
)

func main() {
	dao, err := dao.NewDao(client.GetDBConfig())
	if err != nil {
		log.Fatalf("error creating db: %s", err)
		return
	}
	defer dao.Session.Close()
	r := gin.Default()
	router.InitRouters(r)
	err = r.Run()
	if err != nil {
		log.Fatalf("init gin error, error=%v", err)
	}
}
