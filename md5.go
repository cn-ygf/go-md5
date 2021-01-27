package md5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"time"
)

// 计算文件md5
// fileName 文件名称
func ComputeFile(fileName string) (result string, err error) {
	var (
		f     *os.File
		st    os.FileInfo
		begin time.Time
		// end   time.Time
	)
	begin = time.Now()
	fmt.Printf("md5计算中 %s\n", fileName)
	// fmt.Printf("起始时间: %s\n", begin.Format("2006-01-02 15:04:05"))
	f, err = os.Open(fileName)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	defer f.Close()
	st, err = f.Stat()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	r := NewReader(f, st.Size())
	h := md5.New()
	_, err = io.Copy(h, r)
	fmt.Println()
	if err != nil {
		fmt.Printf("出现错误: %s\n", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println(err.Error())
		return
	}
	result = hex.EncodeToString(h.Sum(nil))
	//	end = time.Now()
	fmt.Println("MD5:", result)
	// fmt.Printf("结束时间: %s\n", end.Format("2006-01-02 15:04:05"))
	fmt.Printf("耗时: %s\n", time.Since(begin))
	return
}
