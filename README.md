# Gin学习笔记

## 一、简介

Gin是使用Go（Golang）语言编写的Web框架，由于httprouter，它具有类似 martini-like API，性能提高了40倍。如果你需要高性能和高生产力，你会喜欢上Gin。

<img src="asstes/gin_logo.png" alt="gin_logo" style="zoom: 10%;" />

## 二、安装Gin

需要安装Gin-Package，您需要先安装Go并设置Go工作区。

1. 您首先要安装Go**（需要的版本 1.16+）**，然后您可以使用下面的Go命令安装Gin。

```shell
go get -u github.com/gin-gonic/gin
```

2. 将其倒入项目或者代码中：

```shell
import "github.com/gin-gonic/gin"
```

3. （可选）导入。例如，如果使用`.net/http`包`http.StatusOK`

```shell
import "net/http"
```

## 三、Hello Gin

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建默认engine
	engine := gin.Default()
	// engine调用Handle方法处理GET请求
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		// 响应请求
		// gin.H 返回的是一个map类型的数据结构
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    "hello gin",
		})
	})
	// engine启动
	engine.Run(":8000")
}
```

<img src="asstes/image-20221211155042002.png" alt="image-20221211155042002" style="zoom:50%;" />

## 四、Gin路由

### 4.1 创建Engine

在Gin中，Engine被定义为一个结构体【struct】数据类型，Engine其中定义了路由组、中间件、页面渲染接口、Gin相关配置设置等内容，默认的Engine可以通过`gin.Default()`或者使用`gin.New()`来创建。

两种方式创建Engine如下所示：

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	/*
		两种创建Engine的方式
		第一种：使用gin.Default()
		第二种：使用gin.New()
	*/
	engine1 := gin.Default()
	engine1.Run(":8000")
	//engine2 := gin.New()
  //engine2.Run(":8000")
}
```

使用`gin.Default()`创建Engine的日志信息如下：

<img src="asstes/image-20221211155254654.png" alt="image-20221211155254654" style="zoom:50%;" />

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	/*
		两种创建Engine的方式
		第一种：使用gin.Default()
		第二种：使用gin.New()
	*/
	//engine1 := gin.Default()
	//engine1.Run(":8000")
	engine2 := gin.New()
  engine2.Run(":8000")
}
```

使用`gin.New()`创建Engine的日志如下：

<img src="asstes/image-20221211155547259.png" alt="image-20221211155547259" style="zoom:50%;" />

使用`gin.Default()`和`gin.New()`创建Engine的区别就在于使用`gin.Default()`创建的Engine会默认使用Logger和Recovery这个两个中间件。

Logger中间件是负责进行打印并输出日志的中间件，方便开发者进行程序调试；Recovery中间件的作用是如果程序执行过程中遇到panic中断信号，则Recovery中间件会恢复Go应用程序的执行，并返回服务器500内部错误，通常情况下，我们都是使用`gin.Default()`来创建Engine实例的。

### 4.2 http请求类型

http请求中一共定义了8种方法或者称为8种http请求类型来表明对网络资源（Request-URI）的不同操作方式，分别是：OPTIONS、HEAD、GET、POST、PUT、DELETE、TRACE、CONNECT。

一共有8种HTTP请求，但是实际开发中常用就：GET、POST、PUT、DELETE这几种HTTP请求。

### 4.3 通用请求处理

Engine实例中可以对http请求进行处理。在Engine实例中可以使用Handle()方法对HTTP请求进行处理。

Handle()方法源码如下：

- httpMethod：第一个参数表示处理的http请求的类型，是http请求分类8种其中的1种。
- relativePath：第二个参数表示要解析的接口路径URI，一般由开发者定义。
- handlers：第三个参数表示处理对应的请求的函数。

```go
// Handle registers a new request handle and middleware with the given path and method.
// The last handler should be the real handler, the other ones should be middleware that can and should be shared among different routes.
// See the example code in GitHub.
//
// For GET, POST, PUT, PATCH and DELETE requests the respective shortcut
// functions can be used.
//
// This function is intended for bulk loading and to allow the usage of less
// frequently used, non-standardized or custom methods (e.g. for internal
// communication with a proxy).
func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) IRoutes {
	if matched := regEnLetter.MatchString(httpMethod); !matched {
		panic("http method " + httpMethod + " is not valid")
	}
	return group.handle(httpMethod, relativePath, handlers)
}
```

处理GET请求示例代码：

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/*
		使用Handle()方法来处理http请求
			第一个参数：表示处理的http请求的类型
			第二个参数：表示解析的接口路径URI
			第三个参数：表示处理对应的请求的函数
	*/

	// Handle()通用处理请求处理get请求
	// 创建默认engine
	engine := gin.Default()
	// Handle方法处理get请求
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		// 获取请求的uri路径【解析请求的uri路径】
		path := context.FullPath()
		fmt.Println(path) // 输出结果：/hello
		// 获取到请求拼接在uri上的参数，例如http://127.0.0.1:8080?name=zhangsan
		// 获取的参数不为nil就返回获取到的数据，否则就返回，参数二设置的defaultValue()的默认值
		query := context.DefaultQuery("name", "zhangsan")
		// 响应请求
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    "hello " + query,
		})
	})
	// 启动engine
	engine.Run(":8000")
}
```

使用postman发起get请求，具体配置如下：

<img src="asstes/image-20221211171628272.png" alt="image-20221211171628272" style="zoom:50%;" />

