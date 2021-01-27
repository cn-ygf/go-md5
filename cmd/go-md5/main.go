package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"go-md5.io/md5"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	if os.Args[1] == "--version" {
		version()
		return
	}
	filePathName := os.Args[1]
	st, err := os.Stat(filePathName)
	if err != nil {
		fmt.Printf("获取文件或路径信息失败！err:%s\n", err.Error())
		return
	}
	// 如果是路径
	if st.IsDir() {
		fi, err := ioutil.ReadDir(filePathName)
		if err != nil {
			fmt.Printf("获取文件或路径信息失败！err:%s\n", err.Error())
			return
		}
		for _, st := range fi {
			if !st.IsDir() {
				fileName := fmt.Sprintf("%s/%s", filePathName, st.Name())
				result, err := md5.ComputeFile(fileName)
				if err == nil {
					Save(result, st.Name())
				}
			}
		}
	} else {
		result, err := md5.ComputeFile(filePathName)
		if err == nil {
			Save(result, st.Name())
		}
	}
}

func usage() {
	version()
	fmt.Println("open source:http://github.com/cn-ygf/go-md5\n")
	fmt.Println("使用方式:	go-md5\t[参数]\t<输出目录>")
	fmt.Println("\t文件路径\t计算单个文件md5")
	fmt.Println("\t目录路径\t计算指定目录所有文件md5")
	fmt.Println("\t--version\t打印程序版本号")
	fmt.Println("示例:")
	fmt.Println("\tgo-md5 d:\\go-md5.txt\t计算单个文件")
	fmt.Println("\tgo-md5 d:\\go-md5\t计算指定目录所有文件")
	fmt.Println("\tgo-md5 .\t\t计算当前目录所有文件md5")
	fmt.Println("\tgo-md5 d:\\go-md5 d:\\result\t保存结果到指定目录(默认当前目录)")
}

func version() {
	fmt.Println("go-md5 v1.0.1")
}

func Save(md5 string, fileName string) {
	if len(os.Args) > 2 {
		fileName = fmt.Sprintf("%s/%s", os.Args[2], fileName)
	}
	fileName = fmt.Sprintf("%s.md5", fileName)
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	f.WriteString(md5)
}
