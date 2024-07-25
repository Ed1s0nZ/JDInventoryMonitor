package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type RequestConfig struct {
	URL         string            `json:"url"`
	QueryParams string            `json:"queryParams"`
	Headers     map[string]string `json:"headers"`
}

func readConfig(filePath string) (*RequestConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &RequestConfig{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

func main() {
	// 每分钟读取一次 req.json
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		config, err := readConfig("req.json")
		if err != nil {
			fmt.Println("Error reading config:", err)
			return
		}

		client := &http.Client{}
		req, err := http.NewRequest("GET", config.URL+config.QueryParams, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		// 设置 Headers
		for key, value := range config.Headers {
			req.Header.Set(key, value)
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()

		var reader io.ReadCloser
		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			reader, err = gzip.NewReader(resp.Body)
			if err != nil {
				fmt.Println("Error creating gzip reader:", err)
				return
			}
			defer reader.Close()
		default:
			reader = resp.Body
		}

		body, err := io.ReadAll(reader)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		// 转换 GBK 到 UTF-8
		utf8Body, _, err := transform.String(simplifiedchinese.GBK.NewDecoder(), string(body))
		if err != nil {
			fmt.Println("Error converting from GBK to UTF-8:", err)
			return
		}

		fmt.Println(finder(utf8Body))
		//如果发现任意一个产品不是无货状态，则立即返回 false。如果所有产品都是无货状态，则返回 true。
		StockStateNameBool, err := checkAllOutOfStock(finder(utf8Body))
		if err != nil {
			fmt.Println(err)
		}
		if strings.Contains(utf8Body, "10086628263328") && !StockStateNameBool {
			Robots(false, finder(utf8Body))
			fmt.Println(finder(utf8Body))

		} else if !strings.Contains(utf8Body, "10086628263328") {
			fmt.Println("cookie失效，请替换cookie")
			RobotsCookie(false, "京东cookie已失效，请替换cookie")
			time.Sleep(1 * time.Hour)
		}

		// 等待下一个时间间隔
		<-ticker.C
	}
}
