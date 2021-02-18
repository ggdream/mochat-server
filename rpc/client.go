package rpc

import (
	"context"
	"github.com/ggdream/mochat-server/api"
	"github.com/ggdream/mochat-server/config"
	"google.golang.org/grpc"
)


func New() (*Rpc, error) {
	conn, err := grpc.Dial(config.RpcURI, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	handle := api.NewRemoteClient(conn)

	return &Rpc{
		RemoteClient: handle,
		ctx: context.Background(),
	}, nil
}
