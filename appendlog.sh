#!/bin/sh

LOGFILE="/var/opt/aidlux/ai/log/muslin/test.log"
RANDOM="ABCDEFGHIGKLMNOPQRSTUVWXYZabcdefghigklmnopqrstuvwxyz0123456789"
# 获取随机字符串
# shellcheck disable=SC2112
#function random_str() {
#    str=$(echo $RANDOM | md5sum | cut -c 1-10)
##    for i in {1..16} ; do
##        str="$str$(echo $RANDOM | md5sum | cut -c 1-10)"
##    done
#    echo $str
#}

for i in {1..1000} ; do
    timestamp=$(date +"%Y%m%d-%H%M%S")
    str=$(echo "$RANDOM" | md5sum | cut -c 1-10)
    echo "$i $timestamp $str" >> "$LOGFILE"
    sleep 1
done