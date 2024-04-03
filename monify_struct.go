package monify

import (
	"net/http"
	"time"
)

// This takes in the MainUrl,ApiKey and also takes into account the transport
type Monify struct {
	BaseUrl string
	ApiKey  string
	Client  *http.Client
}

// The token you need to verify in the application in the log in before you run other things
type GenerateTokenResponse struct {
	AccessToken string    `json:"accessToken"`
	ExpiresIn   time.Time `json:"expiresIn"`
}

// This is the response you will need once you log in to the application
type ResponseGenerate struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      GenerateTokenResponse
}

// This is the Request transacrion Body things you need to send to the body

type TransactionRequestBody struct {
	Amount             float64  `json:"amount"`
	CustomerName       string   `json:"customerName"`
	CustomerEmail      string   `json:"customerEmail"`
	PaymentReference   string   `json:"paymentReference"`
	PaymentDescription string   `json:"paymentDescription"`
	CurrencyCode       string   `json:"currencyCode"`
	RedirectUrl        string   `json:"redirectUrl"`
	PaymentMethod      []string `json:"paymentMethod"`
}
type TransactionResponseBody struct {
	TransactionReference string   `json:"transactionReference"`
	PaymentReference     string   `json:"paymentReference"`
	MerchantName         string   `json:"merchantName"`
	ApiKey               string   `json:"apiKey"`
	EnablePaymentMethod  []string `json:"enablePaymentMethod"`
	CheckOutUrl          string   `json:"checkoutUrl"`
}
type TransactionResponse struct {
	RequestSuccessful bool                   `json:"requestSuccessful"`
	ResponseMessage   string                 `json:"responseMessage"`
	ResponseCode      string                 `json:"responseCode"`
	ResponseBody      TransactionRequestBody `jso:"responseBody"`
}
