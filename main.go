package main

import "flag"

var logfile string
func main() {
	flag.StringVar(&logfile,"logfile","/var/log/test.log","miner日志路径")
	flag.Parse()
	Tail_f()
}
