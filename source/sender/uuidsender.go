package sender

import (
	"../api"
	"context"
	"github.com/google/uuid"
)

type GRPCServer struct {}

func (p *GRPCServer) IdSend(ctx context.Context, req *api.Empty) (*api.UUID, error) {
	id, err := uuid.NewUUID()
	return &api.UUID{Uuid: id.String()}, err
}