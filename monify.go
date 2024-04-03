package monify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	mainUrl       = "https://sandbox.monnify.com"
	EndPointLogin = "/api/v1/auth/login"
)

func NewMonify(ApiKey string) *Monify {
	return &Monify{
		BaseUrl: mainUrl,
		ApiKey:  ApiKey,
		Client:  &http.Client{},
	}
}

func (m *Monify) GenerateToken(ctx context.Context) (*GenerateTokenResponse, error) {
	url := fmt.Sprintf("%s%s", m.BaseUrl, EndPointLogin)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	// still need to work on the api connection key with the header to uderstan what I am doing
	athrHeader := fmt.Sprintf("Basic %s", "")
	req.Header.Add("Authorization", athrHeader)
	resp, err := m.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	generateTokenResponse := new(GenerateTokenResponse)
	err = json.NewDecoder(resp.Body).Decode(generateTokenResponse)
	if err != nil {
		return nil, err
	}
	return generateTokenResponse, nil
}
