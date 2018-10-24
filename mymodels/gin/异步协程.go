package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"net/http"
)



func main() {

	router := gin.Default()

	router.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second) // 阻塞返回
		log.Println("Done! in path" + c.Request.URL.Path)
		c.String(http.StatusOK,"this is sync")
	})

	router.GET("/async", func(c *gin.Context) {
		cCp := c.Copy()  // todo 请求的上下文需要copy到异步的上下文，并且这个上下文是只读的
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path" + cCp.Request.URL.Path)
		}()
		c.String(http.StatusOK,"this is async")
	})


	router.Run(":8001")
}
