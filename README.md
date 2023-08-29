# basicdemo

使用go语言来一个后端



#### 项目目录结构

其中：

- **configs** 目录包含项目的配置文件
- **controllers** 目录包含**控制器**文件
- **middleware** 目录包含**中间件**文件
- **models** 目录包含模型（ORM）文件
- **routes** 目录包含**路由**文件
- **services** 目录包含服务文件
- **utils** 目录包含多个工具文件
- **main.go** 是项目的入口文件，`README.md` 是项目的介绍文件。

```bash
├── configs
│   ├── config.yaml
│   └── db.yaml
├── controllers
│   ├── auth_controller.go
│   └── user_controller.go
├── middleware
│   ├── auth_middleware.go
│   └── logger_middleware.go
├── models
│   ├── db.go
│   ├── user.go
│   └── ...
├── routes
│   └── routes.go
├── services
│   ├── auth_service.go
│   └── user_service.go
├── utils
│   ├── response.go
│   └── ...
├── main.go
└── README.md

```









### 1 类型设计


#### 1.1 go用户字段
```go



```



#### 1.2 go文件字段








#### 方法设计

POST 和 GET 是 HTTP 请求方法，它们的 HTTP 报文格式如下：

POST 请求的 HTTP 报文格式：
```plaintext
POST /path HTTP/1.1
Host: example.com
Content-Type: application/json

{
  "name": "John",
  "age": 30
}
```
在上面的 HTTP 报文中，POST 方法用于向服务器发送数据。请求的路径是 /path，HTTP 版本是 HTTP/1.1。请求报头中包含了 Host 和 Content-Type 两个字段，分别表示请求的主机名和请求数据的类型。请求正文部分是一个 JSON 格式的数据，表示请求的具体内容。

GET 请求的 HTTP 报文格式：
```plaintext
GET /path?key1=value1&key2=value2 HTTP/1.1
Host: example.com
```
在上面的 HTTP 报文中，GET 方法用于从服务器获取数据。请求的路径是 /path，HTTP 版本是 HTTP/1.1。请求报头中包含了 Host 字段，表示请求的主机名。请求的查询字符串是 key1=value1&key2=value2，表示请求的具体参数。
需要注意的是，POST 和 GET 请求的 HTTP 报文格式是基本相同的，区别在于请求方法和请求正文部分。POST 请求的正文部分包含了请求的数据，而 GET 请求的查询字符串包含了请求的参数。





#### 试运行1 什么函数都没有


```bash

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] GET    /main/user/               --> basicdemo/controller.UserInfo (3 handlers)
[GIN-debug] POST   /main/user/register/      --> basicdemo/controller.Register (3 handlers)
[GIN-debug] POST   /main/user/login/         --> basicdemo/controller.Login (3 handlers)
[GIN-debug] POST   /main/publish/action/     --> basicdemo/controller.Publish (3 handlers)
[GIN-debug] GET    /main/publish/list/       --> basicdemo/controller.PublishList (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2023/08/27 - 21:56:53 | 404 |         922ns |    39.129.5.209 | GET      "/"
```




#### 单例模式

```go
package main

import (
    "fmt"
    "sync"
)

var once sync.Once

func doSomething() {
    once.Do(func() {
        fmt.Println("This function will only be executed once.")
    })
}

func main() {
    doSomething()
    do something()
}

```




### 功能


#### token
以下是 Token 的一些常见使用场景：
- 单点登录（SSO）：Token 可以用于实现单点登录，即在多个应用程序或系统之间共享用户身份验证信息。用户只需登录一次，即可访问所有相关的应用程序和系统。
- 会话管理：Token 可以用于管理会话，即在服务器端生成一个唯一的 Token，并将其发送到客户端。客户端可以使用该 Token 访问服务器，并且服务器可以验证该 Token 的有效性，以确保请求来自合法的客户端。
- 资源保护：Token 可以用于保护敏感资源，例如文件、数据库行、API 请求等。只有拥有有效 Token 的用户才能访问这些资源。
- 移动应用程序：Token 在移动应用程序中非常常见，因为它们可以用于保护用户信息，并在多个服务器之间共享用户身份验证信息。

