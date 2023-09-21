// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package heroku // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/heroku"

import (
	"context"
	"os"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/heroku/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "heroku"
<<<<<<< HEAD

	// The time and date the release was created.
	herokuReleaseCreationTimestamp = "heroku.release.creation_timestamp"
	// The commit hash for the current release
	herokuReleaseCommit = "heroku.release.commit"
	// The unique identifier for the application
	herokuAppID = "heroku.app.id"
=======
>>>>>>> main
)

// NewDetector returns a detector which can detect resource attributes on Heroku
func NewDetector(set processor.CreateSettings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	cfg := dcfg.(Config)
	return &detector{
<<<<<<< HEAD
		logger:             set.Logger,
		resourceAttributes: cfg.ResourceAttributes,
=======
		logger: set.Logger,
		rb:     metadata.NewResourceBuilder(cfg.ResourceAttributes),
>>>>>>> main
	}, nil
}

type detector struct {
<<<<<<< HEAD
	logger             *zap.Logger
	resourceAttributes metadata.ResourceAttributesConfig
=======
	logger *zap.Logger
	rb     *metadata.ResourceBuilder
>>>>>>> main
}

// Detect detects heroku metadata and returns a resource with the available ones
func (d *detector) Detect(_ context.Context) (resource pcommon.Resource, schemaURL string, err error) {
<<<<<<< HEAD
	res := pcommon.NewResource()
	dynoID, ok := os.LookupEnv("HEROKU_DYNO_ID")
	if !ok {
		d.logger.Debug("heroku metadata unavailable", zap.Error(err))
		return res, "", nil
	}

	attrs := res.Attributes()
	if d.resourceAttributes.CloudProvider.Enabled {
		attrs.PutStr(conventions.AttributeCloudProvider, "heroku")
	}
	if d.resourceAttributes.ServiceInstanceID.Enabled {
		attrs.PutStr(conventions.AttributeServiceInstanceID, dynoID)
	}
	if d.resourceAttributes.HerokuAppID.Enabled {
		if v, ok := os.LookupEnv("HEROKU_APP_ID"); ok {
			attrs.PutStr(herokuAppID, v)
		}
	}
	if d.resourceAttributes.HerokuAppName.Enabled {
		if v, ok := os.LookupEnv("HEROKU_APP_NAME"); ok {
			attrs.PutStr(conventions.AttributeServiceName, v)
		}
	}
	if d.resourceAttributes.HerokuReleaseCreationTimestamp.Enabled {
		if v, ok := os.LookupEnv("HEROKU_RELEASE_CREATED_AT"); ok {
			attrs.PutStr(herokuReleaseCreationTimestamp, v)
		}
	}
	if d.resourceAttributes.HerokuReleaseVersion.Enabled {
		if v, ok := os.LookupEnv("HEROKU_RELEASE_VERSION"); ok {
			attrs.PutStr(conventions.AttributeServiceVersion, v)
		}
	}
	if d.resourceAttributes.HerokuReleaseCommit.Enabled {
		if v, ok := os.LookupEnv("HEROKU_SLUG_COMMIT"); ok {
			attrs.PutStr(herokuReleaseCommit, v)
		}
	}

	return res, conventions.SchemaURL, nil
=======
	dynoIDMissing := false
	if dynoID, ok := os.LookupEnv("HEROKU_DYNO_ID"); ok {
		d.rb.SetServiceInstanceID(dynoID)
	} else {
		dynoIDMissing = true
	}

	herokuAppIDMissing := false
	if v, ok := os.LookupEnv("HEROKU_APP_ID"); ok {
		d.rb.SetHerokuAppID(v)
	} else {
		herokuAppIDMissing = true
	}
	if dynoIDMissing {
		if herokuAppIDMissing {
			d.logger.Debug("Heroku metadata is missing. Please check metadata is enabled.")
		} else {
			// some heroku deployments will enable some of the metadata.
			d.logger.Debug("Partial Heroku metadata is missing. Please check metadata is supported.")
		}
	}
	if !herokuAppIDMissing {
		d.rb.SetCloudProvider("heroku")
	}
	if v, ok := os.LookupEnv("HEROKU_APP_NAME"); ok {
		d.rb.SetServiceName(v)
	}
	if v, ok := os.LookupEnv("HEROKU_RELEASE_CREATED_AT"); ok {
		d.rb.SetHerokuReleaseCreationTimestamp(v)
	}
	if v, ok := os.LookupEnv("HEROKU_RELEASE_VERSION"); ok {
		d.rb.SetServiceVersion(v)
	}
	if v, ok := os.LookupEnv("HEROKU_SLUG_COMMIT"); ok {
		d.rb.SetHerokuReleaseCommit(v)
	}

	return d.rb.Emit(), conventions.SchemaURL, nil
>>>>>>> main
}
