# 脚本

放脚本


### githubDow.go  

下载github指定仓库的 releases

```
./githubDow.go -r /2dust/v2rayN/
仓库名  -r  必填，其他选填
```

```
//完整参数事例
go run ./githubDow.go -r /Kr328/ClashForAndroid/ -p /www/wwwroot/ -name clashForAndroid.apk -remove clash

```

```
//查看帮助
go run ./githubDow.go -h
Usage of /tmp/go-build311794379/b001/exe/githubDow:
  -n string
    	releases中第几个链接序号，默认0，从0开始，如下载第1个和第6个 0,5  (default "0")
  -name string
    	重命名 名字 ,默认为空，原文件名
  -p string
    	保存路径  /www/wwwroot/download.gojw.xyz/
  -r string
    	仓库名  /2dust/v2rayN/
  -remove value
    	删除文件关键词 逗号分隔 如 clash,v2rayNG  
exit status 2


```
