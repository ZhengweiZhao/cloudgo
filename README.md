# Sevice Computing：开发 web 服务程序Cloudgo

# 1、概述
开发简单 web 服务程序 cloudgo，了解 web 服务器工作原理。

**任务目标**

 1. 熟悉 go 服务器工作原理 
 2. 基于现有 web 库，编写一个简单 web 应用类似 cloudgo。 
 3. 使用 curl 工具访问 web程序 
 4. 对 web 执行压力测试

**基本要求**

1. 编程 web 服务程序 类似 cloudgo 应用。
- 要求有详细的注释
- 是否使用框架、选哪个框架自己决定 请在README.md 说明你决策的依据
2. 使用 curl 测试，将测试结果写入 README.md
3. 使用 ab 测试，将测试结果写入 README.md。并解释重要参数。

# 2、基础知识
关于本次作业的web服务程序开发，主要是学习了潘老师的博客：
 
 [HTTP 协议 与 golang web 应用服务](https://blog.csdn.net/pmlpml/article/details/78404838)
 
 **HTTP协议基础**
 HTTP 协议是一个复杂的协议， 支持虚拟主机、消息路由（负载均衡）、分段下载、缓存服务、安全认证等等。 HTTP 也是非常简单文本协议。 客户端与服务器建立 TCP 连接后，客户端发出 Request 文本， 服务器端返回 Response 文本。
 
**HTTP协议工具**

- **浏览器**

几乎所有现代浏览器都自带开发者工具，Network 标签就是。

- **curl**
curl 才是 web 开发者最常用的利器。它是一个控制台程序，可以精确控制 HTTP 请求的每一个细节。实战中，配合 shell 程序，我们可以简单，重复给服务器发送不同的请求序列，调试程序或分析输出。curl 是 linux 系统自带的命令行工具。

	curl的命令如下：
	```
	$curl -v http://localhost:9000/
	```
	注 ： 后面的网址根据自己要访问的网址更改即可。
	
	输入curl命令后，每行出现的第一个符号分别有如下含义：

- `*`表示 curl 任务；
- `>`发送的信息;
- `<`返回的信息

# 3、开发实践
## 3.1 框架选择
所有的Web框架都是基于net/http包构建的。

本次使用的web开发框架是Martini，Martini框架是使用Go语言作为开发语言的一个强力的快速构建模块化web应用与服务的开发框架。使用 Go 的 net/http 接口开发，类似 Sinatra 或者 Flask 之类的框架，也可使用自己的 DB 层、会话管理和模板。这个框架在GitHub上都有中文的解释以及用法，比较容易上手。

Martini支持URL参数 通配符和正则表达式，路由功能全面灵活。

Martini匹配的参数可通过 map[string]string 获得，它将被注入到处理器的方法参数，Martini提供完整的依赖注入到处理器方法的参数，允许指定全局或请求级别的映射，Martini在数据绑定上非常省事。

Martini没有明显的控制器或上下文概念，但是依赖注入允许你易于创建你自己这样的概念。

**其特性如下：**

> - 使用非常简单
> - 无侵入设计
> - 可与其他 Go 的包配合工作
> - 超棒的路径匹配和路由
> - 模块化设计，可轻松添加工具
> - 大量很好的处理器和中间件
> - 很棒的开箱即用特性
> - 完全兼容 http.HandlerFunc 接口

**安装**
在terminal下输入命令行：
```
go get github.com/codegangsta/martini
```

## 3.2 搭建简单web服务器
main.go的代码基本上就是按照老师教程上的main函数写的

**main.go**
```javascript
package main
import (
    "os"
    "cloudgo/service"
    flag "github.com/spf13/pflag"
)
const (
    PORT string = "8080"  //默认8080端口
)
func main() {
    //默认8080端口，如果没有监听到端口，则设为默认端口。
    port := os.Getenv("PORT") 
    if len(port) == 0 {
        port = PORT
    }
    //对命令行参数进行设置，用-p设置端口，并完成对端口号的解析
    pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
    flag.Parse()
    if len(*pPort) != 0 {
        port = *pPort
    }
    //启动server
    service.NewServer(port)
}


```

**server.go**
```javascript
package service
//使用martini框架，具体定义main.go文件中启动server后的具体操作
import (
    "github.com/go-martini/martini" 
)
func NewServer(port string) {   
    m := martini.Classic()//新建一个经典的martini实例
    
    // 用户自定义路由规则
    m.Get("/", func(params martini.Params) string {
        return "hello world by ZhaoZhengwei!"  //接收对\的GET方法请求，第二个参数是对一请求的处理方法
    })

    m.RunOnAddr(":"+port)  //运行服务器 
}
```
## 3.3 测试运行
完成代码后用 go run 运行它， 这个时候其实已经在9000端口监听http链接请求了。
运行命令：
```
$ go run main.go -p 9000
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113085529245.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
在网页输入网址：
```
http://localhost:9000
```
可以看到我们在server.go中写的启动server后的具体操作，也就是会打印出一行欢迎提示。
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113090024341.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
### 1） 使用 curl 测试
打开另一个控制台，用 curl 命令测试查看结果。

测试用命令：
```
curl -v http://localhost:9000
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113090142659.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
### 2）使用 ab 测试，并解释重要参数。
**ab是什么？**

>- ab的全称是Apache Bench，是Apache自带的网络压力测试工具，相比于LR、JMeter，是我所知道的 Http 压力测试工具中最简单、最通用的。
>- ab命令对发出负载的计算机要求很低，不会占用很高CPU和内存，但也能给目标服务器产生巨大的负载，能实现基础的压力测试。
>- 在进行压力测试时，最好与服务器使用交换机直连，以获取最大的网络吞吐量。

**压力测试程序安装：**

首先，需要通过yum命令安装Apache web 压力测试程序
```
yum -y install httpd-tools
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113090205441.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
可以看到如下图所示的界面，则说明安装成功了：
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113090151841.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
**执行压力测试：**

```
$ ab -n 1000 -c 100 http://localhost:9000/cloudgo
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113090427124.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
其中：

- Document Path 表示请求的资源也就是我们运行的url的代码包的地址，这里是cloudgo
- Document Length 是文档返回的长度，不包括相应头。
- Concurrency Level 是并发个数
- Complete requests 是总请求数，Time taken for tests 显示了总请求时间，下面的per ...则是平均到每秒／每个请求等等。
-  Transfer rate 表示传输速率
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113090444604.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)

	对压力测试的结果重点关注**吞吐率**（Requests per second）、**用户平均请求等待时间**（Time per request）指标：
	
	**1、吞吐率（Requests per second）：**
	
	服务器并发处理能力的量化描述，单位是reqs/s，指的是在某个并发用户数下单位时间内处理的请求数。某个并发用户数下单位时间内能处理的最大请求数，称之为最大吞吐率。吞吐率是基于并发用户数的，这表示：
	
		a、吞吐率和并发用户数相关
		
		b、不同的并发用户数下，吞吐率一般是不同的
	
	计算公式：总请求数/处理完成这些请求数所花费的时间，即
	
	Request per second=Complete requests/Time taken for tests
	
	必须要说明的是，这个数值表示当前机器的整体性能，值越大越好。
	
	**2、用户平均请求等待时间（Time per request）：**
	
	计算公式：处理完成所有请求数所花费的时间/（总请求数/并发用户数），即：
	
	Time per request=Time taken for tests/（Complete requests/Concurrency Level）
	
	**3、服务器平均请求等待时间（Time per request:across all concurrent requests）：**
	
	计算公式：处理完成所有请求数所花费的时间/总请求数，即：
	
	Time taken for/testsComplete requests
	
	可以看到，它是吞吐率的倒数。同时，它也等于用户平均请求等待时间/并发用户数，即
	
	Time per request/Concurrency Level。
	

**ab命令选项：**

ab命令具有如下图所示的参数选项，我们上面的命令用的是最基本的参数-n 和 -c，下面将对重要参数进行一个解释说明：
```
-n 执行的请求数量
-c 并发请求个数
-t 测试所进行的最大秒数
-p 包含了需要POST的数据的文件
-T POST数据所使用的Content-type头信息
-k 启用HTTP KeepAlive功能，即在一个HTTP会话中执行多个请求，默认时，不启用KeepAlive功能
```
全部完整参数含义如下，可以根据自己的需要自行查找和翻译：
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113090529354.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191113090600427.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MDM3NzY5MQ==,size_16,color_FFFFFF,t_70)
