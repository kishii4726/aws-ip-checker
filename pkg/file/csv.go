package file

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCsv(input_file string) [][]string {
	file, err := os.Open(input_file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return rows
}
