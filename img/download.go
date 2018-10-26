package img

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func Dowmload(imgUrl string) {
	//图片正则
	reg, err := regexp.Compile(`(\w|\d|_)*.jpg`)
	name := reg.FindStringSubmatch(imgUrl)[0]
	if err != nil {
		return
	}
	//fmt.Print(name)
	//通过http请求获取图片的流文件
	resp, _ := http.Get(imgUrl)
	body, _ := ioutil.ReadAll(resp.Body)
	out, _ := os.Create(name)
	io.Copy(out, bytes.NewReader(body))
	fmt.Println("save Img success==")
	return
}
