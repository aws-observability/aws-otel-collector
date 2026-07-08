package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const (
	oidcPath = "/v2/oidc"
)

// OIDCService is the interface to interact with the OIDC endpoints on the Vultr API
type OIDCService interface {
	CreateOIDCProvider(ctx context.Context, oidcProviderReq *OIDCProviderReq) (*OIDCProvider, *http.Response, error)
	GetOIDCProvider(ctx context.Context, oidcProviderID string) (*OIDCProvider, *http.Response, error)
	ListOIDCProviders(ctx context.Context) ([]OIDCProvider, *http.Response, error)
	DeleteOIDCProvider(ctx context.Context, oidcProviderID string) error

	CreateOIDCIssuer(ctx context.Context, oidcIssuerReq *OIDCIssuerReq) (*OIDCIssuer, *http.Response, error)
	GetOIDCIssuer(ctx context.Context, oidcIssuerID string) (*OIDCIssuer, *http.Response, error)
	ListOIDCIssuers(ctx context.Context) ([]OIDCIssuer, *http.Response, error)
	DeleteOIDCIssuer(ctx context.Context, oidcIssuerID string) error

	CreateOIDCToken(ctx context.Context, oidcTokenReq *OIDCTokenReq) (*OIDCToken, *http.Response, error)
	GetOIDCToken(ctx context.Context, oidcTokenID string) (*OIDCToken, *http.Response, error)

	AuthorizeOIDC(ctx context.Context, oidcAuthParams *OIDCAuthorizeParameters) error
	DiscoveryOIDC(ctx context.Context, oidcProviderID string) (*OIDCDocument, *http.Response, error)
}

// OIDCServiceHandler handles interaction with the OIDC methods for the Vultr API
type OIDCServiceHandler struct {
	client *Client
}

// OIDCProvider represents an OIDC provider
type OIDCProvider struct {
	ID       string `json:"id"`
	IssuerID string `json:"issuer_id"`
	Name     string `json:"name"`
}

type oidcProvidersBase struct {
	Providers []OIDCProvider `json:"providers"`
}

type oidcProviderBase struct {
	Provider *OIDCProvider `json:"provider"`
}

// OIDCProviderReq represnts a request for OIDC provider creation
type OIDCProviderReq struct {
	Name string `json:"name"`
}

// OIDCIssuer represents an OIDC issuer
type OIDCIssuer struct {
	ID              string `json:"id"`
	Source          string `json:"source"`
	URI             string `json:"uri"`
	KTY             string `json:"kty"`
	KID             string `json:"kid"`
	N               string `json:"n"`
	E               string `json:"e"`
	ALG             string `json:"alg"`
	USE             string `json:"use"`
	JWKSFetchedDate string `json:"jwks_fetched_at"`
	JWKSExpiryDate  string `json:"jwks_expires_at"`
}

type oidcIssuersBase struct {
	Issuers []OIDCIssuer `json:"issuers"`
}

type oidcIssuerBase struct {
	Issuer *OIDCIssuer `json:"issuer"`
}

// OIDCIssuerReq represents a request for an OIDC issuer creation
type OIDCIssuerReq struct {
	Issuer OIDCIssuerReqDetail `json:"issuer"`
}

// OIDCIssuerReqDetail represents the detail for the OIDC issuer request
type OIDCIssuerReqDetail struct {
	Source   string `json:"source"`
	SourceID string `json:"source_id,omitempty"`
	URI      string `json:"uri,omitempty"`
	KTY      string `json:"kty,omitempty"`
	KID      string `json:"kid,omitempty"`
	N        string `json:"n,omitempty"`
	E        string `json:"e,omitempty"`
	ALG      string `json:"alg,omitempty"`
	USE      string `json:"use,omitempty"`
}

