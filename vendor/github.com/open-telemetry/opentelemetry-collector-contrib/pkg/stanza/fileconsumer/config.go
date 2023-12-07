// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fileconsumer // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer"

import (
	"bufio"
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/collector/featuregate"
	"go.uber.org/zap"
	"golang.org/x/text/encoding"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/decode"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/emit"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/fingerprint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/header"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/reader"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/matcher"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/split"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/trim"
)

const (
	defaultMaxLogSize         = 1024 * 1024
	defaultMaxConcurrentFiles = 1024
	defaultEncoding           = "utf-8"
	defaultPollInterval       = 200 * time.Millisecond
	defaultFlushPeriod        = 500 * time.Millisecond
)

var allowFileDeletion = featuregate.GlobalRegistry().MustRegister(
	"filelog.allowFileDeletion",
	featuregate.StageAlpha,
	featuregate.WithRegisterDescription("When enabled, allows usage of the `delete_after_read` setting."),
	featuregate.WithRegisterReferenceURL("https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/16314"),
)

var AllowHeaderMetadataParsing = featuregate.GlobalRegistry().MustRegister(
	"filelog.allowHeaderMetadataParsing",
	featuregate.StageBeta,
	featuregate.WithRegisterDescription("When enabled, allows usage of the `header` setting."),
	featuregate.WithRegisterReferenceURL("https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/18198"),
)

// NewConfig creates a new input config with default values
func NewConfig() *Config {
	return &Config{
		IncludeFileName:         true,
		IncludeFilePath:         false,
		IncludeFileNameResolved: false,
		IncludeFilePathResolved: false,
		PollInterval:            defaultPollInterval,
		Encoding:                defaultEncoding,
		StartAt:                 "end",
		FingerprintSize:         fingerprint.DefaultSize,
		MaxLogSize:              defaultMaxLogSize,
		MaxConcurrentFiles:      defaultMaxConcurrentFiles,
		MaxBatches:              0,
		FlushPeriod:             defaultFlushPeriod,
	}
}

// Config is the configuration of a file input operator
type Config struct {
	matcher.Criteria        `mapstructure:",squash"`
	IncludeFileName         bool            `mapstructure:"include_file_name,omitempty"`
	IncludeFilePath         bool            `mapstructure:"include_file_path,omitempty"`
	IncludeFileNameResolved bool            `mapstructure:"include_file_name_resolved,omitempty"`
	IncludeFilePathResolved bool            `mapstructure:"include_file_path_resolved,omitempty"`
	PollInterval            time.Duration   `mapstructure:"poll_interval,omitempty"`
	StartAt                 string          `mapstructure:"start_at,omitempty"`
	FingerprintSize         helper.ByteSize `mapstructure:"fingerprint_size,omitempty"`
	MaxLogSize              helper.ByteSize `mapstructure:"max_log_size,omitempty"`
	MaxConcurrentFiles      int             `mapstructure:"max_concurrent_files,omitempty"`
	MaxBatches              int             `mapstructure:"max_batches,omitempty"`
	DeleteAfterRead         bool            `mapstructure:"delete_after_read,omitempty"`
	SplitConfig             split.Config    `mapstructure:"multiline,omitempty"`
	TrimConfig              trim.Config     `mapstructure:",squash,omitempty"`
	Encoding                string          `mapstructure:"encoding,omitempty"`
	FlushPeriod             time.Duration   `mapstructure:"force_flush_period,omitempty"`
	Header                  *HeaderConfig   `mapstructure:"header,omitempty"`
}

type HeaderConfig struct {
	Pattern           string            `mapstructure:"pattern"`
	MetadataOperators []operator.Config `mapstructure:"metadata_operators"`
}

// Build will build a file input operator from the supplied configuration
func (c Config) Build(logger *zap.SugaredLogger, emit emit.Callback) (*Manager, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}

	enc, err := decode.LookupEncoding(c.Encoding)
	if err != nil {
		return nil, err
	}

	splitFunc, err := c.SplitConfig.Func(enc, false, int(c.MaxLogSize))
	if err != nil {
		return nil, err
	}

	trimFunc := trim.Nop
	if enc != encoding.Nop {
		trimFunc = c.TrimConfig.Func()
	}

	return c.buildManager(logger, emit, splitFunc, trimFunc)
}

