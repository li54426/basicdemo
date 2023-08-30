package controller

import (
	"fmt"
	// "net/http"
	"net/http"
    "basicdemo/models"

	"github.com/gin-gonic/gin"
)


func decode(c *gin.Context){
    return
}


type UserResponse struct{
    Response
    User User `json: "user"`
}


func UserInfo(c *gin.Context){ 
    token := c.Query("token")

    username, err := ParseToken(token)

    fmt.Println("token=", token)



    if err != nil {
        c.JSON(http.StatusOK, UserResponse{
            Response: Response{StatusCode: 1, 
                              StatusMsg: "User dont exist"},
        })
    }else{
        userModel, err := models.UserDaoInstance().FindUserByName(username)
        fmt.Println("fin dao")
        
        
        if err == nil{
            c.JSON(http.StatusOK, UserResponse{
                Response: Response{StatusCode:1},
                User: User{ userModel.UserId, userModel.Name},
            })            
        }else{
            c.JSON(http.StatusOK, UserResponse{
                Response: Response{StatusCode: 1, 
                                  StatusMsg: "User dont exist"},
                })
        }
        

    }
    
    fmt.Println("Fin UserInfo")
    
}



type UserIdTokenResponse struct{
    UserId int64 `json:"user_id, omitempty"`
    Token string `json:"token"`
}





func Register(c *gin.Context){
    // decode(c)
    // username := c.Query("username")
    // password := c.Query("password")
    // token, err := RegisterService(username, password)    
    fmt.Println("Register")
}

// func RegisterService(username string, password string) (UserIdTokenResponse, error){
    
// }


func Login(c *gin.Context){
    fmt.Println("Login")
}