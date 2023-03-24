package collector

const (
	PodCpuUsage = "sum(rate(container_cpu_usage_seconds_total{pod=~\"%s.*\"}[1m])) by (pod) * 1000"
	PodCpuLimit = "sum by (container) (kube_pod_container_resource_limits{container=\"%s\", resource=\"cpu\"}) * 1000"
	PodCount    = "count(container_cpu_usage_seconds_total{namespace=\"tidb-cluster\",container=\"%s\"})"
	// 这个是以 instance 为单位进行查询，如果有两个tidb的pod，则得到以下结果
	// {instance="basic-tidb-0"} 103.67999999999994 ms
	// {instance="basic-tidb-1"} 58.133333333333375 ms
	//TiDBP99Latency = "histogram_quantile(0.99, sum(rate(tidb_server_handle_query_duration_seconds_bucket{k8s_cluster=\"\", tidb_cluster=~\"tidb-cluster-basic\", instance=~\"%s.*\"}[1m])) by (le, instance)) * 1000"

	// 这个是以集群为单位进行查询 延时单位转化为了ms
	TiDBP99Latency = "histogram_quantile(0.99, sum(rate(tidb_server_handle_query_duration_seconds_bucket{k8s_cluster=\"\", tidb_cluster=~\"tidb-cluster-basic\"}[1m])) by (le)) * 1000"
	StatementOps   = "sum(rate(tidb_executor_statement_total{k8s_cluster=\"\", tidb_cluster=~\"tidb-cluster-basic\"}[1m])) by (type)"
	//PodCpuUsagePercentage     = "sum(rate(container_cpu_usage_seconds_total{pod=~\"%s.*\"}[1m])) by (pod)/sum(cluster:namespace:pod_cpu:active:kube_pod_container_resource_requests) by (pod) *100"
	//PodMemoryUsage            = "sum(container_memory_rss{pod=~\"%s.*\"}) by(pod)" // /1024/1024/1024 = GiB
	//PodMemoryUsagePercentage  = "sum(container_memory_rss{pod=~\"%s-.*\"}) by(pod)/sum(cluster:namespace:pod_memory:active:kube_pod_container_resource_requests) by (pod) *100"
	//ClusterCpuUsagePercentage = "sum(rate(container_cpu_usage_seconds_total{pod=~\"%s-(masters|replicas)-.*\"}[1m]))/sum(cluster:namespace:pod_cpu:active:kube_pod_container_resource_requests{pod=~\"%s-(masters|replicas)-.*\"})*100"
	//MasterCpuUsagePercentage  = "sum(rate(container_cpu_usage_seconds_total{pod=~\"%s-(masters)-.*\"}[1m]))/sum(cluster:namespace:pod_cpu:active:kube_pod_container_resource_requests{pod=~\"%s-(masters)-.*\"})*100"
	//// WorkerCpuUsagePercentage	 = "sum(rate(container_cpu_usage_seconds_total{pod=~\"%s-(replicas)-.*\", image!=\"\", container!=\"POD\"}[1m]))  / sum(cluster:namespace:pod_cpu:active:kube_pod_container_resource_limits{pod=~\"%s-(replicas)-.*\"})*100"
	//WorkerCpuUsagePercentage     = "sum(rate(container_cpu_usage_seconds_total{pod=~\"%s-(replicas)-.*\", image!=\"\", container!=\"POD\"}[1m])) / (sum(container_spec_cpu_quota{image!=\"\", pod=~\"%s-replicas-.*\", node=\"worker211\"})/100000) * 100"
	//ClusterMemoryUsagePercentage = "sum(container_memory_rss{pod=~\"%s-(masters|replicas)-.*\"})/sum(cluster:namespace:pod_memory:active:kube_pod_container_resource_requests{pod=~\"%s-(masters|replicas)-.*\"})*100"
	//ClusterNum                   = "sum(kube_pod_container_info{pod=~\"%s-(masters|replicas)-.\"})"
	//ReplicaMidCount              = "sum(kube_statefulset_replicas{statefulset=~\"%s-replicas-mid\"})"
	//ReplicaMidTwoCpuUsage        = "sum(rate(container_cpu_usage_seconds_total{pod=~\"%s-replicas-mid-(%s|%s).*\", image!=\"\", container!=\"POD\"}[1m]))  / sum(cluster:namespace:pod_cpu:active:kube_pod_container_resource_limits{pod=~\"%s-replicas-mid-(%s|%s).*\"})*100"
	//ReplicaSmallCpuUsage         = "sum(rate(container_cpu_usage_seconds_total{pod=~\"%s-replicas-small-.*\", image!=\"\", container!=\"POD\"}[1m]))  / sum(cluster:namespace:pod_cpu:active:kube_pod_container_resource_limits{pod=~\"%s-replicas-small-.*\"})*100"
	//ReplicaMidCpuUsage           = "sum(rate(container_cpu_usage_seconds_total{pod=~\"%s-replicas-mid-.*\", image!=\"\", container!=\"POD\"}[1m]))  / sum(cluster:namespace:pod_cpu:active:kube_pod_container_resource_limits{pod=~\"%s-replicas-mid-.*\"})*100"
)

//const (
//	//PodCpuUsage        = "sum(rate(container_cpu_usage_seconds_total{pod=~\"%s.*\"}[1m])) by (pod)*100"
//	//PodMemoryUsage     = "sum(container_memory_rss{pod=~\"%s.*\"}) by(pod)"
//	ReplicaAvgCpuUsage = "avg(rate(container_cpu_usage_seconds_total{namespace=\"polardb\",container=~\"%s.*\"}[1m]))*100"
//	QueryReplicaCount  = "count(container_cpu_usage_seconds_total{namespace=\"polardb\",container=\"%s\"})"
//)
