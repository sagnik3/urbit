package web

import (
	"fmt"
	"io"
	"net/http"
)

/*
Package handler for web handler;
contains HTTP method handler to be used for the database

*/

type Server struct {
	//connect to HTTP method handler tp be used for db
	database *db.Location
	shards   *config.Shards
}

func InitNewServer(database *db.Location, shards *config.Shards) *Server {
	//initialize a new server instance with HTTP handlers
	return &Server{
		database: database,
		shards:   shards,
	}
}

func (server *Server) redirectShard(shard int, w http.ResponseWriter, r *http.Request) {
	//function to redirect load to different shards to balance and and copy the data to the new server

	urltoRedirect := "http://" + shard.shards.Address[shard] + r.RequestURI

	fmt.Fprintf(w, "\n--> Redirecting from shard no %d to shard no %d (%q)\n", shard.shards.CurrentShardId, shard, &urltoRedirect)

	response, err := http.Get(urltoRedirect)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "REDIRECTION-ERROR: Error redirecting accordin to request : %v", err)
		return
	}
	defer response.Body.Close()

	io.Copy(w, response.Body)

}
