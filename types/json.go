package types

import pb "github.com/yaitsmesj/gRPC-to-REST/proto"

// UserJSON ...
type UserJSON struct {
	Data *pb.User
}

// UserListJSON ...
type UserListJSON struct {
	Data []*pb.User
}

// CreateJSON ...
type CreateJSON struct {
	Name      string
	Job       string
	ID        int32 `json:"id,string"`
	CreatedAt string
}

// UpdateJSON ...
type UpdateJSON struct {
	Name      string
	Job       string
	UpdatedAt string
}
