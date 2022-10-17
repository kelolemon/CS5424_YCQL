package client

import (
	"github.com/yugabyte/gocql"
)

const (
	Cluster0IP = "xcnd5.comp.nus.edu.sg"
	Cluster1IP = "xcnd6.comp.nus.edu.sg"
	Cluster2IP = "xcnd7.comp.nus.edu.sg"
	Cluster3IP = "xcnd8.comp.nus.edu.sg"
	Cluster4IP = "xcnd50.comp.nus.edu.sg"
	TesterIP   = "127.0.0.1"
	KeySpace   = "CS5424"
)

var (
	DBCluster *gocql.ClusterConfig
)

func InitDB() {
	// connect to the cluster
	DBCluster = gocql.NewCluster(TesterIP)
	DBCluster.Keyspace = KeySpace
	DBCluster.Consistency = gocql.Quorum
}
