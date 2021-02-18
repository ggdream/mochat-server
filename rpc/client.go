package rpc

import (
	"context"
	"github.com/ggdream/mochat-server/api"
	"google.golang.org/grpc"
)


func New() (*Rpc, error) {
	conn, err := grpc.Dial(":9999", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	handle := api.NewRemoteClient(conn)

	return &Rpc{
		RemoteClient: handle,
		ctx: context.Background(),
	}, nil
}
