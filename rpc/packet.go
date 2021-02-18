package rpc

import (
	"context"
	"github.com/ggdream/mochat-server/api"
	"github.com/ggdream/mochat-server/model/msg"
	"google.golang.org/protobuf/types/known/timestamppb"
)


type Rpc struct {
	api.RemoteClient
	ctx		context.Context
}

func (r *Rpc) Send(m *msg.Message) (*msg.Status, error) {
	record := api.Record{
		Id:   m.ID,
		From: m.From,
		To:   m.To,
		Time: timestamppb.Now(),
	}
	res, err := r.RemoteClient.Send(r.ctx, &record)
	if err != nil {
		return nil, err
	}

	return &msg.Status{
		Code:    res.Code,
		Message: res.Message,
	}, nil
}

func (r *Rpc) Query(d *msg.DescMsg) (*api.Data, error)  {
	queryMeta := api.QueryMeta{
		From:   d.From,
		To:     d.To,
		Start:  d.Start,
		Offset: d.Offset,
	}
	return r.RemoteClient.Query(r.ctx, &queryMeta)
}