要是没有携带参数就使用默认的参数，要是携带参数就使用携带的参数。

<img src="asstes/image-20221211171723509.png" alt="image-20221211171723509" style="zoom:50%;" />

处理POST请求示例代码：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/*
		使用Handle()方法来处理http请求
					第一个参数：表示处理的http请求的类型
					第二个参数：表示解析的接口路径URI
					第三个参数：表示处理对应的请求的函数
	*/
	// 创建默认engine
	engine := gin.Default()
	// Handle方法处理get请求
	engine.Handle("POST", "/hello", func(context *gin.Context) {
		// PostForm()可以解析POST请求Body携带的参数
		name := context.PostForm("name")
		password := context.PostForm("password")
		if name == "zhangsan" && password == "123456" {
			// 响应请求
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "success",
				"data":    name + "登录成功!",
			})
		}
	})
	// 启动engine
	engine.Run(":8000")
}
```

使用postman发生post请求，具体配置如下：

<img src="asstes/image-20221211224320606.png" alt="image-20221211224320606" style="zoom:50%;" />

上述案例，通过第一个参数指定解析POST的请求类型，第二个参数指定解析URI接口为`/hello`，POST请求是以form-data【表单】的方式提交数据的，可以通过`context.PostForm()`来获取表单中提交的数据内容，其他类型的http请求也可以通过`engine.Handle()`的方式来处理。

### 4.4 http请求处理分类

Engine实例除了有通用处理请求的Handle()方法外，还有按照http请求类型分类的方法，可以直接按照http请求的类型直接解析，例如Engine实例中的GET()、POST()、PUT()、DELETE()等与http请求类型对应的方法。

GET请求代码示例：

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/*
		使用Handle()方法来处理http请求
			第一个参数：表示处理的http请求的类型
			第二个参数：表示解析的接口路径URI
			第三个参数：表示处理对应的请求的函数
	*/

	// Handle()通用处理请求处理get请求
	// 创建默认engine
	engine := gin.Default()
	// Handle方法处理get请求
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		// 获取请求的uri路径【解析请求的uri路径】
		path := context.FullPath()
		fmt.Println(path) // 输出结果：/hello
		// 获取到请求拼接在uri上的参数，例如http://127.0.0.1:8080?name=zhangsan
		// 获取的参数不为nil就返回获取到的数据，否则就返回，参数二设置的defaultValue()的默认值
		query := context.DefaultQuery("name", "zhangsan")
		// 响应请求
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    "hello " + query,
		})
	})
	// 启动engine
	engine.Run(":8000")
}
```

POST请求代码示例：

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/*
		post请求
	*/
	// 创建engine实例
	engine := gin.Default()
	// 处理post请求
	engine.POST("/hello", func(context *gin.Context) {
		// PostForm方法是用于解析post请求form-data里面的参数，form表单
		UserName := context.PostForm("UserName")
		Password := context.PostForm("Password")
		fmt.Println(UserName)
		fmt.Println(Password)
		// 响应处理
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    UserName + "登录成功",
		})
	})
	// 启动engine
	engine.Run(":8000")
}
```

### 4.5 请求参数解析

http://127.0.0.1:8000/student/params这种类型的参数，可以使用Context.Param()来解析在获得。

例如：http://127.0.0.1:8000/student/zhangsan

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/*
		gin参数解析
		get请求解析的url：http://127.0.0.1:8000/student/zhangsan
	*/

	// 创建engine实例
	engine := gin.Default()
	// 处理get请求
	engine.GET("/student/:name", func(context *gin.Context) {
		name := context.Param("name")
		fmt.Println(name)
		// 响应请求
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "get success",
			"data":    "姓名=" + name,
		})
	})
	// 处理post请求
	engine.POST("/student/:name", func(context *gin.Context) {
		name := context.Param("name")
		fmt.Println(name)
		// 响应请求
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "post success",
			"data":    "姓名=" + name,
		})
	})
	// 启动engine
	engine.Run(":8000")
}
```

postman发起get请求，具体配置如下：

<img src="asstes/image-20221212194745532.png" alt="image-20221212194745532" style="zoom:50%;" />

postman发起post请求，具体配置如下：

<img src="asstes/image-20221212194813794.png" alt="image-20221212194813794" style="zoom:50%;" />

### 4.6 解析URL上的参数

获取URL参数通过`context.DefaultQuery()`或者`context.Query()`来获取URL上的参数。

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建engine实例
	engine := gin.Default()
	// get请求
	engine.GET("/student", func(context *gin.Context) {
		// DefaultQuery()
		name := context.DefaultQuery("name", "zhangsan")
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "get success",
			"data":    name + "登录成功!",
		})
	})
	// post请求
	engine.POST("/student", func(context *gin.Context) {
		// Query()
		name := context.Query("name")
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "post success",
			"data":    name + "登录成功!",
		})
	})
	engine.Run(":8000")
}
```

postman发起get请求，具体配置如下：

<img src="asstes/image-20221212214046390.png" alt="image-20221212214046390" style="zoom:50%;" />

<img src="asstes/image-20221212214111850.png" alt="image-20221212214111850" style="zoom:50%;" />

postman发起post请求，具体配置如下：

<img src="asstes/image-20221212214625191.png" alt="image-20221212214625191" style="zoom:50%;" />

### 4.7 表单参数

表单传输为post请求，常见的传输格式为4种

- application/json
- application/x-www-form-urlencoded
- application/xml
- multipart/form-data

Gin可以通过Context.

