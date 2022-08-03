package hostinfo

import (
	"testing"
)

func TestHostInfo(t *testing.T) {
	t.Log(GetHost())
}

func TestDiskInfo(t *testing.T) {
	t.Log(GetDiskInfo("d:/"))
}

func TestDiskInfo1(t *testing.T) {
	t.Log(GetPartInfo("c:"))
}

func TestDiskParts(t *testing.T) {
	t.Log(GetDiskParts())
}
