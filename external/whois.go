package external

import (
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

type WhoisClient struct{}

func NewWhoisClient() *WhoisClient {
	return &WhoisClient{}
}

func (c *WhoisClient) GetWhoisInfo(domain string) (string, error) {
	domainInfo, err := whois.Whois(domain)
	if err != nil {
		return "", err
	}

	return domainInfo, nil

}

func ParseWhoisData(data string) (whoisparser.WhoisInfo, error) {
	return whoisparser.Parse(data)
}
