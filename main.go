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

// open the file and read the results
func createData(filePath string) [][]string {
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

	return rawData
}

// write the a json to each line of the output document
func writeData(newData [][]string, filePath string) (returnString string) {

	//declaring variables
	var csvHead []string //setting this up to create the head from first column

	// create and open a new document to hold the data
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// iterating through the data to write each line
	for eachLine, record := range newData {
		// build the first row
		if eachLine == 0 {
			csvHead = buildHead(record)

		} else {
			// create a map of the first string and all subsequent lines
			line := createMap(csvHead, record)

			// create the json for the records
			jsonOut, err := json.Marshal(line)
			if err != nil {
				log.Fatal(err)
			}

			// write the json file to a new line in the document
			if _, err := f.Write(jsonOut); err != nil {
				log.Fatal(err)
			}
			// write a /n to the document
			if _, err := f.Write([]byte("\n")); err != nil {
				log.Fatal(err)
			}
		}
	}

	// close the new document that was created
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	returnString = "Your file is ready"

	return returnString
}

// builds the header of a csv file
func buildHead(inputSlice []string) []string {

	header := []string{}

	for i := 0; i < len(inputSlice); i++ {
		header = append(header, strings.TrimSpace(inputSlice[i]))
	}
	return header
}

// create a map
func createMap(header []string, record []string) map[string]int {

	line := map[string]int{}

	for i := 0; i < len(record); i++ {
		// convert the values from string to int
		v := record[i]
		if s, err := strconv.Atoi(v); err == nil {
			line[header[i]] = s
		}
	}
	return line
}

func main() {

	var inputPath, outputPath string = "housesInput.csv", "housesOutput.jsonl"
	if len(os.Args) == 2 {
		inputPath = os.Args[2]
	} else if len(os.Args) > 2 {
		inputPath = os.Args[2]
		outputPath = os.Args[3]
	}

	csvRecords := createData(inputPath)

	fmt.Println(writeData(csvRecords, outputPath))
}
