package main

import (
	"flag"
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

func description() {
	fmt.Println("urbit v0.0.0.1 : A distributed,key-value store.")
}

//definign flags for cli usage

var (
	dbLocation = flag.String("db-location", "", "The path to the bolt DB database")
)

//validate our flags
/*
func parseFlags() {
	flag.Parse()

	if *dbLocation == "" {
		//if empty then not a valid location
		log.Fatal("ERROR: DB Location Provided is empty, must provide a DB location")
	}
}
*/
func main() {
	description()

	//	parseFlags()
	db, err := bolt.Open(*dbLocation, 0666, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

}
