package actiontypes

type (
	PostReq struct {
		Records  []struct {
			RecordId string `json:"recordId"`
			Fields   struct {
				FldWgM7Y77RrU string `json:"fldWgM7Y77RrU"` // Содержание
			}
		}
		FieldKey string `json:"fieldKey"`
	}
)

type (
	PostRes struct {
		Code    int    `json:"code"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    struct {
			Records  []struct {
				RecordId string `json:"recordId"`
				Fields   struct {
					FldGKR0G2BTog int    `json:"fldGKR0G2BTog"` // id
					FldWgM7Y77RrU string `json:"fldWgM7Y77RrU"` // Содержание
				} `json:"records"`
			}
		} `json:"data"`
	}
)
