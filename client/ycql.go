package client

import (
	"github.com/yugabyte/gocql"
	"time"
)

const (
	Cluster0IP = "xcnd5.comp.nus.edu.sg"
	Cluster1IP = "xcnd6.comp.nus.edu.sg"
	Cluster2IP = "xcnd7.comp.nus.edu.sg"
	Cluster3IP = "xcnd8.comp.nus.edu.sg"
	Cluster4IP = "xcnd50.comp.nus.edu.sg"
	TestIP     = "127.0.0.1"
	KeySpace   = "cs5424"
)

var (
	Session *gocql.Session
)

func InitDB() (err error) {
	var dbCluster *gocql.ClusterConfig
	dbCluster = gocql.NewCluster(TestIP)
	dbCluster.Keyspace = KeySpace
	dbCluster.Consistency = gocql.Quorum

	dbCluster.Timeout = 1000000 * time.Millisecond
	dbCluster.Port = 9042
	dbCluster.ConnectTimeout = 100000000 * time.Millisecond
	dbCluster.MaxWaitSchemaAgreement = 100000000 * time.Millisecond

	Session, err = dbCluster.CreateSession()
	return err
}
