package actions

type (
	PostReq struct {
		Records []struct {
			Fields struct {
				IdEmployee []string `json:"fld4bl4ul8fmz"`
				IdAction   []string `json:"fldRGb7pEMUb8"`
				IdLocation []string `json:"fldu9XfPVMGGP"`
			} `json:"fields"`
		} `json:"records"`
		FieldKey string `json:"fieldKey"`
	}
)

type (
	PostRes struct {
		Code    int    `json:"code"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)
