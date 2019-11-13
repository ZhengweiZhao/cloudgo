package service
//使用martini框架，具体定义main.go文件中启动server后的具体操作
import (
    "github.com/go-martini/martini" 
)
func NewServer(port string) {   
    m := martini.Classic()//新建一个经典的martini实例
    
    // 用户自定义路由规则
    m.Get("/", func(params martini.Params) string {
        return "hello world by Zhao Zhengwei!"  //接收对\的GET方法请求，第二个参数是对一请求的处理方法
    })

    m.RunOnAddr(":"+port)  //运行服务器 
}

