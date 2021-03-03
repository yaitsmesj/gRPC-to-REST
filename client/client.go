package main

import (
	"context"
	"log"

	pb "github.com/yaitsmesj/gRPC-to-REST/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	user, err := client.GetUser(context.Background(), &pb.GetUserRequest{Id: 2})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user)

	list, err := client.GetUserList(context.Background(), &pb.GetUserListRequest{Page: 2})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(list)

	nuser, err := client.Create(context.Background(), &pb.CreateRequest{Name: "Suraj", Job: "Engineer"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(nuser)

	muser, err := client.Update(context.Background(), &pb.UpdateRequest{Name: "jj", Job: "Engineer", Id: 2})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(muser)

	duser, err := client.Delete(context.Background(), &pb.DeleteRequest{Id: 2})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(duser)
}
