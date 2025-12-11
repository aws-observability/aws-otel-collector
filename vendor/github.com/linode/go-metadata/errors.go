package metadata

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

// APIError is the error-set returned by the Linode API when presented with an invalid request
type APIError struct {
	Errors []APIErrorReason `json:"errors"`
}

func (e APIError) Error() string {
	result := make([]string, len(e.Errors))

	for i, msg := range e.Errors {
		result[i] = msg.Error()
	}

	return strings.Join(result, "; ")
}

// APIErrorReason is an individual invalid request message returned by the Linode API
type APIErrorReason struct {
	Reason string `json:"reason"`
	Field  string `json:"field"`
}

func (r APIErrorReason) Error() string {
	if len(r.Field) == 0 {
		return r.Reason
	}

	return fmt.Sprintf("[%s] %s", r.Field, r.Reason)
}

// Error wraps the LinodeGo error with the relevant http.Response
type Error struct {
	Response *http.Response
	Code     int
	Message  string
}

func (g Error) Error() string {
	return fmt.Sprintf("[%03d] %s", g.Code, g.Message)
}

func coupleAPIErrors(r *resty.Response, err error) (*resty.Response, error) {
	if err != nil {
		return nil, Error{Message: err.Error()}
	}

	if r.Error() == nil {
		return r, nil
	}

	apiError, ok := r.Error().(*APIError)
	if !ok {
		return nil, fmt.Errorf("returned error type is not *APIError")
	}

	if len(apiError.Errors) == 0 {
		return r, nil
	}

	return nil, &Error{
		Code:     r.RawResponse.StatusCode,
		Message:  apiError.Error(),
		Response: r.RawResponse,
	}
}
