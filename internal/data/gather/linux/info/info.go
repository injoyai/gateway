package info

import (
	"github.com/injoyai/goutil/oss/shell"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"strings"
	"time"
)

func GetCPUInfo() CPUInfo {
	var cpuInfo CPUInfo
	totalPercent, _ := cpu.Percent(0, false)
	if len(totalPercent) == 1 {
		cpuInfo.CPUAvgPercent = totalPercent[0]
		cpuInfo.CPUUsed = cpuInfo.CPUAvgPercent * 0.01 * float64(cpuInfo.CPUNum)
	}
	cpuInfo.CPUPercent, _ = cpu.Percent(0, true)
	return cpuInfo
}

func GetMemoryInfo() MemoryInfo {
	info, _ := mem.VirtualMemory()
	return MemoryInfo{
		Total:   info.Total,
		Used:    info.Used,
		Unused:  info.Available,
		Percent: info.UsedPercent,
	}
}

func LoadCurrentInfo(ioOption string, netOption string) *DashboardCurrent {
	var currentInfo DashboardCurrent
	hostInfo, _ := host.Info()
	currentInfo.RunTime = hostInfo.Uptime
	currentInfo.StartTime = time.Now().Add(-time.Duration(hostInfo.Uptime) * time.Second).Format("2006-01-02 15:04:05")
	currentInfo.ProcessNum = hostInfo.Procs

	currentInfo.CPUNum, _ = cpu.Counts(true)
	totalPercent, _ := cpu.Percent(0, false)
	if len(totalPercent) == 1 {
		currentInfo.CPUAvgPercent = totalPercent[0]
		currentInfo.CPUUsed = currentInfo.CPUAvgPercent * 0.01 * float64(currentInfo.CPUNum)
	}
	currentInfo.CPUPercent, _ = cpu.Percent(0, true)

	loadInfo, _ := load.Avg()
	currentInfo.Load1 = loadInfo.Load1
	currentInfo.Load5 = loadInfo.Load5
	currentInfo.Load15 = loadInfo.Load15
	currentInfo.LoadUsagePercent = loadInfo.Load1 / (float64(currentInfo.CPUNum*2) * 0.75) * 100

	memoryInfo, _ := mem.VirtualMemory()
	currentInfo.MemoryTotal = memoryInfo.Total
	currentInfo.MemoryAvailable = memoryInfo.Available
	currentInfo.MemoryUsed = memoryInfo.Used
	currentInfo.MemoryUsedPercent = memoryInfo.UsedPercent

	currentInfo.DiskData = loadDiskInfo()

	if ioOption == "all" {
		diskInfo, _ := disk.IOCounters()
		for _, state := range diskInfo {
			currentInfo.IOReadBytes += state.ReadBytes
			currentInfo.IOWriteBytes += state.WriteBytes
			currentInfo.IOCount += state.ReadCount + state.WriteCount
			currentInfo.IOReadTime += state.ReadTime
			currentInfo.IOWriteTime += state.WriteTime
		}
	} else {
		diskInfo, _ := disk.IOCounters(ioOption)
		for _, state := range diskInfo {
			currentInfo.IOReadBytes += state.ReadBytes
			currentInfo.IOWriteBytes += state.WriteBytes
			currentInfo.IOCount += state.ReadCount + state.WriteCount
			currentInfo.IOReadTime += state.ReadTime
			currentInfo.IOWriteTime += state.WriteTime
		}
	}

	netInfo, _ := net.IOCounters(false)
	if len(netInfo) != 0 {
		currentInfo.NetBytesSent = netInfo[0].BytesSent
		currentInfo.NetBytesRecv = netInfo[0].BytesRecv
	}

	currentInfo.ShotTime = time.Now()
	return &currentInfo
}

