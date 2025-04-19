package client

import (
	"hahathon/internal/tabs-client/repo"
	"hahathon/internal/tabs-client/repo/actions"
	"log/slog"
)

type Client struct {
	ActionTypes repo.ActionTypesRepo
	Token       string

	log *slog.Logger
}

func NewClient(token string, log *slog.Logger) *Client {
	return &Client{
		ActionTypes: actions.NewActionsTypes(token, log),
		log:         log,
	}
}
