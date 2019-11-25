package sender

import (
	"../api"
	"context"
	"github.com/google/uuid"
	"time"
)

var counter int = 0

type GRPCServer struct {}

func (p *GRPCServer) IdSend(ctx context.Context, req *api.Empty) (*api.UUID, error) {
	counter++
	if counter % 100 == 0 {
		time.Sleep(2 * time.Second)
	}
	id, err := uuid.NewUUID()
	return &api.UUID{Uuid: id.String()}, err
}