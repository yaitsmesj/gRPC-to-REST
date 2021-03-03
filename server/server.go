package server

import (
	"context"

	"github.com/yaitsmesj/gRPC-to-REST/customerror"
	"github.com/yaitsmesj/gRPC-to-REST/http"
	pb "github.com/yaitsmesj/gRPC-to-REST/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserServer that implements user proto
type UserServer struct {
	pb.UnimplementedUserServiceServer
}

// GetUser ...
func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var res pb.GetUserResponse
	r, err := http.GetUser(req.GetId())
	if err != nil {
		if r, ok := err.(*customerror.StatusCodeError); ok {
			return &res, status.Error(codes.Code(r.StatusCode), string(r.Body))
		}
		return &res, status.Error(codes.Internal, err.Error())
	}
	res.User = r.Data
	return &res, nil
}

// GetUserList ...
func (s *UserServer) GetUserList(ctx context.Context, req *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
	var res pb.GetUserListResponse
	r, err := http.GetUserList(req.GetPage())
	if err != nil {
		return &res, rError(err)
	}
	res.Users = r.Data
	return &res, nil
}

// Create ...
func (s *UserServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	var res pb.CreateResponse
	r, err := http.Create(req.GetName(), req.GetJob())
	if err != nil {
		return &res, rError(err)
	}
	res.Name = r.Name
	res.Job = r.Job
	res.Id = r.ID
	res.CreatedAt = r.CreatedAt
	return &res, nil
}

// Update ...
func (s *UserServer) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	var res pb.UpdateResponse
	r, err := http.Update(req.GetId(), req.GetName(), req.GetJob())
	if err != nil {
		return &res, rError(err)
	}
	res.Name = r.Name
	res.Job = r.Job
	res.UpdatedAt = r.UpdatedAt
	return &res, nil
}

// Delete ...
func (s *UserServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	var res pb.DeleteResponse
	err := http.Delete(req.GetId())
	if err != nil {
		return &res, rError(err)
	}
	return &res, nil
}

func rError(err error) error {
	if r, ok := err.(*customerror.StatusCodeError); ok {
		return status.Error(codes.Code(r.StatusCode), string(r.Body))
	}
	return status.Error(codes.Internal, err.Error())
}
