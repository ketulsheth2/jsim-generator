package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

const (
	ES_INTERVAL = 15
	ES_INDEX    = "ghost_jsonsim_test"
)

func initializeESClient() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{"https://elk.etterno.io:8080"},
		Username:  "elastic",
		Password:  "i6ql7meFan3n9BB4VKbU",
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
		return nil, err
	}
	log.Println(elasticsearch.Version)
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	return es, nil
}

func main() {

	//Read app data
	file, err := ioutil.ReadFile("./resources/app_data.json")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
		return
	}
	appData := AppData{}
	err = json.Unmarshal([]byte(file), &appData)
	if err != nil {
		log.Fatalf("Error unmarshalling app data: %s", err)
		return
	}

	// Read user data
	file, err = ioutil.ReadFile("./resources/user_data.json")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
		return
	}
	users := make([]Users, 0)
	err = json.Unmarshal([]byte(file), &users)
	if err != nil {
		log.Fatalf("Error unmarshalling users data: %s", err)
		return
	}

	//Initialize ES client
	es, err := initializeESClient()
	if err != nil {
		log.Fatalf("Error initializing ES client: %s", err)
		return
	}

	//Date Range
	start := time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC)
	end := time.Date(2022, 1, 25, 0, 0, 0, 0, time.UTC)

	//Iterate over time every ES_INTERVAL
	for {
		if start.After(end) {
			break
		}

		fmt.Println(start)

		//Generate ghost data
		ghost := generateGhostData(start, users, appData)

		//Send data to ES
		func() {
			ghostJ, err := json.Marshal(ghost)
			if err != nil {
				log.Fatalf("Error marshalling ghost data: %s", err)
			}
			//log.Printf("%s\n", ghostJ)
			res, err := es.Index(ES_INDEX, strings.NewReader(string(ghostJ)))
			if err != nil {
				log.Fatalf("Error indexing ES data: %s", err)
			}
			defer res.Body.Close()
			//[TODO] Remove print
			log.Println(res)
		}()

		// ghostJson, err := json.Marshal(ghost)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Printf("%s\n", ghostJson)
		start = start.Add(time.Second * ES_INTERVAL)
	}
}
