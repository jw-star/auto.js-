# -*- coding:utf8 -*- 
import urllib 
import urllib2 
import json
import os
import sys
#描述：下载github 最新release的文件
#调用方法： python dow.py repoName localPath
#可以添加定时任务执行脚本
#repoName='/2dust/v2rayN/'
repoName=sys.argv[1]
# localPath='/www/wwwroot/download.gojw.xyz'

localPath=sys.argv[2]

def Schedule(a,b,c):
    per = 100.0 * a * b / c
    if per > 100 :
        per = 100
    print '%.2f%%' % per

html = urllib.urlopen('https://api.github.com/repos'+repoName+'releases/latest')
hjson = json.loads(html.read())
v2coreUrl = hjson['assets'][0]['browser_download_url']
print(v2coreUrl)

LocalPath = os.path.join(localPath, v2coreUrl.strip().split('/')[-1])
urllib.urlretrieve(v2coreUrl,LocalPath,Schedule)
