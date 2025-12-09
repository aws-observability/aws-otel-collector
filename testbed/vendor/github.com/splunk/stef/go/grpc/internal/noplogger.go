package internal

import (
	"context"

	"github.com/splunk/stef/go/grpc/types"
)

var _ types.Logger = &NopLogger{}

type NopLogger struct{}

func (l NopLogger) Debugf(ctx context.Context, format string, v ...interface{}) {}
func (l NopLogger) Errorf(ctx context.Context, format string, v ...interface{}) {}
