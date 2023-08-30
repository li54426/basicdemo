package models

import(
    "testing"
)


// 有个问题, 找不到config/application.yaml:
func TestInitDB(t *testing.T){
    InitDB()
}
