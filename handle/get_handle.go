package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/cihub/seelog"
)

//需要注意一下, 在linux下 &符号会让进程在后台运行所以需要使用转义符号 \&
//curl -v  http://192.168.1.220:8787/get_method?age=1\&name=xiaoming
// **  get 参数的解析实际是form类型的解析
type JsonGet struct {
	Name string `form:"name"`
	Age  uint32 `form:"age"`
}
func GetHandle(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	seelog.Infof("GetHandle,name:%v, age:%v", name, age)
	var jsonD JsonGet
	if err := c.Bind(&jsonD);err != nil{
		seelog.Errorf("json parse error")
	}
	seelog.Infof("GET METHOD PARSE JSOND:%v",jsonD)
}
