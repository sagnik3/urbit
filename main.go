package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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
	/*
		Parse the required cli flags required for the db location and the shard name
		required for parsing.
	*/
	flag.Parse()

	if *dbLocation == "" {
		log.Fatalf("DB-ERROR: Database Location is not provided.")
	}
	if *shardName == "" {
		log.Fatalf("SHARD_NAME-ERROR: Shard Name is not provided.")
	}
}

func main() {
	description()

	parseCLIFlgs()

	c, err := urbitconfig.ParseFile(*shardConfigFile)

	if err != nil {
		log.Fatalf("PARSING-ERROR: Error parsing shard config file. %q: %v", *shardConfigFile, err)
	}

	//adding a web server handle the server
	server := web.CreateNewServer(newDB, shards)

	http.HandleFunc("/get", server.GetHandler)
	http.HandleFunc("/put", server.PutHandler)
	http.HandleFunc("/delete", server.DeleteHandler)
	http.HandleFunc("/next-replica-key", server.GetNextReplicationKey)
	http.HandleFunc("/delete-replica-key", server.DeleteReplicationKey)

	log.Fatal(http.ListenAndServe(*httpAddress, nil))
}
