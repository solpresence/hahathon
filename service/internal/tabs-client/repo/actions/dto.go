package actions

type (
	PostReq struct {
		Records []struct {
			Fields struct {
				IdEmployee []string `json:"id_employee"`
				IdAction   []string `josn:"id_action"`
				IdLocation []string `json:"id_location"`
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
