package client

import (
	"com.azure.iot/iotcentral/iotcgo/config"
	"fmt"

	"com.azure.iot/iotcentral/iotcgo/client/operations"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type (
	// Token containing client credentials
	Token struct {
		Subdomain  string
		Credential string
		Host       string
	}
)

const (
	// DefaultTokenHost IoT Central host
	DefaultTokenHost string = "azureiotcentral.com"
)

// NewFromToken creates a new client for a given app name.
func NewFromToken(appName string) (*AzureIoTCentral, error) {
	appConfig, err := config.GetAppConfig(appName)
	if err != nil {
		return nil, err
	}

	host := DefaultTokenHost
	if appConfig.Domain != "" {
		host = appConfig.Domain
	}
	t := Token{Subdomain: appConfig.Subdomain, Credential: appConfig.ApiToken}
	if t.Host != "" {
		host = t.Host
	}

	transport := httptransport.New(
		fmt.Sprintf("%s.%s", t.Subdomain, host),
		DefaultBasePath,
		nil,
	)

	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", t.Credential)
	cli := new(AzureIoTCentral)
	cli.Transport = transport
	cli.Operations = operations.New(transport, strfmt.Default)
	return cli, nil
}

func GetURL(appName string) (string, error) {
	appConfig, err := config.GetAppConfig(appName)
	if err != nil {
		return "", err
	}

	t := Token{Subdomain: appConfig.Subdomain, Credential: appConfig.ApiToken}
	host := DefaultTokenHost
	if t.Host != "" {
		host = t.Host
	}

	url := fmt.Sprintf("https://%s.%s%s", t.Subdomain, host, DefaultBasePath)
	return url, nil
}