// BuildWithSplitFunc will build a file input operator with customized splitFunc function
func (c Config) BuildWithSplitFunc(logger *zap.SugaredLogger, emit emit.Callback, splitFunc bufio.SplitFunc) (*Manager, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	return c.buildManager(logger, emit, splitFunc, c.TrimConfig.Func())
}

func (c Config) buildManager(logger *zap.SugaredLogger, emit emit.Callback, splitFunc bufio.SplitFunc, trimFunc trim.Func) (*Manager, error) {
	if emit == nil {
		return nil, fmt.Errorf("must provide emit function")
	}
	var startAtBeginning bool
	switch c.StartAt {
	case "beginning":
		startAtBeginning = true
	case "end":
		startAtBeginning = false
	default:
		return nil, fmt.Errorf("invalid start_at location '%s'", c.StartAt)
	}

	enc, err := decode.LookupEncoding(c.Encoding)
	if err != nil {
		return nil, fmt.Errorf("failed to find encoding: %w", err)
	}

	var hCfg *header.Config
	if c.Header != nil {
		hCfg, err = header.NewConfig(c.Header.Pattern, c.Header.MetadataOperators, enc)
		if err != nil {
			return nil, fmt.Errorf("failed to build header config: %w", err)
		}
	}

	fileMatcher, err := matcher.New(c.Criteria)
	if err != nil {
		return nil, err
	}

	return &Manager{
		SugaredLogger: logger.With("component", "fileconsumer"),
		cancel:        func() {},
		readerFactory: reader.Factory{
			SugaredLogger: logger.With("component", "fileconsumer"),
			Config: &reader.Config{
				FingerprintSize:         int(c.FingerprintSize),
				MaxLogSize:              int(c.MaxLogSize),
				Emit:                    emit,
				IncludeFileName:         c.IncludeFileName,
				IncludeFilePath:         c.IncludeFilePath,
				IncludeFileNameResolved: c.IncludeFileNameResolved,
				IncludeFilePathResolved: c.IncludeFilePathResolved,
				DeleteAtEOF:             c.DeleteAfterRead,
				FlushTimeout:            c.FlushPeriod,
			},
			FromBeginning: startAtBeginning,
			Encoding:      enc,
			SplitFunc:     splitFunc,
			TrimFunc:      trimFunc,
			HeaderConfig:  hCfg,
		},
		fileMatcher:       fileMatcher,
		pollInterval:      c.PollInterval,
		maxBatchFiles:     c.MaxConcurrentFiles / 2,
		maxBatches:        c.MaxBatches,
		previousPollFiles: make([]*reader.Reader, 0, c.MaxConcurrentFiles/2),
		knownFiles:        []*reader.Metadata{},
	}, nil
}

func (c Config) validate() error {
	if c.DeleteAfterRead && !allowFileDeletion.IsEnabled() {
		return fmt.Errorf("`delete_after_read` requires feature gate `%s`", allowFileDeletion.ID())
	}

	if c.Header != nil && !AllowHeaderMetadataParsing.IsEnabled() {
		return fmt.Errorf("`header` requires feature gate `%s`", AllowHeaderMetadataParsing.ID())
	}

	if _, err := matcher.New(c.Criteria); err != nil {
		return err
	}

	if c.MaxLogSize <= 0 {
		return fmt.Errorf("`max_log_size` must be positive")
	}

	if c.MaxConcurrentFiles <= 1 {
		return fmt.Errorf("`max_concurrent_files` must be greater than 1")
	}

	if c.FingerprintSize < fingerprint.MinSize {
		return fmt.Errorf("`fingerprint_size` must be at least %d bytes", fingerprint.MinSize)
	}

	if c.DeleteAfterRead && c.StartAt == "end" {
		return fmt.Errorf("`delete_after_read` cannot be used with `start_at: end`")
	}

	if c.Header != nil && c.StartAt == "end" {
		return fmt.Errorf("`header` cannot be specified with `start_at: end`")
	}

	if c.MaxBatches < 0 {
		return errors.New("`max_batches` must not be negative")
	}

	enc, err := decode.LookupEncoding(c.Encoding)
	if err != nil {
		return err
	}

	if c.Header != nil {
		if _, err := header.NewConfig(c.Header.Pattern, c.Header.MetadataOperators, enc); err != nil {
			return fmt.Errorf("invalid config for `header`: %w", err)
		}
	}

	return nil
}
