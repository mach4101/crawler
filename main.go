package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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

	fmt.Println(string(body))

	// fmt.Println("body:", string(body))
	numLinks := strings.Count(string(body), "<a>")

	exist := strings.Contains(string(body), "疫情")

	fmt.Printf("there are %d links!\n", numLinks)

	if exist {
		fmt.Println("there are exist covid-19")
	} else {
		fmt.Println("there are not exist covid-19")
	}
}
