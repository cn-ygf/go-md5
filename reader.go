package md5

import (
	"io"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

// 创建带进度条功能的Reader
// r 原始读io
// size 总数据大小
func NewReader(r io.Reader, size int64) Reader {
	return &reader{
		io:   r,
		size: size,
		b:    NewBar(0, size, "#"),
	}
}

type reader struct {
	io   io.Reader
	size int64
	b    Bar
	cur  int64
}

func (r *reader) Read(p []byte) (n int, err error) {
	n, err = r.io.Read(p)
	r.cur = r.cur + int64(n)
	r.b.Play(r.cur)
	return
}
