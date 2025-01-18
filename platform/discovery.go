package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LinodeAccount struct {
	Email string `json:"email"`
	Euuid string `json:"euuid"`
	State string `json:"state"`
}

// LinodeIntegrationDiscovery fetches Linode account details using the provided token.
func LinodeIntegrationDiscovery(token string) (*LinodeAccount, error) {
	const linodeAPIURL = "https://api.linode.com/v4/account"

	client := &http.Client{}
	req, err := http.NewRequest("GET", linodeAPIURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set the Authorization header with the token.
	req.Header.Set("Authorization", "Bearer "+token)

	// Perform the request.
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status is not OK.
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to fetch account: %s, status: %d", string(body), resp.StatusCode)
	}

	// Decode the JSON response.
	var account LinodeAccount
	if err := json.NewDecoder(resp.Body).Decode(&account); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &account, nil
}
