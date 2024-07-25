package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func Robots(mentionAll bool, data string) {
	// mentionAllStr := "@xxx"

	// if mentionAll {
	// 	mentionAllStr = "@all"
	// }
	dataJsonStr := fmt.Sprintf(`{
		"msgtype": "text",
		"text": {
		  "content":"` + jsonToString(data) + `",
		  "mentioned_list":["@all"],
		}
	  }`)
	// fmt.Println(dataJsonStr)
	resp, err := http.Post(
		"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxx",
		"application/json",
		bytes.NewBuffer([]byte(dataJsonStr)))
	if err != nil {
		// fmt.Println("weworkAlarm request error")
		return
	}
	defer resp.Body.Close()
	// webAlive()
}
func RobotsCookie(mentionAll bool, data string) {
	// mentionAllStr := "@xxx"

	// if mentionAll {
	// 	mentionAllStr = "@all"
	// }
	dataJsonStr := fmt.Sprintf(`{
		"msgtype": "text",
		"text": {
		  "content":"` + data + `",
		  "mentioned_list":["@all"],
		}
	  }`)
	fmt.Println(dataJsonStr)
	resp, err := http.Post(
		"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxxx",
		"application/json",
		bytes.NewBuffer([]byte(dataJsonStr)))
	if err != nil {
		// fmt.Println("weworkAlarm request error")
		return
	}
	defer resp.Body.Close()
	// webAlive()
}

// func timenow() string {
// 	return time.Now().Format("2006-01-02 15:04:05")
// }
