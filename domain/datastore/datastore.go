package datastore

import (
	//	"DCCS/constants"
	"DCCS/constants"
	"fmt"
	"log"
	"time"

	"github.com/couchbase/gocb/v2"
)

func NewCouchbaseSession() (*gocb.Cluster, error) {
	cluster, err := gocb.Connect(
		constants.CouchbaseAddress,
		gocb.ClusterOptions{
			Username: constants.CouchbaseUsername,
			Password: constants.CouchbasePassword,
		},
	)
	if err != nil {
		fmt.Println("error on connecting with db :", err.Error())
		panic(err)
	}
	err = cluster.WaitUntilReady(
		10*time.Second,
		&gocb.WaitUntilReadyOptions{DesiredState: gocb.ClusterStateOnline})
	if err != nil {
		fmt.Println("error on wait until ready:", err.Error())
		panic(err)
	}

	return cluster, nil
}

func CouchbasePing() {
	cluster, err := NewCouchbaseSession()
	if err != nil {
		log.Fatal("error on creating couchbase session")
	}
	fmt.Println("pinging with db ...")
	_, err = cluster.Ping(nil)
	if err != nil {
		log.Println("we have an error on connectiong with database: ", err.Error())
		panic(err)
	}
	log.Println("db ping success ")
}
