package govultr

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
)

const logsPath = "/v2/logs"

// LogsService is the interface to interact with the Logs endpoint
// Link: https://www.vultr.com/api/#tag/logs
type LogsService interface {
	List(ctx context.Context, options LogsOptions) ([]Log, *LogsMeta, *http.Response, error)
}

// LogsServiceHandler handle interaction with the server methods for the Vultr API
type LogsServiceHandler struct {
	client *Client
}

// Log represents a log entry
type Log struct {
	ResourceID   string      `json:"resource_id"`
	ResourceType string      `json:"resource_type"`
	Level        string      `json:"log_level"`
	Message      string      `json:"message"`
	Timestamp    string      `json:"timestamp"`
	Metadata     LogMetadata `json:"metadata"`
}

// LogMetadata represents a log entry's metadata
type LogMetadata struct {
	UserID          string `json:"user_id"`
	IPAddress       string `json:"ip_address"`
	UserName        string `json:"username"`
	HTTPStatusCode  int    `json:"http_status_code"`
	Method          string `json:"method"`
	RequestPath     string `json:"request_path"`
	RequestBody     string `json:"request_body"`
	QueryParameters string `json:"query_parameters"`
}

type logsBase struct {
	Logs []Log    `json:"logs"`
	Meta LogsMeta `json:"meta"`
}

// LogsMeta represent pagination data for log entries
type LogsMeta struct {
	ContinueTime    string `json:"continue_time"`
	ReturnedCount   int    `json:"returned_count"`
	UnreturnedCount int    `json:"unreturned_count"`
	TotalCount      int    `json:"total_count"`
}

// LogsOptions represents the query params for the logs list
type LogsOptions struct {
	StartTime    string `url:"start_time"`
	EndTime      string `url:"end_time"`
	LogLevel     string `url:"log_level,omitempty"`
	ResourceType string `url:"resource_type,omitempty"`
	ResourceID   string `url:"resource_id,omitempty"`
}

// List retrieves logs
func (l *LogsServiceHandler) List(ctx context.Context, options LogsOptions) ([]Log, *LogsMeta, *http.Response, error) { //nolint:gocritic,lll
	req, err := l.client.NewRequest(ctx, http.MethodGet, logsPath, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	params, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = params.Encode()

	logs := new(logsBase)
	resp, err := l.client.DoWithContext(ctx, req, logs)
	if err != nil {
		return nil, nil, resp, err
	}

	return logs.Logs, &logs.Meta, resp, nil
}
