package dao

import (
	"github.com/yugabyte/gocql"
	"log"
)

var dao *Dao

type Dao struct {
	DBCluster *gocql.ClusterConfig
	Session   *gocql.Session
}

func NewDao(dbCluster *gocql.ClusterConfig) (*Dao, error) {
	session, err := dbCluster.CreateSession()
	if err != nil {
		log.Printf("[Error] Get DB session err, err=%v", err)
		return nil, err
	}
	dao = &Dao{
		DBCluster: dbCluster,
		Session:   session,
	}
	return dao, nil
}
