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

func (c *LivespaceClient) prepareAuthorizedRequest() (req model.AuthorizedRequest, err error) {
	tokenO, err := c.GetAuth()
	if err != nil {
		return req, err
	}

	req.ApiKey = c.Config.ApiKey
	req.ApiSession = tokenO.Data.SessionID
	//sha1 function on the character string created from the following concatenation: API_KEY, TOKEN,
	//and API_SECRET
	h := sha1.New()
	h.Write([]byte(c.Config.ApiKey + tokenO.Data.Token + c.Config.ApiSecret))
	shaBytes := h.Sum(nil)
	req.ApiSHA = hex.EncodeToString(shaBytes)

	return req, err
}

func (c *LivespaceClient) GetAuth() (tokenO model.Token, err error) {
	responseBytes, err := c.makeRequest("_Api/auth_call/_api_method/getToken", model.GetToken{
		ApiKey:  c.Config.ApiKey,
		ApiAuth: "key",
	}, false)
	if err != nil {
		return tokenO, err
	}

	_ = json.Unmarshal(responseBytes, &tokenO)

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
