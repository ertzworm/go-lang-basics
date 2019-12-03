package main

import ("os"
				"log"
				"encoding/csv"
)


var data = [][]string{{"Company Number 1", "RP Innotech"}, {"Company Number 2", "Twist Resources"}}

func main(){
	//@ Creation of CSV File
	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error){
	if err != nil {
		log.Fatal(message, err)
	}
}