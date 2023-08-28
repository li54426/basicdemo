package controller

import (
   "github.com/gin-gonic/gin"
  "fmt"

)



func UserInfo(c *gin.Context){ 
    fmt.Println("UserInfo")
    
}

func Register(c *gin.Context){
    fmt.Println("Register")
}



func Login(c *gin.Context){
    fmt.Println("Login")
}