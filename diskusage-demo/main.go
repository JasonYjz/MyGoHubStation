package main

import (
	"fmt"
	"io/ioutil"
	"syscall"
)

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

// disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func main() {
	//disk := DiskUsage("/")
	//fmt.Printf("All: %.2f GB\n", float64(disk.All)/float64(GB))
	//fmt.Printf("Used: %.2f GB\n", float64(disk.Used)/float64(GB))
	//fmt.Printf("Free: %.2f GB\n", float64(disk.Free)/float64(GB))
	//
	/////media/sdg1
	//upan := DiskUsage("/media/sdg1")
	//fmt.Printf("All: %.2f GB\n", float64(upan.All)/float64(GB))
	//fmt.Printf("Used: %.2f GB\n", float64(upan.Used)/float64(GB))
	//fmt.Printf("Free: %.2f GB\n", float64(upan.Free)/float64(GB))
	//COMMAND := "df -h |grep -w '/media\\|/' |awk '{print $2,$3,$4,$5,$6}'"

	//timeout, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancelFunc()
	////CommandContext
	////output, err := exec.Command(COMMAND).Output()
	////output, err := exec.Command("df", "-h", "|", "grep", "-w", "/media\\|/'").Output()
	////output, err := exec.Command("/bin/sh", "-c", "df -h |grep -w '/media\\|/' |awk '{print $2,$3,$4,$5,$6}'").Output()
	//output, err := exec.CommandContext(timeout, "/bin/sh", "-c", "df -h |grep -w '/media\\|/' |awk '{print $2,$3,$4,$5,$6}'").Output()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//fmt.Println(string(output))

	fmt.Printf("%v\n", WalkDir("/media"))

}

func WalkDir(dir string) []string {

	var files []string
	fileInfos, _ := ioutil.ReadDir(dir)

	for _, info := range fileInfos {
		if info.IsDir() {

			fmt.Println(info.Name())
			files = append(files, info.Name())
		}
	}

	return files
}
