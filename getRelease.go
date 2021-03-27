package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
    "encoding/json"
    "io/ioutil"
    "strings"
    "flag"
)
	var repo = flag.String("r", "", "仓库名  /2dust/v2rayN/")
	var localPath = flag.String("p", "", "仓库名  /www/wwwroot/download.gojw.xyz/")
func main() {
     flag.Parse()
	// 自动文件下载，比如自动下载图片、压缩包
	var rep=*repo
    r, err := http.Get("https://api.github.com/repos"+rep+"releases/latest")
	if err != nil {
		panic(err)
	}
	defer func() {_ = r.Body.Close()}()
	body, _ := ioutil.ReadAll(r.Body)
	
	var xxm mybody
	err = json.Unmarshal(body, &xxm)
	
	url := xxm.Assets[0].BrowserDownloadURL
	filenameArr := strings.Split(url,"/")
    filename := filenameArr[len(filenameArr)-1]
 	var localPathstr=*localPath
 	DownloadFileProgress(url,localPathstr+filename)
}



func downloadFile(url, filename string) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() {_ = r.Body.Close()}()

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() {_ = f.Close()}()

	n, err := io.Copy(f, r.Body)
	fmt.Println(n, err)
}

type Reader struct {
	io.Reader
	Total int64
	Current int64
}

func (r *Reader) Read(p []byte) (n int, err error){
	n, err = r.Reader.Read(p)

	r.Current += int64(n)
	fmt.Printf("\r进度 %.2f%%", float64(r.Current * 10000/ r.Total)/100)

	return
}

func DownloadFileProgress(url, filename string) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() {_ = r.Body.Close()}()

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() {_ = f.Close()}()

	reader := &Reader{
		Reader: r.Body,
		Total: r.ContentLength,
	}

	_, _ = io.Copy(f, reader)
}


type mybody struct {
	Assets []Assets `json:"assets"`
}


type Assets struct {
	BrowserDownloadURL string `json:"browser_download_url"`
}
