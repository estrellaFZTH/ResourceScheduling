package scheduling

import (
	"ResourceScheduling/cmd/collector"
	"ResourceScheduling/global"
	"ResourceScheduling/internal/scale"
	"log"
	"time"
)

func NormalSchedule() {
	log.Printf("NormalSchedule...")
	t := time.Tick(global.ScalingInterval)
	var lastScaleOut = time.Now().Unix()
	for {
		<-t
		// res[0] cpuUsage
		// res[1] cpuLimit
		// res[2] cpuUsagePercentage
		// res[3] podCount
		// res[4] avgCpuUsage
		log.Printf("NormalSchedule ::: GetCpuUsageAndPodCount")
		metrics := collector.GetCpuUsageAndPodCount()
		podCount := metrics[3]
		avgCpuUsage := metrics[4]

		if avgCpuUsage >= global.MaxThreshold && podCount <= global.MaxScalePodCount {
			var now = time.Now().Unix()
			// 距离上一次伸缩时间大于3分钟才进行伸缩
			if now-lastScaleOut > 180 {
				lastScaleOut = now
				scale.ScaleOut(int(podCount)+1, global.NameSpace, global.ClusterName)
			}

		} else if avgCpuUsage <= global.MinThreshold && podCount > global.MinScalePodCount {
			var now = time.Now().Unix()
			if now-lastScaleOut > 180 {
				lastScaleOut = now
				scale.ScaleIn(int(podCount)-1, global.NameSpace, global.ClusterName)
			}
		}
	}
}
