package main


import "net"

const (
	port = ":8080"
)

func main () {
	lis,err := net.Listen("tcp",port)
	if err !=nil{
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()
