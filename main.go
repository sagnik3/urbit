package main

import (
	"flag"
	"fmt"
	"log"
)

func description() {
	fmt.Println("urbit v0.0.0.1 : A distributed,key-value store.")
}

//global flgas for setup for urbit

var dbLocation = flag.String("dbLocation", "", "Path to the Bolt DB database")
var httpAddress = flag.String("httpAddress", "127.0.0.1:8080", "HTTP host and port")
var shardConfigFile = flag.String("shardConfigFile", "shardConfig.toml", "File that describes the configuration file for static sharding")
var shardName = flag.String("shardName", "", "Name of the shard instance for the data")
var replicaName = flag.Bool("replicaName", false, "Whether to use only as read-only replica or to write to it also.")

func parseCLIFlgs() {
	flag.Parse()

	if *dbLocation == "" {
		log.Fatalf("DB-ERROR: Database Location is not provided.")
	}
	if *shardName == "" {
		log.Fatalf("SHARD_NAME-ERROR: Shard Name is not provided.")
	}
}
