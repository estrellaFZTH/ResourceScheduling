package model

type InPut struct {
	Tps            float32
	P99Latency     float32
	AvgCpuUsage    float32
	AvgCpuGradient float32
	PodCount       float32
	TiDBScaleTime  float32
}
