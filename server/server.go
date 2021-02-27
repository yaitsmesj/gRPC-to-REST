package server

import (
	"context"

	pb "github.com/yaitsmesj/gRPC-to-REST/proto"
)

// UserServer that implements user proto
type UserServer struct {
	pb.UnimplementedUserServiceServer
}

// GetUser ...
func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	return nil, nil
}

// GetUserList ...
func (s *UserServer) GetUserList(ctx context.Context, req *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {

	return nil, nil
}

// Create ...
func (s *UserServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {

	return nil, nil
}

// Update ...
func (s *UserServer) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {

	return nil, nil
}

// Delete ...
func (s *UserServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {

	return nil, nil
}
