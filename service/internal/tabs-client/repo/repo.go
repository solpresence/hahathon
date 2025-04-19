package repo

import actiontypes "hahathon/internal/tabs-client/repo/action_types"

type ActionTypesRepo interface {
	Create(actiontypes.PostReq) *actiontypes.PostRes
}
