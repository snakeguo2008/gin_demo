package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin_demo/handle"
	"github.com/cihub/seelog"
)

func pings(c *gin.Context) {
	name := c.Query("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func main() {
	defer seelog.Flush()
	r := gin.Default()
	r.GET("/ping", pings)
	r.GET("/get_method",handle.GetHandle)
	r.POST("/post_method_form",handle.PostFormHandle)
	r.POST("/post_json_form",handle.PostJsonHandle)

	//默认监听8080
	r.Run(":8787")

}
