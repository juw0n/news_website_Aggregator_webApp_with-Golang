package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sitemapindex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://news.sky.com/sitemap-index.xml")
	//resp, _ := http.Get("https://news.sky.com/robots.txt")
	bytes, _ := ioutil.ReadAll(resp.Body)
	/*stringBody := string(bytes)
	fmt.Println(stringBody) //The output here should be a string of the source code of the website you put in the http.Get()
	resp.Body.Close()*/
	var s Sitemapindex
	xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)

	for _, Location := range s.Locations {
		fmt.Printf("%s\n", Location)
	}
}
