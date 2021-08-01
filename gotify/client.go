package gotify

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	Client HTTPClient
	URL    string
}

func New(http_client HTTPClient, base_url string, token string) Client {
	if http_client == nil {
		http_client = http.DefaultClient
	}
	return Client{
		Client: http_client,
		URL:    base_url + "/message?token=" + token,
	}
}

type Message struct {
	Title    string
	Message  string
	Priority int
}

func (c Client) Notify(m Message) error {

	client := c.Client
	if client == nil {
		client = http.DefaultClient
	}

	formData := strings.NewReader(url.Values{
		"title":    {m.Title},
		"message":  {m.Message},
		"priority": {strconv.Itoa(m.Priority)},
	}.Encode())

	req, err := http.NewRequest("POST", c.URL, formData)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Golang_Gotify/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
