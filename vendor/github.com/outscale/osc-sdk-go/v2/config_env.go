package osc

import (
	"context"
	"crypto/tls"
	b64 "encoding/base64"
	"errors"
	"net/http"
	"os"
)

type ConfigEnv struct {
	AccessKey           *string
	SecretKey           *string
	OutscaleApiEndpoint *string
	ProfileName         *string
	Region              *string
	X509ClientCert      *string
	X509ClientCertB64   *string
	X509ClientKey       *string
	X509ClientKeyB64    *string
}

func NewConfigEnv() *ConfigEnv {
	var configEnv ConfigEnv
	if value, present := os.LookupEnv("OSC_ACCESS_KEY"); present {
		configEnv.AccessKey = &value
	}
	if value, present := os.LookupEnv("OSC_SECRET_KEY"); present {
		configEnv.SecretKey = &value
	}
	if value, present := os.LookupEnv("OSC_ENDPOINT_API"); present {
		configEnv.OutscaleApiEndpoint = &value
	}
	if value, present := os.LookupEnv("OSC_PROFILE"); present {
		configEnv.ProfileName = &value
	}
	if value, present := os.LookupEnv("OSC_REGION"); present {
		configEnv.Region = &value
	}
	if value, present := os.LookupEnv("OSC_X509_CLIENT_CERT"); present {
		configEnv.X509ClientCert = &value
	}
	if value, present := os.LookupEnv("OSC_X509_CLIENT_CERT_B64"); present {
		configEnv.X509ClientCertB64 = &value
	}
	if value, present := os.LookupEnv("OSC_X509_CLIENT_KEY"); present {
		configEnv.X509ClientKey = &value
	}
	if value, present := os.LookupEnv("OSC_X509_CLIENT_KEY_B64"); present {
		configEnv.X509ClientKeyB64 = &value
	}
	return &configEnv
}

func (configEnv *ConfigEnv) Configuration() (*Configuration, error) {
	var config *Configuration

	if configEnv.AccessKey == nil && configEnv.SecretKey == nil {
		if configEnv.ProfileName == nil {
			value := "default"
			configEnv.ProfileName = &value
		}
		configFile, err := LoadDefaultConfigFile()
		if err != nil {
			return nil, err
		}
		config, err = configFile.Configuration(*configEnv.ProfileName)
		if err != nil {
			return nil, err
		}
	} else {
		config = NewConfiguration()
		default_region := "eu-west-2"
		config.Servers = ServerConfigurations{
			{
				URL:         "https://api.{region}.outscale.com/api/v1",
				Description: "Loaded from profile",
				Variables: map[string]ServerVariable{
					"region": ServerVariable{
						Description:  "Loaded from env variables",
						DefaultValue: default_region,
						EnumValues:   []string{default_region},
					},
				},
			},
		}
	}

	// Overload with remaining environement variable values
	if configEnv.OutscaleApiEndpoint != nil {
		config.Servers[0].URL = *configEnv.OutscaleApiEndpoint
	}

	if configEnv.Region != nil && len(config.Servers) > 0 {
		config.Servers[0].Variables["region"] = ServerVariable{
			Description:  "Loaded from env variables",
			DefaultValue: *configEnv.Region,
			EnumValues:   []string{*configEnv.Region},
		}
	}

	tlsConfigured := false
	if configEnv.X509ClientCert != nil && configEnv.X509ClientKey != nil {
		tlsConfigured = true
		cert, err := tls.LoadX509KeyPair(*configEnv.X509ClientCert, *configEnv.X509ClientKey)
		if err != nil {
			return nil, errors.New("error while loading client certificate and key")
		}

		tlsconfig := &tls.Config{
			InsecureSkipVerify: false,
			Certificates:       []tls.Certificate{cert},
		}

		httpClient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsconfig,
				Proxy:           http.ProxyFromEnvironment,
			},
		}

		config.HTTPClient = httpClient
	}

	if configEnv.X509ClientCertB64 != nil && configEnv.X509ClientKeyB64 != nil {
		if tlsConfigured {
			return nil, errors.New("cannot configure client certificate with both file and base64")
		}

		clientCertificate, err := b64.StdEncoding.DecodeString(*configEnv.X509ClientCertB64)
		if err != nil {
			return nil, errors.New("error while decoding client certificate from base64")
		}

		clientKey, err := b64.StdEncoding.DecodeString(*configEnv.X509ClientKeyB64)
		if err != nil {
			return nil, errors.New("error while decoding client key from base64")
		}
		cert, err := tls.X509KeyPair(clientCertificate, clientKey)
		if err != nil {
			return nil, errors.New("error while loading client certificate and key")
		}

		tlsconfig := &tls.Config{
			InsecureSkipVerify: false,
			Certificates:       []tls.Certificate{cert},
		}

		httpClient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsconfig,
				Proxy:           http.ProxyFromEnvironment,
			},
		}

		config.HTTPClient = httpClient
	}

	return config, nil
}

func (configEnv *ConfigEnv) Context(ctx context.Context) (context.Context, error) {
	var accessKey *string
	var secretKey *string

	if configEnv.AccessKey == nil && configEnv.SecretKey == nil {
		if configEnv.ProfileName == nil {
			value := "default"
			configEnv.ProfileName = &value
		}
		configFile, err := LoadDefaultConfigFile()
		if err != nil {
			return nil, err
		}
		profile, ok := configFile.profiles[*configEnv.ProfileName]
		if !ok {
			return nil, errors.New("profile not found for creating Context, did you load config file?")
		}
		accessKey = &profile.AccessKey
		secretKey = &profile.SecretKey
	}
	// Overload with environement variable values
	if configEnv.AccessKey != nil {
		accessKey = configEnv.AccessKey
	}
	if configEnv.SecretKey != nil {
		secretKey = configEnv.SecretKey
	}

	if accessKey == nil || len(*accessKey) == 0 || secretKey == nil || len(*secretKey) == 0 {
		return ctx, nil
	}
	ctx = context.WithValue(ctx, ContextAWSv4, AWSv4{
		AccessKey: *accessKey,
		SecretKey: *secretKey,
	})
	return ctx, nil
}
