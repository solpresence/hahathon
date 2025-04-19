package actionsEndpoints

type (
	PostReq struct {
		IdEmployee string `json:"fld4bl4ul8fmz" validate:"required"`
		IdAction   string `json:"fldRGb7pEMUb8" validate:"required"`
		IdLocation string `json:"fldu9XfPVMGGP" validate:"required"`
	}
)

type (
	PostRes struct {
		Code    int    `json:"code"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)
