package global

import "time"

const (
	// Prometheus server address & controller server address
	K8sPrometheusAddress  = "http://10.77.110.132:31090"
	TiDBPrometheusAddress = "http://10.77.110.132:19090"

	CollectMetricsInterval = 5 * time.Second
	ScalingInterval        = 15 * time.Second

	NameSpace   = "tidb-cluster"
	ClusterName = "basic"

	MaxScalePodCount = 6
	MinScalePodCount = 3

	MaxThreshold = 80
	MinThreshold = 40
)

//func GetPrometheusAddress() string {
//	return prometheusAddress
//}
