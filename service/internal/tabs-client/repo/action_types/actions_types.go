package actiontypes

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
)

const (
	url = "https://true.tabs.sale/fusion/v1/datasheets/dsteLahm1rxF86r9fC/records?viewId=viwUwwAUFHpP6&fieldKey=id"
)

type ActionsTypes struct {
	token string
	log *slog.Logger
}

func (at *ActionsTypes) Create(body PostReq) *PostRes {
	jsonData, err := json.Marshal(body)
	if err != nil {
		at.log.Warn("error encoding JSON", "err", err)
		return nil
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		at.log.Warn("failed to creating request", "err", err)
		return nil
	}

	req.Header.Set("Authorization", "Bearer "+at.token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		at.log.Warn("error sending request", "err", err)
		return nil
	}
	defer resp.Body.Close()

	var result PostRes
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		at.log.Warn("error decoding response", "err", err)
	}

	return &result
}

func NewActionsTypes(token string, log *slog.Logger) *ActionsTypes {
	return &ActionsTypes{
		token: token,
		log: log,
	}
}