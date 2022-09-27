// Copyright 2022 Nexient LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

// Basic imports
import (
	"regexp"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

var (
	// A simple regex pattern to match an ARN
	// https://regex101.com/library/pOfxYN
	// Go does not allow regexp to be constants
	ARN_PATTERN = regexp.MustCompile("^arn:(?P<Partition>[^:\n]*):(?P<Service>[^:\n]*):(?P<Region>[^:\n]*):(?P<AccountID>[^:\n]*):(?P<Ignore>(?P<ResourceType>[^:/\n]*)[:/])?(?P<Resource>.*)$")
)

// All methods that begin with "Test" are run as tests within a suite.
func (suite *TerraTestSuite) TestOutput() {
	output := terraform.Output(suite.T(), suite.TerraformOptions, "bucket_arn")
	// Output contains only alphanumeric characters
	suite.Regexp(ARN_PATTERN, output)
}
