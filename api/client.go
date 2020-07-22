package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/detectify/dtfycli/models"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	client *http.Client
	token  string
}

func NewClient(token string, client *http.Client) *Client {
	return &Client{
		client: client,
		token:  token,
	}
}

// See: https://developer.detectify.com/#asset-inventory-manage-assets-get
func (c *Client) GetDomains() ([]*models.Domain, error) {
	req, _ := http.NewRequest("GET", "https://api.detectify.com/rest/v2/domains/", nil)
	req.Header.Set("X-Detectify-Key", c.token)
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var response []*models.Domain
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// See: https://developer.detectify.com/#asset-monitoring-get-single-finding-get
func (c *Client) GetFinding(domainToken, uuid string) (*models.Finding, error) {
	url := strings.Builder{}
	url.WriteString("https://api.detectify.com/rest/v2/domains/")
	url.WriteString(domainToken)
	url.WriteString("/findings/")
	url.WriteString(uuid)
	url.WriteString("/")
	req, _ := http.NewRequest("GET", url.String(), nil)
	req.Header.Set("X-Detectify-Key", c.token)
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var finding models.Finding
	err = json.Unmarshal(body, &finding)
	if err != nil {
		return nil, err
	}
	return &finding, nil
}

// See: https://developer.detectify.com/#asset-monitoring-get-finding-uuids-for-asset-get
func (c *Client) GetFindingUUIDs(domainToken string, from, to *time.Time, severity models.Severity) ([]string, error) {

	// Figure if any optional parameters should be sent
	params := []string{}
	switch severity {
	case models.SeverityHigh:
		params = append(params, "severity=high")
	case models.SeverityMedium:
		params = append(params, "severity=medium")
	case models.SeverityLow:
		params = append(params, "severity=low")
	case models.SeverityInformation:
		params = append(params, "severity=information")
	}
	if from != nil {
		ts := from.Unix()
		params = append(params, fmt.Sprintf("from=%d", ts))
	}
	if to != nil {
		ts := to.Unix()
		params = append(params, fmt.Sprintf("to=%d", ts))
	}

	// Build up the URI to be sent
	url := strings.Builder{}
	url.WriteString("https://api.detectify.com/rest/v2/domains/")
	url.WriteString(domainToken)
	url.WriteString("/findinguuids/")
	if len(params) > 0 {
		url.WriteString("?")
		url.WriteString(strings.Join(params, "&"))
	}

	// Build up and transmit the request
	req, _ := http.NewRequest("GET", url.String(), nil)
	req.Header.Set("X-Detectify-Key", c.token)
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var uuids []string
	err = json.Unmarshal(body, &uuids)
	if err != nil {
		return nil, err
	}
	return uuids, nil
}
