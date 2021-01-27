package md5

import (
	"fmt"
)

type Bar interface {
	Play(cur int64)
	Finish()
}

// 创建进度条
// start 起始位置
// total 总大小
func NewBar(start, total int64, graph string) Bar {
	b := &bar{
		cur:   start,
		total: total,
		graph: graph,
	}
	b.percent = b.getPercent()
	for i := 0; i < int(b.percent); i += 2 {
		b.rate += b.graph //初始化进度条位置
	}
	return b
}

type bar struct {
	percent int64  // 百分比
	cur     int64  // 当前进度位置
	total   int64  // 总进度
	rate    string // 进度条
	graph   string // 显示符号
}

func (b *bar) getPercent() int64 {
	return int64(float32(b.cur) / float32(b.total) * 100)
}

func (b *bar) Play(cur int64) {
	b.cur = cur
	last := b.percent
	b.percent = b.getPercent()
	if b.percent != last && b.percent%2 == 0 {
		b.rate += b.graph
	}
	fmt.Printf("\r[%-50s]%3d%%  %8d/%d", b.rate, b.percent, b.cur, b.total)

}

func (b *bar) Finish() {
	fmt.Println()
}
