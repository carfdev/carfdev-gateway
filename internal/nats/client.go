package nats

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/nats-io/nats.go"
)

type NatsClient struct {
	conn *nats.Conn
}

type Envelope struct {
	Data  json.RawMessage `json:"data,omitempty"`
	Error *struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func NewNatsClient(url string) (*NatsClient, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsClient{conn: conn}, nil
}

func (c *NatsClient) RequestWithContext(ctx context.Context, subject string, data []byte) ([]byte, error) {
	msg, err := c.conn.RequestWithContext(ctx, subject, data)
	if err != nil {
		return nil, err
	}

	var env Envelope
	if err := json.Unmarshal(msg.Data, &env); err != nil {
		return nil, err
	}

	if env.Error != nil {
		return nil, errors.New(env.Error.Message)
	}

	return env.Data, nil
}
