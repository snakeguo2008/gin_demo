package handle

import (
	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

/*
http头格式 application/x-www-form-urlencoded
*/
type PostJson struct {
	Name string `json:"name"`
	Age uint32 `json:"age"`
	Sex string `json:"sex"`
}
func PostJsonHandle(c *gin.Context) {
	var jsonData PostJson
	if err := c.BindJSON(&jsonData); err != nil{
		seelog.Infof("json parse failed, err:%v",err)
	}

	seelog.Infof("PostsHandle,name:%v, age:%v， sex：%v", jsonData.Name, jsonData.Age, jsonData.Sex)
}
