#!/bin/bash

# This is the key script file for the standard application template,
# and it's best not to modify its content to avoid affecting normal usage later on.

INTERVAL=1
PROJECT_HOME="$(dirname "$0")"
cd ${PROJECT_HOME}

# Main Process Name
PROCESS_NAME="svapp"


function start() {
    local pid=$(ps -ef | grep -v grep | grep "${PROCESS_NAME}" | awk '{print $2}' |head -n 1)

    if [[ -n ${pid} ]]; then
        echo "has started at pid : ${pid}"
        exit 0
    fi

    # a command to make process running background
    # todo here just a sample, change it in production
    nohup "./${PROCESS_NAME}" 1 >/dev/null 2>&1 &

    sleep $INTERVAL

    pid=$(ps -ef | grep -v grep | grep "${PROCESS_NAME}" | awk '{print $2}')

    if [[ -z "${pid}" ]]; then
      while((count <= 10)); do
          pid=$(ps -ef | grep -v grep | grep "${PROCESS_NAME}" | awk '{print $2}')
          if [[ -n "${pid}" ]]; then
              echo "${PROCESS_NAME} started with pid : ${pid}"
          fi
          ((count++))
          sleep $INTERVAL
      done
      echo "${PROCESS_NAME} start failed"
    fi
    echo "${pid}"
}

function stop() {
    local count=0
    local pids=$(ps -ef | grep -v grep | grep "${PROCESS_NAME}" | awk '{print $2}')

    for pid in ${pids}
    do
      echo "${pid}"

      if [[ -n ${pid} ]]; then
          #send term singe
          kill ${pid}

          sleep 3

          while kill -0 ${pid} > /dev/null 2>&1; do
              if ((count <= 5)); then
                  kill -9 ${pid}
                  echo "After 5 seconds, force kill ${pid}"
                  break
              fi
              ((count++))
              sleep 1
          done
      fi
    done
    echo "stopped"
}

function main() {
    local param=""
    while (($# > 0)); do
        param=$1
        case ${param} in
           "start")
               start
               exit $?
           ;;
           "stop")
               stop
               exit $?
           ;;
           "restart")
               stop
               start
               exit $?
           ;;
           *)
             echo "Usage $0 <start|stop|restart>"
             exit 1
           ;;
        esac
    done
}

if (($# <= 0 || $# > 2)); then
    echo "Usage $0 <start|stop|restart>"
    exit 1
fi

main "$@"

exit 0
