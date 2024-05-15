package monify

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

func NewMonify(ApiKey, SrtKey, MainUrl string) *Monify {
	return &Monify{
		BaseUrl:   MainUrl,
		ApiKey:    ApiKey,
		SecretKey: SrtKey,
		Client:    &http.Client{},
	}
}

func (m *Monify) GenerateToken(ctx context.Context) (*GenerateTokenResponse, error) {
	url := fmt.Sprintf("%s%s", m.BaseUrl, EndPointLogin)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	// still need to work on the api connection key with the header to uderstan what I am doing
	athrHeader := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", m.ApiKey, mainUrl))))
	req.Header.Add("Authorization", athrHeader)
	resp, err := m.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	defer resp.Body.Close()
	var generateTokenResponse GenerateTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&generateTokenResponse)
	if err != nil {
		return nil, err
	}
	return &generateTokenResponse, nil
}

// Writing the code to initialize the transaction

func (m *Monify) Transactions(ctx context.Context, bodyLoad TransactionRequestBody) error {

}
