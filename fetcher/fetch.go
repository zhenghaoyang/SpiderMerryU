package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)
//处理URL,返回文本
func Fetch(url string) ([]byte, error) {
	<-rateLimiter //让每个worker 来抢
	//获取网页信息
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//在次判断请求码是否正确
	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code :%d", resp.StatusCode)
	}
	////乱码处理，讲GBK流转成UTF8,不具通用性
	//utf8Buffer := transform.NewReader(resp.Body,	simplifiedchinese.GBK.NewDecoder())

	//reader包装成bufio,避免readerPeek数据缺失
	buffer := bufio.NewReader(resp.Body)
	//检测编码类型
	encodeType := determineEncoding(buffer)
	utf8Buffer := transform.NewReader(buffer, encodeType.NewDecoder())
	//读取
	content, err := ioutil.ReadAll(utf8Buffer)
	if err != nil {
		return nil, err
	}
	//[]byte类型用s打印,自动转string
	//fmt.Printf("%s\n", content)
	return content, nil
}

//检测编码函数 放回编码类型
func determineEncoding(
	r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetcher err = ", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
