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
    "strconv"
)
	var repo = flag.String("r", "", "仓库名  /2dust/v2rayN/")
	var localPath = flag.String("p", "保存路径", "  /www/wwwroot/download.gojw.xyz/")
	var indexs = flag.String("n", "0", "默认0，releases中第几个链接序号，从0开始，如下载第1个和第6个 0,5  ")
	var remvStrs = flag.String("remove", "", "删除文件关键词 如 clash  ")
func main() {
     flag.Parse()
	// 自动文件下载，比如自动下载图片、压缩包
	var rep=*repo
	var localPathstr=*localPath
	//先删除特殊关键字文件 v2rayNG
	files, _ := ioutil.ReadDir(localPathstr)
	var remvStr =[...]string{*remvStrs}
	
    for _, f := range files {
         for _,r := range remvStr {
              if  r!="" && strings.Contains(f.Name(),r) {
                   os.Remove(localPathstr+f.Name())
              }
         }
    }

	
	
    r, err := http.Get("https://api.github.com/repos"+rep+"releases/latest")
	if err != nil {
		panic(err)
	}
	defer func() {_ = r.Body.Close()}()
	body, _ := ioutil.ReadAll(r.Body)
	
	var xxm mybody
	err = json.Unmarshal(body, &xxm)
    var indexs=*indexs
    indexArr :=	strings.Split(indexs,",")
	for _,i := range indexArr {
	    //字符串转int
	   	i,_ := strconv.Atoi(i)
    	url := xxm.Assets[i].BrowserDownloadURL
    	filenameArr := strings.Split(url,"/")
        filename := filenameArr[len(filenameArr)-1]
     	fmt.Println("正在下载---------"+filename)
         DownloadFileProgress(url,localPathstr+filename)
	}

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


