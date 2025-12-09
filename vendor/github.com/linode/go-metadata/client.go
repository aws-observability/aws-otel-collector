package metadata

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	APIHost    = "169.254.169.254"
	APIProto   = "http"
	APIVersion = "v1"
)

type Logger = resty.Logger

// Client represents an instance of a Linode Metadata Service client.
type Client struct {
	managedTokenExpiry time.Time
	resty              *resty.Client
	apiBaseURL         string
	apiProtocol        string
	apiVersion         string
	userAgent          string
	managedTokenOpts   []TokenOption
	managedToken       bool
}

// NewClient creates a new Metadata API client configured
// with the given options.
func NewClient(ctx context.Context, opts ...ClientOption) (*Client, error) {
	clientOpts := clientCreateConfig{
		HTTPClient:      nil,
		BaseURLOverride: "",
		VersionOverride: "",
		UserAgentPrefix: "",
		ManagedToken:    true,
		StartingToken:   "",
	}

	for _, opt := range opts {
		opt(&clientOpts)
	}

	var result Client

	result.managedToken = clientOpts.ManagedToken
	result.managedTokenOpts = clientOpts.ManagedTokenOpts

	userAgent := DefaultUserAgent

	if clientOpts.BaseURLOverride != "" {
		result.SetBaseURL(clientOpts.BaseURLOverride)
	}

	if clientOpts.VersionOverride != "" {
		result.SetVersion(clientOpts.VersionOverride)
	}

	if clientOpts.UserAgentPrefix != "" {
		userAgent = fmt.Sprintf("%s %s", clientOpts.UserAgentPrefix, userAgent)
	}

	if clientOpts.HTTPClient != nil {
		result.resty = resty.NewWithClient(clientOpts.HTTPClient)
	} else {
		result.resty = resty.New()
	}

	if debugEnv, ok := os.LookupEnv("LINODE_DEBUG"); ok {
		debugBool, err := strconv.ParseBool(debugEnv)
		if err != nil {
			return nil, fmt.Errorf("failed to parse debug bool: %s", err)
		}
		result.resty.SetDebug(debugBool)
	}

	result.updateHostURL()
	result.setUserAgent(userAgent)

	if clientOpts.ManagedToken && clientOpts.StartingToken == "" {
		if _, err := result.RefreshToken(ctx, result.managedTokenOpts...); err != nil {
			return nil, fmt.Errorf("failed to refresh metadata token: %s", err)
		}
	} else if clientOpts.StartingToken != "" {
		result.UseToken(clientOpts.StartingToken)
	}

	result.resty.SetDebug(clientOpts.Debug)
	if clientOpts.DebugLogger != nil {
		result.resty.SetLogger(clientOpts.DebugLogger)
	}

	result.resty.OnBeforeRequest(result.middlewareTokenRefresh)

	return &result, nil
}

// SetBaseURL configures the target URL for metadata API this client accesses.
func (c *Client) SetBaseURL(baseURL string) *Client {
	baseURLPath, _ := url.Parse(baseURL)

	c.apiBaseURL = path.Join(baseURLPath.Host, baseURLPath.Path)
	c.apiProtocol = baseURLPath.Scheme

	c.updateHostURL()

	return c
}

// SetVersion configures the target metadata API version for this client.
func (c *Client) SetVersion(version string) *Client {
	c.apiVersion = version

	c.updateHostURL()

	return c
}

func (c *Client) updateHostURL() {
	apiProto := APIProto
	baseURL := APIHost
	apiVersion := APIVersion

	if c.apiBaseURL != "" {
		baseURL = c.apiBaseURL
	}

	if c.apiVersion != "" {
		apiVersion = c.apiVersion
	}

	if c.apiProtocol != "" {
		apiProto = c.apiProtocol
	}

	c.resty.SetBaseURL(fmt.Sprintf("%s://%s/%s", apiProto, baseURL, apiVersion))
}

// middlewareTokenRefresh handles automatically refreshing managed tokens.
func (c *Client) middlewareTokenRefresh(_ *resty.Client, r *resty.Request) error {
	// Don't run this middleware when generating tokens
	if r.URL == "token" {
		return nil
	}

	if !c.managedToken || time.Now().Before(c.managedTokenExpiry) {
		return nil
	}

	// Token needs to be refreshed
	if _, err := c.RefreshToken(r.Context(), c.managedTokenOpts...); err != nil {
		return err
	}

	return nil
}

// R wraps resty's R method
func (c *Client) R(ctx context.Context) *resty.Request {
	return c.resty.R().
		ExpectContentType("application/json").
		SetHeader("Content-Type", "application/json").
		SetContext(ctx).
		SetError(APIError{})
}

func (c *Client) setUserAgent(userAgent string) *Client {
	c.userAgent = userAgent
	c.resty.SetHeader("User-Agent", c.userAgent)

	return c
}

type clientCreateConfig struct {
	HTTPClient *http.Client

	BaseURLOverride string
	VersionOverride string

	UserAgentPrefix string

	ManagedToken     bool
	ManagedTokenOpts []TokenOption

	StartingToken string

	Debug       bool
	DebugLogger Logger
}

// ClientOption is an option that can be used
// during client creation.
type ClientOption func(options *clientCreateConfig)

// ClientWithHTTPClient configures the underlying HTTP client
// to communicate with the Metadata API.
func ClientWithHTTPClient(client *http.Client) ClientOption {
	return func(options *clientCreateConfig) {
		options.HTTPClient = client
	}
}

// ClientWithBaseURL configures the target host of the
// Metadata API this client points to.
// Default: "169.254.169.254"
func ClientWithBaseURL(baseURL string) ClientOption {
	return func(options *clientCreateConfig) {
		options.BaseURLOverride = baseURL
	}
}

// ClientWithVersion configures the Metadata API version this
// client should target.
// Default: "v1"
func ClientWithVersion(version string) ClientOption {
	return func(options *clientCreateConfig) {
		options.VersionOverride = version
	}
}

// ClientWithUAPrefix configures the prefix for user agents
// on API requests made by this client.
func ClientWithUAPrefix(uaPrefix string) ClientOption {
	return func(options *clientCreateConfig) {
		options.UserAgentPrefix = uaPrefix
	}
}

// ClientWithManagedToken configures the metadata client
// to automatically generate and refresh the API token
// for the Metadata client.
func ClientWithManagedToken(opts ...TokenOption) ClientOption {
	return func(options *clientCreateConfig) {
		options.ManagedToken = true
		options.ManagedTokenOpts = opts
	}
}

// ClientWithoutManagedToken configures the metadata client
// to disable automatic token management.
func ClientWithoutManagedToken() ClientOption {
	return func(options *clientCreateConfig) {
		options.ManagedToken = false
	}
}

// ClientWithToken configures the starting token
// for the metadata client.
// If this option is specified and managed tokens
// are enabled for a client, the client will not
// generate an initial Metadata API token.
func ClientWithToken(token string) ClientOption {
	return func(options *clientCreateConfig) {
		options.StartingToken = token
	}
}

// ClientWithDebug enables debug mode for the metadata client.
// If this option is specified, all metadata service API requests
// and responses will be written to the client logger (default: stderr).
//
// To override the client logger, refer to ClientWithLogger.
func ClientWithDebug() ClientOption {
	return func(options *clientCreateConfig) {
		options.Debug = true
	}
}

// ClientWithLogger specifies the logger that should be used
// when outputting debug logs. The logger argument should implement
// the Logger interface.
// Default: stderr
func ClientWithLogger(logger Logger) ClientOption {
	return func(options *clientCreateConfig) {
		options.DebugLogger = logger
	}
}
