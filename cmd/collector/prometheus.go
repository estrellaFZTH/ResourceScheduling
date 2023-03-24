package collector

import (
	"ResourceScheduling/global"
	"ResourceScheduling/internal/collector"
	"ResourceScheduling/internal/utils"
	"github.com/prometheus/common/model"
	"log"
	"time"
)

func CollectMetrics() {
	_, queryClientK8s, err := collector.GetPrometheusClient(global.K8sPrometheusAddress)
	if err != nil {
		log.Fatalf("Cannot connect to K8sPrometheus: %s, %s", global.K8sPrometheusAddress, err.Error())
		return
	}
	log.Printf("Success to connect K8sPrometheus\n")

	_, queryClientTiDB, err := collector.GetPrometheusClient(global.TiDBPrometheusAddress)
	if err != nil {
		log.Fatalf("Cannot connect to TiDBPrometheus: %s, %s", global.TiDBPrometheusAddress, err.Error())
		return
	}
	log.Printf("Success to connect TiDBPrometheus\n")

	t := time.Tick(global.CollectMetricsInterval)
	start := 0
	var resArr []collector.Result
	for {
		<-t
		// sum(rate(container_cpu_usage_seconds_total{pod=~"basic-tikv.*"}[1m])) by (pod)
		// basic-tikv-0  basic-tikv-1
		cpuUsageResult, err := collector.QueryPodCpuUsage("basic-tikv", queryClientK8s)
		if err != nil {
			log.Fatalf("Cannot query cpusage prometheus: %s, %s", global.K8sPrometheusAddress, err.Error())
		}
		// sum by (container) (kube_pod_container_resource_limits{container="tikv", resource="cpu"}) * 1000
		cpuLimitResult, err := collector.QueryPodCpuLimit("tikv", queryClientK8s)
		if err != nil {
			log.Fatalf("Cannot query cpusage prometheus: %s, %s", global.K8sPrometheusAddress, err.Error())
		}
		// count(container_cpu_usage_seconds_total{namespace="tidb-cluster",container="tikv"})
		podCountResult, err := collector.QueryPodCountSelect("tikv", queryClientK8s)
		if err != nil {
			log.Fatalf("Cannot query pod from prometheus: %s, %s", global.K8sPrometheusAddress, err.Error())
		}
		cpuUsageArray := collector.ExtractValue(&cpuUsageResult)
		temp, err := utils.Sum(cpuUsageArray)
		if err != nil {
			log.Fatalf("ParseFloat cpuUsageArray failed")
		}
		cpuUsage := float32(temp)
		log.Printf("TiDB Cluster Current Total CPU Usage(m): %f\n", cpuUsage)

		cpuLimitArray := collector.ExtractValue(&cpuLimitResult)
		temp, err = utils.Sum(cpuLimitArray)
		if err != nil {
			log.Fatalf("ParseFloat cpuLimitArray failed")
		}
		cpuLimit := float32(temp)
		log.Printf("TiDB Cluster Current Total CPU Limit(m): %f\n", cpuLimit)

		cpuUsagePercentage := cpuUsage / cpuLimit
		log.Printf("TiDB Cluster Current Total CPU Usage Percentage: %f\n", cpuUsagePercentage)

		podCountArray := collector.ExtractValue(&podCountResult)
		temp, err = utils.Sum(podCountArray)
		if err != nil {
			log.Fatalf("ParseFloat podCountArray failed")
		}
		podCount := float32(temp)
		log.Printf("TiDB Cluster pod Count : %f\n", podCount)

		var avgCpuUsage = (cpuUsagePercentage / podCount)
		log.Printf("AvgCpuUsage: %f", avgCpuUsage)

		//=================================================
		// statementOpsResult ==== TPS
		statementOpsResult, err := collector.QueryStatementOps("basic-tidb", queryClientTiDB)
		if err != nil {
			log.Fatalf("Cannot query cpusage prometheus: %s, %s", global.K8sPrometheusAddress, err.Error())
		}
		statementOpsArray := collector.ExtractValue(&statementOpsResult)
		temp, err = utils.Sum(statementOpsArray)
		if err != nil {
			log.Fatalf("ParseFloat statementOpsArray failed")
		}
		statementOps := float32(temp)                                   // 单位是 k个
		log.Printf("TiDB Cluster Current TPS: %f\n", statementOps*1000) // 乘 1000 将单位转为个

		p99LatencyResult, err := collector.QueryTiDBP99Latency("basic-tidb", queryClientTiDB)
		if err != nil {
			log.Fatalf("Cannot query cpusage prometheus: %s, %s", global.K8sPrometheusAddress, err.Error())
		}
		p99LatencyArray := collector.ExtractValue(&p99LatencyResult)
		temp, err = utils.Sum(p99LatencyArray)
		if err != nil {
			log.Fatalf("ParseFloat p99LatencyArray failed")
		}
		p99Latency := float32(temp)
		log.Printf("TiDB Cluster Current p99Latency(ms): %f\n", p99Latency)

		var r collector.Result
		r.Time = int64(start)
		r.CpuUsage = cpuUsage
		r.CpuLimit = cpuLimit
		r.CpuUsagePercentage = cpuUsagePercentage
		r.PodCount = int64(podCount)
		r.AvgCpu = avgCpuUsage
		r.StatementOps = statementOps
		r.P99Latency = p99Latency
		resArr = append(resArr, r)

		start += 5
		if start == 5*3000 {
			collector.WriteOutputs(resArr)
			break
		}

	}
}

