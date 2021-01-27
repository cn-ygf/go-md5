package md5

import (
	"testing"
)

func TestComputeFile(t *testing.T) {
	r, err := ComputeFile("/Volumes/Data/软件/deepin-desktop-community-1002-amd64.iso")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(r)
}
