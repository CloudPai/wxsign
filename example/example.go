package main

import (
	"fmt"
	"wxsign"

	//"github.com/CloudPai/wxsign"
	"gopkg.in/redis.v3"
)

func init() {
	// 初始化缓存access_token及ticket的redis
	rdsClient := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	wxsign.WxSignRdsInit(rdsClient)
}

func main() {
	ws := wxsign.New(
		//"appid",
		//"secret",
		//// 缓存access_token使用的redis key
		//"wxsign:token",
		//// 缓存ticket使用的redis key
		//"wxsign:ticket",
		"",
		//公众号秘钥
		"",
		// 缓存access_token使用的redis key
		"wxsign:token",
		// 缓存ticket使用的redis key
		"wxsign:ticket",
	)
	sign, err := ws.GetJsSign("http://cloudpai.nat300.top/v1/api")
	if err != nil {
		fmt.Print("Get js sign err-> %#v", err)
		return
	}
	fmt.Print("Js Sign: %#v", sign)
}
