package util

import "C"
import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var (
	Uptime   int64 // 系统启动时间戳
	scClkTck = int64(C.sysconf(C._SC_CLK_TCK))
)

func init() {
	sys := syscall.Sysinfo_t{}
	syscall.Sysinfo(&sys)
	Uptime = time.Now().Unix() - sys.Uptime
}

func ProcessStartTime(pid int) (ts time.Time) {
	buf, err := ioutil.ReadFile(fmt.Sprintf("/proc/%v/stat", pid))
	if err != nil {
		return time.Unix(0, 0)
	}
	if fields := strings.Fields(string(buf)); len(fields) > 22 {
		start, err := strconv.ParseInt(fields[21], 10, 0)
		if err == nil {
			if scClkTck > 0 {
				return time.Unix(Uptime+(start/scClkTck), 0)
			}
			return time.Unix(Uptime+(start/100), 0)
		}
	}
	return time.Unix(0, 0)
}
