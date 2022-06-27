package main

import (
	"github.com/Diva2504/Assignment-3-GLNG-KS03-003/controller"

	//"github.com/Diva2504/Assignment-3-GLNG-KS03-003/models"
	"github.com/gin-gonic/gin"
	//"net/http"
)

func main() {
	go controller.UpdateStatus()
	r := gin.Default()
	//r.Use(cors.New(cors.Config{AllowAllOrigins: true}))
	r.GET("/status", controller.GetStatus)
	r.LoadHTMLFiles("index.html")
	r.Run("localhost:8080")

}