// OIDCDocument represents the OIDC document for a provider
type OIDCDocument struct {
	Issuer                            string   `json:"issuer"`
	AuthorizeEndpoint                 string   `json:"authorization_endpoint"`
	TokenEndpoint                     string   `json:"token_endpoint"`
	JWKSURI                           string   `json:"jwks_uri"`
	UserInfoEndpoint                  string   `json:"userinfo_endpoint"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	SubjectTypesSupported             []string `json:"subject_types_supported"`
	IDTokenSigningValuesSupported     []string `json:"id_token_signing_alg_values_supported"`
	ScopesSupported                   []string `json:"scopes_supported"`
	ClaimsSupported                   []string `json:"claims_supported"`
	GrantTypesSupported               []string `json:"grant_types_supported"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
}

// OIDCToken represents an OIDC token
type OIDCToken struct {
	AccessToken    string `json:"access_token"`
	TokenType      string `json:"token_type"`
	ExpiresSeconds string `json:"expires_in"`
	RefreshToken   string `json:"refresh_token"`
	IDToken        string `json:"id_token"`
	Scope          string `json:"scope"`
}

type oidcTokenBase struct {
	Token *OIDCToken `json:"token"`
}

// OIDCTokenReq represents an OIDC token request
type OIDCTokenReq struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code,omitempty"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// OIDCJWTToken represents an OIDC JWT token
type OIDCJWTToken struct {
	ISS string          `json:"iss"`
	SUB string          `json:"sub"`
	AUD string          `json:"aud"`
	EXP string          `json:"exp"`
	NBF string          `json:"nbf"`
	IAT string          `json:"iat"`
	JTI string          `json:"jti"`
	JWK OIDCJWTTokenJWK `json:"jwk"`
}

// OIDCJWTTokenJWK represents an OIDC JWT token JWK
type OIDCJWTTokenJWK struct {
	KID string `json:"kid"`
	KTY string `json:"kty"`
	N   string `json:"n"`
	E   string `json:"e"`
	ALG string `json:"alg"`
	USE string `json:"use"`
}

// OIDCAuthorizeParameters represents the query parameters used for OIDC authorize
type OIDCAuthorizeParameters struct {
	ClientID     string `url:"client_id"`
	RedirectURI  string `url:"redirect_uri"`
	ResponseType string `url:"response_type"`
	Scope        string `url:"scope,omitempty"`
	State        string `url:"state,omitempty"`
	Nonce        string `url:"nonce,omitempty"`
	ResponseMode string `url:"response_mode,omitempty"`
}

func (o *OIDCServiceHandler) CreateOIDCProvider(ctx context.Context, oidcProviderReq *OIDCProviderReq) (*OIDCProvider, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/provider", oidcPath)

	req, err := o.client.NewRequest(ctx, http.MethodPost, uri, oidcProviderReq)
	if err != nil {
		return nil, nil, err
	}

	provider := new(oidcProviderBase)
	resp, err := o.client.DoWithContext(ctx, req, provider)
	if err != nil {
		return nil, resp, err
	}

	return provider.Provider, resp, nil
}

