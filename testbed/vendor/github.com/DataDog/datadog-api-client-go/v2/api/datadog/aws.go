// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadog

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

type Credentials struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

// SigningData is the data structure that represents the Data used to generate and AWS Proof
type SigningData struct {
	HeadersEncoded string `json:"iam_headers_encoded"`
	BodyEncoded    string `json:"iam_body_encoded"`
	URLEncoded     string `json:"iam_url_encoded"`
	Method         string `json:"iam_method"`
}

const (
	// Common Headers
	// orgIdHeader is the header we use to specify the name of the org we request a token for
	orgIdHeader     = "x-ddog-org-id"
	hostHeader      = "host"
	applicationForm = "application/x-www-form-urlencoded; charset=utf-8"

	// AWS specific constants
	AWSAccessKeyIdName     = "AWS_ACCESS_KEY_ID"
	AWSSecretAccessKeyName = "AWS_SECRET_ACCESS_KEY"
	AWSSessionTokenName    = "AWS_SESSION_TOKEN"

	amzDateHeader         = "X-Amz-Date"
	amzTokenHeader        = "X-Amz-Security-Token"
	amzDateFormat         = "20060102"
	amzDateTimeFormat     = "20060102T150405Z"
	defaultRegion         = "us-east-1"
	defaultStsHost        = "sts.amazonaws.com"
	regionalStsHost       = "sts.%s.amazonaws.com"
	service               = "sts"
	algorithm             = "AWS4-HMAC-SHA256"
	aws4Request           = "aws4_request"
	getCallerIdentityBody = "Action=GetCallerIdentity&Version=2011-06-15"
)

const ProviderAWS = "aws"

type AWSAuth struct {
	AwsRegion string
}

func (a *AWSAuth) Authenticate(ctx context.Context, config *DelegatedTokenConfig) (*DelegatedTokenCredentials, error) {
	// Get local AWS Credentials
	creds := a.GetCredentials(ctx)

	if config == nil || config.OrgUUID == "" {
		return nil, fmt.Errorf("missing org UUID in config")
	}

	// Use the credentials to generate the signing data
	data, err := a.GenerateAwsAuthData(config.OrgUUID, creds)
	if err != nil {
		return nil, err
	}

	// Generate the auth string passed to the token endpoint
	authString := data.BodyEncoded + "|" + data.HeadersEncoded + "|" + data.Method + "|" + data.URLEncoded

	authResponse, err := GetDelegatedToken(ctx, config.OrgUUID, authString)
	return authResponse, err
}

func (a *AWSAuth) GetCredentials(ctx context.Context) *Credentials {
	keys := ctx.Value(ContextAWSVariables)
	if keys != nil {
		keysMap := keys.(map[string]string)
		creds := Credentials{}
		if accessKey, ok := keysMap[AWSAccessKeyIdName]; ok {
			creds.AccessKeyID = accessKey
		}
		if secretKey, ok := keysMap[AWSSecretAccessKeyName]; ok {
			creds.SecretAccessKey = secretKey
		}
		if sessionToken, ok := keysMap[AWSSessionTokenName]; ok {
			creds.SessionToken = sessionToken
		}
		return &creds
	} else {
		accessKey := os.Getenv(AWSAccessKeyIdName)
		secretKey := os.Getenv(AWSSecretAccessKeyName)
		sessionToken := os.Getenv(AWSSessionTokenName)
		return &Credentials{
			AccessKeyID:     accessKey,
			SecretAccessKey: secretKey,
			SessionToken:    sessionToken,
		}
	}
}

func (a *AWSAuth) getConnectionParameters() (string, string, string) {
	region := a.AwsRegion
	var host string
	// Default to the default global STS Host (see here: https://docs.aws.amazon.com/general/latest/gr/sts.html)
	if region == "" {
		region = defaultRegion
		host = defaultStsHost
	} else {
		// If the region is not empty, use the regional STS host
		host = fmt.Sprintf(regionalStsHost, region)
	}
	stsFullURL := fmt.Sprintf("https://%s", host)
	return stsFullURL, region, host
}

