package scale

import (
	"ResourceScheduling/internal/utils"
	"fmt"
	"log"
)

func ScaleOut(podCount int, nameSpace string, clusterName string) {
	log.Fatalf("ScaleOut")
	go UpdateDatabaseConfigYaml(podCount, nameSpace, clusterName)
}

func ScaleIn(podCount int, nameSpace string, clusterName string) {
	log.Fatalf("ScaleIn")
	go UpdateDatabaseConfigYaml(podCount, nameSpace, clusterName)

}

func UpdateDatabaseConfigYaml(podCount int, nameSpace string, clusterName string) {
	log.Fatalf("UpdateDatabaseConfigYaml")
	// TiDB 运维文档 https://docs.pingcap.com/zh/tidb-in-kubernetes/stable/scale-a-tidb-cluster
	// kubectl patch -n ${namespace} tc ${cluster_name} --type merge --patch '{"spec":{"pd":{"replicas":3}}}'
	cmdStr := fmt.Sprintf("kubectl patch -n %s tc %s --type merge --patch '{\"spec\":{\"pd\":{\"replicas\":%d}}}'", nameSpace, clusterName, podCount)
	utils.CmdExecute(cmdStr)
}
