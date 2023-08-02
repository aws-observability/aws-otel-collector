// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ec2 // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/aws/ec2"

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	conventions "go.opentelemetry.io/collector/semconv/v1.6.1"
	"go.uber.org/zap"

	ec2provider "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/aws/ec2"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/aws/ec2/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr   = "ec2"
	tagPrefix = "ec2.tag."
)

var _ internal.Detector = (*Detector)(nil)

type Detector struct {
	metadataProvider ec2provider.Provider
	tagKeyRegexes    []*regexp.Regexp
	logger           *zap.Logger
	rb               *metadata.ResourceBuilder
}

func NewDetector(set processor.CreateSettings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	cfg := dcfg.(Config)
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	tagKeyRegexes, err := compileRegexes(cfg)
	if err != nil {
		return nil, err
	}

	return &Detector{
		metadataProvider: ec2provider.NewProvider(sess),
		tagKeyRegexes:    tagKeyRegexes,
		logger:           set.Logger,
		rb:               metadata.NewResourceBuilder(cfg.ResourceAttributes),
	}, nil
}

func (d *Detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	if _, err = d.metadataProvider.InstanceID(ctx); err != nil {
		d.logger.Debug("EC2 metadata unavailable", zap.Error(err))
		return pcommon.NewResource(), "", nil
	}

	meta, err := d.metadataProvider.Get(ctx)
	if err != nil {
		return pcommon.NewResource(), "", fmt.Errorf("failed getting identity document: %w", err)
	}

	hostname, err := d.metadataProvider.Hostname(ctx)
	if err != nil {
		return pcommon.NewResource(), "", fmt.Errorf("failed getting hostname: %w", err)
	}

	d.rb.SetCloudProvider(conventions.AttributeCloudProviderAWS)
	d.rb.SetCloudPlatform(conventions.AttributeCloudPlatformAWSEC2)
	d.rb.SetCloudRegion(meta.Region)
	d.rb.SetCloudAccountID(meta.AccountID)
	d.rb.SetCloudAvailabilityZone(meta.AvailabilityZone)
	d.rb.SetHostID(meta.InstanceID)
	d.rb.SetHostImageID(meta.ImageID)
	d.rb.SetHostType(meta.InstanceType)
	d.rb.SetHostName(hostname)
	res := d.rb.Emit()

	if len(d.tagKeyRegexes) != 0 {
		client := getHTTPClientSettings(ctx, d.logger)
		tags, err := connectAndFetchEc2Tags(meta.Region, meta.InstanceID, d.tagKeyRegexes, client)
		if err != nil {
			return res, "", fmt.Errorf("failed fetching ec2 instance tags: %w", err)
		}
		for key, val := range tags {
			res.Attributes().PutStr(tagPrefix+key, val)
		}
	}

	return res, conventions.SchemaURL, nil
}

func getHTTPClientSettings(ctx context.Context, logger *zap.Logger) *http.Client {
	client, err := internal.ClientFromContext(ctx)
	if err != nil {
		client = http.DefaultClient
		logger.Debug("Error retrieving client from context thus creating default", zap.Error(err))
	}
	return client
}

func connectAndFetchEc2Tags(region string, instanceID string, tagKeyRegexes []*regexp.Regexp, client *http.Client) (map[string]string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:     aws.String(region),
		HTTPClient: client},
	)
	if err != nil {
		return nil, err
	}
	e := ec2.New(sess)

	return fetchEC2Tags(e, instanceID, tagKeyRegexes)
}

func fetchEC2Tags(svc ec2iface.EC2API, instanceID string, tagKeyRegexes []*regexp.Regexp) (map[string]string, error) {
	ec2Tags, err := svc.DescribeTags(&ec2.DescribeTagsInput{
		Filters: []*ec2.Filter{{
			Name: aws.String("resource-id"),
			Values: []*string{
				aws.String(instanceID),
			},
		}},
	})
	if err != nil {
		return nil, err
	}
	tags := make(map[string]string)
	for _, tag := range ec2Tags.Tags {
		matched := regexArrayMatch(tagKeyRegexes, *tag.Key)
		if matched {
			tags[*tag.Key] = *tag.Value
		}
	}
	return tags, nil
}

func compileRegexes(cfg Config) ([]*regexp.Regexp, error) {
	tagRegexes := make([]*regexp.Regexp, len(cfg.Tags))
	for i, elem := range cfg.Tags {
		regex, err := regexp.Compile(elem)
		if err != nil {
			return nil, err
		}
		tagRegexes[i] = regex
	}
	return tagRegexes, nil
}

func regexArrayMatch(arr []*regexp.Regexp, val string) bool {
	for _, elem := range arr {
		matched := elem.MatchString(val)
		if matched {
			return true
		}
	}
	return false
}
