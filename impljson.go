package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func toJSON(pb proto.Message) string {
	option := protojson.MarshalOptions{
		Multiline: true,
	}

	content, err := option.Marshal(pb)
	if err != nil {
		log.Fatalln("Cannot convert given data to json format!", err)
		return ""
	}

	return string(content)
}

func fromJSON(givenData string, pb proto.Message) {
	if err := protojson.Unmarshal([]byte(givenData), pb); err != nil {
		log.Fatalln("Failed to convert data to protobuf!", err)
		return
	}
}
