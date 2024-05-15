package monify

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	EndPointLogin = "/api/v1/auth/login"
)

func (m *Monify) createToken() {
	athrHeader := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", m.ApiKey, m.SecretKey)))
	m.BasicToken = athrHeader
}

func (m *Monify) generateToken() error {
	url := fmt.Sprintf("%s%s", m.BaseUrl, EndPointLogin)
	m.createToken()
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil
	}
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", m.BasicToken))

	resp, err := m.Client.Do(req)

	if err != nil {
		return err
	}
	var response ResponseGenerate
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}
	m.BearerToken = response.ResponseBody.AccessToken
	return nil
}
func (m *Monify) regenerateToken() error {
	if err := m.generateToken(); err != nil {
		return err
	}
	return nil
}
