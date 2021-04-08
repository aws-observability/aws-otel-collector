// Copyright  OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ecsobserver

import (
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"go.uber.org/zap"
)

type ServiceConfig struct {
	CommonExporterConfig `mapstructure:",squash" yaml:",inline"`

	// NamePattern is mandetory, empty string means service name based match is skipped.
	NamePattern string `mapstructure:"name_pattern" yaml:"name_pattern"`
	// ContainerNamePattern is optional, empty string means all containers in that service would be exported.
	// Otherwise both service and container name petterns need to metch.
	ContainerNamePattern string `mapstructure:"container_name_pattern" yaml:"container_name_pattern"`

	nameRegex          *regexp.Regexp
	containerNameRegex *regexp.Regexp
}

func (s *ServiceConfig) Init() error {
	if s.NamePattern == "" {
		return nil
	}

	r, err := regexp.Compile(s.NamePattern)
	if err != nil {
		return fmt.Errorf("invalid name pattern %w", err)
	}
	s.nameRegex = r
	if s.ContainerNamePattern != "" {
		r, err = regexp.Compile(s.ContainerNamePattern)
		if err != nil {
			return fmt.Errorf("invalid container name pattern %w", err)
		}
		s.containerNameRegex = r
	}
	return nil
}

func (s *ServiceConfig) NewMatcher(opts MatcherOptions) (Matcher, error) {
	return &ServiceMatcher{
		logger: opts.Logger,
		cfg:    *s,
	}, nil
}

func servicConfigsToMatchers(cfgs []ServiceConfig) []MatcherConfig {
	var matchers []MatcherConfig
	for _, cfg := range cfgs {
		// NOTE: &cfg points to the temp var, whose value would end up be the last one in the slice.
		copied := cfg
		matchers = append(matchers, &copied)
	}
	return matchers
}

type ServiceNameFilter func(name string) bool

func serviceConfigsToFilter(cfgs []ServiceConfig) (ServiceNameFilter, error) {
	// If no service config, don't descibe any services
	if len(cfgs) == 0 {
		return func(name string) bool {
			return false
		}, nil
	}
	var regs []*regexp.Regexp
	for _, cfg := range cfgs {
		if cfg.NamePattern == "" {
			continue
		}
		r, err := regexp.Compile(cfg.NamePattern)
		if err != nil {
			return nil, fmt.Errorf("invalid service name pattern %q: %w", cfg.NamePattern, err)
		}
		regs = append(regs, r)
	}
	return func(name string) bool {
		for _, r := range regs {
			if r.MatchString(name) {
				return true
			}
		}
		return false
	}, nil
}

type ServiceMatcher struct {
	logger *zap.Logger
	cfg    ServiceConfig
}

func (s *ServiceMatcher) Type() MatcherType {
	return MatcherTypeService
}

func (s *ServiceMatcher) ExporterConfig() CommonExporterConfig {
	return s.cfg.CommonExporterConfig
}

func (s *ServiceMatcher) MatchTargets(t *Task, c *ecs.ContainerDefinition) ([]MatchedTarget, error) {
	if s.cfg.NamePattern == "" {
		return nil, errNotMatched
	}

	// Service info is only attached for tasks whoses services are included in config.
	// Howver, Match is called on tasks so we need to guard nil pointer.
	if t.Service == nil {
		return nil, errNotMatched
	}
	if !s.cfg.nameRegex.MatchString(aws.StringValue(t.Service.ServiceName)) {
		return nil, errNotMatched
	}
	// The rest is same as TaskDefinitionMatcher
	return matchContainerTargets(s.cfg.containerNameRegex, s.cfg.CommonExporterConfig, c)
}
