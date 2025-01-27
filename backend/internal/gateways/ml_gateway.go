package gateways

import (
	"Christmas_prj/backend/internal/payload"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MLGateway struct {
	MLURL string
}

func NewMLGateway(MLURL string) *MLGateway {
	return &MLGateway{MLURL: MLURL}
}

func (mlg *MLGateway) GetGiftRecommendation(giftRequest payload.GiftRequest) (*payload.GiftResponse, error) {
	giftRequestJSON, err := json.Marshal(giftRequest)
	if err != nil {
		return nil, fmt.Errorf("unable to convert user struct to json: %s", err)
	}

	r := bytes.NewBuffer(giftRequestJSON)
	resp, err := http.Post(mlg.MLURL, "application/json", r)
	if err != nil {
		return nil, fmt.Errorf("unable to send request to ml: %s", err)
	}
	defer resp.Body.Close()

	return validateMLResponse(resp)
}

func validateMLResponse(response *http.Response) (*payload.GiftResponse, error) {
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ML model responded with status: %d", response.StatusCode)
	}

	var giftResponse payload.GiftResponse
	if err := json.NewDecoder(response.Body).Decode(&giftResponse); err != nil {
		return nil, fmt.Errorf("failed to decode ML response: %s", err)
	}
	return &giftResponse, nil
}
