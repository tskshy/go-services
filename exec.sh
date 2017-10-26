#!/usr/bin/env bash

APPNAME=$2

function init() {
	if [ "$APPNAME" == "" ]
	then
		APPNAME=app_abc_def
	fi
}

function usage() {
	echo "Usage: $0 [--build | --start | --stop | --restart | --clean | --show] [optional(app name)]"
}

function build() {
	go build -o $APPNAME src/app/main/main.go
}

function clean() {
	rm -rf ./$APPNAME
}

function start() {
#nohup ./$APPNAME 1>/dev/null 2>&1 &
	./$APPNAME --log="/root/devel/golang/go-services/src/resources/log.json" --conf="/root/devel/golang/go-services/src/resources/app.json"
}

function stop() {
	kill -9 `ps -ef| grep "$APPNAME" |grep -v "grep" |grep -v "$0" |awk '{print $2}'`
}

function restart() {
	stop
	start
}

function show() {
	ps -ef| grep "$APPNAME" |grep -v "grep" |grep -v "$0"
}

init
#### main
case "$1" in
	--build)
		build
		;;
	--clean)
		clean
		;;
	--start)
		start
		;;
	--stop)
		stop
		;;
	--restart)
		restart
		;;
	--show)
		show
		;;
	*)
		usage
		;;
esac
