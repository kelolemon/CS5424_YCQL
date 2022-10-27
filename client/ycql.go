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
	KeySpace   = "cs5424"
)

var (
	Session *gocql.Session
)

func InitDB() (err error) {
	var dbCluster *gocql.ClusterConfig
	dbCluster = gocql.NewCluster(TesterIP)
	dbCluster.Keyspace = KeySpace
	dbCluster.Consistency = gocql.Quorum
	Session, err = dbCluster.CreateSession()
	return err
}
