package main

import (
	"bufio"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type person struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type students struct {
	id  string                 `json:"id"`
	per map[string]interface{} `json:"per"`
}

func main() {

	//getCpuInfo()
	//getCpuLoad()
	//getDiskInfo()
	//getHostInfo()
	//getMemInfo()
	//getNetInfo()
	getBootTime()
	fmt.Println("#########################")
	//getDiskUsage(cmd.C.Path)

}

// cpu info
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, true)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}

func getCpuLoad() {
	info, _ := load.Avg()
	fmt.Printf("%v\n", info)
}

func getMemInfo() {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("mem info:%v\n", memInfo)
}

func getHostInfo() {
	fmt.Printf("start to get host info.")
	hInfo, _ := host.Info()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}

func getBootTime() {
	timestamp, _ := host.BootTime()
	t := time.Unix(int64(timestamp), 0)
	fmt.Println(t.Local().Format("2006-01-02 15:04:05"))

	sys := syscall.Sysinfo_t{}
	syscall.Sysinfo(&sys)
	Uptime := time.Now().Unix() - sys.Uptime

	t = time.Unix(int64(Uptime), 0)
	fmt.Println(t.Local().Format("2006-01-02 15:04:05"))

	t = ProcessUptime()
	fmt.Println(t.Local().Format("2006-01-02 15:04:05"))

}

func ProcessUptime() (ts time.Time) {
	lines, err := ReadLinesOffsetN("/proc/uptime", 0, -1)
	if err != nil {
		fmt.Println("failed to read lines offset n")
		return
	}

	if len(lines) != 1 {
		fmt.Println("wrong uptime format")
		return
	}
	f := strings.Fields(lines[0])
	b, err := strconv.ParseFloat(f[0], 64)
	if err != nil {
		fmt.Println("parse float failed")
		return
	}
	timestamp := uint64(time.Now().Unix()) - uint64(b)
	return time.Unix(int64(timestamp), 0)
}

func ReadLinesOffsetN(filename string, offset uint, n int) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{""}, err
	}
	defer f.Close()

	var ret []string

	r := bufio.NewReader(f)
	for i := 0; i < n+int(offset) || n < 0; i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if i < int(offset) {
			continue
		}
		ret = append(ret, strings.Trim(line, "\n"))
	}

	return ret, nil
}

// disk info
func getDiskInfo() {
	fmt.Printf("start to calculate Partitions.")
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}
	for _, part := range parts {
		fmt.Printf("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}
	fmt.Printf("start to calculate IOCounters.")
	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("%v:%v\n", k, v)
	}
}

func getDiskUsage(path string) {
	usage, err := disk.Usage(path)
	if err != nil {
		fmt.Printf("error disk usage:%s\n", err.Error())
		return
	}

	fmt.Printf("disk usage:%s\n", usage.String())
}

func getNetInfo() {
	info, _ := net.IOCounters(true)
	for index, v := range info {
		fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	}
}

func printTmp() {
	var p = make([]person, 0)

	fmt.Printf("variable address:%p\n", p)
	fmt.Printf("len(p) is %d\n", len(p))
	//fmt.Println(fmt.Sprintf("variable address:%v", p))
	//fmt.Println(fmt.Sprintf("len(p) is %d", len(p)))

	p = make([]person, 1)
	fmt.Println(fmt.Sprintf("variable address:%p", p))
	fmt.Println(fmt.Sprintf("len(p) is %d", len(p)))

	var h []person
	fmt.Println(fmt.Sprintf("variable address:%p", h))
	fmt.Println(fmt.Sprintf("len(h) is %d", len(h)))

	h = make([]person, 1)
	fmt.Println(fmt.Sprintf("variable address:%p", h))
	fmt.Println(fmt.Sprintf("len(h) is %d", len(h)))

	stu := &students{}
	fmt.Printf("variable address:%p\n", stu)
	fmt.Printf("in variable per address:%p\n", stu.per)

	mp := make(map[string]interface{})
	mp["test"] = "abc"

	stu1 := &students{
		id:  "001",
		per: mp,
	}
	fmt.Printf("variable address:%p\n", stu1)
	fmt.Printf("in variable per address:%p\n", stu1.per)
}
