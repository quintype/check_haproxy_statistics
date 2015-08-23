package main

import (
	"github.com/quintype/check-haproxy-statistics"
	"os"
	"fmt"
)

func readEntitiesFromSocket(filename string) ([]map[string]string) {
	stream, err := haproxy.StatsStream(filename)
	if (err != nil) {
		fmt.Println("Couldn't connect to Socket")
		os.Exit(2)
	}
	defer stream.Close()

	reader := haproxy.NewCSVReader(stream)

	entities, err := reader.ReadAll()
	if (err != nil) {
		fmt.Println("Malformed CSV")
		os.Exit(2)
	}

	return entities
}

func findEntity(entities []map[string]string, f func(map[string]string) bool) (map[string]string) {
	for _, row := range entities {
		if(f(row)) {
			return row
		}
	}
	fmt.Println("The Server is not in HAProxy")
	os.Exit(2)
	return nil
}

func main() {
	entities := readEntitiesFromSocket("/tmp/haproxysock")
	entity := findEntity(entities, func (row map[string]string) bool {
		return row["pxname"] == "listen9001" && row["svname"] == "server-1"
	})

	fmt.Println("The Server is", entity["status"])

	if(entity["status"] != "UP") {
		os.Exit(2)
	}
}
