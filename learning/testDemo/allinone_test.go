package testDemo

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

// 被测函数直接写在 test 文件里
func GetUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("-----失败------")
		fmt.Println(err)
		return
	}
	fmt.Println("-----成功------")
	fmt.Println(resp)
}

func GetUrlWithRetry(url string, retry int) {
	timeout := time.Duration(retry) * time.Minute
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("-----失败------")
			fmt.Println(err)
			time.Sleep(time.Second * 5)
			continue
		}
		fmt.Println("-----成功------")
		fmt.Println(resp)
		return
	}
}

func TestGetUrl(t *testing.T) {
	// if Add(1, 2) != 3 {
	// 	t.Fatal("fail")
	// }
	GetUrl("http://www.baidu.com")
	GetUrl("http://www.adhfansfdasf")
	GetUrlWithRetry("http://www.adhfansfdasf", 1)
}
