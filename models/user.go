package models

import(
    // "fmt"
    "sync"
    "time"
)

type User struct{
    UserId     int64 `gorm;"primary_key"`
    Name     string
    Password    string
    signature  string
    CreateAt time.Time
    // DeleteAt time.Time
}


type UserDao struct{}

var userOnce sync.Once
var userDao *UserDao

// interface
func UserDaoInstance() *UserDao{
    userOnce.Do(
        func(){
            userDao = &UserDao{}
        })
    return userDao
}



// 根据用户名和密码，创建一个新的User，返回UserId
func (*UserDao) CreateUser(user *User) (int64, error){
    result := SqlSession.Create(&user)
    if result.Error != nil {
        return -1, result.Error
    }

    return user.UserId, nil
}


/* 
* 根据用户名，查找用户实体
*/

func (*UserDao) FindUserByName(username string)(*User, error){
    user := User{Name: username }
    result := SqlSession.Where("name=?", username).First(&user)
    err := result.Error

    if err != nil {
        return nil, err
    }

    return &user, err
}



/* 
* 根据用户id，查找用户实体
*/

func (*UserDao) FindUserById(id int64)(*User, error){
    user := User{UserId: id }
    result := SqlSession.Where("user_id=?", id).First(&user)
    err := result.Error

    if err != nil {
        return nil, err
    }

    return &user, err
}









