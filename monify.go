package monify

import (
	"bytes"
	"encoding/json"
	"io"
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
func (m *Monify) run(method string, url string, payload, response interface{}) error {
	var body io.Reader
	if payload != nil {
		val, _ := json.Marshal(payload)
		body = bytes.NewReader(val)
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	res, err := m.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&response)
	return nil
}