func (o *OIDCServiceHandler) GetOIDCProvider(ctx context.Context, oidcProviderID string) (*OIDCProvider, *http.Response, error) {
	uri := fmt.Sprintf("%s/provider/%s", oidcPath, oidcProviderID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	provider := new(oidcProviderBase)
	resp, err := o.client.DoWithContext(ctx, req, provider)
	if err != nil {
		return nil, resp, err
	}

	return provider.Provider, resp, nil
}

func (o *OIDCServiceHandler) ListOIDCProviders(ctx context.Context) ([]OIDCProvider, *http.Response, error) {
	uri := fmt.Sprintf("%s/provider", oidcPath)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	providers := new(oidcProvidersBase)
	resp, err := o.client.DoWithContext(ctx, req, providers)
	if err != nil {
		return nil, resp, err
	}

	return providers.Providers, resp, nil
}

func (o *OIDCServiceHandler) DeleteOIDCProvider(ctx context.Context, oidcProviderID string) error {
	uri := fmt.Sprintf("%s/provider/%s", oidcPath, oidcProviderID)

	req, err := o.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	if _, err = o.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (o *OIDCServiceHandler) CreateOIDCIssuer(ctx context.Context, oidcIssuerReq *OIDCIssuerReq) (*OIDCIssuer, *http.Response, error) {
	uri := fmt.Sprintf("%s/issuer", oidcPath)

	req, err := o.client.NewRequest(ctx, http.MethodPost, uri, oidcIssuerReq)
	if err != nil {
		return nil, nil, err
	}

	issuer := new(oidcIssuerBase)
	resp, err := o.client.DoWithContext(ctx, req, issuer)
	if err != nil {
		return nil, resp, err
	}

	return issuer.Issuer, resp, nil
}

func (o *OIDCServiceHandler) GetOIDCIssuer(ctx context.Context, oidcIssuerID string) (*OIDCIssuer, *http.Response, error) {
	uri := fmt.Sprintf("%s/issuer/%s", oidcPath, oidcIssuerID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	issuer := new(oidcIssuerBase)
	resp, err := o.client.DoWithContext(ctx, req, issuer)
	if err != nil {
		return nil, resp, err
	}

	return issuer.Issuer, resp, nil
}

func (o *OIDCServiceHandler) ListOIDCIssuers(ctx context.Context) ([]OIDCIssuer, *http.Response, error) {
	uri := fmt.Sprintf("%s/issuer", oidcPath)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	issuers := new(oidcIssuersBase)
	resp, err := o.client.DoWithContext(ctx, req, issuers)
	if err != nil {
		return nil, resp, err
	}

	return issuers.Issuers, resp, nil
}

func (o *OIDCServiceHandler) DeleteOIDCIssuer(ctx context.Context, oidcIssuerID string) error {
	uri := fmt.Sprintf("%s/issuer/%s", oidcPath, oidcIssuerID)

	req, err := o.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	if _, err = o.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (o *OIDCServiceHandler) CreateOIDCToken(ctx context.Context, oidcTokenReq *OIDCTokenReq) (*OIDCToken, *http.Response, error) {
	uri := fmt.Sprintf("%s/issuer/token", oidcPath)

	req, err := o.client.NewRequest(ctx, http.MethodPost, uri, oidcTokenReq)
	if err != nil {
		return nil, nil, err
	}

	token := new(oidcTokenBase)
	resp, err := o.client.DoWithContext(ctx, req, token)
	if err != nil {
		return nil, resp, err
	}

	return token.Token, resp, nil
}
func (o *OIDCServiceHandler) GetOIDCToken(ctx context.Context, oidcTokenID string) (*OIDCToken, *http.Response, error) {
	uri := fmt.Sprintf("%s/issuer/token/%s", oidcPath, oidcTokenID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	token := new(oidcTokenBase)
	resp, err := o.client.DoWithContext(ctx, req, token)
	if err != nil {
		return nil, resp, err
	}

	return token.Token, resp, nil
}

func (o *OIDCServiceHandler) AuthorizeOIDC(ctx context.Context, oidcAuthParams *OIDCAuthorizeParameters) error {
	uri := fmt.Sprintf("%s/authorize", oidcPath)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return err
	}

	newValues, err := query.Values(oidcAuthParams)
	if err != nil {
		return err
	}

	req.URL.RawQuery = newValues.Encode()

	if _, err = o.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (o *OIDCServiceHandler) DiscoveryOIDC(ctx context.Context, oidcProviderID string) (*OIDCDocument, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/.well-known/openid-configuration", oidcPath, oidcProviderID)

	req, err := o.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	doc := new(OIDCDocument)
	resp, err := o.client.DoWithContext(ctx, req, doc)
	if err != nil {
		return nil, resp, err
	}

	return doc, resp, nil
}
