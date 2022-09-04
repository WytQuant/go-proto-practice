package main

import (
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

func WriteToFile(fileName string, pb proto.Message) {
	dataOut, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Cannot serialize input data!", err)
		return
	}

	if err := ioutil.WriteFile(fileName, dataOut, 0644); err != nil {
		log.Fatalln("Cannot writing file with given data!", err)
		return
	}
}

func ReadFromFile(fileName string, pb proto.Message) {
	readData, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Cannot read data from given file!", err)
		return
	}

	if err := proto.Unmarshal(readData, pb); err != nil {
		log.Fatalln("Failed to parse to protobuf!", err)
		return
	}
}
