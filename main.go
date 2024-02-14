package main

import "fmt"

func main() {
    // 读取数据库配置
    config := ReadConfig("config.json")
    fmt.Println("[info]加载配置:")
    fmt.Println(config)
    // 连接数据库
    db := 
    fmt.Println("[info]连接数据库:")
    fmt.Println(db)
}