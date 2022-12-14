package ftdClient

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// FTDURL - Default Hashicups URL
const FTDURL string = "https://127.0.0.1"

// Client -
type Client struct {
	FTDURL       string
	HTTPClient   *http.Client
	AuthResponse AuthResponse
	AuthTime     time.Time
	Auth         AuthRequest
}

// AuthStruct -
type AuthRequest struct {
	GrantType string `json:"grant_type"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

// AuthResponse -
type AuthResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	TokenType        string `json:"token_type"`
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
}

// NewClient -
func NewClient(host, username, password *string) (*Client, error) {

	// for test porpuses
	insecureTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c := Client{
		HTTPClient: &http.Client{Timeout: 20 * time.Second, Transport: insecureTransport},
		FTDURL:     FTDURL,
	}

	if host != nil {
		c.FTDURL = *host
	}

	// If username or password not provided, return empty client
	if username == nil || password == nil {
		return &c, nil
	}

	c.Auth = AuthRequest{
		GrantType: "password",
		Username:  *username,
		Password:  *password,
	}

	ar, err := c.LogIn()
	if err != nil {
		return nil, err
	}
	c.AuthTime = time.Now()
	c.AuthResponse = *ar

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	token := c.AuthResponse.AccessToken

	if c.AuthTime.Add(time.Second*time.Duration(c.AuthResponse.ExpiresIn)).UnixMicro() < time.Now().UnixMicro() {
		c.LogIn() //switch to token renew later
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && (res.StatusCode != http.StatusNoContent && req.Method != "DELETE") {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

// func doFTDRequest[T FtdModel](m *T, name string, method string, c *Client) error {
func doFTDRequest[T any](m *T, name string, method string, c *Client) error {
	URL := fmt.Sprintf("%s/api/fdm/v6/%s", c.FTDURL, name)
	//fmt.Printf("Method: %s %s \n", method, URL)
	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	//remove empty structs from marshalled json. Need other json marshaler to remove empty (default values) struct
	mustc := regexp.MustCompile(`,"([a-zA-Z])*":{}`)
	for i := 1; i < 3; i++ {
		rb = []byte(mustc.ReplaceAllString(string(rb), ""))
	}
	fmt.Println(string(rb))

	req, err := http.NewRequest(method, URL, strings.NewReader(string(rb)))
	if err != nil {
		fmt.Println(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	//fmt.Println(string(body)) //for debug usage
	// Cisco FTD does not return data on delete
	if method != "DELETE" {
		err = json.Unmarshal(body, m)
		if err != nil {
			return err
		}
	}

	return nil
}
