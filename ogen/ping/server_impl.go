package ping

import "context"

type PingService struct{}

func NewPingService() *PingService {
	return new(PingService)
}

func (s *PingService) Ping(ctx context.Context) (*Pong, error) {
	resp := Pong{
		Ping: "pong",
	}

	return &resp, nil
}
