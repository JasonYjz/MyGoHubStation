import os
import time

'''
cpu  1050008 127632 898432 43828767 37203 63 99244 0 0 0
cpu0 212383 20476 204704 8389202 7253 42 12597 0 0 0
cpu1 224452 24947 215570 8372502 8135 4 42768 0 0 0
cpu2 222993 17440 200925 8424262 8069 9 17732 0 0 0
cpu3 186835 8775 195974 8486330 5746 3 8360 0 0 0
cpu4 107075 32886 48854 8688521 3995 4 5758 0 0 0
cpu5 90733 20914 27798 1429573 2984 1 11419 0 0 0

user（62124）：从系统启动开始累计到当前时刻，用户态的CPU时间，不包含 nice值为负进程。
nice（11）：从系统启动开始累计到当前时刻。
system（47890）：从系统启动开始累计到当前时刻，nice值为负的进程所占用的CPU时间。
idle（8715270）：从系统启动开始累计到当前时刻，除硬盘IO等待时间以外其它等待时间。
iowait（84729）：从系统启动开始累计到当前时刻，硬盘IO等待时间。
irq（0）：从系统启动开始累计到当前时刻，硬中断时间。
softirq（1483）：从系统启动开始累计到当前时刻，软中断时间。

CPU时间=user+nice+system+idle+iowait+irq+softirq。
CPU利用率=1-(idle2-idle1)/(cpu2-cpu1)*100。
'''


def get_cpu_stat():
    f = open("/proc/stat", "r")
    stat_text = f.readlines()
    cpu_info = {}
    for i in stat_text:
        if i.startswith('cpu'):
            cpu_info[i.split(' ')[0]] = list(map(int, i.split()[1:8]))
    return cpu_info


def get_cpu_use(cpu_stat_old: dict, cpu_stat: dict) -> dict:
    cpu_use_info = {}
    for key, value in cpu_stat.items():
        cpu_old = sum(cpu_stat_old[key])
        cpu = sum(value)
        idle_old = int(cpu_stat_old[key][3])
        idle = int(value[3])
        if (cpu - cpu_old) != 0:
            usage = 10 - int((idle - idle_old) / (cpu - cpu_old) * 10)
        else:
            usage = 0
        cpu_use_info[key] = usage
    return cpu_use_info

cpu_info_old = get_cpu_stat()
cpu_info = get_cpu_stat()
soc_info['cpu_use'] = get_cpu_use(cpu_info_old, cpu_info)

{'cpu': 1, 'cpu0': 1, 'cpu1': 3, 'cpu2': 1, 'cpu3': 2, 'cpu4': 0, 'cpu5': 0, 'cpu6': 0, 'cpu7': 0}

# cat /sys/kernel/gpu/