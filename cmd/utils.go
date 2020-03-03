package cmd

import (
	"github.com/imroc/req"
	"os"
	"strings"
)

func makeBarkUrl() {
	var barkServer string
	if temp := os.Getenv("barkServer"); temp == "" {
		barkServer = "https://api.day.app/"
	} else {
		if strings.HasSuffix(temp, "/") {
			barkServer = temp
		} else {
			barkServer = temp + "/"
		}
	}
	barkApiUrl = barkServer + barkApi + "/"
}

var barkApiUrl string

func barkNotify(message string, title string) {
	params := req.Param{
		"title": title,
		"body":  message,
	}
	req.Post(barkApiUrl, params)
}

func barkUrl(url, title string) {
	fullUrl := barkApiUrl + title + "?url=" + url
	req.Get(fullUrl)
}
