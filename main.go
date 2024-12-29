package main

import (
	"flag"
	"fmt"
	"runtime"
)

// /  创建全局变量
var (
	version  = "0.0.1"
	codename = "l2go"
	///  用于解析输入参数
	f = flag.String("f", "client.json", "config file name")
)

// /  答应版本信息
func printVersion() {
	fmt.Printf("name: %v \nversion: %v\n", codename, version)
	fmt.Printf("runtime: %v %v %v\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

/*
	https://go.dev/blog/laws-of-reflection
*/

// /
type Config struct {
	Local  string `json:"local"`
	Route  string `json:"route"`
	Remote string `json:"remote"`
}

// /
/* func loadConfig(configFileName string) (*Config, error) {

} */

type Shape interface{}

func main() {
	//
	printVersion()
	//
	flag.Parse()

	var a Config = Config{"1", "2", "3"}

	b := Config{"2", "2", "2"}

	var c Shape = 11

	fmt.Println(a.Local, b.Local, c)

}
