package config

import (
	"fmt"
	"hash/fnv"

	"github.com/BurntSushi/toml"
)

/*
	Shard struct describes the structure which holds appropriate set of keys
	Each shard instance will have a unique set of keys
*/
type Shard struct {
	ShardName    string
	ShardIdx     int
	ShardAddress string
}

// config for shrading config
type Config struct {
	Shards []Shard
}

/*

Shards are reprsentingg an easier-to-represent of sharding config:
the shards will count. current index and the addresses of all other shards too
*/

type Shards struct {
	Count      int
	CurrentIdx int
	Addrs      map[int]string
}

//parse file will parse the file and will return it upon success

func ParseFile(filename string) (Config, error) {
	var conf Config
	_, err := toml.DecodeFile(filename, &conf)
	if err != nil {
		return Config{}, err
	}
	return conf, nil
}

//parseShards will convert and verify the list of the shards specified in the config
// into a form that can be used to route the files

func ParseShards(shards []Shard, currentShardName string) (*Shards, error) {
	shardCount := len(shards)
	shardIndex := -1
	addrs := make(map[int]string)

	for _, s := range shards {
		_, locErr := addrs[s.CurrentIdx]
		if locErr != nil {
			return nil, fmt.Errorf("Duplicate of the Shard Index: %d", s.ShardIdx)
		}

		addrs[s.CurrentIdx] = s.Address
	}
	return &Shards{
		Addrs:      addrs,
		Count:      shardCount,
		CurrentIdx: shardIndex,
	}, nil
}

//Index will returns the shard number for the corresponding key

func (s *Shards) Index(key string) int {
	h := fnv.New64()
	h.Write([]byte(key))
	return int(h.Sum64() % uint64(s.Count))
}
