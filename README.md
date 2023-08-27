# basicdemo

使用go语言来一个后端




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