# Assignment 3 - Creating a Command-Line Application
This program takes a csv file that is read in to the program as strings, and then outputs a jsonl file. The first line of the csv file is used as the key for all lines in the jsonl file, and all of the entries are converted to numbers (ints). 

The user can provide two different arguments: an input file path and an output file path, but this is optional. If the user does not specify inputs or outputs, then the program will use housesInput.csv as the default input and housesOutput.jsonl as the default output. If a user wants to specify an output, they must also specify an input.

## Data
The housesInput.csv has been included in this repository to verify that this program functions with the default settings.

The second csv file, testData.csv, is used to verify that the TestCreateData test works.

## How to Test
To test the functions in this program, run **go test** with the terminal set to the same directory as the main.go and main_test.go files.

## Application
Use **go build** to compile to program into an executable. For more information, see [this entry](https://www.markdownguide.org/basic-syntax/) in the Go documentation. This program was initially made with Windows.

