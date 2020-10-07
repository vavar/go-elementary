package main

import (
	"log"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./actors.csv")
	if err != nil {
		log.Fatal(err)
	}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	ranking := map[string]int{}
	for _, record := range records {
		ranking[record[3]]++
	}

	all := 0
	for name, count := range ranking {
		if count > 1 {
			fmt.Printf(" %s\n", name)
			all++
		}
	}
	fmt.Printf("Total Actors: %d\n", all)
}
