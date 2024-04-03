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
