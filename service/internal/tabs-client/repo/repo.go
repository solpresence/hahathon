package repo

import "hahathon/internal/tabs-client/repo/actions"

type ActionTypesRepo interface {
	Create(actions.PostReq) (*actions.PostRes, error)
}
