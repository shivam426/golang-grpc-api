package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "totality_corp/gen/proto"

	"google.golang.org/grpc"
)

var db = []User{}

type User struct {
	id      int32
	name    string
	city    string
	phone   int64
	height  float32
	married bool
}

type CorpapiServer struct {
	pb.UnimplementedCorpapiServer
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

func (serve *CorpapiServer) CreateUser(ctx context.Context, req *pb.CreateRequestResponse) (*pb.CreateRequestResponse, error) {
	data := User{req.User.GetId(), req.User.GetName(), req.User.GetCity(), req.User.GetPhone(), req.User.GetHeight(), req.User.GetMarried()}
	db = append(db, data)
	return &pb.CreateRequestResponse{User: req.User}, nil
}

func (serve *CorpapiServer) FetchUser(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	var data *pb.User
	for _, value := range db {
		if value.id == req.Id {
			data = getData(&value)
		}
	}
	return &pb.ReadResponse{User: data}, nil
}

func main() {

	listner, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	pb.RegisterCorpapiServer(s, &CorpapiServer{})
	err = s.Serve(listner)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(db)
}
