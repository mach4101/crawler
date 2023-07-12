package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(err)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func DeterminEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		fmt.Printf("fetch error: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func main() {
	url := "https://www.thepaper.cn/"

	body, err := Fetch(url)
	if err != nil {
		fmt.Printf("read content fail %v", err)
		return
	}

	// fmt.Println(string(body))

	// fmt.Println("body:", string(body))
	// numLinks := strings.Count(string(body), "<a>")
	//
	//	exist := strings.Contains(string(body), "疫情")
	//
	//	fmt.Printf("there are %d links!\n", numLinks)
	//
	//	if exist {
	//		fmt.Println("there are exist covid-19")
	//	} else {
	//		fmt.Println("there are not exist covid-19")
	//	}

	// headerRE := regexp.MustCompile(
	//
	//	`<li class=".*?" .*?><a target="_blank" href="/newsDetail_forward_.*?" class="index_inherit__A1ImK">(.*?)</a></li>`,
	//
	// )
	//
	// matches := headerRE.FindAllSubmatch(body, -1)
	//
	//	for _, m := range matches {
	//		if !strings.HasPrefix(string(m[1]), "<img") {
	//			fmt.Println("fetch card news:", string(m[1]))
	//		}
	//	}

	// XPATH parse

	doc, err := htmlquery.Parse(bytes.NewReader(body))
	if err != nil {
		fmt.Printf("htmlquery.Parse fail: %v", err)
	}

	nodes := htmlquery.Find(
		doc,
		`//*[@id="__next"]/main/div[2]/div[3]/div[2]/div[1]/div[1]/div/div[3]/div[2]/ul/li/a[@target="_blank"]`,
	)

	for _, node := range nodes {
		fmt.Println("fetch card ", node.FirstChild.Data)
	}
}
