package main

import (
	"fmt"

	"github.com/mach4101/crawler/collect"
)

func main() {
	url := "https://book.douban.com/subject/1007305"

	var f collect.Fetcher = collect.BrowserFetch{}
	body, err := f.Get(url)
	if err != nil {
		fmt.Printf("there are exist error %v\n", err)
		return
	}

	fmt.Println(string(body))
}
