package email

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/carfdev/carfdev-gateway/internal/nats"
)

const (
	SubjectSendContact = "email.send_contact"
)

type EmailService interface {
	SendContact(ctx context.Context, req SendContactRequest) (*SendResponse, error)
}

type emailService struct {
	nats *nats.NatsClient
}

func NewEmailService(nc *nats.NatsClient) EmailService {
	return &emailService{nats: nc}
}

func (s *emailService) SendContact(ctx context.Context, req SendContactRequest) (*SendResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.nats.RequestWithContext(ctx, SubjectSendContact, data)
	if err != nil {
		return nil, fmt.Errorf("failed to send NATS request on %s: %w", SubjectSendContact, err)
	}

	var response SendResponse
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal NATS response: %w", err)
	}
	return &response, nil
}
