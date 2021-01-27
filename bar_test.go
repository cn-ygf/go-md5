package md5

import (
	"testing"
	"time"
)

func TestBar(t *testing.T) {
	b := NewBar(0, 100, "#")
	for i := 0; i <= 100; i++ {
		time.Sleep(10 * time.Millisecond)
		b.Play(int64(i))
	}
	b.Finish()
}