// res[0] cpuUsage
// res[1] cpuLimit
// res[2] cpuUsagePercentage
// res[3] podCount
// res[4] avgCpuUsage
func GetCpuUsageAndPodCount() []float32 {
	var res []float32
	_, queryClientK8s, err := collector.GetPrometheusClient(global.K8sPrometheusAddress)
	if err != nil {
		log.Fatalf("Cannot connect to K8sPrometheus: %s, %s", global.K8sPrometheusAddress, err.Error())
		return nil
	}
	// sum(rate(container_cpu_usage_seconds_total{pod=~"basic-tikv.*"}[1m])) by (pod)
	// basic-tikv-0  basic-tikv-1
	cpuUsageResult, err := collector.QueryPodCpuUsage("basic-tikv", queryClientK8s)
	if err != nil {
		log.Fatalf("Cannot query cpusage prometheus: %s, %s", global.K8sPrometheusAddress, err.Error())
	}
	// sum by (container) (kube_pod_container_resource_limits{container="tikv", resource="cpu"}) * 1000
	cpuLimitResult, err := collector.QueryPodCpuLimit("tikv", queryClientK8s)
	if err != nil {
		log.Fatalf("Cannot query cpusage prometheus: %s, %s", global.K8sPrometheusAddress, err.Error())
	}
	// count(container_cpu_usage_seconds_total{namespace="tidb-cluster",container="tikv"})
	podCountResult, err := collector.QueryPodCountSelect("tikv", queryClientK8s)
	if err != nil {
		log.Fatalf("Cannot query pod from prometheus: %s, %s", global.K8sPrometheusAddress, err.Error())
	}
	cpuUsageArray := collector.ExtractValue(&cpuUsageResult)
	temp, err := utils.Sum(cpuUsageArray)
	if err != nil {
		log.Fatalf("ParseFloat cpuUsageArray failed")
	}
	cpuUsage := float32(temp)
	//log.Printf("TiDB Cluster Current Total CPU Usage(m): %f\n", cpuUsage)
	res = append(res, cpuUsage)

	cpuLimitArray := collector.ExtractValue(&cpuLimitResult)
	temp, err = utils.Sum(cpuLimitArray)
	if err != nil {
		log.Fatalf("ParseFloat cpuLimitArray failed")
	}
	cpuLimit := float32(temp)
	//log.Printf("TiDB Cluster Current Total CPU Limit(m): %f\n", cpuLimit)
	res = append(res, cpuLimit)

	cpuUsagePercentage := cpuUsage / cpuLimit
	//log.Printf("TiDB Cluster Current Total CPU Usage Percentage: %f\n", cpuUsagePercentage)
	res = append(res, cpuUsagePercentage)

	podCountArray := collector.ExtractValue(&podCountResult)
	temp, err = utils.Sum(podCountArray)
	if err != nil {
		log.Fatalf("ParseFloat podCountArray failed")
	}
	podCount := float32(temp)
	//log.Printf("TiDB Cluster pod Count : %f\n", podCount)
	res = append(res, podCount)

	var avgCpuUsage = (cpuUsagePercentage / podCount)
	//log.Printf("AvgCpuUsage: %f", avgCpuUsage)
	res = append(res, avgCpuUsage)
	return res

}

// res[0]  statementOpsResult ==== TPS
// res[1]  p99Latency
func GetTpSAndLatency() []float32 {
	var res []float32
	_, queryClientTiDB, err := collector.GetPrometheusClient(global.TiDBPrometheusAddress)
	if err != nil {
		log.Fatalf("Cannot connect to TiDBPrometheus: %s, %s", global.TiDBPrometheusAddress, err.Error())
		return nil
	}
	// statementOpsResult ==== TPS
	statementOpsResult, err := collector.QueryStatementOps("basic-tidb", queryClientTiDB)
	if err != nil {
		log.Fatalf("Cannot query cpusage prometheus: %s, %s", global.K8sPrometheusAddress, err.Error())
	}
	statementOpsArray := collector.ExtractValue(&statementOpsResult)
	temp, err := utils.Sum(statementOpsArray)
	if err != nil {
		log.Fatalf("ParseFloat statementOpsArray failed")
	}
	statementOps := float32(temp) // 单位是 k个
	//log.Printf("TiDB Cluster Current TPS: %f\n", statementOps*1000) // 乘 1000 将单位转为个
	res = append(res, statementOps)

	p99LatencyResult, err := collector.QueryTiDBP99Latency("basic-tidb", queryClientTiDB)
	if err != nil {
		log.Fatalf("Cannot query cpusage prometheus: %s, %s", global.K8sPrometheusAddress, err.Error())
	}
	p99LatencyArray := collector.ExtractValue(&p99LatencyResult)
	temp, err = utils.Sum(p99LatencyArray)
	if err != nil {
		log.Fatalf("ParseFloat p99LatencyArray failed")
	}
	p99Latency := float32(temp)
	//log.Printf("TiDB Cluster Current p99Latency(ms): %f\n", p99Latency)
	res = append(res, p99Latency)

	return res
}

func GetP99LatencyInFiveMin() model.Value {
	_, queryClientTiDB, err := collector.GetPrometheusClient(global.TiDBPrometheusAddress)
	if err != nil {
		log.Fatalf("Cannot connect to TiDBPrometheus: %s, %s", global.TiDBPrometheusAddress, err.Error())
		return nil
	}

	result, _ := collector.QueryP99LatencyInFiveMin("basic-tidb", queryClientTiDB)
	return result
}
