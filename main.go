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

