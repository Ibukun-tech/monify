package monify

import "net/http"

func (m *Monify) ListTransaction(trb TransactionRequestBody) (*TransactionResponseBody, error) {
	var response TransactionResponseBody
	if err := m.run(http.MethodPost, "", trb, response); err != nil {
		if err := m.regenerateToken(); err != nil {
			return nil, err
		}
	}
	return &response, nil
}
