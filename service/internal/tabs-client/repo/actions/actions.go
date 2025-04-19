package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

const (
	url = "https://true.tabs.sale/fusion/v1/datasheets/dstByDoDJcx86DCame/records?viewId=viw8w1ZCpTYEt&fieldKey=name"
)

type ActionsTypes struct {
	token string
	log *slog.Logger
}

func (at *ActionsTypes) Create(body PostReq) (*PostRes, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		at.log.Warn("error encoding JSON", "err", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		at.log.Warn("failed to creating request", "err", err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+at.token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		at.log.Warn("error sending request", "err", err)
		return nil, err
	}
	defer resp.Body.Close()

	var result PostRes
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		at.log.Warn("error decoding response", "err", err)
		return nil, err
	}

	if !result.Success {
		at.log.Warn(result.Message)
		return nil, fmt.Errorf(result.Message)
	}

	return &result, nil
}

func NewActionsTypes(token string, log *slog.Logger) *ActionsTypes {
	return &ActionsTypes{
		token: token,
		log: log,
	}
}