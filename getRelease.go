package main
    /* 
        下载github 最新 releases 文件到服务器
    */
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
var (
         repo        string
         localPath   string
         indexs      string
         remvStrs    ArrayValue
         newName     string
         filename string
         xxm mybody
        )   

func main() {
 
    flag.StringVar(&repo, "r", "", "仓库名  /2dust/v2rayN/")
    flag.StringVar(&localPath, "p", "", "保存路径  /www/wwwroot/download.gojw.xyz/")
    flag.StringVar(&indexs, "n", "0", "releases中第几个链接序号，默认0，从0开始，如下载第1个和第6个 0,5 ")
    flag.Var(&remvStrs, "remove", "删除文件关键词 逗号分隔 如 clash,v2rayNG  ")
    flag.StringVar(&newName, "name", "", "重命名 名字 ,默认为空，原文件名")    
    flag.Parse()
	// 自动文件下载，比如自动下载图片、压缩包

	//先删除特殊关键字文件 v2rayNG
	files, _ := ioutil.ReadDir(localPath)
	 
    for _, f := range files {
         for _,r := range remvStrs {
              if  r!="" && strings.Contains(f.Name(),r) {
                   fmt.Println("删除关键词"+r+"文件")
                   os.Remove(localPath+f.Name())
              }
         }
    }

	
	
    r, err := http.Get("https://api.github.com/repos"+repo+"releases/latest")
	if err != nil {
		panic(err)
	}
	defer func() {_ = r.Body.Close()}()
	body, _ := ioutil.ReadAll(r.Body)
	
	err = json.Unmarshal(body, &xxm)
    indexArr :=	strings.Split(indexs,",")

	for _,i := range indexArr {
	    //字符串转int
	   	i,_ := strconv.Atoi(i)
    	url := xxm.Assets[i].BrowserDownloadURL
    	if newName != ""{
    	    filename = newName
    	}else {
    	    filename =xxm.Assets[i].Name
    	}
        
     	fmt.Println("正在下载---------"+filename)
        DownloadFileProgress(url,localPath+filename)
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

// 重写 flag的 Set  String接口 
func (a *ArrayValue) Set(s string) error {
    *a = strings.Split(s, ",")
    return nil
}
func (s *ArrayValue) String() string {
    return fmt.Sprintf("%v", *s)
}

type mybody struct {
	Assets []Assets `json:"assets"`
}


type Assets struct {
	BrowserDownloadURL string `json:"browser_download_url"`
	Name string `json:"name"`
}

// 自定义类型实现命令行接收数组参数
type ArrayValue []string


