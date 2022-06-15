package replica

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

//Repsonse for the get next key for replication in form of struct

type NextKeyStruct struct {
	Key   string
	Value string
	Err   error
}

//client struct that handles the database instance and the leader address

type clientStruct struct {
	database      *db.Database
	leaderAddress string
}

// function to downloads new keys from the main and apply them to sharding principle

func ClientLookupLoop(database *db.Database, leaderAddress string) {
	C := &clientStruct{
		database: database, leaderAddress: leaderAddress,
	}

	for {
		presentInstance, err := C.loopAround()

		if err != nil {
			log.Printf("Error-loop: Looping Error to look around %v", err)
			time.Sleep(time.Second)
			continue
		}

		if !presentInstance {
			//wait to free and then lookup again
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func (client *clientStruct) loopAround() (presentInstace bool, err error) {
	resp, err := http.Get("http://" + c.leaderAddress + "/next-replica-key")
	if err != nil {
		return false, err
	}
	var keyInstance NextKeyStruct

	if err := json.NewDecoder(resp.Body).Decode(&keyInstance); err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if keyInstance.Err != nil {
		//edge case: if some keyInstance causes some error
		return false, err
	}

	if keyInstance.Key == "" {
		//edge case :- no key present in the keyInstance
		return false, nil
	}

	if err := client.db.SettingKeyOnReplica(keyInstance.Key, []byte(keyInstance.Value)); err != nil {
		return false, nil
	}

	if err := client.DeleteKeyFromReplicationQueue(keyInstance.Key, keyInstance.Value); err != nil {
		log.Printf("DELETE-ERROR: Error Deleting the replication key: %v", err)
	}

	return true, nil
}

func (client *clientStruct) DeleteKeyFromReplicationQueue(key, value string) error {
	urlVal := url.Values{}
	urlVal.Set("key", key)
	urlVal.Set("value", value)

	log.Printf("Deleting key=%q, value=%q from the replication queue on %q", key, value, client.leaderAddress)

	response, err := http.Get("http://" + client.leaderAddress + "/delete-replica-key" + urlVal.Encode())

	if err != nil {
		return err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if !bytes.Equal(result, []byte("success")) {
		return errors.New(string(result))
	}
	return nil
}
