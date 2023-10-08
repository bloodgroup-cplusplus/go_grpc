package main

// its like the client is making functional request on the server side

import (
	"fmt"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/bloodgroup-cplusplus/go_grpc/proto"
)

const (
	port = ":8080"
)



func main(){
	// how to make http request to 
	// server request from lcient in golang 
	// we have a golang program in client to make a request 
	conn,err := grpc.Dial("localhost"+port,grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err !=nil {
		log.Fatalf("did not connect :%v",err)
	}


	// when everything is over client closes

	defer conn.Close()

//	client :=pb.NewGreetServiceClient(conn)

//	names :=&pb.NamesList{
//		Names: []string {"Chad","Alice","Bob"},
//	}

//	callSayHello(client)
}
