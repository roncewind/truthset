package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/roncewind/szrecord"
)

// ----------------------------------------------------------------------------
func readJSONLResource(jsonURL string, recordchan chan *szrecord.Record) {
	response, err := http.Get(jsonURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	scanner.Split(bufio.ScanLines)

	i := 0
	for scanner.Scan() {
		i++
		str := strings.TrimSpace(scanner.Text())
		// ignore blank lines
		if len(str) > 0 {
			record, err := szrecord.NewRecord(str)
			if err != nil {
				//something went wrong with a record.
				fmt.Println("Line", i, ": id =", record.Id)
			} else {
				recordchan <- record
			}
		}
	}
	close(recordchan)
}

// ----------------------------------------------------------------------------
func main() {
	recordchan := make(chan *szrecord.Record)
	outfilename := "snippet.go"

	file, err := os.OpenFile(outfilename, os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		//Just bail out if we can't create the file.
		panic(err)
	}

	// start reading the remote truth set.
	go readJSONLResource("https://s3.amazonaws.com/public-read-access/TestDataSets/SenzingTruthSet/truth-set-3.0.0.jsonl", recordchan)

	// write the file header
	file.WriteString("// A list of test records.\n")
	file.WriteString("var ReferenceRecords = map[string]struct {\n")
	file.WriteString("\tDataSource string\n")
	file.WriteString("\tId         string\n")
	file.WriteString("\tData       string\n")
	file.WriteString("\tLoadId     string\n")
	file.WriteString("}{\n")

	// loop over the records and write them
	for {
		select {
		case record, open := <-recordchan:
			if !open && len(recordchan) == 0 {
				//channel is empty and closed, we're done.
				file.WriteString("}\n")
				os.Exit(0)
			}
			file.WriteString(fmt.Sprintf("\"%s\": {\n", record.Id))
			file.WriteString(fmt.Sprintf("\tDataSource: \"%s\",\n", record.DataSource))
			file.WriteString(fmt.Sprintf("\tId:         \"%s\",\n", record.Id))
			file.WriteString(fmt.Sprintf("\tData:       `%s`,\n", record.Json))
			file.WriteString("\tLoadId:     \"TRUTHSET_REFERENCE_LOAD\",\n")
			file.WriteString("},\n")

		}
	}
}
