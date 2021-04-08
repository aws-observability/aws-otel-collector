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
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

type Matcher interface {
	Type() MatcherType
	// MatchTargets returns targets fond from the specific container.
	// One container can have multiple targets because it may have multiple ports.
	MatchTargets(task *Task, container *ecs.ContainerDefinition) ([]MatchedTarget, error)
}

type MatcherConfig interface {
	// Init validates the configuration and initializes some internal strcutrues like regexp.
	Init() error
	NewMatcher(options MatcherOptions) (Matcher, error)
}

type MatcherOptions struct {
	Logger *zap.Logger
}

type MatcherType int

const (
	MatcherTypeService MatcherType = iota + 1
	MatcherTypeTaskDefinition
	MatcherTypeDockerLabel
)

func (t MatcherType) String() string {
	switch t {
	case MatcherTypeService:
		return "service"
	case MatcherTypeTaskDefinition:
		return "task_definition"
	case MatcherTypeDockerLabel:
		return "docker_label"
	default:
		// Give it a _matcher_type suffix so people can find it by string search.
		return "unknown_matcher_type"
	}
}

type MatchResult struct {
	// Tasks are index for tasks that include matched containers
	Tasks []int
	// Containers are index for matched containers. containers should show up in the original order of the task list and container definitions.
	Containers []MatchedContainer
}

type MatchedContainer struct {
	TaskIndex      int // Index in task list
	ContainerIndex int // Index within a tasks defintion's container list
	Targets        []MatchedTarget
}

func (mc *MatchedContainer) MergeTargets(newTargets []MatchedTarget) {
	for _, newt := range newTargets {
		for _, old := range mc.Targets {
			// If port and metrics_path are same, then we treat them as same target and keep the existing one
			if old.Port == newt.Port && old.MetricsPath == newt.MetricsPath {
				continue
			}
			mc.Targets = append(mc.Targets, newt)
		}
	}
}

type MatchedTarget struct {
	MatcherType  MatcherType
	MatcherIndex int // Index within a specific matcher type
	Port         int
	MetricsPath  string
	Job          string
}

func matcherOrders() []MatcherType {
	return []MatcherType{
		MatcherTypeService,
		MatcherTypeTaskDefinition,
		MatcherTypeDockerLabel,
	}
}

func newMatchers(c Config, mOpt MatcherOptions) (map[MatcherType][]Matcher, error) {
	// We can have a registry or factory methods etc. but since we only have three type of metchers in filter.
	matcherConfigs := map[MatcherType][]MatcherConfig{
		MatcherTypeService:        servicConfigsToMatchers(c.Services),
		MatcherTypeTaskDefinition: taskDefintionConfigsToMatchers(c.TaskDefinitions),
		MatcherTypeDockerLabel:    dockerLabelConfigToMatchers(c.DockerLabels),
	}
	matchers := make(map[MatcherType][]Matcher)
	matcherCount := 0
	for tpe, cfgs := range matcherConfigs {
		for i, cfg := range cfgs {
			if err := cfg.Init(); err != nil {
				return nil, fmt.Errorf("init matcher config failed type %s index %d: %w", tpe, i, err)
			}
			matcher, err := cfg.NewMatcher(mOpt)
			if err != nil {
				return nil, fmt.Errorf("create matcher failed type %s index %d: %w", tpe, i, err)
			}
			matchers[tpe] = append(matchers[tpe], matcher)
			matcherCount++
		}
	}
	if matcherCount == 0 {
		return nil, fmt.Errorf("no matcher specified in config")
	}
	return matchers, nil
}

// a global instance because we don't care about why the container didn't match (for now).
var errNotMatched = fmt.Errorf("container not matched")

func matchContainers(tasks []*Task, matcher Matcher, matcherIndex int) (*MatchResult, error) {
	var (
		matchedTasks      []int
		matchedContainers []MatchedContainer
	)
	var merr error
	tpe := matcher.Type()
	for tIndex, t := range tasks {
		var matched []MatchedContainer
		for cIndex, c := range t.Definition.ContainerDefinitions {
			targets, err := matcher.MatchTargets(t, c)
			// NOTE: we don't stop when there is an error becaause it could be one task has invalid docker label.
			if err != nil {
				// Keep track of unexpected error
				if err != errNotMatched {
					multierr.AppendInto(&merr, err)
				}
				continue
			}
			for i := range targets {
				targets[i].MatcherType = tpe
				targets[i].MatcherIndex = matcherIndex
			}
			matched = append(matched, MatchedContainer{
				TaskIndex:      tIndex,
				ContainerIndex: cIndex,
				Targets:        targets,
			})
		}
		if len(matched) > 0 {
			matchedTasks = append(matchedTasks, tIndex)
			matchedContainers = append(matchedContainers, matched...)
		}
	}
	return &MatchResult{
		Tasks:      matchedTasks,
		Containers: matchedContainers,
	}, merr
}

// matcherContainerTargets is used by TaskDefinitionMatcher and ServiceMatcher.
// The only exception is DockerLabelMatcher because it get ports from docker lebel.
func matchContainerTargets(nameRegex *regexp.Regexp, expCfg CommonExporterConfig, container *ecs.ContainerDefinition) ([]MatchedTarget, error) {
	if nameRegex != nil && !nameRegex.MatchString(aws.StringValue(container.Name)) {
		return nil, errNotMatched
	}
	// Match based on port
	var targets []MatchedTarget
	// Only export container if it has at least one matching port.
	for _, portMapping := range container.PortMappings {
		for _, port := range expCfg.MetricsPorts {
			if aws.Int64Value(portMapping.ContainerPort) == int64(port) {
				targets = append(targets, MatchedTarget{
					Port:        port,
					MetricsPath: expCfg.MetricsPath,
					Job:         expCfg.JobName,
				})
			}
		}
	}
	return targets, nil
}
