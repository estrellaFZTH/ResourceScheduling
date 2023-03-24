package scheduling

import (
	"ResourceScheduling/cmd/collector"
	"ResourceScheduling/global"
	"ResourceScheduling/internal/model"
	"ResourceScheduling/internal/scale"
	"log"
	"time"
)

func SmartSchedule() {
	log.Printf("SmartSchedule...")
	t := time.Tick(global.ScalingInterval)
	var lastScaleOut = time.Now().Unix()
	for {
		<-t
		// res[0] cpuUsage
		// res[1] cpuLimit
		// res[2] cpuUsagePercentage
		// res[3] podCount
		// res[4] avgCpuUsage
		log.Printf("SmartSchedule ::: GetCpuUsageAndPodCount")
		k8sMetrics := collector.GetCpuUsageAndPodCount()
		podCount := k8sMetrics[3]
		avgCpuUsage := k8sMetrics[4]

		// res[0]  statementOpsResult ==== TPS
		// res[1]  p99Latency
		log.Printf("SmartSchedule ::: GetTpSAndLatency")
		tidbMetrics := collector.GetTpSAndLatency()
		tps := tidbMetrics[0]
		p99Latency := tidbMetrics[1]
		log.Printf("tps: %s, p99Latency: %s.", tps, p99Latency)

		//	Tps            float32
		//	P99Latency     float32
		//	AvgCpuUsage    float32
		//	AvgCpuGradient float32
		//	PodCount       float32
		//	TiDBScaleTime  float32
		var inPut = model.InPut{}
		inPut.Tps = tps
		inPut.P99Latency = p99Latency
		inPut.AvgCpuUsage = avgCpuUsage
		inPut.PodCount = podCount
		inPut.TiDBScaleTime = float32(global.TiDBScaleTime)

		// (暂定)Input [tps, p99Latency, avgCpuUsage, avgCpuGradient, podCount, TiDBScaleTime]
		// AI Model
		outPut := model.AIModel(inPut)
		// Output [max_threshold, min_threshold, ScaleOutPodCount, ScaleInPodCount]
		log.Printf("outPut: %s", outPut)

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
