// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tailsamplingprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/sampling"
)

func getNewAndPolicy(settings component.TelemetrySettings, config *AndCfg) (sampling.PolicyEvaluator, error) {
<<<<<<< HEAD
	var subPolicyEvaluators []sampling.PolicyEvaluator
=======
	subPolicyEvaluators := make([]sampling.PolicyEvaluator, len(config.SubPolicyCfg))
>>>>>>> main
	for i := range config.SubPolicyCfg {
		policyCfg := &config.SubPolicyCfg[i]
		policy, err := getAndSubPolicyEvaluator(settings, policyCfg)
		if err != nil {
			return nil, err
		}
<<<<<<< HEAD
		subPolicyEvaluators = append(subPolicyEvaluators, policy)
=======
		subPolicyEvaluators[i] = policy
>>>>>>> main
	}
	return sampling.NewAnd(settings.Logger, subPolicyEvaluators), nil
}

// Return instance of and sub-policy
func getAndSubPolicyEvaluator(settings component.TelemetrySettings, cfg *AndSubPolicyCfg) (sampling.PolicyEvaluator, error) {
	return getSharedPolicyEvaluator(settings, &cfg.sharedPolicyCfg)
}
