package main
import (
  "github.com/gin-gonic/gin"
  "basicdemo/controller"
)




func initRouter(r *gin.Engine) {
  // 设置静态文件, . 表示当前目录下的public文件夹
  r.Static("/static", "./public")

  // 根节点, 记得加上单引号 "/"
  apiRouter := r.Group("/main")

  apiRouter.GET("/user/", controller.UserInfo)
  apiRouter.POST("/user/register/", controller.Register)
  apiRouter.POST("/user/login/", controller.Login)

  // 上传文件
  apiRouter.POST("publish/action/", controller.Publish)
  // 返回上传列表
  apiRouter.GET("/publish/list/", controller.PublishList)
  


}




func main(){

  // 创建一个 有 log 和rec 的服务器
  r:= gin.Default()
  initRouter(r)
  r.run()
  
}