func loadDiskInfo() []DiskInfo {
	var datas []DiskInfo
	stdout, err := shell.Exec("df -hT -P|grep '/'|grep -v tmpfs|grep -v 'snap/core'|grep -v udev")
	if err != nil {
		return datas
	}
	lines := strings.Split(stdout, "\n")

	var mounts []diskInfo
	var excludes = []string{"/mnt/cdrom", "/boot", "/boot/efi", "/dev", "/dev/shm", "/run/lock", "/run", "/run/shm", "/run/user"}
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 7 {
			continue
		}
		if fields[1] == "tmpfs" {
			continue
		}
		if strings.Contains(fields[2], "M") || strings.Contains(fields[2], "K") {
			continue
		}
		if strings.Contains(fields[6], "docker") {
			continue
		}
		isExclude := false
		for _, exclude := range excludes {
			if exclude == fields[6] {
				isExclude = true
			}
		}
		if isExclude {
			continue
		}
		mounts = append(mounts, diskInfo{Type: fields[1], Device: fields[0], Mount: fields[6]})
	}

	for i := 0; i < len(mounts); i++ {
		state, err := disk.Usage(mounts[i].Mount)
		if err != nil {
			continue
		}
		var itemData DiskInfo
		itemData.Path = mounts[i].Mount
		itemData.Type = mounts[i].Type
		itemData.Device = mounts[i].Device
		itemData.Total = state.Total
		itemData.Free = state.Free
		itemData.Used = state.Used
		itemData.UsedPercent = state.UsedPercent
		itemData.InodesTotal = state.InodesTotal
		itemData.InodesUsed = state.InodesUsed
		itemData.InodesFree = state.InodesFree
		itemData.InodesUsedPercent = state.InodesUsedPercent
		datas = append(datas, itemData)
	}
	return datas
}

type DiskInfo struct {
	Path        string  `json:"path"`
	Type        string  `json:"type"`
	Device      string  `json:"device"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`

	InodesTotal       uint64  `json:"inodesTotal"`
	InodesUsed        uint64  `json:"inodesUsed"`
	InodesFree        uint64  `json:"inodesFree"`
	InodesUsedPercent float64 `json:"inodesUsedPercent"`
}

type diskInfo struct {
	Type   string
	Mount  string
	Device string
}

type DashboardCurrent struct {
	RunTime    uint64 `json:"runTime"`    //运行时间,秒
	StartTime  string `json:"startTime"`  //开始时间
	ProcessNum uint64 `json:"processNum"` //进程数

	Load1            float64 `json:"load1"`  //近1分钟负载
	Load5            float64 `json:"load5"`  //近5分钟负载
	Load15           float64 `json:"load15"` //近15分钟负载
	LoadUsagePercent float64 `json:"loadUsagePercent"`

	CPUPercent    []float64 `json:"cpuPercent"`    //cpu各核占用百分比
	CPUAvgPercent float64   `json:"cpuAvgPercent"` //cpu平均占用百分比
	CPUUsed       float64   `json:"cpuUsed"`
	CPUNum        int       `json:"cpuNum"` //cpu核数

	MemoryTotal       uint64  `json:"memoryTotal"`
	MemoryAvailable   uint64  `json:"memoryAvailable"`
	MemoryUsed        uint64  `json:"memoryUsed"`
	MemoryUsedPercent float64 `json:"MemoryUsedPercent"`

	IOReadBytes  uint64 `json:"ioReadBytes"`
	IOWriteBytes uint64 `json:"ioWriteBytes"`
	IOCount      uint64 `json:"ioCount"`
	IOReadTime   uint64 `json:"ioReadTime"`
	IOWriteTime  uint64 `json:"ioWriteTime"`

	DiskData []DiskInfo `json:"diskData"`

	NetBytesSent uint64 `json:"netBytesSent"`
	NetBytesRecv uint64 `json:"netBytesRecv"`

	ShotTime time.Time `json:"shotTime"`
}

type CPUInfo struct {
	CPUPercent    []float64 `json:"cpuPercent"`    //cpu各核占用百分比
	CPUAvgPercent float64   `json:"cpuAvgPercent"` //cpu平均占用百分比
	CPUUsed       float64   `json:"cpuUsed"`
	CPUNum        int       `json:"cpuNum"` //cpu核数
}

type MemoryInfo struct {
	Total   uint64  `json:"total"`   //内存总量
	Unused  uint64  `json:"unused"`  //可用内存
	Used    uint64  `json:"used"`    //已使用内存
	Percent float64 `json:"percent"` //内存占用百分比
}
