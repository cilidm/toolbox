package hostinfo

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
)

// GetHost 主机信息
func GetHost() *host.InfoStat {
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Println("host.Info() failed: ", err)
		return nil
	}
	return hostInfo
}

type DiskStatus struct {
	Size uint64
	Used uint64
	Free uint64
}

func GetDiskInfo(dir string) *DiskStatus {
	info, err := disk.Usage(dir)
	if err != nil {
		return nil
	}
	return &DiskStatus{
		Size: info.Total,
		Used: info.Used,
		Free: info.Free,
	}
}

func GetDiskParts() ([]disk.PartitionStat, error) {
	return disk.Partitions(true)
}

func GetPartInfo(path string) (*disk.UsageStat, error) {
	return disk.Usage(path)
}
