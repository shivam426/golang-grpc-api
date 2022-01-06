package main

import (
	"context"
	"fmt"
	"log"
	pb "totality_corp/gen/proto"

	"google.golang.org/grpc"
)

type User struct {
	id      int32
	name    string
	city    string
	phone   int64
	height  float32
	married bool
}

func getData(data *User) *pb.User {
	return &pb.User{
		Id:      data.id,
		Name:    data.name,
		Phone:   data.phone,
		Height:  data.height,
		Married: data.married,
	}
}

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	client := pb.NewCorpapiClient(conn)
	data := User{1, "shivam", "bbs", 123456789, 5.1, true}
	req1 := &pb.CreateRequestResponse{User: getData(&data)}
	resp1, err := client.CreateUser(context.Background(), req1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Created Data  >>>>", resp1)

	req2 := &pb.ReadRequest{Id: resp1.GetUser().GetId()}
	resp2, err := client.FetchUser(context.Background(), req2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Fetching Data  >>>>", resp2)
}
