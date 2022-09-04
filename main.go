package main

import (
	"fmt"
	"reflect"

	pb "github.com/WytQuant/go-proto-practice/proto"
	"google.golang.org/protobuf/proto"
)

func doAddrbook() *pb.Person {
	return &pb.Person{
		Name:  "wittawas",
		Id:    1,
		Email: "wittawas@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "091231132", Type: pb.Person_MOBILE},
			{Number: "087223718", Type: pb.Person_WORK},
			{Number: "080123331", Type: pb.Person_HOME},
		},
	}
}

func implementFile(p proto.Message) {
	fileName := "test_file.txt"

	WriteToFile(fileName, p)
	fmt.Println("Writing given data to file successfully!")

	message := &pb.Person{}
	ReadFromFile(fileName, message)
	fmt.Println(message)
}

func implToJSON(pb proto.Message) string {
	return toJSON(pb)
}

func implFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonString, message)
	return message
}

func main() {
	// fmt.Println(doAddrbook())
	// implementFile(doAddrbook())
	jsonString := implToJSON(doAddrbook())
	fmt.Println(jsonString)

	protoMessage := implFromJSON(jsonString, reflect.TypeOf(pb.Person{}))
	fmt.Println(protoMessage)
}
