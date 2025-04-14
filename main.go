package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// checks to make sure that the user is providing an input file
// and an output file
func checkInput() {
	if len(os.Args) == 1 {
		fmt.Println("Please include an input file and output file name.")
		os.Exit(0)
	} else if len(os.Args) == 2 {
		fmt.Println("Please include an output file name.")
		os.Exit(0)
	} else if len(os.Args) > 3 {
		fmt.Println("Too many arguments.")
		os.Exit(0)
	}
}

// read a CSV file and return the records
func readCsvFile(filePath string) (returnMap []map[string]int) {
	f, err := os.Open(filePath)
	// error handling
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	rawData, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	// doing the mapping starts here
	header := []string{} // to hold the first row of the csv

	for eachLine, record := range rawData {
		// build the first row and use it as keys
		if eachLine == 0 {
			for i := 0; i < len(record); i++ {
				header = append(header, strings.TrimSpace(record[i]))
			}
		} else {
			// add the values based on the keys from first line
			line := map[string]int{}
			for i := 0; i < len(record); i++ {
				// convert the values from string to int
				v := record[i]
				if s, err := strconv.Atoi(v); err == nil {
					line[header[i]] = s
				}
			}
			returnMap = append(returnMap, line)
		}
	}
	return returnMap
}

func main() {
	checkInput()
	inputPath := os.Args[1]
	outputPath := os.Args[2]
	csvRecords := readCsvFile(inputPath)

	// create the json for the records
	jsonOut, err := json.Marshal(csvRecords)
	if err != nil {
		log.Fatal(err)
	}

	//writing the json output
	err_new := os.WriteFile(outputPath, jsonOut, 0664)
	if err_new != nil {
		log.Fatal(err)
	}
}
