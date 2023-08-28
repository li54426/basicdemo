package controller

import(
    "testing"
    "fmt"
)

func TestGenerateToken(t *testing.T){
    got, _ := GenerateToken("testusername")
    srcstring, _ := ParseToken(got)
    fmt.Println("srcstring=", srcstring)
    fmt.Println("got=", got)
}
