package collector

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Result struct {
	Time               int64
	CpuUsage           float32
	CpuLimit           float32
	AvgCpu             float32
	CpuUsagePercentage float32
	PodCount           int64
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
	header := "time,cpuUsage, cpuLimit, avgCpu,podCount\n"
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
		bt.WriteString(strconv.FormatFloat(float64(val.AvgCpu), 'f', 6, 32))
		bt.WriteString(",")
		bt.WriteString(strconv.FormatFloat(float64(val.CpuUsagePercentage), 'f', 6, 32))
		bt.WriteString(",")
		bt.WriteString(strconv.FormatInt(val.PodCount, 10))
		bt.WriteString("\n")
		dstFile.WriteString(bt.String())
	}
	fmt.Println("写入文档" + fileName + "成功!")

}
