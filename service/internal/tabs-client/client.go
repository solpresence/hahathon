package client

import (
	"hahathon/internal/tabs-client/repo"
	"hahathon/internal/tabs-client/repo/actions"
	"log/slog"
)

type client struct {
	ActionTypes repo.ActionTypesRepo
	
	log *slog.Logger
}

func NewClient(token string, log *slog.Logger) *client {
	return &client{
		ActionTypes: actions.NewActionsTypes(token, log),
		log:         log,
	}
}