func (a *AWSAuth) GenerateAwsAuthData(orgUUID string, creds *Credentials) (*SigningData, error) {
	if orgUUID == "" {
		return nil, fmt.Errorf("missing org UUID")
	}
	if creds == nil || (creds.AccessKeyID == "" && creds.SecretAccessKey == "") || creds.SessionToken == "" {
		return nil, fmt.Errorf("missing AWS credentials")
	}
	stsFullURL, region, host := a.getConnectionParameters()

	now := time.Now().UTC()

	requestBody := getCallerIdentityBody
	h := sha256.Sum256([]byte(requestBody))
	payloadHash := hex.EncodeToString(h[:])

	// Create the headers that factor into the signing algorithm
	headerMap := map[string][]string{
		contextLengthHeader: {
			fmt.Sprintf("%d", len(requestBody)),
		},
		contentTypeHeader: {
			applicationForm,
		},
		amzDateHeader: {
			now.Format(amzDateTimeFormat),
		},
		orgIdHeader: {
			orgUUID,
		},
		amzTokenHeader: {
			creds.SessionToken,
		},
		hostHeader: {
			host,
		},
	}

	headerArr := make([]string, len(headerMap), len(headerMap))
	signedHeadersArr := make([]string, len(headerMap), len(headerMap))
	headerIdx := 0
	for k, v := range headerMap {
		loweredHeaderName := strings.ToLower(k)
		headerArr[headerIdx] = fmt.Sprintf("%s:%s", loweredHeaderName, strings.Join(v, ","))
		signedHeadersArr[headerIdx] = loweredHeaderName
		headerIdx++
	}
	sort.Strings(headerArr)
	sort.Strings(signedHeadersArr)
	signedHeaders := strings.Join(signedHeadersArr, ";")

	canonicalRequest := strings.Join([]string{
		http.MethodPost,
		"/",
		"", // No query string
		strings.Join(headerArr, "\n") + "\n",
		signedHeaders,
		payloadHash,
	}, "\n")

	// Create the string to sign
	hashCanonicalRequest := sha256.Sum256([]byte(canonicalRequest))
	credentialScope := strings.Join([]string{
		now.Format(amzDateFormat),
		region,
		service,
		aws4Request,
	}, "/")
	stringToSign := a.makeSignature(
		now,
		credentialScope,
		hex.EncodeToString(hashCanonicalRequest[:]),
		region,
		service,
		creds.SecretAccessKey,
		algorithm,
	)

	// Create the authorization header
	credential := strings.Join([]string{
		creds.AccessKeyID,
		credentialScope,
	}, "/")
	authHeader := fmt.Sprintf("%s Credential=%s, SignedHeaders=%s, Signature=%s",
		algorithm, credential, signedHeaders, stringToSign)

	headerMap["Authorization"] = []string{authHeader}
	headerMap["User-Agent"] = []string{GetUserAgent()}
	headersJSON, err := json.Marshal(headerMap)
	if err != nil {
		return nil, err
	}

	return &SigningData{
		HeadersEncoded: base64.StdEncoding.EncodeToString(headersJSON),
		BodyEncoded:    base64.StdEncoding.EncodeToString([]byte(requestBody)),
		Method:         http.MethodPost,
		URLEncoded:     base64.StdEncoding.EncodeToString([]byte(stsFullURL)),
	}, nil
}

func (a *AWSAuth) makeSignature(t time.Time, credentialScope, payloadHash, region, service, secretAccessKey, algorithm string) string {
	// Create the string to sign
	stringToSign := strings.Join([]string{
		algorithm,
		t.Format(amzDateTimeFormat),
		credentialScope,
		payloadHash,
	}, "\n")

	// Create the signing key
	kDate := hmac256(t.Format(amzDateFormat), []byte("AWS4"+secretAccessKey))
	kRegion := hmac256(region, kDate)
	kService := hmac256(service, kRegion)
	kSigning := hmac256(aws4Request, kService)

	// Sign the string
	signature := hex.EncodeToString(hmac256(stringToSign, kSigning))

	return signature
}

func hmac256(data string, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return h.Sum(nil)
}
