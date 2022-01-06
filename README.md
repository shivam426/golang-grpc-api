# golang-grpc-api

installation of protoc plugins

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26

$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

Installation of grpc

https://grpc.io/blog/installation/

create the .proto file for the grpc
After creating the .proto file , we need to run the protoc,
before which we will needing to  export the file

export PATH="$PATH:$(go env GOPATH)/bin"

protoc --proto_path=proto proto/*.proto --go_out=gen

protoc --proto_path=proto proto/*.proto --go-grpc_out=gen

if in any case we need to delete the .pb we can do this using this command

rm gen/proto/*.go

After this we have to create the server part.
we need to test the file before creating the client part.

go run server/server.go

we will create a client part and test through both ends 
sever and client.we need to run both file sever.go and client.go both in different terminal
Response->>>>>>

Created Data  >>>> user:{id:1  name:"shivam"  phone:123456789  height:5.1  married:true}

Fetching Data  >>>> user:{id:1  name:"shivam"  phone:123456789  height:5.1  married:true}

