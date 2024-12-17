package main

import (
	"flag"
	"fmt"
	"net/http"
)

const (
	YD = "%40fzcmcc"  // 移动
	DX = "@fztelecom" // 电信
	LT = "@fzunicom"  // 联通
)

// ParseStation 解析运营商
func ParseStation(station string) string {
	switch station {
	case "移动":
		return YD
	case "电信":
		return DX
	case "联通":
		return LT
	default:
		return YD
	}
}

// NotNull 判断命令行参数是否为空
func NotNull(args ...*string) {
	for _, arg := range args {
		if *arg == "" {
			fmt.Println("参数不能为空")
			flag.PrintDefaults()
			return
		}
	}
}

func main() {

	username := flag.String("u", "", "用户名")
	password := flag.String("p", "", "密码")
	station := ParseStation(*flag.String("s", "移动", "运营商"))

	flag.Parse()

	NotNull(username, password)

	url := "http://172.30.255.105:801/eportal/?c=Portal&a=login&callback=dr1&login_method=1&user_account=" + *username + station + "&user_password=" + *password + "&wlan_user_ip=&wlan_user_mac=000000000000&wlan_ac_ip=&wlan_ac_name=&jsVersion=3.0&_=1"

	fmt.Println("请求URL:", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求错误:", err)
	}

	content := make([]byte, 1024)
	resp.Body.Read(content)
	fmt.Println("状态码:", resp.StatusCode)
	fmt.Println("内容:", string(content))
	defer resp.Body.Close()

	fmt.Println("按下Ctrl + C退出")

	select {}
}
