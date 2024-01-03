package nextsms

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var (
	singleURL   = "https://messaging-service.co.tz/api/sms/v1/text/single"
	multipleURL = "https://messaging-service.co.tz/api/sms/v1/text/multiple"
)

type Client struct {
	http *http.Client
	auth string
}

func New() *Client {
	encoded := authorizationHeader()
	return &Client{
		http: http.DefaultClient,
		auth: encoded,
	}
}

// authorizationHeader generates the authorization string
// for accessing the api endpoints
func authorizationHeader() string {
	username := os.Getenv("NEXT_USERNAME")
	password := os.Getenv("NEXT_PASSWORD")

	str := fmt.Sprintf("%s:%s", username, password)
	encoded := base64.StdEncoding.EncodeToString([]byte(str))

	encoded = fmt.Sprintf("Basic %s", encoded)
	return encoded
}

// SendSSMS sends a single sms to a particular phone number provided
// for defaults senderid, put from ""
func (c *Client) SendSSMS(to, message, from string) (*http.Response, error) {
	if from == "" {
		from = "N-SMS"
	}

	payload := strings.NewReader(fmt.Sprintf(`{"from":%s, "to":%s,  "text": %s,}`, from, to, message))
	req, err := http.NewRequest("POST", singleURL, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", c.auth)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// TODO: Implement multiple messaging
