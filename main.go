package main

import (
	"ResourceScheduling/cmd/collector"
	"ResourceScheduling/cmd/scheduling"
)

func main() {
	// 指标收集
	go collector.CollectMetrics()
	// 常规资源调度
	scheduling.NormalSchedule()
	// 智能资源调度
	//scheduling.SmartSchedule()
}
