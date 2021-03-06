package fetcher

import (
	"bufio"
	"fmt"
	"github.com/qiuye2015/go_dev/crawler/config"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var rateLimiter = time.Tick(time.Second / config.Qps)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	log.Printf("Fetching url %s", url)
	//方法一
	//resp, err := http.Get(url)
	//方法二
	client := &http.Client{}
	newUrl := strings.Replace(url, "http://", "https://", 1)
	request, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Mobile Safari/537.36")
	request.Header.Add("cookie", "bid=fm4PKdjgiNA; __yadk_uid=r7M2Qh2hWEw2BANNW6FBOItyG3RWkSut; douban-fav-remind=1; __gads=ID=f9fd7e6257f29723-22bdbd6763c400a1:T=1603517537:RT=1603517537:S=ALNI_MYgq6XWJ3HrnxZsDPcET7E_L4HoLg; __utmv=30149280.17427; douban-profile-remind=1; __utmz=30149280.1607432424.8.4.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); __utmc=30149280; dbcl2=\"174274438:f0bZfxVp4u8\"; ck=Jsi7; push_noty_num=0; push_doumail_num=0; _pk_ref.100001.8cb4=%5B%22%22%2C%22%22%2C1608121073%2C%22https%3A%2F%2Fwww.google.com.hk%2F%22%5D; _pk_ses.100001.8cb4=*; ap_v=0,6.0; __utma=30149280.1204411877.1602129025.1608106627.1608121076.13; _pk_id.100001.8cb4=6f862408e0e29613.1602129024.13.1608121585.1608106716.; __utmb=30149280.12.6.1608121586136")
	request.Header.Add("Referer", newUrl)
	//处理返回结果
	resp, err := client.Do(request)
	//方法结束

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		//fmt.Println("Error: status code", resp.StatusCode)
		return nil, fmt.Errorf("wrong status code:%d", resp.StatusCode)
	}
	//检测html编码&转化为UTF-8
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error:%v\n", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
