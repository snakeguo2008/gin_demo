package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/cihub/seelog"
)

/*
http头格式 application/x-www-form-urlencoded
c.PostForm 是解析Form格式的post数据
*/

type PostForm struct {
	Name string `form:"name"`
	Age string `form:"age"`
	Sex string `form:"sex"`
}
func PostFormHandle(c *gin.Context) {

	name := c.PostForm("name")
	age := c.PostForm("age")
	sex := c.DefaultPostForm("sex","默认值是男")

	var postForm PostForm
	if err := c.Bind(&postForm);err != nil{
		seelog.Errorf("post form parse failed", )
	}
	seelog.Infof("postForm:%v",postForm)

	seelog.Infof("PostsHandle,name:%v, age:%v， sex：%v", name, age, sex)
}
