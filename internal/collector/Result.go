package collector

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

//		r.Time = int64(start)
//		r.CpuUsage = cpuUsage
//		r.CpuLimit = cpuLimit
//		r.CpuUsagePercentage = cpuUsagePercentage
//		r.PodCount = int64(podCount)
//		r.AvgCpu = avgCpuUsage
//		r.StatementOps = statementOps
//		r.P99Latency = p99Latency
type Result struct {
	Time               int64
	CpuUsage           float32
	CpuLimit           float32
	CpuUsagePercentage float32
	PodCount           int64
	AvgCpu             float32
	StatementOps       float32
	P99Latency         float32
}

func WriteOutputs(resArr []Result) {
	log.Println("=====================")

	ti := time.Now().Unix()
	fileName := "result/" + strconv.FormatInt(ti, 10) + ".csv"

	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	header := "Time, CpuUsage, CpuLimit, CpuUsagePercentage, PodCount, AvgCpu, StatementOps, P99Latency \n"
	dstFile.WriteString(header)
	//dstFile.WriteString("time ")
	//dstFile.WriteString("avgCpu ")
	//dstFile.WriteString("pods\n")

	for _, val := range resArr {
		var bt bytes.Buffer
		bt.WriteString(strconv.FormatInt(val.Time, 10))
		bt.WriteString(",")
		bt.WriteString(strconv.FormatFloat(float64(val.CpuUsage), 'f', 6, 32))
		bt.WriteString(",")
		bt.WriteString(strconv.FormatFloat(float64(val.CpuLimit), 'f', 6, 32))
		bt.WriteString(",")
		bt.WriteString(strconv.FormatFloat(float64(val.CpuUsagePercentage), 'f', 6, 32))
		bt.WriteString(",")
		bt.WriteString(strconv.FormatInt(val.PodCount, 10))
		bt.WriteString(",")
		bt.WriteString(strconv.FormatFloat(float64(val.AvgCpu), 'f', 6, 32))
		bt.WriteString(",")
		bt.WriteString(strconv.FormatFloat(float64(val.StatementOps), 'f', 6, 32))
		bt.WriteString(",")
		bt.WriteString(strconv.FormatFloat(float64(val.P99Latency), 'f', 6, 32))
		bt.WriteString("\n")
		dstFile.WriteString(bt.String())
	}
	fmt.Println("写入文档" + fileName + "成功!")

}
