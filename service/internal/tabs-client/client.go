package client

import (
	"hahathon/internal/tabs-client/repo"
	actiontypes "hahathon/internal/tabs-client/repo/action_types"
	"log/slog"
)

type Client struct {
	ActionTypes repo.ActionTypesRepo
	Token       string

	log *slog.Logger
}

func NewClient(token string, log *slog.Logger) *Client {
	return &Client{
		ActionTypes: actiontypes.NewActionsTypes(token, log),
		log:         log,
	}
}
