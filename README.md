# go-md5 一个go语言写的md5计算工具

## 下载编译
```shell
git clone https://github.com/cn-ygf/go-md5
cd go-md5/cmd
# 编译当前目标平台
go build
# 编译所有平台
chmod +x build.sh
./build.sh
```

## 使用说明
```shell
go-md5 v1.0.1
open source:http://github.com/cn-ygf/go-md5

使用方式:	go-md5	[参数]	<输出目录>
	文件路径	计算单个文件md5
	目录路径	计算指定目录所有文件md5
	--version	打印程序版本号
示例:
	go-md5 d:\go-md5.txt	计算单个文件
	go-md5 d:\go-md5	计算指定目录所有文件
	go-md5 .		计算当前目录所有文件
	go-md5 d:\go-md5 d:\result	保存结果到指定目录(默认当前目录)
 ```
