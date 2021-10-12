package dcutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bytixo/go-discord-utils/logger"
)

type Experiments struct {
	Fingerprint string `json:"fingerprint"`
}

// GetCfCookie : Thanks sympthey I kinda stole it
func GetCfCookie() string {
	req, err := http.NewRequest("GET", "https://discord.com", nil)
	if err != nil {
		logger.Fatal(err)
	}
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		logger.Fatal(err)
	}
	defer resp.Body.Close()

	respCookies := resp.Cookies()
	dcf := respCookies[0].Value
	sdc := respCookies[1].Value
	return fmt.Sprintf("__dcfduid=%s; __sdcfduid=%s; locale=en-GB", dcf, sdc)
}

// GetFingerprint Will Make a request to get fingerprint header and will return it
func GetFingerprint() string {
	req, err := http.NewRequest("GET", "https://discordapp.com/api/v9/experiments", nil)
	if err != nil {
		logger.Fatal(err)
	}
	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		logger.Fatal(err)
	}
	defer resp.Body.Close()

	// the fingerprint is in the body of the response
	var exp Experiments
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
	}
	err = json.Unmarshal(body, &exp)
	if err != nil {
		logger.Error(err)
	}
	return exp.Fingerprint
}

// CheckSingleToken Will send a request with the token and depending on which status code we get
// we can determine if it is invalid
func CheckSingleToken(token string) bool {
	request, err := http.NewRequest("GET", "https://discord.com/api/v9/users/@me/guild-events", nil)
	if err != nil {
		logger.Fatal(err)
	}
	request.Header = http.Header{
		"Authorization": []string{token},
		"Content-type":  []string{"application/json"},
	}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Do(request)
	if err != nil {
		logger.Fatal(err)
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		logger.Info("Token provided is Valid")
		return true
	case 401:
		logger.Error("Token provided is Invalid !")
		return false
	case 403:
		logger.Error("Token provided is phone locked !")
	default:
		return false
	}
	return false
}
