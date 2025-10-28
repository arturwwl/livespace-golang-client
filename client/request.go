package livespaceclient

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/arturwwl/livespace-golang-client/model"
	"github.com/go-playground/form"
	"io"
	"net/http"
	"strings"
)

func (c *LivespaceClient) getUrl(path string) string {
	if c.Config.IsProd {
		return fmt.Sprintf("https://%s.livespace.io/api/public/json/%s", c.Config.Subdomain, path)
	}

	return fmt.Sprintf("%s/%s", c.Config.DevUrl, path)
}

func (c *LivespaceClient) prepareAuthorizedRequest() (model.AuthorizedRequest, error) {
	tokenO, err := c.GetAuth()
	if err != nil {
		return model.AuthorizedRequest{}, err
	}

	req := model.AuthorizedRequest{
		ApiKey:     c.Config.ApiKey,
		ApiSession: tokenO.Data.SessionID,
	}

	//sha1 function on the character string created from the following concatenation: API_KEY, TOKEN,
	//and API_SECRET
	h := sha1.New()
	_, err = h.Write([]byte(c.Config.ApiKey + tokenO.Data.Token + c.Config.ApiSecret))
	if err != nil {
		return model.AuthorizedRequest{}, err
	}
	shaBytes := h.Sum(nil)
	req.ApiSHA = hex.EncodeToString(shaBytes)

	return req, nil
}

func (c *LivespaceClient) GetAuth() (model.Token, error) {
	var tokenO model.Token
	var err error

	responseBytes, err := c.makeRequest("_Api/auth_call/_api_method/getToken", model.GetToken{
		ApiKey:  c.Config.ApiKey,
		ApiAuth: "key",
	}, false)
	if err != nil {
		return tokenO, err
	}

	err = json.Unmarshal(responseBytes, &tokenO)
	if err != nil {
		return tokenO, err
	}

	if !tokenO.Status { //invalid response
		return tokenO, fmt.Errorf("error result %d", tokenO.Result)
	}

	if tokenO.Data == nil || tokenO.Data.Token == "" {
		return tokenO, fmt.Errorf("no token data in response")
	}

	return tokenO, nil
}

func (c *LivespaceClient) makeRequest(path string, data interface{}, isJson bool) ([]byte, error) {
	var requestBody io.Reader
	var err error
	var req *http.Request

	if data != nil {
		if isJson {
			postBody, _ := json.Marshal(data)
			requestBody = bytes.NewBuffer(postBody)
		} else {
			encoder := form.NewEncoder()
			values, _ := encoder.Encode(&data)
			requestBody = strings.NewReader(values.Encode())
		}
	}

	req, err = http.NewRequest(http.MethodPost, c.getUrl(path), requestBody)
	if err != nil {
		return nil, err
	}

	if isJson {
		req.Header.Add("Content-type", "application/json")
	} else {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 204 {
		return nil, fmt.Errorf("invalid resposne status code")
	}

	return bodyBytes, err
}
