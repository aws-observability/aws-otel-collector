/*
 * Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package extraconfig

import "os"

const (
	EnvKeyRunInContainer = "RUN_IN_CONTAINER"
	EnvValTrue           = "True"
)

// IsRunningInContainer checks EnvKeyRunInContainer (i.e. RUN_IN_CONTAINER) to determine
// if the collector is running as a container.
//
// Following behaviour changes when running in container:
// - log writes to stderr instead of rotate in local file under /opt/aws #339
// - switch user based on config is ignored
func IsRunningInContainer() bool {
	return os.Getenv(EnvKeyRunInContainer) == EnvValTrue
}
